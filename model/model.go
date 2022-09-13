package model

import (
	"time"
)

//usecase 1
type Login struct {
	EmailorUsername string `json:"email/username" form:"email/username" query:"email/username"`
	Password        string `json:"password" form:"password" query:"password"`
}

type ResponseToken struct {
	TokenAccess  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Response struct {
	StatusCode   int         `json:"statusCode,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Token        interface{} `json:"token_sekarang,omitempty"`
	RefreshToken interface{} `json:"token_refresh,omitempty"`
}

type Users struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	IsLogin      bool   `json:"is_login"`
	RefreshToken string `json:"refresh_token"`
}

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
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}

//usecase 4
type ResponseMessage struct {
	MessageResponse []ResponseInputData `json:"Response"`
}

type ResponseInputData struct {
	Message         string    `json:"message"`
	NumOfDataStored int       `json:"num_of_data_stored"`
	DataStored      []Product `json:"data_stored"`
}

type ResponseInputDataSingle struct {
	Message         string             `json:"message"`
	NumOfDataStored int                `json:"num_of_data_stored"`
	DataStored      []ProductrResponse `json:"data_stored"`
}

type Product struct {
	ID              string
	ProductsourceId int       `json:"id,"`
	Title           string    `json:"title"`
	Price           int       `json:"price"`
	Stock           int       `json:"stock"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	ProductSource   string    `json:"product_source"`
	Rate            float64   `json:"rating"`
}

type Products struct {
	ID              string
	ProductsourceId int       `json:"id"`
	Title           string    `json:"title"`
	Price           int       `json:"price"`
	Stock           int       `json:"stock"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
	ProductSource   string    `json:"product_source"`
	Rate            float64   `json:"rating"`
}

type ResponseInputDataDb struct {
	Message         string             `json:"message"`
	NumOfDataStored int                `json:"num_of_data_stored"`
	DataStored      []ProductrResponse `json:"data_stored"`
}

type ProductrResponse struct {
	ProductsourceId int     `json:"id"`
	Title           string  `json:"title"`
	Price           int     `json:"price"`
	Stock           int     `json:"stock"`
	Rate            float64 `json:"rating"`
}

type ResponseInputDataS struct {
	Message         string     `json:"message"`
	NumOfDataStored int        `json:"num_of_data_stored"`
	DataStored      []Products `json:"data_stored"`
}

type RequestProduct struct {
	Data  Product    `json:"product"`
	Datas []Products `json:"products"`
}
