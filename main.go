package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	  criteriaPrice   = "price"
	  criteriaArrTime = "arrival-time"
	  criteriaDepTime = "departure-time"
)

const (
	errCriteria         = "unsupported criteria"
	errEmptyDepStation  = "empty departure station"
	errEmptyArrStation  = "empty arrival station"
	errInputDepStation  = "bad departure station input"
	errInputArrStation  = "bad arrival station input"
)

var path = "data.json"

type Trains []Train

type Train struct {
	TrainID            int       `json:"trainId"`
	DepartureStationID int       `json:"departureStationId"`
	ArrivalStationID   int       `json:"arrivalStationId"`
	Price              float32   `json:"price"`
	ArrivalTime        time.Time `json:"arrivalTime"`
	DepartureTime      time.Time `json:"departureTime"`
}

func main() {
	
	var departureStation, arrivalStation, criteria string
	
	fmt.Println("Enter the station of departure: ")
	fmt.Scanln(&departureStation)
	
	fmt.Println("Enter the station of arrival: ")
	fmt.Scanln(&arrivalStation)

	fmt.Println("Enter the criteria: ")
	fmt.Scanln(&criteria)
	
	criteria = strings.ToLower(criteria)

	result, err := FindTrains(departureStation, arrivalStation, criteria)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	if result == nil{
		fmt.Println("No trains found")	
	}

	for _, t := range result{
		fmt.Println(t)	
	}

}

//
func (t *Train) UnmarshalJSON(data []byte) error {
	
	type copyTrain Train

	req := &struct {
		ArrivalTime   string `json:"arrivalTime"`
		DepartureTime string `json:"departureTime"`
		*copyTrain
	}{
		copyTrain: (*copyTrain)(t),
	}

	err := json.Unmarshal(data, req)
	if err != nil {
		return err
	}

	layout := "15:04:05"

	t.ArrivalTime, err = time.Parse(layout, req.ArrivalTime)
	if err != nil {
		return err
	}

	t.DepartureTime, err = time.Parse(layout, req.DepartureTime)
	if err != nil {
		return err
	}

	return nil
}


func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	
	depStationNumber, arrStationNumber, err := inputValidation(departureStation, arrivalStation, criteria)

	if err != nil {
		return nil, err
	}
	
	trains, err := readJSON()

	if err != nil {
		return nil, err
	}
	
	result := make(Trains, 0)

	for i := range trains{
		if (trains[i].DepartureStationID == depStationNumber && trains[i].ArrivalStationID == arrStationNumber){
			result = append(result, trains[i])	
		}		
	}
	
	if len(result) == 0 {
		return nil, nil
	}

	switch criteria {
	case criteriaPrice:
		sort.SliceStable(result, func(i, j int) bool { return result[i].Price < result[j].Price })	
	case criteriaArrTime:
		sort.SliceStable(result, func(i, j int) bool { return result[i].ArrivalTime.Before(result[j].ArrivalTime) })
	case criteriaDepTime:
		sort.SliceStable(result, func(i, j int) bool { return result[i].DepartureTime.Before(result[j].DepartureTime)  })	
	}
	
	if len(result) > 3 {
		result = result[0:3]
	}
	
	return result, nil 
}

func readJSON() (Trains, error){

	byteValue, err := ioutil.ReadFile(path)
	
	if err != nil {
		return nil, err
	}

	trains := make(Trains, 0)

	err = json.Unmarshal(byteValue, &trains)

	if err != nil {
		return nil, err
	}

	return trains, nil

}

func inputValidation(departureStation, arrivalStation, criteria string)(int, int, error){

	if departureStation == "" {
		return 0, 0, errors.New(errEmptyDepStation)
	}

	depStationNumber, err := strconv.Atoi(departureStation)

	if err != nil{
		return 0, 0, errors.New(errInputDepStation) 
	}
	
	if arrivalStation == "" {
		return 0, 0, errors.New(errEmptyArrStation)
	}

	arrStationNumber, err := strconv.Atoi(arrivalStation)

	if err != nil{
		return 0, 0, errors.New(errInputArrStation) 
	}
	
	if criteria != criteriaPrice && criteria != criteriaArrTime && criteria != criteriaDepTime{
		return 0, 0, errors.New(errCriteria)
	}

	return depStationNumber, arrStationNumber, nil

}