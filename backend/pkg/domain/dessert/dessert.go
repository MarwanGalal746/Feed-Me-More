package dessert

type Dessert struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DessertRepository interface {
	Create(Dessert) (int, error)
	Read() ([]Dessert, error)
	Update(Dessert, string) error
	Delete(string) error
}
