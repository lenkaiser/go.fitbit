package main

import (
	"errors"
	"io"
)

// getData is used to perform basic connection checks and GET data from the given requestURL
// It returns the Body of the response or an error if one occours
func (c *Client) getData(requestURL string) (io.ReadCloser, error) {
	//Check for OAuth consumer
	if c.oc == nil {
		return nil, errors.New("no consumer")
	}

	//Retrieve data from URL
	response, err := c.oc.Get(c.api.apiURL+requestURL, map[string]string{}, c.accessToken)
	if err != nil {
		return nil, err
	}

	//Return body
	return response.Body, nil
}

// postData is used to perform basic connection checks and POST data to the given requestURL
// It returns the Body of the response or an error if one occours
func (c *Client) postData(requestURL string, dataArguments map[string]string) (io.ReadCloser, error) {
	//Check for OAuth consumer
	if c.oc == nil {
		return nil, errors.New("no consumer")
	}

	//Put data to URL
	response, err := c.oc.Post(c.api.apiURL+requestURL, dataArguments, c.accessToken)
	if err != nil {
		return nil, err
	}

	//Return response
	return response.Body, nil
}

// deleteData is used to perform basic connection checks and DELETE data to the given requestURL
// It returns the Body of the response or an error if one occours
func (c *Client) deleteData(requestURL string, dataArguments map[string]string) (io.ReadCloser, error) {
	//Check for OAuth consumer
	if c.oc == nil {
		return nil, errors.New("no consumer")
	}

	//Put data to URL
	response, err := c.oc.Delete(c.api.apiURL+requestURL, dataArguments, c.accessToken)
	if err != nil {
		return nil, err
	}

	//Return response
	return response.Body, nil
}

// putData is used to perform basic connection checks and PUT data to the given requestURL
// It returns the Body of the response or an error if one occours
func (c *Client) putData(requestURL, body string, dataArguments map[string]string) (io.ReadCloser, error) {
	//Check for OAuth consumer
	if c.oc == nil {
		return nil, errors.New("no consumer")
	}

	//Put data to URL
	response, err := c.oc.Put(c.api.apiURL+requestURL, body, dataArguments, c.accessToken)
	if err != nil {
		return nil, err
	}

	//Return response
	return response.Body, nil
}
