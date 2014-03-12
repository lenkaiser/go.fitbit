package main

import (
	"errors"
)

// GetBloodpressure gets the bloodpressure of the user for a specific date
// It returns an collection of Bloodpressure or an error if one occours
func (c *Client) GetBloodpressure(date string) (*Bloodpressure, error) {
	return nil, errors.New("not implemented yet")
}

// LogBloodpressure logs the bloodpresure of the given user
// It returns an object Bloodpressure or an error if one occours
func (c *Client) LogBloodpressure(date, time string, systolic, distolic uint64) (*Bloodpressure, error) {
	return nil, errors.New("not implemented yet")
}

// DeleteBloodpressure removes a record from the user's Fitbit account
// It returns an error if one occours
func (c *Client) DeleteBloodpressure(bloodpressureId uint64) error {
	return errors.New("not implemented yet")
}
