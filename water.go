package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// WaterUnit is an object that contains one log of water consumption
type WaterUnit struct {
	Amount uint64 `json:"amount"`
	LogID  uint64 `json:"logId"`
}

// WaterSummary is a summary of the requested water logs
type WaterSummary struct {
	Water uint64 `json:"water"`
}

// Water object contains all the data of the logged water
type Water struct {
	Summary *WaterSummary `json:"summary"`
	Water   []*WaterUnit  `json:"water"`
}

// GetWater gets the amount of water taken on a specific date
// It returns an collection of Water or an error if one occours
func (c *Client) GetWater(date time.Time) (*Water, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/foods/log/water/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	waterData := &Water{}
	err = json.NewDecoder(responseBody).Decode(waterData)
	if err != nil {
		return nil, err
	}

	return waterData, nil
}

// LogWater adds a certain amount of water taken on a specific date
// It returns an error if one occours
func (c *Client) LogWater(measurementType string, amount float64, date time.Time) (*Water, error) {
	//Supported unit types
	measurementUnitTypes := map[string]string{"ml": "", "fl oz": "", "cup": ""}

	//Build arguments map
	dataArguments := map[string]string{
		"amount": strconv.FormatFloat(amount, 'f', 2, 32),
		"date":   date.Format("2006-01-02"),
	}

	//Check parameters
	if amount == 0 {
		return nil, errors.New("missing paramters")
	}

	//Check if correct measurement is passed
	_, ok := measurementUnitTypes[measurementType]
	if ok {
		dataArguments["unit"] = measurementType
	}

	//Build an put request-URL
	responseBody, err := c.postData("user/-/foods/log/water.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logWater := &Water{}
	err = json.NewDecoder(responseBody).Decode(logWater)
	if err != nil {
		return nil, err
	}

	return logWater, nil
}

// DeleteWater removes a record of water taken based on the waterId
// It returns an error if one occours
func (c *Client) DeleteWater(waterId uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("usr/-/foods/log/water/%i.json", waterId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
