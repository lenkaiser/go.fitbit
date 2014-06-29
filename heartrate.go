package fitbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Average holds all the average data for the heartrate
type Average struct {
	HeartRate uint64 `json:"heartRate"`
	Tracker   string `json:"tracker"`
}

// Heart holds all the data for the logged heart rates
type Heart struct {
	HeartRate uint64 `json:"heartrate"`
	LogID     uint64 `json:"logId"`
	Tracker   string `json:"tracker"`
	Time      string `json:"time"`
}

// HeartRate is an object that holds all the data for the users heart condition
type Heartrate struct {
	Average []*Average `json:"average"`
	Heart   []*Heart   `json:"heart"`
}

// GetHeartRate gets the heartrate levels of the given user
// It returns an collection HeartRate or an error if one occours
func (c *Client) GetHeartRate(date time.Time) (*Heartrate, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/heart/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	heartData := &Heartrate{}
	err = json.NewDecoder(responseBody).Decode(heartData)
	if err != nil {
		return nil, err
	}

	return heartData, nil
}

type LogHeartRate struct {
	HeartLog *Heart `json:"heartLog"`
}

// LogHeartRate adds a record with the heartrate to the user's Fitbit account
// It returns an object HeartRate or an error if one occours
func (c *Client) LogHeartRate(trackerName string, heartRate uint64, date time.Time) (*LogHeartRate, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"time": date.Format("15:04"),
		"date": date.Format("2006-01-02"),
	}

	//Check parameters
	if len(trackerName) == 0 && heartRate > 0 {
		return nil, errors.New("missing paramters")
	} else {
		//Set tracker name
		dataArguments["tracker"] = trackerName
		//Set heartRate value
		dataArguments["heartRate"] = strconv.FormatUint(heartRate, 10)
	}

	//Build an put request-URL
	responseBody, err := c.postData("user/-/heart.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logHeartRate := &LogHeartRate{}
	err = json.NewDecoder(responseBody).Decode(logHeartRate)
	if err != nil {
		return nil, err
	}

	return logHeartRate, nil
}

// DeleteHeartRate removes a record from the user's Fitbit account
// It returns an error if on occours
func (c *Client) DeleteHeartRate(heartRateId uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("user/-/heart/%d.json", heartRateId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
