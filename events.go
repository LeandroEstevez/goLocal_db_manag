package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var events Weather

func getEvents(ch chan int, c *gin.Context) {
	size := "size=5"
	apiKey := "apikey=RTNVR3nhCo1rADYAlnfzECKd2VKJGcY0"
	zipCode := "postalcode=" + c.Param("zipCode")
	params := size + "&" + zipCode + "&" + apiKey
	endPoint := "https://app.ticketmaster.com/discovery/v2/events.json?" + params

	client := &http.Client{}
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		ch <- 1
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		ch <- 1
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- 1
		return
	}
	json.Unmarshal(body, &weather)

	responseObj.Weather = weather
	ch <- 0
}
