package main

import (
	"errors"
)

// GetMeals gets all the meals of the given user
// It returns an collection of Meal or an error if one occours
func (c *Client) GetMeals() (*Meals, error) {
	return nil, errors.New("not implemented yetE")
}

// GetFoods gets all the foods for a specific date
// It returns an collection of Food or an error if one occours
func (c *Client) GetFoods(date string) (*Foods, error) {
	return nil, errors.New("not implemented yet")
}

// GetRecentFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetRecentFoods() (*RecentFoods, error) {
	return nil, errors.New("not implemented yet")
}

// GetFrequentFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetFrequentFoods() (*FrequentFoods, error) {
	return nil, errors.New("not implemented yet")
}

// GetFavoriteFoods gets all the recent foods for the given user
// It returns an collection of Food or an error if one occours
func (c *Client) GetFavoriteFoods() (*FavoriteFoods, error) {
	return nil, errors.New("not implemented yet")
}

// LogFood makes it possible to add food to the user's Fitbit account
// It returns an error if one occours
func (c *Client) LogFood(date, foodName, brandName string, foodId, mealTypeId, unitId, calories, nutrition uint64) error {
	return errors.New("not implemented yet")
}

// DeleteFood makes it possible to remove food from the user's Fitbit account
// It returns an error if one occours
func (c *Client) DeleteFood(foodId uint64) error {
	return errors.New("not implemented yet")
}

// SearchFood makes it possible to search the user's food list with a query
// It returns an collection of Food or an error if one occours
func (c *Client) SearchFood(query string) (*Foods, error) {
	return nil, errors.New("not implemented yet")
}

// AddFavoriteFood adds a record to the list of favorite foods of the given user
// It returns an error if one occours
func (c *Client) AddFavoriteFood(foodId uint64) error {
	return errors.New("not implemented yet")
}

// DeleteFavoriteFood removes a record from the list of favorite foods of the given user
// It returns an error if one occours
func (c *Client) DeleteFavoriteFood(foodId uint64) error {
	return errors.New("not implemented yet")
}
