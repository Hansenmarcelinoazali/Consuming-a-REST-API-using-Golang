package service

import (
	"external_api/api/repository"
	"external_api/dbcon"
	"fmt"
	"strconv"
)

func ServiceUsecase6() (*dbcon.ResponseCalculated, error) {

	result, err := repository.RepoUsecase6()
	if err != nil {
		return nil, err
	}

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

	var hold []string

	var modelResult []dbcon.Calculated

	for _, v := range result {
		if v.Rate < 1.00 {
			// result_count :=
			// xResult = "Rate 0,00 - 1,00"
			xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
			hold = append(hold, xValueConversion)
			xValueint := len(hold)
			xValue := strconv.Itoa(xValueint)
			modelResult = append(modelResult, dbcon.Calculated{
				X:     "Rate 0,00 - 1,00",
				Value: xValue,
				Fill:  "pink",
			})
		} else if v.Rate >= 1.01 && v.Rate <= 2.00 {
			// xResult = "Rate 1,01 - 2,00"
			xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
			hold = append(hold, xValueConversion)
			xValueint := len(hold)
			xValue := strconv.Itoa(xValueint)
			modelResult = append(modelResult, dbcon.Calculated{
				X:     "Rate 1,01 - 2,00",
				Value: xValue,
				Fill:  "Green",
			})
		} else if v.Rate >= 2.01 && v.Rate <= 3.00 {
			// xResult = "Rate 2,01 - 3,00"
			xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
			hold = append(hold, xValueConversion)
			xValueint := len(hold)
			xValue := strconv.Itoa(xValueint)
			modelResult = append(modelResult, dbcon.Calculated{
				X:     "Rate 2,01 - 3,00",
				Value: xValue,
				Fill:  "red",
			})
		} else if v.Rate >= 3.01 && v.Rate <= 4.00 {
			// xResult = "Rate 3,01 - 4,00"
			xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
			hold = append(hold, xValueConversion)
			xValueint := len(hold)
			xValue := strconv.Itoa(xValueint)
			fmt.Println("@@@@3-4", xValueint)
			modelResult = append(modelResult, dbcon.Calculated{
				X:     "Rate 3,01 - 4,00",
				Value: xValue,
				Fill:  "blue",
			})
			// fmt.Println("ini rate 4:", v.Rate)
		} else if v.Rate > 4.00 {
			// xResult = "Rate more than 4"
			xValueConversion := strconv.FormatFloat(v.Rate, 'E', -1, 64)
			hold = append(hold, xValueConversion)
			xValueint := len(hold)
			xValue := strconv.Itoa(xValueint)
			modelResult = append(modelResult, dbcon.Calculated{
				X:     "More than 4",
				Value: xValue,
				Fill:  "black",
			})
		}
	}

	return &dbcon.ResponseCalculated{
		Data: modelResult,
	}, nil
}
