package main

//type MealService interface {
//	Create(Meal) (int, error)
//	Read() ([]Meal, error)
//	Update(Meal, string) error
//	Delete(string) error
//}
//
//type DefaultMealService struct {
//	repo MealRepository
//}
//
//func (s DefaultMealService) Create(newMeal Meal) (int, error) {
//	return s.repo.Create(newMeal)
//}
//
//func (s DefaultMealService) Read() ([]Meal, error) {
//	return s.repo.Read()
//}
//
//func (s DefaultMealService) Update(targetedMeal Meal, id string) error {
//	return s.repo.Update(targetedMeal, id)
//}
//
//func (s DefaultMealService) Delete(id string) error {
//	return s.repo.Delete(id)
//}
//
//func NewMealService(repository MealRepository) DefaultMealService {
//	return DefaultMealService{repo: repository}
//}
