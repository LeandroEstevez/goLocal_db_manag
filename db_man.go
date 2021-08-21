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

type restaurant struct {
	Name string `json:"name"`
	Cuisine string `json:"cuisine"`
}

var restaurants = []restaurant{
	{Name: "El Coco", Cuisine: "Italian"},
	{Name: "El Taco", Cuisine: "Mexican"},
}

func main() {
	router := gin.Default()
	router.GET("/restaurants/:zipCode", getRestaurants)
	router.Run("localhost:8080")
}

func getRestaurants(c *gin.Context) {
	// zipCode := c.Param("zipCode")
	c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	c.IndentedJSON(http.StatusOK, restaurants)
}

// func ConnectDB() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"dbname=%s sslmode=disable", host, port, user, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Conected!!")
// }
