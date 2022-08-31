package service

import (
	"external_api/api/repository"
	"external_api/model"
	"fmt"
	"time"
)

func ServiceSavetoDb(bodyInput *model.Product, refreshToken string) (*model.ResponseInputData, error) {

	data := model.Product{
		// ID:    bodyInput.ID,
		Title: bodyInput.Title,
		Price: bodyInput.Price,
		Stock: bodyInput.Stock,
		Rate:  bodyInput.Rate,
	}

	err := repository.RepoCreateProduct(&data, refreshToken)
	if err != nil {
		return nil, err
	}
	return &model.ResponseInputData{
		Message:         "store data success",
		NumOfDataStored: 0,
		DataStored:      []model.Product{*bodyInput},
	}, nil
}

func ServiceInputDatas(bodys *model.RequestProduct, refreshToken string) (*model.ResponseInputData, error) {

	err := repository.RepoCreateProducts(bodys, refreshToken)
	fmt.Println("service = = =", bodys)
	if err != nil {
		fmt.Println("ini error service", err)
		return nil, err
	}
	var respondata model.ResponseInputData
	for _, v := range bodys.Datas {
		respondata.DataStored = append(respondata.DataStored, model.Product{
			// ID:              v.ID,
			ProductsourceId: v.ProductsourceId,
			Title:           v.Title,
			Price:           v.Price,
			Stock:           v.Stock,
			CreatedAt:       time.Now(),
			ProductSource:   "https://dummyjson.com/products?limit=100&skip=0",
		})

	}

	return &model.ResponseInputData{
		Message:         "store data success",
		NumOfDataStored: 0,
		DataStored:      respondata.DataStored,
	}, nil
}
