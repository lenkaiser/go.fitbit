package main

import (
	"time"
)

type AverageBP struct {
	Condition string `json:"condition"`
	Diastolic uint64 `json:"diastolic"`
	Systolic  uint64 `json:"systolic"`
}

type Bloodpressure struct {
	Diastolic uint64 `json:"diastolic"`
	Systolic  uint64 `json:"systolic"`
	LogId     uint64 `json:"logId"`
	time      string `json:"time"`
}

// Bloodpressure holds all the details for the user's BP
type GetBloodpressure struct {
	Avarage *AverageBP       `json:"average"`
	Bp      []*Bloodpressure `json:"bp"`
}

// GetBloodpressure gets the bloodpressure of the user for a specific date
// It returns an collection of Bloodpressure or an error if one occours
// Date has to be specific is following format: 2006-02-25
func (c *Client) GetBloodpressure(date time.Time) (GetBloodpressure, error) {
	//Build and GET requestURL
	requestURL := fmt.Sprintf("user/-/bp/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	bloodPressureData := GetBloodpressure{}
	err = json.NewDecoder(responseBody).Decode(&bloodPressureData)
	if err != nil {
		return nil, err
	}

	return bloodPressureData, nil
}

// LogBloodpressure logs the bloodpresure of the given user
// It returns an object Bloodpressure or an error if one occours
func (c *Client) LogBloodpressure(date time.Time, systolic, diastolic uint64) (Bloodpressure, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"systolic":  systolic,
		"diastolic": diastolic,
		"date":      date.Format("2006-01-02"),
		"time":      date.Format("15:04"),
	}

	//Build and POST requestURL
	responseBody, err := c.postData("user/-/bp.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	bloodPressureData = Bloodpressure{}
	err = json.NewDecoder(responseBody).Decode(&bloodPressureData)
	if err != nil {
		return nil, err
	}

	return bloodPressureData, nil
}

// DeleteBloodpressure removes a record from the user's Fitbit account
// It returns an error if one occours
func (c *Client) DeleteBloodpressure(bloodpressureId uint64) error {
	//Build and DELETE requestURL
	requestURL := fmt.Sprintf("user/-/bp/%i.json", bloodpressureId)
	responseBody, err := c.deleteData(requestURL, nil)

	return nil
}
