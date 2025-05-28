package service

import (
	"context"
	"time"

	"github.com/TheTeemka/TaskNameManager/internal/repo"
)

type PersonService struct {
	repo repo.PersonRepository
}

func NewPersonService(rep repo.PersonRepository) *PersonService {
	return &PersonService{
		repo: rep,
	}
}

func (s *PersonService) CreatePerson(req *CreatePersonReq) (*repo.Person, error) {
	p := &repo.Person{
		Name:    req.Name,
		Surname: req.Surname,
	}

	var err error
	p.Age, err = fetchAge(p.Name)
	if err != nil {
		return nil, err
	}

	p.Gender, err = fetchGender(p.Name)
	if err != nil {
		return nil, err
	}

	p.Nationality, err = fetchNationality(p.Name)
	if err != nil {
		return nil, err
	}

	p, err = s.repo.Create(context.Background(), p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PersonService) GetByFilters(filters *repo.Filter) ([]*repo.Person, error) {
	p, err := s.repo.GetByFilters(context.Background(), filters)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) GetByID(id int64) (*repo.Person, error) {
	p, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) DeleteByID(id int64) error {
	err := s.repo.DeleteByID(context.Background(), id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonService) UpdateByID(id int64, req *UpdatePersonReq) (*repo.Person, error) {
	p, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Surname != "" {
		p.Surname = req.Surname
	}
	if req.Nationality != "" {
		p.Nationality = req.Nationality
	}
	if req.Gender != "" {
		p.Gender = req.Gender
	}
	if req.Age != 0 {
		p.Age = req.Age
	}
	p.UpdatedAt = time.Now()

	err = s.repo.Update(context.Background(), p)
	return p, nil
}
