package service

import "backend/pkg/domain/drink"

type DrinkService interface {
	Create(drink.Drink) (int, error)
	Read() ([]drink.Drink, error)
	Update(drink.Drink, string) error
	Delete(string) error
}

type DefaultDrinkService struct {
	repo drink.DrinkRepository
}

func (s DefaultDrinkService) Create(newDrink drink.Drink) (int, error) {
	return s.repo.Create(newDrink)
}

func (s DefaultDrinkService) Read() ([]drink.Drink, error) {
	return s.repo.Read()
}

func (s DefaultDrinkService) Update(targetedDrink drink.Drink, id string) error {
	return s.repo.Update(targetedDrink, id)
}

func (s DefaultDrinkService) Delete(id string) error {
	return s.repo.Delete(id)
}

func NewDrinkService(repository drink.DrinkRepository) DefaultDrinkService {
	return DefaultDrinkService{repo: repository}
}
