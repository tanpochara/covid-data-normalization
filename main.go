package main

import (
	"github.com/gin-gonic/gin"
	covidService "github.com/tanpochara/covid-data-normalization/covid"
)

func main() {
	router := gin.Default()

	covid := router.Group("/covid")
	{
		covid.GET("/summary", covidService.GetCovidSummary)
	}

	router.Run("localhost:8888")
}
