package repository

import (
	"external_api/db"
	"external_api/dbcon"
)

func RepoUsecase6() ([]*dbcon.Products, error) {

	db := db.DbManager()
	var Result []*dbcon.Products

	err := db.Raw("SELECT id, productsource_id, title, rate, price, stock FROM products").Scan(&Result).Error

	if err != nil {
		return nil, err
	}
	// fmt.Println("ini repo usecase 6",Result)
	return Result, nil
}
