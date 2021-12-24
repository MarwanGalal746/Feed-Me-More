package drink

type Drink struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DrinkRepository interface {
	Create(Drink) (int, error)
	Read() ([]Drink, error)
	Update(Drink, string) error
	Delete(string) error
}
