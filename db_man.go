package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
	Restaurants []Restaurant `json:"data"`
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

var responseObj Response

func main() {
	router := gin.Default()
	router.GET("/restaurants/:zipCode", getRestaurants)
	router.Run("localhost:8080")
}

func getRestaurants(c *gin.Context) {
	zipCode := c.Param("zipCode")
	size := "?size=10"
	params := "/restaurants/zip_code/" + zipCode + size
	endPoint := "https://api.documenu.com/v2" + params

	err := executeGetRestaurants(endPoint, c)
	if err != nil {
		failedDepend(c)
	}
}

func executeGetRestaurants(endPoint string, c *gin.Context) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return errors.New("failed dependency")
	}
	req.Header.Add("x-api-key", "6f69bcc291f26de6e81350fe6535f846")

	resp, err := client.Do(req)
	if err != nil {
		return errors.New("failed dependency")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed dependency")
	}
	json.Unmarshal(body, &responseObj)

	c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	c.IndentedJSON(http.StatusOK, responseObj)

	return nil
}

func failedDepend(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	c.IndentedJSON(http.StatusFailedDependency, gin.H{"message": "error: failed dependency"})
}
