package service

import (
	"external_api/api/repository"
	"external_api/dbcon"
	"fmt"
	"strconv"
)

func ServiceUsecase6() (*dbcon.ResponseCalculated, error) {

	rate0 := float32(0.00)
	rate1 := float32(1.00)

	result, err := repository.RepoUsecase6(rate0, rate1)
	if err != nil {
		return nil, err
	}

	resultStr := strconv.Itoa(0)
	rate0str := strconv.Itoa(int(rate0))
	rate1Str := strconv.Itoa(int(rate1))
	x := fmt.Sprint(rate0str, " sampai ", rate1Str)

	modelService := dbcon.Calculated{
		X:     x,
		Value: resultStr,
		Fill:  "",
	}

	fmt.Println(modelService, result)

	rate2 := float32(1.01)
	rate3 := float32(2.00)

	result2, err := repository.RepoUsecase6(rate2, rate3)
	if err != nil {
		return nil, err
	}

	result2Str := strconv.Itoa(0)
	z := fmt.Sprint(rate2, " sampai ", rate3)

	modelService2 := dbcon.Calculated{
		X:     z,
		Value: result2Str,
		Fill:  "",
	}
	fmt.Println(modelService2, result2)

	//
	ratea := float32(2.01)
	rateb := float32(3.00)

	result2to3s, err := repository.RepoUsecase6(ratea, rateb)
	if err != nil {
		return nil, err
	}

	result2to3 := strconv.Itoa(0)
	p := fmt.Sprint(ratea, " sampai ", rateb)

	modelService2to3 := dbcon.Calculated{
		X:     p,
		Value: result2to3,
		Fill:  "",
	}

	fmt.Println(modelService2to3, result2to3s)

	//
	rate3a := float32(3.01)
	rate4b := float32(4.00)

	result3to4a, err := repository.RepoUsecase6(rate3a, rate4b)
	if err != nil {
		return nil, err
	}

	result3to4s := strconv.Itoa(0)
	f := fmt.Sprint(ratea, " sampai ", rateb)

	modelService3to4 := dbcon.Calculated{
		X:     f,
		Value: result3to4s,
		Fill:  "",
	}

	fmt.Println(modelService3to4, result3to4a)

	//
	ratemorethan4 := float32(4.00)
	resultmorethan4, err := repository.RepoUsecase6morethan4(ratemorethan4)
	if err != nil {
		return nil, err
	}
	resultm4 := strconv.Itoa(0)
	c := fmt.Sprint(" lebih dari ", ratemorethan4)

	modelServicemore4 := dbcon.Calculated{
		X:     c,
		Value: resultm4,
		Fill:  "",
	}

	fmt.Println(modelServicemore4, resultmorethan4)

	// var xResult1 string
	// var xValue1 string

	// var xResult2 string
	// var xValue2 string

	// var xResult3 string
	// var xValue3 string

	// var xResult4 string
	// var xValue4 string

	// var xResultmore string
	// var xValuemore string
	// var hold0 []string
	// var hold1 []string
	// var hold2 []string
	// var hold3to4 []string

	// var holdmore4 []string
	// var count int

	// var modelResult []dbcon.Calculated

	// // var xValueint int

	// for _, v := range result {
	// 	if v.Rate < 1.00 {
	// 		// result_count :=
	// 		// xResult = "Rate 0,00 - 1,00"
	// 		xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
	// 		hold0 = append(hold0, xValueConversion)
	// 		xValueint := len(hold0)
	// 		xValue := strconv.Itoa(xValueint)
	// 		modelResult = append(modelResult, dbcon.Calculated{
	// 			X:     "Rate 0,00 - 1,00",
	// 			Value: xValue,
	// 			Fill:  "pink",
	// 		})
	// 	} else if v.Rate >= 1.01 && v.Rate <= 2.00 {
	// 		// xResult = "Rate 1,01 - 2,00"
	// 		xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
	// 		hold1 = append(hold1, xValueConversion)
	// 		xValueint := len(hold1)
	// 		xValue := strconv.Itoa(xValueint)
	// 		modelResult = append(modelResult, dbcon.Calculated{
	// 			X:     "Rate 1,01 - 2,00",
	// 			Value: xValue,
	// 			Fill:  "Green",
	// 		})
	// 	} else if v.Rate >= 2.01 && v.Rate <= 3.00 {
	// 		// xResult = "Rate 2,01 - 3,00"
	// 		xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
	// 		hold2 = append(hold2, xValueConversion)
	// 		xValueint := len(hold2)
	// 		xValue := strconv.Itoa(xValueint)
	// 		modelResult = append(modelResult, dbcon.Calculated{
	// 			X:     "Rate 2,01 - 3,00",
	// 			Value: xValue,
	// 			Fill:  "red",
	// 		})
	// 	} else if v.Rate >= 3.01 && v.Rate <=count := 0 4.00 {
	// 		// xResult = "Rate 3,01 - 4,00"
	// 		xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
	// 		hold3to4 = append(hold3to4, xValueConversion)
	// 		xValueint := len(hold3to4)
	// 		xValue := strconv.Itoa(xValueint)
	// 		fmt.Println("@@@@3-4", xValueint)
	// 		modelResult = append(modelResult, dbcon.Calculated{
	// 			X:     "Rate 3,01 - 4,00",
	// 			Value: xValue,
	// 			Fill:  "blue",
	// 		})
	// 		// fmt.Println("ini rate 4:", v.Rate)
	// 	} else if v.Rate > 4.00 {
	// 		// xResult = "Rate more than 4"
	// 		xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
	// 		holdmore4 = append(holdmore4, xValueConversion)
	// 		count += 1
	// 		total := strconv.Itoa(count)
	// 		fmt.Println(total)
	// 		// xValueint = len(holdmore4) //
	// 		// xValue := strconv.Itoa(xValueint)
	// 		// fmt.Println("@@@@3-more", xValue)
	// 		// modelResult = append(modelResult, dbcon.Calculated{
	// 		// 	X:     "More than 4",
	// 		// 	Value: xValue,
	// 		// 	Fill:  "black",
	// 		// })
	// 	}
	// }

	return nil, nil
}
