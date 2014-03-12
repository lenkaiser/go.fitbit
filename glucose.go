package main

import (
	"errors"
)

// GetGlucose gets the bloodpressure of the user for a specific date
// It returns an collection of Glucose or an error if one occours
func (c *Client) GetGlucose(date string) (*Glucose, error) {
	return nil, errors.New("not implemented yet")
}

// LogGlucose logs the bloodpresure of the given user
// It returns an object Bloodpressure or an error if one occours
func (c *Client) LogGlucose(date, trackerName, glucose, hba1c string) (*Bloodpressure, error) {
	return nil, errors.New("not implemented yet")
}
