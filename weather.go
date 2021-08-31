package main

import (
	"encoding/json"
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
	Temp string
	Feels_like string
	Temp_min string
	Temp_max string
	Humidity string
}

type WindDescription struct {
	Speed string
}

type CloudsDescription struct {
	All string
}

func getWeather(ch chan int, c *gin.Context) {
	zipCode := c.Param("zipCode")
	appid := "f1ef95e37251350ef36c400fcb3d3d73"
	params := zipCode + "&" + appid
	endPoint := "api.openweathermap.org/data/2.5/weather?zip=" + params

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
}
