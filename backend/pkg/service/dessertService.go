package service

import "backend/pkg/domain/dessert"

type DessertService interface {
	Create(dessert.Dessert) (int, error)
	Read() ([]dessert.Dessert, error)
	Update(dessert.Dessert, string) error
	Delete(string) error
}

type DefaultDessertService struct {
	repo dessert.DessertRepository
}

func (s DefaultDessertService) Create(newDessert dessert.Dessert) (int, error) {
	return s.repo.Create(newDessert)
}

func (s DefaultDessertService) Read() ([]dessert.Dessert, error) {
	return s.repo.Read()
}

func (s DefaultDessertService) Update(targetedDessert dessert.Dessert, id string) error {
	return s.repo.Update(targetedDessert, id)
}

func (s DefaultDessertService) Delete(id string) error {
	return s.repo.Delete(id)
}

func NewDessertService(repository dessert.DessertRepository) DefaultDessertService {
	return DefaultDessertService{repo: repository}
}
