package main

import (
	"encoding/json"
)

type Device struct {
	Battery       string `json:"battery"`
	Id            string `json:"id"`
	LastSyncTime  string `json:"lastSyncTime"`
	Type          string `json:"type"`          //Types: {TRACKER|SCALE}
	DeviceVersion string `json:"deviceVersion"` //Types: {Flex|One|Zip|Ultra|Classic|Aria}
}

type GetDevices []*Device

// GetDevices is a method that returns all the devices connected to a specific user
// It returns an collection of 'Device' when succeful
func (c *Client) GetDevices() (GetDevices, error) {
	//Build requestURL and GET data
	responseBody, err := c.getData("user/-/devices.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	devicesData := GetDevices{}
	err = json.NewDecoder(responseBody).Decode(&devicesData)
	if err != nil {
		return nil, err
	}

	return devicesData, nil
}
