package fitbit

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// GetWeightGoal is a object that is used to accept the return data
type GetWeightGoal struct {
	Goal struct {
		StartDate   string `json:"startDate"`
		StartWeight uint64 `json:"startWeight"`
		Weight      uint64 `json:"weight"`
	} `json:"goal"`
}

// GetBodyWeightGoal is used to get the body weight goal for the given user
// The object GetWeightGoal is returned or an error if one occours
func (c *Client) GetBodyWeightGoals() (*GetWeightGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/body/log/weight/goal.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	weightData := &GetWeightGoal{}
	err = json.NewDecoder(responseBody).Decode(weightData)
	if err != nil {
		return nil, err
	}

	return weightData, nil
}

// UpdateBodyWeightGoal is used to update the body weight goal for the given user
// The object GetWeightGoal is returned or an error if one occours
func (c *Client) UpdateBodyWeightGoal(date time.Time, startWeight, weight float64) (*GetWeightGoal, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"startDate":   date.Format("2006-01-02"),
		"startWeight": strconv.FormatFloat(startWeight, 'f', 2, 32),
	}

	//Check for missing parameters
	if startWeight == 0 && weight == 0 {
		return nil, errors.New("missing parameters")
	}

	//Add weight
	if weight > 0 {
		dataArguments["weight"] = strconv.FormatFloat(weight, 'f', 2, 32)
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/body/log/weight/goal.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	weightData := &GetWeightGoal{}
	err = json.NewDecoder(responseBody).Decode(weightData)
	if err != nil {
		return nil, err
	}

	return weightData, nil
}

// GetFatGoal is a container used to receive the response from the server
type GetFatGoal struct {
	Goal struct {
		BodyFat uint64 `json:"bodyFat"`
	} `json:"goal"`
}

// GetBodyFatGoal is used to get the body fat goal for the given user
// The object GetFatGoal is returned or an error if one occours
func (c *Client) GetBodyFatGoals() (*GetFatGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/body/log/fat/goal.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	fatData := &GetFatGoal{}
	err = json.NewDecoder(responseBody).Decode(fatData)
	if err != nil {
		return nil, err
	}

	return fatData, nil
}

// UpdateBodyFatGoal is used to update the body fat goal for the given user
// The object GetFatGoal is returned or an error if one occours
func (c *Client) UpdateBodyFatGoal(fat float64) (*GetFatGoal, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"fat": strconv.FormatFloat(fat, 'f', 2, 32),
	}

	//Check for missing parameters
	if fat == 0 {
		return nil, errors.New("missing parameters")
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/body/log/fat/goal.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	fatData := &GetFatGoal{}
	err = json.NewDecoder(responseBody).Decode(fatData)
	if err != nil {
		return nil, err
	}

	return fatData, nil
}

// GetActivityDailyGoal is a container used to receive the response from the server
type GetActivityGoal struct {
	Goals struct {
		CaloriesOut   uint64  `json:"caloriesOut"`
		Distance      float64 `json:"distance"`
		ActiveMinutes uint64  `json:"activeMinutes"`
		Floors        uint64  `json:"floors"`
		Steps         uint64  `json:"steps"`
	} `json:"goals"`
}

// GetActivityDailyGoals is used to get the daily activities goal for the given user
// The object GetActivityGoal is returned or an error if one occours
func (c *Client) GetActivityDailyGoals() (*GetActivityGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/activities/goals/daily.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	activityData := &GetActivityGoal{}
	err = json.NewDecoder(responseBody).Decode(activityData)
	if err != nil {
		return nil, err
	}

	return activityData, nil
}

// UpdateActivityDailyGoals is used to update the body fat goal for the given user
// The object GetActivityGoal is returned or an error if one occours
func (c *Client) UpdateActivityDailyGoals(caloriesOut, activeMinutes, floors, steps uint64, distance float64) (*GetActivityGoal, error) {
	//Build arguments map
	dataArguments := map[string]string{
		"caloriesOut":   strconv.FormatUint(caloriesOut, 10),
		"activeMinutes": strconv.FormatUint(activeMinutes, 10),
		"floors":        strconv.FormatUint(floors, 10),
		"steps":         strconv.FormatUint(steps, 10),
		"distance":      strconv.FormatFloat(distance, 'f', 2, 32),
	}

	//Check for missing parameters
	if caloriesOut == 0 || activeMinutes == 0 || floors == 0 || steps == 0 || distance == 0 {
		return nil, errors.New("missing parameters")
	}

	//Buid and POST requestURL
	responseBody, err := c.postData("user/-/activities/goals/daily.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	activityData := &GetActivityGoal{}
	err = json.NewDecoder(responseBody).Decode(activityData)
	if err != nil {
		return nil, err
	}

	return activityData, nil
}

// GetActivityWeeklyGoals is used to get the daily activities goal for the given user
// The object GetActivityGoal is returned or an error if one occours
func (c *Client) GetActivityWeeklyGoals() (*GetActivityGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/activities/goals/weekly.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	activityData := &GetActivityGoal{}
	err = json.NewDecoder(responseBody).Decode(activityData)
	if err != nil {
		return nil, err
	}

	return activityData, nil
}

// GetFoodGoal is a object that is used to receive the server resposne
type GetFoodGoal struct {
	Goals struct {
		Calories uint64 `json:"calories"`
	} `json:"calories"`
	FoodPlan struct {
		Intensity     string `json:"intensity"`
		EstimatedDate string `json:"estimatedDate"`
		Personalized  bool   `json:"personalized"`
	} `json:"foodPlan"`
}

// GetFoodGoals is used to get the food goal for the given user
// The object GetFoodGoal is returned or an error if one occours
func (c *Client) GetFoodGoals() (*GetFoodGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/foods/log/goal.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	foodData := &GetFoodGoal{}
	err = json.NewDecoder(responseBody).Decode(foodData)
	if err != nil {
		return nil, err
	}

	return foodData, nil
}

// GetWaterGoal is a object that is used to receive the server resposne
type GetWaterGoal struct {
	Goal uint64 `json:"water"`
}

// GetWaterGoals is used to get the water goal for the given user
// The object GetWaterGoal is returned or an error if one occours
func (c *Client) GetWaterGoals() (*GetWaterGoal, error) {
	//Build and GET requestURL
	responseBody, err := c.getData("user/-/foods/log/water/goal.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	waterData := &GetWaterGoal{}
	err = json.NewDecoder(responseBody).Decode(waterData)
	if err != nil {
		return nil, err
	}

	return waterData, nil
}
