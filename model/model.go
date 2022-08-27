package model

type ResponseProduct struct {
	Products []struct {
		ID                 int      `json:"id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		Price              int      `json:"price"`
		DiscountPercentage float64  `json:"discountPercentage"`
		Rating             float64  `json:"rating"`
		Stock              int      `json:"stock"`
		Brand              string   `json:"brand"`
		Category           string   `json:"category"`
		Thumbnail          string   `json:"thumbnail"`
		Images             []string `json:"images"`
	}
}

type ResponseGetUrl struct {
	Data []Data `json:"data"`
}

type Data struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type InjectToRedis struct {
	Inject []Inject `json:"data"`
}

type Inject struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}
