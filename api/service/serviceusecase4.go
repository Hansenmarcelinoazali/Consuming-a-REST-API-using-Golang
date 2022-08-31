package service

import (
	"external_api/api/repository"
	"external_api/model"
	"fmt"
)

func ServiceSavetoDb(bodyInput *model.Product, refreshToken string) (*model.Product, error) {

	data := model.Product{
		ID:    bodyInput.ID,
		Title: bodyInput.Title,
		Price: bodyInput.Price,
		Stock: bodyInput.Stock,
	}

	err := repository.RepoCreateProduct(&data, refreshToken)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ServiceInputDatas(bodys *[]model.Product, refreshToken string) (*[]model.Product, error) {

	err := repository.RepoCreateProducts(bodys, refreshToken)
	fmt.Println("service = = =",bodys)
	if err != nil {
		fmt.Println("ini error service",err)
		return nil, err
	}
	return bodys, nil
}
