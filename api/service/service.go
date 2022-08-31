package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"external_api/api/repository"
	"external_api/model"
	"external_api/redis"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

//usecase 1
//Is_login
func ServiceIsLogin(email, username, password string) (string, error) {

	result, err := repository.RepositoryIsLogin(email, username, password)
	if err != nil {
		return " ", err
	}
	if result.IsLogin == true {
		return "Sudah Login", nil
	}
	return " ", nil

}

func ServiceLoginUser(email, username, password string) (*model.Users, error) {

	result, err := repository.RepositoryLogin(email, username, password)
	if err != nil {
		return nil, err
	}
	return result, nil

}

//generate token access token
func GenerateAccessToken(email string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(hash[0:20]), err
}

//generate refreshtoken
func GenerateRefreshToken(email string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(hash[0:20]), err
}

//mengubah status login menjadi true saat login
func TurnToTrue(email, username, password, refreshToken string) error {
	// token, _ := GenerateRefreshToken(email)

	err := repository.RepoUpdateToTrue(email, username, password, refreshToken)
	if err != nil {
		return err
	}
	return nil
}

//set redis
func ServiceSetRedis(tokenAccess, identity string) error {

	repository.SetRedis(tokenAccess, string(identity))
	return nil
}

//endpoint checkhealtz
func ServiceCheckHealth(accesstoken string) (string, error) {

	responseCheckHealth, err := repository.RepoCheckHealth(accesstoken)
	if err != nil {
		return " ", err
	}
	return responseCheckHealth, nil
}

//service logout
func ServiceLogout(tokenAccess, refreshToken string) error {

	errDb := repository.RepoLogoutDb(refreshToken)
	if errDb != nil {
		return errDb
	}

	errRedis := repository.RepoLogoutRedis(tokenAccess)
	if errRedis != nil {
		return errRedis
	}
	return nil
}

//usecase 2
func GetRedisData(url, limit, skip string) (*model.ResponseGetUrl, error) {

	var res string
	rdb := redis.NewRedisClient()

	op2 := rdb.Get(url)
	//jika terjadi error kaan masuke ke service get data dan set
	err := op2.Err()
	if err != nil {
		res, err = ServiceGetUrl(url)
		if err != nil {
			return nil, err
		}
	} else { //jika datanya ada maka akan masuk  dan dilanjutkan untuk dipilah sesuai requirment
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

	response := resultGetRedis.Data[skips:offset]
	resultGetRedis.Data = response

	//return sebagai hasil
	return &resultGetRedis, nil
}

/* get url dan langsung set redis */
func ServiceGetUrl(url string) (string, error) {
	response, _ := http.Get(url)
	if response.StatusCode == 500 {
		return " ", errors.New(" Internal Server Error")
	} else if response.StatusCode == 400 {
		return " ", errors.New(" Bad Request")
	} else if response.StatusCode == 504 {
		return " ", errors.New(" Time Out")
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
