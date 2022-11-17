package covid

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	covidType "github.com/tanpochara/covid-data-normalization/types"
)

func TestEmptyCase(T *testing.T) {

	//case empty array
	emptyCase := []byte(`[]`)
	var emptyCaseJson []covidType.CovidCase

	json.Unmarshal(emptyCase, &emptyCaseJson)

	groupByAgeResults := map[string]int{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	groupByProvinceResults := make(map[string]int)

	for _, v := range emptyCaseJson {
		age := v.Age
		province := v.Province
		UpdateAge(age, groupByAgeResults)
		UpdateProvince(&province, groupByProvinceResults)
	}

	expectedAgeGroupTestResult := map[string]int{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	ExpectedProvinceGroupTestResult := make(map[string]int)

	assert.Equal(T, expectedAgeGroupTestResult, groupByAgeResults, "the result age group suppost to be the same as initial")
	assert.Equal(T, ExpectedProvinceGroupTestResult, groupByProvinceResults, "the result province group suppose to be the same as initial")
}

func TestFloorCase(T *testing.T) {

	//case empty array
	floorCase := []byte(`[{
		"ConfirmDate": "",
		"No": null,
		"Age": null,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": "",
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }, {
		"ConfirmDate": "",
		"No": null,
		"Age": null,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": null,
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }, {
		"ConfirmDate": "",
		"No": null,
		"Age": null,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": "",
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }]`)
	var floorCaseJson []covidType.CovidCase

	json.Unmarshal(floorCase, &floorCaseJson)

	groupByAgeResults := map[string]int{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	groupByProvinceResults := make(map[string]int)

	for _, v := range floorCaseJson {
		age := v.Age
		province := v.Province
		UpdateAge(age, groupByAgeResults)
		UpdateProvince(&province, groupByProvinceResults)
	}

	expectedAgeGroupTestResult := map[string]int{
		"N/A":   3,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	ExpectedProvinceGroupTestResult := map[string]int{
		"N/A": 3,
	}

	//check if there is any patient missing in group by age
	var totalPatientByAge int
	for _, v := range groupByAgeResults {
		totalPatientByAge += v
	}

	//check if there is any patient missing in group by Province
	var totalPatientByProvince int
	for _, v := range groupByProvinceResults {
		totalPatientByProvince += v
	}

	assert.Equal(T, expectedAgeGroupTestResult, groupByAgeResults, "the result age group suppost to be `N/A`: 3")
	assert.Equal(T, ExpectedProvinceGroupTestResult, groupByProvinceResults, "the result province group suppose to be to be `N/A`: 3")
	assert.Equal(T, 3, totalPatientByAge, "there is patient missing in group by age function")
	assert.Equal(T, 3, totalPatientByProvince, "there is patient missing in group by province function")
}

func TestNormalCase(T *testing.T) {
	//case empty array
	floorCase := []byte(`[{
		"ConfirmDate": "",
		"No": null,
		"Age": 10,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": "Phrae",
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }, {
		"ConfirmDate": "",
		"No": null,
		"Age": 67,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": "Suphan Buri",
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }, {
		"ConfirmDate": "",
		"No": null,
		"Age": 40,
		"Gender": "",
		"GenderEn": "",
		"Nation": null,
		"NationEn": null,
		"Province": "Roi Et",
		"ProvinceId": 0,
		"District": null,
		"ProvinceEn": "",
		"StatQuarantine": 0
	  }]`)
	var floorCaseJson []covidType.CovidCase

	json.Unmarshal(floorCase, &floorCaseJson)

	groupByAgeResults := map[string]int{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}

	groupByProvinceResults := make(map[string]int)

	for _, v := range floorCaseJson {
		age := v.Age
		province := v.Province
		UpdateAge(age, groupByAgeResults)
		UpdateProvince(&province, groupByProvinceResults)
	}

	expectedAgeGroupTestResult := map[string]int{
		"N/A":   0,
		"0-30":  1,
		"31-60": 1,
		"61+":   1,
	}

	ExpectedProvinceGroupTestResult := map[string]int{
		"Phrae":       1,
		"Roi Et":      1,
		"Suphan Buri": 1,
	}

	//check if there is any patient missing in group by age
	var totalPatientByAge int
	for _, v := range groupByAgeResults {
		totalPatientByAge += v
	}

	//check if there is any patient missing in group by Province
	var totalPatientByProvince int
	for _, v := range groupByProvinceResults {
		totalPatientByProvince += v
	}

	assert.Equal(T, expectedAgeGroupTestResult, groupByAgeResults, "the result age group suppost to be 1 for each age group")
	assert.Equal(T, ExpectedProvinceGroupTestResult, groupByProvinceResults, "the result province group suppose to be 1 for Phrea , Roi Et , Suphan Buri province")
	assert.Equal(T, 3, totalPatientByAge, "there is patient missing in group by age function")
	assert.Equal(T, 3, totalPatientByProvince, "there is patient missing in group by province function")
}
