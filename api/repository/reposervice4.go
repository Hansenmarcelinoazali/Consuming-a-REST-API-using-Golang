package repository

import (
	"external_api/db"
	"external_api/dbcon"
	"external_api/model"
	"fmt"
	"strconv"
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
		CreatedBy:       data.CreatedBy,
		ProductSource:   "https://dummyjson.com/products?limit=100&skip=0",
		Rate:            data.Rate, //
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func RepoCreateProducts(data *model.RequestProduct, refresh_token string) (*[]model.Products, error) {
	db := db.DbManager()

	tx := db.Begin()

	for _, v := range data.Datas {
		fmt.Println("INI REPO", v)

		v.CreatedAt = time.Now()
		v.ID = uuid.NewString()
		v.CreatedBy = refresh_token
		v.ProductSource = "https://dummyjson.com/products?limit=100&skip=0"

		if err := tx.Create(v).Error; err != nil {
			fmt.Println("ini repo", err)
			tx.Rollback()
			return nil, err
		}

	}

	return &data.Datas, tx.Commit().Error
}

func RepoGetDatafromDB(limit, skip string) (*[]dbcon.Products, error) {

	db := db.DbManager()

	limitcon, _ := strconv.Atoi(limit)
	skipcon, _ := strconv.Atoi(skip)

	var products []dbcon.Products

	err := db.Limit(limitcon).Offset(skipcon).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func RepoGetFilter(filter string) (*dbcon.Products, error) {

	db := db.DbManager()

	var products dbcon.Products

	db.Where("title= ? OR id = ?", filter, filter).First(&products)

	return &products, nil

}
