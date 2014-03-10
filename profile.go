package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
)

//Get profile if the an oauth connection is available
func (c *Client) getProfile() {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/profile.json", c.api.apiURL)
		log.Println("requestURL >> %s", requestURL)

		//Retrieve data from URL
		response, err := c.oc.Get(requestURL, map[string]string{}, c.accessToken)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		bits, err := ioutil.ReadAll(response.Body)
		spew.Dump(bits) //Dump data
	}
}

func (c *Client) updateProfile(fullName, nickname, gender, birthday, height, timezone string) {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/profile.json", c.api.apiURL)
		log.Println("requestURL >> %s", requestURL)
	}
}
