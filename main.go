package main

import (
	"covid/api"
	"covid/config"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	godotenv.Load(".env")
	url := os.Getenv("covidUrl") + os.Getenv("covidKey")

	database := config.DbConnection()
	defer database.Close()

	data := api.CovidCall(url).Response.Result[0]
	dataDate := "2022-" + strings.Split(data.Mmddhh, ".")[0] + "-" + strings.Split(data.Mmddhh, ".")[1]
	data.Mmddhh = dataDate

	var query = `insert into covidstatus (RateHospitalizations,RateConfirmations,CntSevereSymptoms,CntDeaths,RateSevereSymptoms,CntHospitalizations,CntConfirmations,Mmddhh,RateDeaths) values ("` + data.RateHospitalizations + `", "` + data.RateConfirmations + `", "` + data.CntSevereSymptoms + `", "` + data.CntDeaths + `", "` + data.RateSevereSymptoms + `", "` + data.CntHospitalizations + `", "` + data.CntConfirmations + `", "` + data.Mmddhh + `", "` + data.RateDeaths + `")`

	fmt.Println(query)

	_, err := database.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("database upload 성공")
}
