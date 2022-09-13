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
	fmt.Println("@@@@@@@@@@@@@", result)

	resultStr := strconv.Itoa(result)
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

	result2Str := strconv.Itoa(result2)
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

	result2to3 := strconv.Itoa(result2to3s)
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

	result3to4s := strconv.Itoa(result3to4a)
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
	fmt.Println("\n@@@@@@@@@@@@@", result3to4s)
	resultm4 := strconv.Itoa(resultmorethan4)
	c := fmt.Sprint(" lebih dari ", ratemorethan4)

	modelServicemore4 := dbcon.Calculated{
		X:     c,
		Value: resultm4,
		Fill:  "",
	}

	fmt.Println(modelServicemore4, resultmorethan4)

	var modelResponseRes dbcon.ResponseCalculated

	modelResponseRes.Data = append(modelResponseRes.Data, dbcon.Calculated{
		X:     modelService.X,
		Value: modelService.Value,
		Fill:  modelService.Fill,
	})
	modelResponseRes.Data = append(modelResponseRes.Data, dbcon.Calculated{
		X:     modelService2.X,
		Value: modelService2.Value,
		Fill:  modelService2.Fill,
	})
	modelResponseRes.Data = append(modelResponseRes.Data, dbcon.Calculated{
		X:     modelService2to3.X,
		Value: modelService2to3.Value,
		Fill:  modelService2to3.Fill,
	})
	modelResponseRes.Data = append(modelResponseRes.Data, dbcon.Calculated{
		X:     modelService3to4.X,
		Value: modelService3to4.Value,
		Fill:  modelService3to4.Fill,
	})
	modelResponseRes.Data = append(modelResponseRes.Data, dbcon.Calculated{
		X:     modelServicemore4.X,
		Value: modelServicemore4.Value,
		Fill:  modelServicemore4.Fill,
	})

	return &dbcon.ResponseCalculated{modelResponseRes.Data}, nil
}
