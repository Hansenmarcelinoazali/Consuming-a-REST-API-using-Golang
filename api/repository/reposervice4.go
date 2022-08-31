package repository

import (
	"external_api/db"
	"external_api/model"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		ID:              uuid.NewString(),
		ProductsourceId: data.ProductsourceId,
		Title:           data.Title,
		Price:           data.Price,
		Stock:           data.Stock,
		CreatedAt:       time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func RepoCreateProducts(data *[]model.Product, refresh_token string) error {
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

	for _, v := range *data {
		v.CreatedAt = time.Now()
		if err := tx.Create(v).Error; err != nil {
			fmt.Println("ini repo", err)
			tx.Rollback()
			return err
		}

	}

	return tx.Commit().Error
}
