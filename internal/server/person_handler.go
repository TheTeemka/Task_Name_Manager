package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/TheTeemka/hhChat/internal/repo"
	"github.com/TheTeemka/hhChat/internal/service"
	"github.com/TheTeemka/hhChat/pkg/utils"
	"github.com/TheTeemka/hhChat/pkg/validator"
	"github.com/go-chi/chi"
)

var ErrNotFound = errors.New("resourse not found")

type PersonHandler struct {
	personService *service.PersonService
}

func NewPersonHandler(PersonService *service.PersonService) *PersonHandler {
	return &PersonHandler{
		personService: PersonService,
	}
}

// @Summary Create a person
// @Description Create a new person in the database
// @Tags People
// @Accept json
// @Produce json
// @Param person body service.CreatePersonReq true "Person to create"
// @Success 201 {object} repo.Person
// @Failure 400 {object} ErrorWrapper "Bad Request"
// @Failure 500 {object} ErrorWrapper "Internal Server Error"
// @Router /people [post]
func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	req, err := utils.DecodeJson[service.CreatePersonReq](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	req.Validate(v)
	if !v.Valid() {
		ErrorResponseMap(w, v.Errors, http.StatusBadRequest)
		return
	}

	p, err := h.personService.CreatePerson(req)
	if err != nil {
		slog.Error("CreatePerson", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.EncodeJson(w, p, true)
}

// @Summary Get people by filters
// @Description Retrieve a list of people based on query parameters
// @Tags People
// @Accept json
// @Produce json
// @Param name query string false "Filter by name"
// @Param surname query string false "Filter by surname"
// @Param age query int false "Filter by age"
// @Param gender query string false "Filter by gender"
// @Param nationality query string false "Filter by nationality"
// @Success 200 {array} repo.Person
// @Failure 400 {object} ErrorWrapper "Bad Request"
// @Failure 500 {object} ErrorWrapper "Internal Server Error"
// @Router /people [get]
func (h *PersonHandler) GetByFilters(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	filter := repo.NewFilters()
	v := validator.New()
	filter.ParseURL(vals, v)
	if !v.Valid() {
		http.Error(w, v.String(), http.StatusBadRequest)
		return
	}

	people, err := h.personService.GetByFilters(filter)
	if err != nil {
		slog.Error("GetByFilters", "error", err)
		ErrorResponse(w, "Failed to get people", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, people, true)

}

// @Summary Get a person by ID
// @Description Retrieve a single person by their ID
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} repo.Person
// @Failure 400 {object} ErrorWrapper "Bad Request"
// @Failure 404 {object} ErrorWrapper "Not Found"
// @Failure 500 {object} ErrorWrapper "Internal Server Error"
// @Router /people/{id} [get]
func (h *PersonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := h.personService.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ErrorResponse(w, ErrNotFound.Error(), http.StatusNotFound)
		} else {
			slog.Error("GetByID", "error", err)
			ErrorResponse(w, "failed to get person", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, person, true)

}

// @Summary Delete a person by ID
// @Description Delete a person from the database by their ID
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorWrapper "Bad Request"
// @Failure 404 {object} ErrorWrapper "Not Found"
// @Failure 500 {object} ErrorWrapper "Internal Server Error"
// @Router /people/{id} [delete]
func (h *PersonHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.personService.DeleteByID(id)
	if err != nil {
		slog.Error("DeleteByID", "error", err)
		ErrorResponse(w, "failed to delete person", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Update a person by ID
// @Description Update the details of a person by their ID
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param person body service.UpdatePersonReq true "Updated person details"
// @Success 200 {object} repo.Person
// @Failure 400 {object} ErrorWrapper "Bad Request"
// @Failure 404 {object} ErrorWrapper "Not Found"
// @Failure 500 {object} ErrorWrapper "Internal Server Error"
// @Router /people/{id} [patch]
func (h *PersonHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := utils.DecodeJson[service.UpdatePersonReq](r.Body)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := validator.New()
	req.Validate(v)
	if !v.Valid() {
		ErrorResponseMap(w, v.Errors, http.StatusBadRequest)
		return
	}

	p, err := h.personService.UpdateByID(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ErrorResponse(w, ErrNotFound.Error(), http.StatusNotFound)
		} else {
			slog.Error("GetByID", "error", err)
			ErrorResponse(w, "failed to update person", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, p, true)
}

type ErrorWrapper struct {
	Errors map[string]string `json:"errors"`
}

func ErrorResponse(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorWrapper{
		Errors: map[string]string{
			"msg": msg,
		},
	})
}

func ErrorResponseMap(w http.ResponseWriter, mp map[string]string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorWrapper{
		Errors: mp,
	})
}
