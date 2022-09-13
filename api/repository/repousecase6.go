package repository

import (
	"external_api/db"
	"external_api/dbcon"
	"fmt"
)

func RepoUsecase6(dari, sampai float32) (int, error) {

	db := db.DbManager()
	var Result *dbcon.Products
	var results []*dbcon.Products
	// var Count int
	fmt.Println("masuk", dari, sampai)

	err := db.Model(Result).Where("rate BETWEEN ? AND ?", dari, sampai).Find(&results).Error
	if err != nil {
		return 0, nil
	}
	fmt.Printf("hasil: %+v/n", results)

	return len(results), nil
}

func RepoUsecase6morethan4(dari float32) (int, error) {

	db := db.DbManager()
	var Result *dbcon.Products
	var results []*dbcon.Products
	// var Count int
	// fmt.Println("masuk", dari, sampai)

	err := db.Model(Result).Where("rate > ?", dari).Find(&results).Error
	if err != nil {
		return 0, nil
	}
	fmt.Printf("hasil: %+v/n", results)

	return len(results), nil
}
