package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type price struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Prices = []price{
	{ID: "btc", Name: "bitcoin", Price: 19844.2},
}

func getPrices(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Prices)
}

func main() {
	parse()

	router := gin.Default()
	router.GET("/prices", getPrices)
	router.Run("localhost:9090")
}
