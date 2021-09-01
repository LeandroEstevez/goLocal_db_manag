package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var restaurants Restaurants

type Restaurants struct {
	Data []Restaurant
}

type Restaurant struct {
	Name string `json:"restaurant_name"`
	Phone string `json:"restaurant_phone"`
	Website string `json:"restaurant_website"`
	Cuisines []string `json:"cuisines"`
	Address Address `json:"address"`
}

type Address struct {
	FullAddress string `json:"formatted"`
}

func getRestaurants(ch chan int, c *gin.Context) {
	zipCode := c.Param("zipCode")
	size := "?size=10"
	params := "/restaurants/zip_code/" + zipCode + size
	endPoint := "https://api.documenu.com/v2" + params

	client := &http.Client{}
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		ch <- 1
		return
	}
	req.Header.Add("x-api-key", "6f69bcc291f26de6e81350fe6535f846")

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
	json.Unmarshal(body, &restaurants)

	responseObj.Restaurants = restaurants
	ch <- 0
}
