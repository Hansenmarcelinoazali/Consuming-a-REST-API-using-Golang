package service

import (
	"external_api/api/repository"
	"external_api/dbcon"
	"external_api/model"
	"fmt"
	"strconv"
)

func ServiceSavetoDb(bodyInput *model.Product, refreshToken string) (*model.ResponseInputDataSingle, error) {
	fmt.Println("====================", refreshToken)
	data := model.Product{
		// ID:    bodyInput.ID,
		ProductsourceId: bodyInput.ProductsourceId,
		Title:           bodyInput.Title,
		Price:           bodyInput.Price,
		Stock:           bodyInput.Stock,
		Rate:            bodyInput.Rate,
		CreatedBy:       refreshToken,
	}

	err := repository.RepoCreateProduct(&data, refreshToken)
	if err != nil {
		return nil, err
	}

	var modelResult model.ResponseInputDataSingle

	modelResult.DataStored = append(modelResult.DataStored, model.ProductrResponse{
		ProductsourceId: data.ProductsourceId,
		Title:           data.Title,
		Price:           data.Price,
		Stock:           data.Stock,
		Rate:            data.Rate,
	})
	return &model.ResponseInputDataSingle{
		Message:         "store data success",
		NumOfDataStored: 0,
		DataStored:      modelResult.DataStored,
	}, nil
}

func ServiceInputDatas(bodys *model.RequestProduct, refreshToken string) (*model.ResponseInputDataDb, error) {

	// inputId := model.InputUuid{
	// 	ID: uuid.NewString(),
	// }
	res, err := repository.RepoCreateProducts(bodys, refreshToken)

	fmt.Println("+++++++++++++++++++++++++++++++++++++", bodys)
	fmt.Println("service = = =", bodys)
	if err != nil {
		fmt.Println("ini error service", err)
		return nil, err
	}

	// var respondata model.ResponseInputData
	// for _, v := range bodys.Datas {

	// 	respondata.DataStored = append(respondata.DataStored, model.Product{
	// 		ID:              uuid.NewString(),
	// 		ProductsourceId: v.ID,
	// 		Title:           v.Title,
	// 		Price:           v.Price,
	// 		Stock:           v.Stock,
	// 		CreatedAt:       time.Now(),
	// 		ProductSource:   "https://dummyjson.com/products?limit=100&skip=0",
	// 	})

	// }

	var modelResult model.ResponseInputDataDb

	for _, v := range *res {
		modelResult.DataStored = append(modelResult.DataStored, model.ProductrResponse{
			ProductsourceId: v.ProductsourceId,
			Title:           v.Title,
			Price:           v.Price,
			Stock:           v.Stock,
			Rate:            v.Rate,
		})

	}

	return &model.ResponseInputDataDb{
		Message:         "store data success",
		NumOfDataStored: len(modelResult.DataStored),
		DataStored:      modelResult.DataStored,
	}, nil
}

func ServiceGetDatafromDb(filter, limit, skip string) (*dbcon.ResponseDBGet, error) {

	fmt.Println("di sini service")
	if filter != "" {
		resultfilter, err := repository.RepoGetFilter(filter)
		if err != nil {
			return nil, err
		}

		var filterResult dbcon.ResponseDBGet

		filterResult.Data = append(filterResult.Data, dbcon.ResponseServiceGetData{
			ID:    resultfilter.ID,
			Title: resultfilter.Title,
			Price: int(resultfilter.Price),
			Stock: resultfilter.Stock,
		})
		return &dbcon.ResponseDBGet{
			Data: filterResult.Data,
			Filter: dbcon.Filter{
				Total: 1,
				Limit: 0,
			},
		}, nil
	}

	resultGet, err := repository.RepoGetDatafromDB(limit, skip)
	if err != nil {
		return nil, err
	}
	fmt.Println(resultGet)
	var responseData dbcon.ResponseDBGet

	for _, v := range *resultGet {
		responseData.Data = append(responseData.Data, dbcon.ResponseServiceGetData{
			ID:    v.ID,
			Title: v.Title,
			Price: int(v.Price),
			Stock: v.Stock,
		})
	}

	conver, _ := strconv.Atoi(limit)
	return &dbcon.ResponseDBGet{
		Data: responseData.Data,
		Filter: dbcon.Filter{
			Total: len(responseData.Data),
			Limit: conver,
		},
		// Filter: []dbcon.Filter{},
	}, nil
}
