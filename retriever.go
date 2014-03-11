package main

import (
	"errors"
	"io"
)

func (c *Client) getData(requestURL string) (io.Reader, error) {
	//Check for consumer
	if c.oc != nil {
		return nil, errors.New("no consumer")
	}

	//Retrieve data from URL
	response, err := c.oc.Get(requestURL, map[string]string{}, c.accessToken)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	//Return body
	return response.Body, nil
}
