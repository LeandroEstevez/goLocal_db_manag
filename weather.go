package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var weather Weather

type Weather struct {
	Weather []WeatherDescription
	Main MainDescription
	Wind WindDescription
	Clouds CloudsDescription
	Name string
}

type WeatherDescription struct {
	Main string
	Description string
	Icon string
}

type MainDescription struct {
	Temp float64
	Feels_like float64
	Temp_min float64
	Temp_max float64
	Humidity int
}

type WindDescription struct {
	Speed float64
}

type CloudsDescription struct {
	All int
}

func getWeather(ch chan int, c *gin.Context) {
	zipCode := c.Param("zipCode")
	appid := "appid=286be3842d1da260d7e15a5cdf394d2f"
	params := zipCode + ",us" + "&" + appid
	endPoint := "https://api.openweathermap.org/data/2.5/weather?zip=" + params
	fmt.Println(endPoint)
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
