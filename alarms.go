package fitbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// TrackerAlarm is an object containing all the data about an alarm
type TrackerAlarm struct {
	AlarmId        uint64   `json:"alarmId"`
	Deleted        bool     `json:"deleted"`
	Enabled        bool     `json:"enabled"`
	Label          string   `json:"label"`
	Recurring      bool     `json:"recurring"`
	SnoozeCount    uint64   `json:"snoozeCount"`
	SnoozeLength   uint64   `json:"snoozeLength"`
	SyncedToDevice bool     `json:"syncedToDevice"`
	Time           string   `json:"time"`
	Vibe           string   `json:"vibe"`
	WeekDays       []string `json:"weekDays"`
}

// GetTrackerAlarms is an container which holds all the alarms connected to a device
type GetTrackerAlarms struct {
	TrackerAlarms []*TrackerAlarm `json:"trackerAlarms"`
}

// GetAlarms gets all the alarms for the trackers connected to the user
// It returns an collection of TrackerAlarm or an error if one occours
func (c *Client) GetAlarms(deviceID uint64) (*GetTrackerAlarms, error) {
	//Build and GET requestURL
	requestURL := fmt.Sprintf("user/-/devices/tracker/%d/alarms.json", deviceID)
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	alarmsData := &GetTrackerAlarms{}
	err = json.NewDecoder(responseBody).Decode(alarmsData)
	if err != nil {
		return nil, err
	}

	return alarmsData, nil
}

type AddAlarm struct {
	TrackerAlarm TrackerAlarm `json:"trackerAlarm"`
}

// AddAlarm is used to add a alarm to the users profile
// This method returns an object containing the created alarm if succesful
func (c *Client) AddAlarm(date time.Time, enabled, recurring bool, weekDays []string, label string, snoozeLength, snoozeCount, deviceID uint64) (*AddAlarm, error) {
	/*
	 * NOTE:
	 * This method is currently not available in the API
	 * When this method is called the server returns a 500 Internal server error
	 * The method is for the time begin unavailable
	 */
	// return nil, errors.New("not implemented yet")

	//Build arguments map
	dataArguments := map[string]string{
		"time":      date.Format("15:04-0700"),
		"enabled":   strconv.FormatBool(enabled),
		"recurring": strconv.FormatBool(recurring),
		"vibe":      "DEFAULT", //Only value for now
	}

	//Check for label
	if len(label) > 0 {
		dataArguments["label"] = label
	}

	//Check for snooze length
	if snoozeLength > 0 {
		dataArguments["snoozeLength"] = strconv.FormatUint(snoozeLength, 10)
	}

	//Check for snooze count
	if snoozeCount > 0 {
		dataArguments["snoozeCount"] = strconv.FormatUint(snoozeCount, 10)
	}

	splitWeekdays := ""
	if len(weekDays) > 0 {
		for i, day := range weekDays {
			if i > 0 {
				splitWeekdays += ","
			}
			splitWeekdays += day
		}
	}
	dataArguments["weekDays"] = "MONDAY"

	//Build an put request-URL
	requestURL := fmt.Sprintf("user/-/devices/tracker/%d/alarms.json", deviceID)
	responseBody, err := c.postData(requestURL, dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	addAlarm := &AddAlarm{}
	err = json.NewDecoder(responseBody).Decode(addAlarm)
	if err != nil {
		return nil, err
	}

	return addAlarm, nil

}

// UpdateAlarm is used to edit/update an existing alarm
// It returns an object containing the edited alarm if succesful
func (c *Client) UpdateAlarm(date time.Time, enabled, recurring bool, weekDays []string, label string, snoozeLength, snoozeCount, deviceID, alarmID uint64) (*AddAlarm, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"time":         date.Format("15:04-0700"),
		"enabled":      strconv.FormatBool(enabled),
		"recurring":    strconv.FormatBool(recurring),
		"vibe":         "DEFAULT", //Only value for now
		"snoozeLength": strconv.FormatUint(snoozeLength, 10),
		"snoozeCount":  strconv.FormatUint(snoozeCount, 10),
	}

	//Check parameters
	if snoozeLength == 0 && snoozeCount == 0 {
		return nil, errors.New("missing paramters")
	}

	//Check for label
	if len(label) > 0 {
		dataArguments["label"] = label
	}

	splitWeekdays := ""
	if len(weekDays) > 0 {
		for i, day := range weekDays {
			if i > 0 {
				splitWeekdays += ","
			}
			splitWeekdays += day
		}
	}
	dataArguments["weekDays"] = splitWeekdays

	//Build an put request-URL
	requestURL := fmt.Sprintf("devices/tracker/%d/alarms/%d.json", deviceID, alarmID)
	responseBody, err := c.postData(requestURL, dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	updateAlarm := &AddAlarm{}
	err = json.NewDecoder(responseBody).Decode(updateAlarm)
	if err != nil {
		return nil, err
	}

	return updateAlarm, nil
}

// DeleteAlarm removes an alarm from the user's Fitbit account
// It returns an error if on occours
func (c *Client) DeleteAlarm(deviceID, alarmID uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("devices/tracker/%d/alarms/%d.json", deviceID, alarmID)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
