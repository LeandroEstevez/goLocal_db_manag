package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "gotest"
)

type Response struct {
	Restaurants Restaurants
	Weather Weather
	State string
}

var responseObj Response

func main() {
	router := gin.Default()
	router.GET("/restaurants/:zipCode", getInfo)
	router.Run("localhost:8080")
}

func getInfo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	ch := make(chan int)
	err := 0

	go getRestaurants(ch, c)
	go getWeather(ch, c)

	for i := 0; i < 2; i++ {
		if (<- ch == 1) {
			err++
		}
	}

	if err == 2 {
		c.IndentedJSON(http.StatusFailedDependency, gin.H{"message": "error: all dependencies failed"})
	} else if err > 0 {
		responseObj.State = "Incomplete info, part of the dependencies failed"
		c.IndentedJSON(http.StatusOK, responseObj)
	} else {
		responseObj.State = "Info is complete"
		c.IndentedJSON(http.StatusOK, responseObj)
	}
}
