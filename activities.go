package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"time"
)

func (c *Client) getActivities(date time.Time) {
	if c.oc != nil {
		//Convert date to string
		strDate := ""

		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-//activities/date/%s.json", c.api.apiURL, strDate)
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

func (c *Client) getRecentActivities() {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/activities/recent.json", c.api.apiURL)
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

func (c *Client) getFrequentActivities() {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/activities/frequent.json", c.api.apiURL)
		log.Println("requestURL >> %s", requestURL)
	}
}

func (c *Client) getFavoriteActivities() {
	if c.oc != nil {
		//Build request-URL
		requestURL := fmt.Sprintf("%suser/-/activities/favorite.json", c.api.apiURL)
		log.Println("requestURL >> %s", requestURL)
	}
}
