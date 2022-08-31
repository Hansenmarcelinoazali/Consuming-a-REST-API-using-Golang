package repository

import (
	"external_api/db"
	"external_api/model"
	"fmt"
	"time"
	// "github.com/google/uuid"
)

func RepoCreateProduct(data *model.Product, refresh_token string) error {
	db := db.DbManager()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&model.Product{
		ID:              data.ID,
		ProductsourceId: data.ProductsourceId, //
		Title:           data.Title,
		Price:           data.Price,
		Stock:           data.Stock,
		CreatedAt:       time.Now(),
		CreatedBy:       refresh_token,
		ProductSource:   "https://dummyjson.com/products?limit=100&skip=0",
		Rate:            data.Rate, //
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func RepoCreateProducts(data *model.RequestProduct, refresh_token string) error {
	db := db.DbManager()

	tx := db.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()
	if err := tx.Error; err != nil {
		return err
	}

	for _, v := range data.Datas {
		v.CreatedAt = time.Now()
		if err := tx.Create(v).Error; err != nil {
			fmt.Println("ini repo", err)
			tx.Rollback()
			return err
		}

	}

	return tx.Commit().Error
}
