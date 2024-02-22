package model

type MyData struct {
	Name  string `json:"name"`
	Place int    `json:"place"`
	Price int    `json:"price"`
}

func (m *MyData) Names() []string {
	return []string{"Name", "Place", "Price"}
}

func (m *MyData) Iterator() []interface{} {
	return []interface{}{m.Name, m.Place, m.Price}
}
