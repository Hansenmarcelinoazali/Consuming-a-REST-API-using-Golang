package service

import (
	"encoding/json"
	"external_api/api/repository"
	"external_api/model"
	"external_api/redis"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetRedisData(url, limit, skip string) (*model.ResponseGetUrl, error) {

	var res string
	rdb := redis.NewRedisClient()

	op2 := rdb.Get(url)
	//jika terjadi error kaan masuke ke service get data dan set
	err := op2.Err()
	if err != nil {
		fmt.Println("INI MASUK ERROR IF 1")
		res, err = ServiceGetUrl(url)
		if err != nil {
			return nil, err
		}
	} else { //jika datanya ada maka akan masuk  dan dilanjutkan untuk dipilah sesuai requirment
		fmt.Println("INI MASUK BUKAN ERROR")
		res, err = op2.Result()
		if err != nil {

			fmt.Printf("unable to GET data. error: %av", err)
			return nil, nil
		}
	}

	//karena data yang masuk dalam bentuk string json maka di ubah dlu ke struct dengan unmarshal
	var unmarsStruct model.ResponseProduct
	json.Unmarshal([]byte(res), &unmarsStruct)

	var resultGetRedis model.ResponseGetUrl

	for _, v := range unmarsStruct.Products {
		resultGetRedis.Data = append(resultGetRedis.Data, model.Data{
			ID:    v.ID,
			Title: v.Title,
			Price: v.Price,
			Stock: v.Stock,
		})

	}

	//pagination
	limits, _ := strconv.Atoi(limit)
	skips, _ := strconv.Atoi(skip)
	offset := skips + limits
	fmt.Println(limits)
	fmt.Println(skips)

	//pemilahan berdasarkan skip dan limit
	response := resultGetRedis.Data[skips:offset]
	resultGetRedis.Data = response
	fmt.Println(response)

	//return sebagai hasil
	return &resultGetRedis, nil
}

/* get url dan langsung set redis */
func ServiceGetUrl(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {

		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ResponseObject model.ResponseProduct
	json.Unmarshal(responseData, &ResponseObject)

	marshaltoRedis, _ := json.Marshal(ResponseObject)

	repository.SetRedis(url, string(marshaltoRedis))

	return string(responseData), nil
}
