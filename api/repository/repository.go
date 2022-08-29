package repository

import (
	// "ECHO-GORM/db"
	"errors"
	"external_api/db"
	"external_api/model"
	"external_api/redis"
	"fmt"
	"log"
	"time"
)

func RepositoryLogin(email, username, password string) (*model.Users, error) {
	db := db.DbManager()
	note := model.Users{}

	err := db.Raw("SELECT id, username, email, password, name, is_login, refresh_token FROM users WHERE email = ? OR username = ? AND password = ?", email, username, password).Scan(&note).Error

	return &note, err
}

func RepositoryIsLogin(email, username, password string) (*model.Users, error) {
	db := db.DbManager()

	var result model.Users

	err := db.Raw("SELECT is_login FROM users WHERE email = ? OR username = ? AND password = ?", email, username, password).Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil

}

func RepoUpdateToTrue(email, username, password, token string) error {
	db := db.DbManager()

	var kirim error
	db.Model(&model.Users{}).Where("email = ? or username = ? AND password = ?", email, username, password).Updates(map[string]interface{}{"is_login": true, "refresh_token": token})

	return kirim
}

// repository check health
func RepoCheckHealth(accessToken string) (string, error) {
	rdb := redis.NewRedisClient()
	fmt.Println("redis client initialized")

	op2 := rdb.Get(accessToken)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return "", nil
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return "", nil
	}
	log.Println("get operation success. result:", res)

	return res, nil

}

//logoutDB
func RepoLogoutDb(refreshToken string) error {
	db := db.DbManager()
	updaterefresh := db.Model(&model.Users{}).Where("refresh_token = ?", refreshToken).Update("is_login", false)
	if updaterefresh != nil {
		fmt.Println("update error", updaterefresh)
	}
	return nil
}

func RepoLogoutRedis(accessToken string) error {
	rdb := redis.NewRedisClient()
	err := rdb.Del(accessToken)
	if err != nil {
		return nil
	}
	return errors.New("anda berhasil logout")
}
func SetRedis(url string, productdata string) {

	rdb := redis.NewRedisClient()
	fmt.Println("redis client initialized")

	key := url
	data := productdata
	ttl := time.Duration(600) * time.Second

	// store data using SET command
	op1 := rdb.Set(key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

}
