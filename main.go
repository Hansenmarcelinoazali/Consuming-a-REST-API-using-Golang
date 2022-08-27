package main

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

import (
	"external_api/route"
)

func main() {
	// db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
