package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var domain string = "coinmarketcap.com"

func generateUrl(d string, cryptoName string) string {
	return "https://" + d + "/currencies/" + cryptoName + "/"
}

func parse(cryptoName string, id int) {
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("div.priceValue", func(e *colly.HTMLElement) {
		unformPrice := e.Text
		formPrice := strings.ReplaceAll(unformPrice, "$", "")
		formPrice = strings.ReplaceAll(formPrice, ",", "")
		newPrice, err := strconv.ParseFloat(formPrice, 32)
		if err != nil {
			log.Fatal(err)
		}
		newPrice = math.Round(newPrice*100) / 100
		Prices[id].Price = newPrice
	})

	url := generateUrl(domain, cryptoName)

	c.Visit(url)
}
