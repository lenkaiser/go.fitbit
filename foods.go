package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// MealUnit describes the unit of a meal
type MealUnit struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Plural string `json:"plural"`
}

// Meal holds all the details for a specific meal containing one or more food items
type Meal struct {
	AccessLevel string    `json:"accessLevel"`
	Amount      uint64    `json:"amount"`
	Brand       string    `json:"brand"`
	Calories    uint64    `json:"calories"`
	FoodId      uint64    `json:"foodId"`
	MealTypeId  uint64    `json:"mealTypeId"`
	Locale      string    `json:"locale"`
	Name        string    `json:"name"`
	Unit        *MealUnit `json:"unit"`
	Units       []int     `json:"units"`
}

// MealFood holds the details about a meal
type MealFood struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Id          uint64  `json:"id"`
	MealsFoods  []*Meal `json:"mealFoods"`
}

// GetMeals is a collection of all the meals of the given user
type GetMeals struct {
	Meals []*MealFood `json:"mealFoods"`
}

// GetMeals gets all the meals of the given user
// It returns an collection of Meal or an error if one occours
func (c *Client) GetMeals() (*GetMeals, error) {
	//Build requestURL and GET data
	responseBody, err := c.getData("user/-/meals.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	mealData := &GetMeals{}
	err = json.NewDecoder(responseBody).Decode(mealData)
	if err != nil {
		return nil, err
	}

	return mealData, nil
}

type NutritionValues struct {
	Biotin            uint64  `json:"biotin"`
	Calcium           uint64  `json:"calcium"`
	Calories          uint64  `json:"calories"`
	CaloriesFromFat   uint64  `json:"caloriesFromFat"`
	Cholesterol       uint64  `json:"cholesterol"`
	Copper            uint64  `json:"copper"`
	DietaryFiber      uint64  `json:"dietaryFiber"`
	FolicAcid         uint64  `json:"folicAcid"`
	Iodine            uint64  `json:"iodine"`
	Iron              uint64  `json:"iron"`
	Magnesium         uint64  `json:"magnesium"`
	Niacin            uint64  `json:"niacin"`
	PantothenicAcid   uint64  `json:"pantothenicAcid"`
	Phosphorus        uint64  `json:"phosphorus"`
	Potassium         uint64  `json:"potassium"`
	Protein           uint64  `json:"protein"`
	Riboflavin        uint64  `json:"riboflavin"`
	SaturatedFat      float64 `json:"saturatedFat"`
	Sodium            uint64  `json:"sodium"`
	Sugars            uint64  `json:"sugars"`
	Thiamin           uint64  `json:"thiamin"`
	TotalCarbohydrate uint64  `json:"totalCarbohydrate"`
	TotalFat          uint64  `json:"totalFat"`
	TransFat          uint64  `json:"transFat"`
	VitaminA          uint64  `json:"vitaminA"`
	VitaminB12        uint64  `json:"vitaminB12"`
	VitaminB6         uint64  `json:"vitaminB6"`
	VitaminC          uint64  `json:"vitaminC"`
	VitaminD          uint64  `json:"vitaminD"`
	VitaminE          uint64  `json:"vitaminE"`
	Zinc              uint64  `json:"zinc"`
}

// Food holds the details about a piece of food
type Food struct {
	IsFavorite      bool             `json:"isFavorite"`
	LogDate         string           `json:"logDate"`
	LogID           uint64           `json:"logId"`
	LoggedFood      *Meal            `json:"loggedFood"`
	NutritionValues *NutritionValues `json:"nutritionValues"`
}

// FoodGoals holds all the goals for food the user wants to achieve
type FoodGoals struct {
	Calories             uint64 `json:"calories"`
	EstimatedCaloriesOut uint64 `json:"estamatedCaloriesOut"`
}

// GetFoods holds all the details about the foods for a specific date
type GetFoods struct {
	Foods   []*Food          `json:"foods"`
	Summary *NutritionValues `json:"summary"`
	Goals   *FoodGoals       `json:"goals"`
}

// GetFoods gets all the foods for a specific date
// It returns an collection of Food or an error if one occours
func (c *Client) GetFoodLogs(date time.Time) (*GetFoods, error) {
	//Build requestURL and GET data
	requestURL := fmt.Sprintf("user/-/foods/log/date/%s.json", date.Format("2006-01-02"))
	responseBody, err := c.getData(requestURL)
	if err != nil {
		return nil, err
	}

	//Parse data
	foodsData := &GetFoods{}
	err = json.NewDecoder(responseBody).Decode(foodsData)
	if err != nil {
		return nil, err
	}

	return foodsData, nil
}

// RecentFoods holds the details for the recent pieces of food taken by the user
type RecentFoods []*Meal

// GetRecentFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetRecentFoods() (*RecentFoods, error) {
	//Build requestURL and GET data
	responseBody, err := c.getData("user/-/foods/log/recent.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	recentFoodData := &RecentFoods{}
	err = json.NewDecoder(responseBody).Decode(recentFoodData)
	if err != nil {
		return nil, err
	}

	return recentFoodData, nil
}

type FrequentFoods []*Meal

// GetFrequentFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetFrequentFoods() (*FrequentFoods, error) {
	//Build requestURL and GET data
	responseBody, err := c.getData("user/-/foods/log/frequent.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	frequentFoodData := &FrequentFoods{}
	err = json.NewDecoder(responseBody).Decode(frequentFoodData)
	if err != nil {
		return nil, err
	}

	return frequentFoodData, nil
}

type Serving struct {
	Multiplier  uint64    `json:"multiplier"`
	ServingSize uint64    `json:"servingSize"`
	UnitId      uint64    `json:"unitId"`
	Unit        *MealUnit `json:"unit"`
}

// FavoriteFoods holds all the details for every favorite piece of food of the given user
type FavoriteFoods struct {
	AccessLevel        string           `json:"accessLevel"`
	Brand              string           `json:"brand"`
	Calories           uint64           `json:"calories"`
	DefaultServingSize uint64           `json:"defaultServingSize"`
	FoodId             uint64           `json:"foodId"`
	Name               string           `json:"name"`
	Servings           []*Serving       `json:"servings"`
	Units              []int            `json:"units"`
	NutritionValues    *NutritionValues `json:"nutritionValues"`
}

type FavoriteFoodsArray []*FavoriteFoods

// GetFavoriteFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetFavoriteFoods() (*FavoriteFoodsArray, error) {
	//Build requestURL and GET data
	responseBody, err := c.getData("user/-/foods/log/favorite.json")
	if err != nil {
		return nil, err
	}

	//Parse data
	favoriteFoodData := &FavoriteFoodsArray{}
	err = json.NewDecoder(responseBody).Decode(favoriteFoodData)
	if err != nil {
		return nil, err
	}

	return favoriteFoodData, nil
}

type LogFood struct {
	FoodLog *Food `json:"foodLog"`
}

// LogFood makes it possible to add food to the user's Fitbit account
// It returns an error if one occours
func (c *Client) LogFood(date time.Time, foodName, brandName string, foodId, mealTypeId, unitId, calories uint64, amount float64, favorite bool) (*LogFood, error) {
	//Build dataArguments
	dataArguments := map[string]string{
		"mealTypeId": strconv.FormatUint(mealTypeId, 10),
		"unitId":     strconv.FormatUint(unitId, 10),
		"amount":     strconv.FormatFloat(amount, 'f', 2, 32),
		"date":       date.Format("2006-01-02"),
		"favorite":   strconv.FormatBool(favorite),
	}

	if foodId == 0 && len(foodName) == 0 {
		return nil, errors.New("missing parameters")
	}
	if foodId > 0 {
		dataArguments["foodId"] = strconv.FormatUint(foodId, 10)
	} else {
		dataArguments["foodName"] = foodName

		//Set calories
		dataArguments["calories"] = strconv.FormatUint(calories, 10)

		//Set brandname
		if len(brandName) > 0 {
			dataArguments["brandName"] = brandName
		}
	}

	//Buid requestURL and POST data
	responseBody, err := c.postData("user/-/foods/log.json", dataArguments)
	if err != nil {
		return nil, err
	}

	//Parse data
	logFoodData := &LogFood{}
	err = json.NewDecoder(responseBody).Decode(logFoodData)
	if err != nil {
		return nil, err
	}

	return logFoodData, nil
}

// DeleteFood makes it possible to remove food from the user's Fitbit account
// It returns an error if one occours
func (c *Client) DeleteFood(foodId uint64) error {
	//Build requestURL and DELETE data
	requestURL := fmt.Sprintf("user/-/foods/log/%d.json", foodId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}

// SearchFood contains all the results of a search query in the user's database
type SearchFood struct {
	Foods []*Meal `json:"foods"`
}

// SearchFood makes it possible to search the user's food list with a query
// It returns an collection of Food or an error if one occours
func (c *Client) SearchFood(query string) (*SearchFood, error) {
	//Build requestURL and GET data
	requestURL := fmt.Sprintf("foods/search.json?query=%s", query)
	responseBody, err := c.getData(requestURL)

	//Parse data
	searchFoodData := &SearchFood{}
	err = json.NewDecoder(responseBody).Decode(searchFoodData)
	if err != nil {
		return nil, err
	}

	return searchFoodData, nil
}

// AddFavoriteFood adds a record to the list of favorite foods of the given user
// It returns an error if one occours
func (c *Client) AddFavoriteFood(foodId uint64) error {
	//Build requestURL and POST data
	requestURL := fmt.Sprintf("user/-/foods/log/favorite/%d.json", foodId)
	_, err := c.postData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFavoriteFood removes a record from the list of favorite foods of the given user
// It returns an error if one occours
func (c *Client) DeleteFavoriteFood(foodId uint64) error {
	//Build requestURL and POST data
	requestURL := fmt.Sprintf("user/-/foods/log/favorite/%d.json", foodId)
	_, err := c.deleteData(requestURL, nil)
	if err != nil {
		return err
	}

	return nil
}
