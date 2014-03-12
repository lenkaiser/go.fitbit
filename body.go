package main

import (
	"time"
)

type GetBody struct {
	body  *Body      `json:"body"`
	goals *BodyGoals `json:"goals"`
}

type Body struct {
	Bicep   uint64  `json:"bicep"`
	Bmi     float64 `json:"bmi"`
	Calf    float64 `json:"calf"`
	Chest   uint64  `json:"chest"`
	Fat     uint64  `json:"fat"`
	Forearm float64 `json:"forearm"`
	Hips    uint64  `json:"hips"`
	Neck    uint64  `json:"neck"`
	Thigh   uint64  `json:"thigh"`
	Waist   uint64  `json:"waist"`
	Weight  float64 `json:"weight"`
}

type BodyGoals struct {
	Weight float64 `json:"weight"`
}

// GetBody gets all the details of the body for the given user
// It returns an object Body or an error if one occours
func (c *Client) GetBody(date time.Time) (*GetBody, error) {
	//Build and GET requestURL
	requestURL := fmt.Sprintf("user/-/body/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	bodyData = GetBody{}
	err = json.NewDecoder(responseBody).Decode(&bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

// LogBody contains all the details of body measurements
type LogBody struct {
	Body *Body `json:"body"`
}

// LogBody adds all the body measurements for the given user
// It returns an error if one occours
func (c *Client) LogBody(date time.Time, bicep, calf, chest, fat, forearm, hips, neck, thigh, waist, weight float64) (LogBody, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"bicep":   bicep,
		"calf":    calf,
		"chest":   chest,
		"fat":     fat,
		"forearm": forearm,
		"hips":    hips,
		"neck":    neck,
		"thigh":   thigh,
		"waist":   waist,
		"weight":  weight,
		"date":    date.Format("2006-01-02"),
	}

	//Build and POST requestURL
	responseBody, err := c.postData("user/-/body.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logBodyData = LogBody{}
	err = json.NewDecoder(responseBody).Decode(&logBodyData)
	if err != nil {
		return nil, err
	}

	return logBodyData, nil

}

// Weighting contains the details of the user's weight
type Weight struct {
	Bmi    float64 `json:"bmi"`
	Date   string  `json:"date"`
	LogId  uint64  `json:"logId"`
	Time   string  `json:"time"`
	Weight uint64  `json:"weight"`
}

type LogWeight struct {
	WeightLog *Weight `json:"weightLog"`
}

// LogWeight logs user's weight
// It returns an object Weight or an error if one occours
func (c *Client) LogWeight(date time.Time, weight float64) (LogWeight, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"weight": weight,
		"date":   date.Format("2006-01-02"),
		"time":   date.Format("15:04:05"),
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/body/log/weight.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	weightingData = LogBody{}
	err = json.NewDecoder(responseBody).Decode(&weightingData)
	if err != nil {
		return nil, err
	}

	return weightingData, nil
}
