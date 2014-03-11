package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Activity holds all the basic information for a single measured activity
type Activity struct {
	ActivityId       uint64 `json:"activityId"`
	ActivityParentId uint64 `json:"activityParentId"`
	Calories         uint64 `json:"calories"`
	Description      string `json:"description"`
	Duration         uint64 `json:"duration"`
	HasStartTime     bool   `json:"hasStartTime"`
	IsFavorite       bool   `json:"isFavorite"`
	LogId            uint64 `json:"logId"`
	Name             string `json:"name"`
	StartTime        string `json:"startTime"`
	Steps            uint64 `json:"steps"`
}

// Goal represents all data reached to a given date
type Goal struct {
	CaloriesOut uint64  `json:"caloriesOut"`
	Distance    float64 `json:"distance"`
	Floors      uint64  `json:"floors"`
	Steps       uint64  `json:"steps"`
}

// Distance holds different distances per activity (tracker, total, veryActive, etc.)
type Distance struct {
	Activity string  `json:"activity"`
	Distance float64 `json:"distance"`
}

// Summary holds a summary of all the activities of a given date
type Summary struct {
	ActivityCalories     uint64      `json:"activityCalories"`
	CaloriesBMR          uint64      `json:"caloriesBMR"`
	CaloriesOut          uint64      `json:"caloresOut"`
	Distances            []*Distance `json:"distances"`
	Elevation            float64     `json:"elevation"`
	FairlyActiveMinutes  uint64      `json:"fairlyActiveMinutes"`
	Floors               uint64      `json:"floors"`
	LightlyActiveMinutes uint64      `json:"lightlyActiveMinutes"`
	MarginalCalories     uint64      `json:"marginalCalories"`
	SedentaryMinutes     uint64      `json:"sedentaryMinutes"`
	Steps                uint64      `json:"steps"`
	VeryActiveMinutes    uint64      `json:"veryActiveMinutes"`
}

// Activities for a specific given date
type Activities struct {
	activities []*Activity `json:"activities"`
	goals      *Goal       `json:"goals"`
	summary    *Summary    `json:"summary"`
}

// GetActivitiesByUnixTimestamp returns Activity struct
// It accepts an time.Time format and translates it to the required format for the API
func (c *Client) GetActivitiesByUnixTimestamp(date time.Time) (*Activities, error) {
	//Convert timestamp to string
	strDate := date.Format("2006-01-02")
	return c.getActivities(strDate)
}

// GetActivitiesByString returns Activity struct
// It accepts an string and forwards it to the API
func (c *Client) GetActivitiesByString(date string) (*Activities, error) {
	return c.getActivities(date)
}

func (c *Client) getActivities(date string) (*Activities, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/date/%s.json", date)
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	activitiesData := &Activities{}
	err = json.NewDecoder(io.TeeReader(responseBody, os.Stdout)).Decode(activitiesData)
	if err != nil {
		return nil, err
	}

	return activitiesData, nil
}

type RecentActivities []*Activity

// GetRecentActivities retrieves all the activities and returns an object array
func (c *Client) GetRecentActivities() (RecentActivities, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/recent.json")
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	recentActivitiesData := RecentActivities{}
	err = json.NewDecoder(io.TeeReader(responseBody, os.Stdout)).Decode(&recentActivitiesData)
	if err != nil {
		return nil, err
	}

	return recentActivitiesData, nil
}

type FrequentActivities []*Activity

// GetFrequentActivities retrieves all the frequent activities of the provided userID
func (c *Client) GetFrequentActivities() (FrequentActivities, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/frequent.json")
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	frequentActivitiesData := FrequentActivities{}
	err = json.NewDecoder(io.TeeReader(responseBody, os.Stdout)).Decode(&frequentActivitiesData)
	if err != nil {
		return nil, err
	}

	return frequentActivitiesData, nil
}

// FavoriteActivity holds basic data about the favorite activity
type FavoriteActivity struct {
	activityId  uint64
	description string
	mets        uint64
	name        string
}

// FavoriteActivities is data collection of FavoriteActivity
type FavoriteActivities []*FavoriteActivity

// GetFavoriteActivities retrieves all the favorite activities provided by the user
// It returns and object array of FavoriteActivity
func (c *Client) GetFavoriteActivities() (FavoriteActivities, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/favorite.json")
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	favoriteActivitiesData := FavoriteActivities{}
	err = json.NewDecoder(io.TeeReader(responseBody, os.Stdout)).Decode(&favoriteActivitiesData)
	if err != nil {
		return nil, err
	}

	return favoriteActivitiesData, nil
}
