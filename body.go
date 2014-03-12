package main

import (
	"errors"
)

// GetBody gets all the details of the body for the given user
// It returns an object Body or an error if one occours
func (c *Client) GetBody(date string) (*Body, error) {
	return nil, errors.New("not implemented yet")
}

// LogBody adds all the body measurements for the given user
// It returns an error if one occours
func (c *Client) LogBody(date string, bicep, calf, chest, fat, forearm, hips, neck, thigh, waist, weight float64) (*Body, error) {
	return nil, errors.New("not implemented yet")
}

// LogWeight logs user's weight
// It returns an object Weight or an error if one occours
func (c *Client) LogWeight(date string, weight float64) {
	return nil, errors.New("not implemented yet")
}
