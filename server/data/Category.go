package data

type Category struct {
	ID 		int
	Name 	string
}

func NewCategory(ID int, Name string) *Category {
	c := new(Category)
	c.ID = ID
	c.Name = Name
	return c
}