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

func parse() {
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
		Prices[0].Price = newPrice
	})

	url := generateUrl(domain, Prices[0].Name)

	c.Visit(url)
}
