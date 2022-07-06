package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type price struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Prices = []price{
	{ID: "btc", Name: "bitcoin", Price: 0.0},
	{ID: "eth", Name: "ethereum", Price: 0.0},
	{ID: "near", Name: "near-protocol", Price: 0.0},
}

func getPriceByID(id string) (*price, error) {
	for i, t := range Prices {
		if t.ID == id {
			return &Prices[i], nil
		}
	}
	return nil, errors.New("currency not found")
}

func getPrices(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Prices)
}

func getPrice(context *gin.Context) {
	id := context.Param("id")
	price, err := getPriceByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "currency not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, price)
}

func main() {
	for i := 0; i < len(Prices); i++ {
		parse(Prices[i].Name, i)
	}

	router := gin.Default()
	router.GET("/prices", getPrices)
	router.GET("/prices/:id", getPrice)
	router.Run("localhost:9090")
}
