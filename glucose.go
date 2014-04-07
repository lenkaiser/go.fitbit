package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// GlucoseUnit is a container that holds data of a single glucose log
type GlucoseUnit struct {
	Glucose float64 `json:"glucose"`
	Tracker string  `json:"tracker"`
	Time    string  `json:"time"`
}

// Glucose is a container that holds all the glucose logs and the HBA1C
type Glucose struct {
	Glucose []*GlucoseUnit `json:"glucose"`
	Hba1c   float64        `json:"hba1c"`
}

// GetGlucose gets the bloodpressure of the user for a specific date
// It returns an collection of Glucose or an error if one occours
func (c *Client) GetGlucose(date time.Time) (*Glucose, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/glucose/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	glucoseData := &Glucose{}
	err = json.NewDecoder(responseBody).Decode(glucoseData)
	if err != nil {
		return nil, err
	}

	return glucoseData, nil
}

// LogGlucose logs the bloodpresure of the given user
// It returns an object Bloodpressure or an error if one occours
func (c *Client) LogGlucose(trackerName, glucose, hba1c string, date time.Time) (*Glucose, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"time": date.Format("15:04"),
		"date": date.Format("2006-01-02"),
	}

	//Check parameters
	if len(trackerName) == 0 && len(hba1c) == 0 {
		return nil, errors.New("missing paramters")
	} else {
		if len(trackerName) > 0 {
			//Set tracker name
			dataArguments["tracker"] = trackerName

			//Set glucose value
			if len(glucose) > 0 {
				dataArguments["glucose"] = glucose
			}
		}
		if len(hba1c) > 0 {
			//Set hba1c value
			dataArguments["hba1c"] = hba1c
		}
	}

	//Build an put request-URL
	responseBody, err := c.postData("user/-/glucose.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logGlucose := &Glucose{}
	err = json.NewDecoder(responseBody).Decode(logGlucose)
	if err != nil {
		return nil, err
	}

	return logGlucose, nil
}
