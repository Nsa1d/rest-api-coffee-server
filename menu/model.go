package menu

type Drink struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	InStock          bool   `json:"inStock"`
	ContainsCaffeine bool   `json:"containsCaffeine"`
	Volume           int    `json:"volume"`
	Description      string `json:"description"`
}

type DrinkList struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
