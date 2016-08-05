package fitbit

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	bodyData := &GetBody{}
	err = json.NewDecoder(responseBody).Decode(bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

// LogBody contains all the details of body measurements
type LogBody struct {
	Body *Body `json:"body"`
}

// GetBodyMeasurements gets all the body measurements for the given user
// It returns an error if one occours
func (c *Client) GetBodyMeasurements(date time.Time) (*LogBody, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/body/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	bodyData := &LogBody{}
	err = json.NewDecoder(responseBody).Decode(bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

// LogBody adds all the body measurements for the given user
// It returns an error if one occours
func (c *Client) LogBodyMeasurements(date time.Time, bicep, calf, chest, fat, forearm, hips, neck, thigh, waist, weight float64) (*LogBody, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"bicep":   strconv.FormatFloat(bicep, 'f', 2, 32),
		"calf":    strconv.FormatFloat(calf, 'f', 2, 32),
		"chest":   strconv.FormatFloat(chest, 'f', 2, 32),
		"fat":     strconv.FormatFloat(fat, 'f', 2, 32),
		"forearm": strconv.FormatFloat(forearm, 'f', 2, 32),
		"hips":    strconv.FormatFloat(hips, 'f', 2, 32),
		"neck":    strconv.FormatFloat(neck, 'f', 2, 32),
		"thigh":   strconv.FormatFloat(thigh, 'f', 2, 32),
		"waist":   strconv.FormatFloat(waist, 'f', 2, 32),
		"weight":  strconv.FormatFloat(weight, 'f', 2, 32),
		"date":    date.Format("2006-01-02"),
	}

	//Build and POST requestURL
	responseBody, err := c.postData("user/-/body.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logBodyData := &LogBody{}
	err = json.NewDecoder(responseBody).Decode(logBodyData)
	if err != nil {
		return nil, err
	}

	return logBodyData, nil

}

// Weighting contains the details of the user's weight
type Weight struct {
	Bmi     float64 `json:"bmi"`
	Date    string  `json:"date"`
	LogID   uint64  `json:"logId"`
	Time    string  `json:"time"`
	Weight  float64 `json:"weight"`
	BodyFat float64 `json:"fat"`
}

type LogWeight struct {
	WeightLog *Weight `json:"weightLog"`
}

// GetBodyWeight gets the body weight for the given user on a specific date
// It returns an error if one occours
func (c *Client) GetBodyWeight(date time.Time) (*LogWeight, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/body/log/weight/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	bodyData := &LogWeight{}
	err = json.NewDecoder(responseBody).Decode(bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

// LogWeight logs user's weight
// It returns an object Weight or an error if one occours
func (c *Client) LogBodyWeight(date time.Time, weight float64) (*LogWeight, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"weight": strconv.FormatFloat(weight, 'f', 2, 32),
		"date":   date.Format("2006-01-02"),
		"time":   date.Format("15:04:05"),
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/body/log/weight.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	weightingData := &LogWeight{}
	err = json.NewDecoder(responseBody).Decode(weightingData)
	if err != nil {
		return nil, err
	}

	return weightingData, nil
}

// DeleteBodyWeight removes an body weight measurement
// It returns an error if on occours
func (c *Client) DeleteBodyWeight(weightID uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("user/-/body/log/weight/%d.json", weightID)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}

// Weighting contains the details of the user's weight
type Fat struct {
	BMI   float64 `json:"bmi"`
	Date  string  `json:"date"`
	Fat   float64 `json:"fat"`
	LogID uint64  `json:"logId"`
	Time  string  `json:"time"`
}

// GetFat struct that contains all the measurements of body fat
type GetFat struct {
	GetFat []*Fat `json:"fat"`
}

// LogFat is a object that is returned by the server when a fat measurement is logged
type LogFat struct {
	FatLog *Fat `json:"fatLog"`
}

// GetBodyFat gets the body fat measurements for the given user on a specific date
// It returns an error if one occours
func (c *Client) GetBodyFat(date time.Time) (*GetFat, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/body/log/fat/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	bodyData := &GetFat{}
	err = json.NewDecoder(responseBody).Decode(bodyData)
	if err != nil {
		return nil, err
	}

	return bodyData, nil
}

// LogWeight logs user's weight
// It returns an object Weight or an error if one occours
func (c *Client) LogBodyFat(date time.Time, fat float64) (*LogFat, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"fat":  strconv.FormatFloat(fat, 'f', 2, 32),
		"date": date.Format("2006-01-02"),
		"time": date.Format("15:04:05"),
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/body/log/fat.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	weightingData := &LogFat{}
	err = json.NewDecoder(responseBody).Decode(weightingData)
	if err != nil {
		return nil, err
	}

	return weightingData, nil
}

// DeleteBodyFat removes an body fat measurement
// It returns an error if on occours
func (c *Client) DeleteBodyFat(fatID uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("user/-/body/log/fat/%d.json", fatID)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
