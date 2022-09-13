package dbcon

import (
	"time"

	"gorm.io/gorm"
)

//Get data form db
type ResponseDBGet struct {
	Data   []ResponseServiceGetData `json:"data"`
	Filter Filter                   `json:"filter"`
}

type Filter struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type Products struct {
	ID              string  `json:"id,omitempty"`
	ProductsourceId int     `json:"productsource_id,omitempty"`
	Title           string  `json:"title"`
	Rate            float64 `json:"rating"`
	Price           float64 `json:"price"`
	Stock           int     `json:"stock"`
	ProductSource   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	CreatedBy       string
	UpdatedBy       string
	DeletedBy       string
}

// type ResponseRepoU6 struct {
// 	Data []Products `json:"data"`
// }

// type RepoProducts struct {
// 	Id               string
// 	Productsource_id int
// 	Title            string
// 	Rate             float64
// 	Price            float64
// 	Stock            int
// 	ProductSource    string
// 	CreatedAt        time.Time
// 	UpdatedAt        time.Time
// 	DeletedAt        gorm.DeletedAt
// 	CreatedBy        string
// 	UpdatedBy        string
// 	DeletedBy        string
// }
type ResponseServiceGetData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type ResponseCalculated struct {
	Data []Calculated `json:"data"`
}

type Calculated struct {
	X     string `json:"x"`
	Value string `json:"value"`
	Fill  string `json:"fill"`
}
