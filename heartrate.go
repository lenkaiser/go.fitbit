package main

import (
	"errors"
)

// GetHeartRate gets the heartrate levels of the given user
// It returns an collection HeartRate or an error if one occours
func (c *Client) GetHeartRate(date string) (*HeartRate, error) {
	return nil, errors.New("not implemented yet")
}

// LogHeartRate adds a record with the heartrate to the user's Fitbit account
// It returns an object HeartRate or an error if one occours
func (c *Client) LogHeartRate(date, tracker, time string, heartRate uint64) (*HeartRate, error) {
	return nil, errors.New("not implemented yet")
}

// DeleteHeartRate removes a record from the user's Fitbit account
// It returns an error if on occours
func (c *Client) DeleteHeartRate(heartRateId uint64) error {
	return errors.New("not implemented yet")
}
