package dbcon

import (
	"time"

	"gorm.io/gorm"
)

//Get data form db
type ResponseDBGet struct {
	Data []ResponseServiceGetData `json:"data"`
	Filter Filter                 `json:"filter"`
}

type Filter struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type Products struct {
	ID               string `json:"id,omitempty"`
	Productsource_id int
	Title            string `json:"title"`
	Rate             float64
	Price            float64 `json:"price"`
	Stock            int     `json:"stock"`
	ProductSource    string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	CreatedBy        string
	UpdatedBy        string
	DeletedBy        string
}

type ResponseServiceGetData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
