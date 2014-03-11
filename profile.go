package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Get profile if the an oauth connection is available
type Profile struct {
	aboutMe             string
	avatar              string
	avatar150           string
	city                string
	country             string
	dateOfBirth         string
	displayName         string
	distanceUnit        string
	encodeId            string
	foodsLocale         string
	fullName            string
	gender              string
	glucoseUnit         string
	height              string
	heightUnit          string
	locale              string
	memberSince         string
	nickname            string
	offsetFromUTCMillis string
	state               string
	strideLengthRunning string
	strideLengthWalking string
	timezone            string
	waterUnit           string
	weight              string
	weightUnit          string
}

func (c *Client) getProfile() (*Profile, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("%suser/-/profile.json", c.api.apiURL)
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	profileData := &Profile{}
	err = json.NewDecoder(responseBody).Decode(profileData)
	if err != nil {
		return nil, err
	}

	return profileData, nil
}

//This method is able to update a few parameters of the users profile
func (c *Client) updateProfile(fullName, nickname, gender, birthday, height, timezone string) {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/profile.json", c.api.apiURL)
		log.Println("requestURL >> %s", requestURL)
	}
}
