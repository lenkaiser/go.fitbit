package main

import (
	"errors"
)

// LogSleep adds a record with minutes of sleep for a specific date
// It returns an error if one occours
func (c *Client) LogSleep(date string, duration float64) error {
	return errors.New("not implemented yet")
}

// DeleteSleep removes a record of sleep based on the sleepId
// It returns an error if one occours
func (c *Client) DeleteSleep(sleepId uint64) error {
	return errors.New("not implemented yet")
}
