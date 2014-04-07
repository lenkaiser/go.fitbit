package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Activity holds all the basic information for a single measured activity
type Activity struct {
	ActivityId       uint64  `json:"activityId"`
	ActivityParentId uint64  `json:"activityParentId"`
	Calories         uint64  `json:"calories"`
	Description      string  `json:"description"`
	Duration         uint64  `json:"duration"`
	Distance         float64 `json:"distance"`
	HasStartTime     bool    `json:"hasStartTime"`
	IsFavorite       bool    `json:"isFavorite"`
	LogID            uint64  `json:"logId"`
	Name             string  `json:"name"`
	StartTime        string  `json:"startTime"`
	Steps            uint64  `json:"steps"`
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
type ActivitySummary struct {
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
	Activities []*Activity      `json:"activities"`
	Goals      *Goal            `json:"goals"`
	Summary    *ActivitySummary `json:"summary"`
}

// GetActivity
type GetActivity struct {
	Activity *Activity `json:"activity"`
}

// GetActivity uses the activityId to get the details of this specific activity
// It returns an object containing all the details or an error if one occours
func (c *Client) GetActivity(activityId uint64) (*GetActivity, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/%d.json", activityId)
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	activityData := &GetActivity{}
	err = json.NewDecoder(responseBody).Decode(activityData)
	if err != nil {
		return nil, err
	}

	return activityData, nil
}

// BrowseActivites gets a collection of all the public and private activities of the user
// It returns an collection of Activity or an error if on occours
func (c *Client) BrowseActivities() ([]*Activity, error) {
	return nil, errors.New("not implemented yet")
}

func (c *Client) GetActivities(date time.Time) (*Activities, error) {
	//Build and get request-URL
	requestURL := fmt.Sprintf("user/-/activities/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	activitiesData := &Activities{}
	err = json.NewDecoder(responseBody).Decode(activitiesData)
	if err != nil {
		return nil, err
	}

	return activitiesData, nil
}

type RecentActivities []*Activity

// GetRecentActivities retrieves all the activities and returns an object array
func (c *Client) GetRecentActivities() (RecentActivities, error) {
	//Build and get request-URL
	responseBody, err := c.getData("user/-/activities/recent.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	recentActivitiesData := RecentActivities{}
	err = json.NewDecoder(responseBody).Decode(&recentActivitiesData)
	if err != nil {
		return nil, err
	}

	return recentActivitiesData, nil
}

type FrequentActivities []*Activity

// GetFrequentActivities retrieves all the frequent activities of the provided userID
func (c *Client) GetFrequentActivities() (FrequentActivities, error) {
	//Build and get request-URL
	responseBody, err := c.getData("user/-/activities/frequent.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	frequentActivitiesData := FrequentActivities{}
	err = json.NewDecoder(responseBody).Decode(&frequentActivitiesData)
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
	responseBody, err := c.getData("user/-/activities/favorite.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	favoriteActivitiesData := FavoriteActivities{}
	err = json.NewDecoder(responseBody).Decode(&favoriteActivitiesData)
	if err != nil {
		return nil, err
	}

	return favoriteActivitiesData, nil
}

type LogActivity struct {
	ActivityLog *Activity `json:"activityLog"`
}

// LogActivity makes it possible to write an activity to the user's Fitbit account
// It returns an error if one occours
// Date has to be specific is following format: 2006-02-25
func (c *Client) LogActivity(date time.Time, activityName, distanceUnit string, activityId, durationMilis, manualCalories uint64, distance float64) (*LogActivity, error) {
	//Supported unit types
	distanceUnitTypes := map[string]string{"Centimeter": "", "Foot": "", "Inch": "", "Kilometer": "", "Meter": "", "Mile": "", "Millimeter": "", "Steps": "", "Yards": ""}

	//Build arguments map
	dataArguments := map[string]string{
		"startTime":      date.Format("15:04"),
		"durationMillis": strconv.FormatUint(durationMilis, 10),
		"date":           date.Format("2006-01-02"),
	}

	//Check parameters
	if activityId == 0 && len(activityName) == 0 {
		return nil, errors.New("missing paramters")
	}

	if activityId > 0 {
		//Set activityId
		dataArguments["activityId"] = strconv.FormatUint(activityId, 10)
	} else {
		//Set activityName
		dataArguments["activityName"] = activityName
		dataArguments["manualCalories"] = strconv.FormatUint(manualCalories, 10)
	}

	if distance > 0 {
		dataArguments["distance"] = strconv.FormatFloat(distance, 'f', 2, 32) //TODO: Check if last parameter is correct (bitsize)
	}

	_, ok := distanceUnitTypes[distanceUnit]
	if ok {
		dataArguments["distanceUnit"] = distanceUnit
	}

	//Build an put request-URL
	responseBody, err := c.postData("user/-/activities.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logActivity := &LogActivity{}
	err = json.NewDecoder(responseBody).Decode(logActivity)
	if err != nil {
		return nil, err
	}

	return logActivity, nil
}

// DeleteActivity makes it possible to remove an activity from the user's Fitbit account
// It returns an error if one occours
func (c *Client) DeleteActivity(activityId uint64) error {
	requestURL := fmt.Sprintf("user/-/activities/%d.json", activityId)
	_, err := c.deleteData(requestURL, map[string]string{})
	if err != nil {
		return err
	}

	return nil
}

// AddFavoriteActivity uses the activityId to mark a specific activity as a favorite one for the user
// It returns an error if one occours
func (c *Client) AddFavoriteActivity(activityId uint64) error {
	requestURL := fmt.Sprintf("user/-/activities/favorite/%d.json", activityId)
	_, err := c.postData(requestURL, map[string]string{})
	if err != nil {
		return err
	}

	return nil
}

// DeleteFavoriteActivity uses the activityId to remove a specific activity from the users favorite activities list
// It returns an error if one occours
func (c *Client) DeleteFavoriteActivity(activityId uint64) error {
	requestURL := fmt.Sprintf("user/-/activities/favorite/%d.json", activityId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
