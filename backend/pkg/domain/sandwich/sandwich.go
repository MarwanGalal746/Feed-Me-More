package sandwich

type Sandwich struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SandwichRepository interface {
	Create(Sandwich) (int, error)
	Read() ([]Sandwich, error)
	Update(Sandwich, string) error
	Delete(string) error
}
