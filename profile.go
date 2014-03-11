package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Profile holds all the basic information about the user
type Profile struct {
	User struct {
		AboutMe             string  `json:"aboutMe"`
		Avatar              string  `json:"avatar"`
		Avatar150           string  `json:"avatar150"`
		City                string  `json:"city"`
		Country             string  `json:"country"`
		DateOfBirth         string  `json:"dateOfBirth"`
		DisplayName         string  `json:"displayName"`
		DistanceUnit        string  `json:"distanceUnit"`
		EncodeId            string  `json:"encodeId"`
		FoodsLocale         string  `json:"foodsLocale"`
		FullName            string  `json:"fullName"`
		Gender              string  `json:"gender"`
		GlucoseUnit         string  `json:"glucoseUnit"`
		Height              float64 `json:"height"`
		HeightUnit          string  `json:"heightUnit"`
		Locale              string  `json:"locale"`
		MemberSince         string  `json:"memberSince"`
		Nickname            string  `json:"nickname"`
		OffsetFromUTCMillis int64   `json:"offsetFromUTCMillis"`
		State               string  `json:"state"`
		StrideLengthRunning float64 `json:"strideLengthRunning"`
		StrideLengthWalking float64 `json:"strideLengthWalking"`
		Timezone            string  `json:"timezone"`
		WaterUnit           string  `json:"waterUnit"`
		Weight              float64 `json:"weight"`
		WeightUnit          string  `json:"weightUnit"`
	} `json:"user"`
}

func (c *Client) GetProfile() (*Profile, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/profile.json")
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}
	defer responseBody.Close()

	//Parse data
	profileData := &Profile{}
	err = json.NewDecoder(responseBody).Decode(profileData)
	if err != nil {
		return nil, err
	}

	return profileData, nil
}

// updateProfile can be used to update a few parameters of the users profile and returns a bool
func (c *Client) UpdateProfile(fullName, nickname, gender, birthday, height, timezone string) error {
	// if c.oc != nil {
	// 	//Build request-URL
	// 	requestURL := fmt.Sprintf("user/-/profile.json")
	// 	log.Println("requestURL >> %s", requestURL)
	// }

	return errors.New("not implemented yet")
}
