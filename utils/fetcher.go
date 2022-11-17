package utils

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/tanpochara/covid-data-normalization/types"
)

func GetCovidData() ([]types.CovidCase, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://static.wongnai.com/devinterview/covid-cases.json")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var covidData types.CovidResponse

	err = json.Unmarshal(body, &covidData)

	if err != nil {
		return nil, err
	}

	return covidData.Data, nil
}
