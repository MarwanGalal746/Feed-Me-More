package service

import "backend/pkg/domain/sandwich"

type SandwichService interface {
	Create(sandwich.Sandwich) (int, error)
	Read() ([]sandwich.Sandwich, error)
	Update(sandwich.Sandwich, string) error
	Delete(string) error
}

type DefaultSandwichService struct {
	repo sandwich.SandwichRepository
}

func (s DefaultSandwichService) Create(newSandwich sandwich.Sandwich) (int, error) {
	return s.repo.Create(newSandwich)
}

func (s DefaultSandwichService) Read() ([]sandwich.Sandwich, error) {
	return s.repo.Read()
}

func (s DefaultSandwichService) Update(targetedSandwich sandwich.Sandwich, id string) error {
	return s.repo.Update(targetedSandwich, id)
}

func (s DefaultSandwichService) Delete(id string) error {
	return s.repo.Delete(id)
}

func NewSandwichService(repository sandwich.SandwichRepository) DefaultSandwichService {
	return DefaultSandwichService{repo: repository}
}
