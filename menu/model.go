package menu

type Drink struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Price            int `json:"price"`
	InStock          bool   `json:"inStock"`
	ContainsCaffeine bool   `json:"containsCaffeine"`
	Volume           int `json:"volume"`
	Description      string `json:"description"`
}
type DrinkCreate struct {
	Name             string `json:"name" binding:"required"`
	Price            int `json:"price" binding:"required"`
	InStock          bool   `json:"inStock"`
	ContainsCaffeine bool   `json:"containsCaffeine"`
	Volume           int `json:"volume" binding:"required"`
	Description      string `json:"description,omitempty"`
}

type DrinkUpdate struct {
	ID               int    `json:"id"`
	Name             string `json:"name" binding:"required"`
	Price            int `json:"price" binding:"required"`
	InStock          bool   `json:"inStock"`
	ContainsCaffeine bool   `json:"containsCaffeine"`
	Volume           int `json:"volume" binding:"required"`
	Description      string `json:"description,omitempty"`
}

type DrinkList struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int `json:"price"`
}
