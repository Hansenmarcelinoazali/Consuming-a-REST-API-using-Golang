package service

import (
	"encoding/json"
	"external_api/model"
	"external_api/redis"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetRedisData(url, limit, skip string) (*model.ResponseGetUrl, error) {

	var res string
	rdb := redis.NewRedisClient()
	// fmt.Println("redis client initialized")

	op2 := rdb.Get(url)
	err := op2.Err()
	if err != nil {
		fmt.Println("INI MASUK ERROR IF 1")
		res, err = ServiceGetUrl(url)
		if err != nil {
			return nil, err
		}
		// fmt.Println(res)
	} else {
		fmt.Println("INI MASUK BUKAN ERROR")
		res, err = op2.Result()
		if err != nil {

			fmt.Printf("unable to GET data. error: %av", err)
			return nil, nil
		}

		// log.Println("get operation success. result:", res)
	}
	// fmt.Println("================", res)

	var unmarsStruct model.ResponseProduct
	json.Unmarshal([]byte(res), &unmarsStruct)
	// fmt.Println(unmars)

	var resultGetRedis model.ResponseGetUrl

	for _, v := range unmarsStruct.Products {
		resultGetRedis.Data = append(resultGetRedis.Data, model.Data{
			ID:    v.ID,
			Title: v.Title,
			Price: v.Price,
			Stock: v.Stock,
		})

	}

	// limits, _ := strconv.Atoi(limit)
	// skips, _ := strconv.Atoi(skip)
	// offset := skips * limits

	// response := resultGetRedis.Data[skips:offset]
	// resultGetRedis.Data = response

	return &resultGetRedis, nil
}

func ServiceGetUrl(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		// fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ResponseObject model.ResponseProduct
	json.Unmarshal(responseData, &ResponseObject)

	// var ResponseGetUrl model.ResponseGetUrl
	// for _, v := range ResponseObject.Products {
	// 	ResponseGetUrl.Data = append(ResponseGetUrl.Data, model.Data{
	// 		ID:    v.ID,
	// 		Title: v.Title,
	// 		Price: v.Price,
	// 		Stock: v.Stock,
	// 	})
	// }

	marshaltoRedis, _ := json.Marshal(ResponseObject)
	// fmt.Println(marshaltoRedis)
	// fmt.Println(url)

	redis.SetRedis(url, string(marshaltoRedis))

	return string(responseData), nil
}
