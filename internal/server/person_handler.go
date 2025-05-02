package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/TheTeemka/hhChat/internal/repo"
	"github.com/TheTeemka/hhChat/internal/service"
	"github.com/TheTeemka/hhChat/pkg/utils"
	"github.com/TheTeemka/hhChat/pkg/validator"
	"github.com/go-chi/chi"
)

type PersonHandler struct {
	personService *service.PersonService
}

func NewPersonHandler(PersonService *service.PersonService) *PersonHandler {
	return &PersonHandler{
		personService: PersonService,
	}
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	req, err := utils.DecodeJson[service.CreatePersonReq](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	req.Validate(v)
	if !v.Valid() {
		http.Error(w, v.String(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, people, true)

}

func (h *PersonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	person, err := h.personService.GetByID(id)
	if err != nil {
		slog.Error("GetByID", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, person, true)

}

func (h *PersonHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.personService.DeleteByID(id)
	if err != nil {
		slog.Error("DeleteByID", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PersonHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := utils.DecodeJson[service.UpdatePersonReq](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	req.Validate(v)
	if !v.Valid() {
		http.Error(w, v.String(), http.StatusBadRequest)
		return
	}

	p, err := h.personService.UpdateByID(id, req)
	if err != nil {
		slog.Error("UpdateByID", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.EncodeJson(w, p, true)
}
