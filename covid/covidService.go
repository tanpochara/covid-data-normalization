package covid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanpochara/covid-data-normalization/utils"
)

func GetCovidSummary(c *gin.Context) {

	covidData, err := utils.GetCovidData()

	groupByProvince := make(map[string]int)

	groupByAge := map[string]int{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	for _, covidCase := range covidData {
		province := covidCase.Province
		age := covidCase.Age

		//the province data may have null value
		//update count of each province
		UpdateProvince(&province, groupByProvince)

		//update count of each age group
		UpdateAge(age, groupByAge)
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Province": groupByProvince, "AgeGroup": groupByAge})
}

func UpdateProvince(province *string, provinceCounter map[string]int) {
	if len(*province) == 0 {
		provinceCounter["N/A"]++
	} else {
		provinceCounter[*province]++
	}
}

func UpdateAge(age *int, ageCounter map[string]int) {
	if age == nil {
		ageCounter["N/A"]++
	} else if *age <= 30 {
		ageCounter["0-30"]++
	} else if *age > 30 && *age <= 60 {
		ageCounter["31-60"]++
	} else {
		ageCounter["61+"]++
	}
}
