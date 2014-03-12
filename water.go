package main

import (
	"errors"
)

// GetWater gets the amount of water taken on a specific date
// It returns an collection of Water or an error if one occours
func (c *Client) GetWater(date string) (*Water, error) {
	return nil, errors.New("not implemented yet")
}

// LogWater adds a certain amount of water taken on a specific date
// It returns an error if one occours
func (c *Client) LogWater(date, waterUnit string, amount float64) error {
	return errors.New("not implemented yet")
}

// DeleteWater removes a record of water taken based on the waterId
// It returns an error if one occours
func (c *Client) DeleteWater(waterId uint64) error {
	return errors.New("not implemented yet")
}
