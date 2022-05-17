package api

import (
	"covid/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CovidCall(url string) models.CovidData {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	var apiObj models.CovidData
	err = json.Unmarshal(data, &apiObj)
	if err != nil {
		fmt.Println("Failed to json.Unmarshal : ", err)
	}

	return apiObj
}
