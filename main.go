package main

import (
	// "ECHO-GORM/db"
	"external_api/db"
	"external_api/route"
	// "github.com/Lionparcel/common-backend/sql/db"
)

// func main() {
// 	url := "https://dummyjson.com/products?limit=100&skip=0"
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}

// 	responseData, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var ResponseObject model.ResponseProduct
// 	json.Unmarshal(responseData, &ResponseObject)

// 	data := string(responseData)

// 	redis.Redis(url, data)

// 	fmt.Println(responseData)
// }

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
