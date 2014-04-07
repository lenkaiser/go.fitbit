package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Summary of the requested sleep data
type SleepSummary struct {
	TotalSleepRecords  uint64 `json:"totalSleepRecords"`
	TotalMinutesAsleep uint64 `json:"totalMinutesAsleep"`
	TotalTimeInBed     uint64 `json:"totalTimeInBed"`
}

type MinuteData struct {
	DateTime string `json:"dateTime"`
	Value    uint64 `json:"value"`
}

// SleepUnit object contains all the data gather during one sleep session
type SleepUnit struct {
	IsMainSleep         bool   `json:"isMainSleep"`
	LogId               uint64 `json:"logId"`
	Efficiency          uint64 `json:"efficiency"`
	StartTime           string `json:"startTime"`
	Duration            uint64 `json:"duration"`
	MinutesToFallAsleep uint64 `json:"minutesToFallAsleep"`
	MinutesAsleep       uint64 `json:"minutesAsleep"`
	MinutesAwake        uint64 `json:"minutesAwake"`
	MinutesAfterWakeup  uint64 `json:"minutesAfterWakeup"`
	// AwakeningsCount     uint64        `json:"awakeningsCount"` -- deprecated
	AwakeCount       uint64        `json:"awakeCount"`
	AwakeDuration    uint64        `json:"awakeDuration"`
	RestlessCount    uint64        `json:"restlessCount"`
	RestlessDuration uint64        `json:"restlessDuration"`
	TimeInBed        uint64        `json:"timeInBed"`
	MinuteData       []*MinuteData `json:"minuteData"`
}

// Sleep is a container that holds a collection of sleep sessions and a summary
type Sleep struct {
	Sleep   []*SleepUnit  `json:"sleep"`
	Summary *SleepSummary `json:"summary"`
}

// GetSleep gets the sleep data for a specific date
// It returns an error if one occours
func (c *Client) GetSleep(date time.Time) (*Sleep, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/sleep/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	sleepData := &Sleep{}
	err = json.NewDecoder(responseBody).Decode(sleepData)
	if err != nil {
		return nil, err
	}

	return sleepData, nil
}

// LogSleep adds a record with minutes of sleep for a specific date
// It returns an error if one occours
func (c *Client) LogSleep(duration float64, date time.Time) (*Sleep, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"startTime": date.Format("15:04"),
		"date":      date.Format("2006-01-02"),
	}

	//Check parameters
	if duration == 0 {
		return nil, errors.New("missing paramters")
	} else {
		dataArguments["duration"] = strconv.FormatFloat(duration, 'f', 2, 32)
	}

	//Build an put request-URL
	responseBody, err := c.postData("user/-/sleep.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logSleep := &Sleep{}
	err = json.NewDecoder(responseBody).Decode(logSleep)
	if err != nil {
		return nil, err
	}

	return logSleep, nil
}

// DeleteSleep removes a record of sleep based on the sleepId
// It returns an error if one occours
func (c *Client) DeleteSleep(sleepId uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("usr/-/sleep/%i.json", sleepId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
