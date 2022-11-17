package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	covidService "github.com/tanpochara/covid-data-normalization/covid"
)

type testRoutes struct {
	endpoint               string
	expectedResponseStatus int
}

func TestRouter(T *testing.T) {

	routes := []testRoutes{
		{
			endpoint:               "/",
			expectedResponseStatus: 404,
		}, {
			endpoint:               "/covid/summary",
			expectedResponseStatus: 200,
		},
	}

	//make routing the same way as main.go
	router := gin.Default()
	covid := router.Group("/covid")
	{
		covid.GET("/summary", covidService.GetCovidSummary)
	}

	for _, route := range routes {
		req := httptest.NewRequest("GET", route.endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		// assert.Nil(T, err, "it should not be an error in requesting / , page")
		assert.Equal(T, route.expectedResponseStatus, recorder.Result().StatusCode)
	}

}

func TestCovidApi(T *testing.T) {

	//make routing the same way as main.go
	router := gin.Default()
	covid := router.Group("/covid")
	{
		covid.GET("/summary", covidService.GetCovidSummary)
	}

	req := httptest.NewRequest("GET", "/covid/summary", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	covidSummaryResponse, err := ioutil.ReadAll(recorder.Body)

	assert.Nil(T, err, "it should not be any error in the requested route")

	var covidSummaryJson map[string]interface{}
	json.Unmarshal(covidSummaryResponse, &covidSummaryJson)

	//getting expected result

	mockfile, err := os.ReadFile("./mock/summarized_covid.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var expectedResonponse map[string]interface{}

	json.Unmarshal(mockfile, &expectedResonponse)

	assert.Equal(T, expectedResonponse, covidSummaryJson, "The actual response mismatch the expected response")
}
