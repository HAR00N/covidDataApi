package models

type CovidData struct {
	Response Response `json:"response"`
}

type Response struct {
	ResultMsg  string   `json:"resultMsg"`
	Result     []Result `json:"result"`
	ResultCnt  string   `json:"resultCnt"`
	ResultCode string   `json:"resultCode"`
}

type Result struct {
	RateHospitalizations string `json:"rate_hospitalizations"`
	RateConfirmations    string `json:"rate_confirmations"`
	CntSevereSymptoms    string `json:"cnt_severe_symptoms"`
	CntDeaths            string `json:"cnt_deaths"`
	RateSevereSymptoms   string `json:"rate_severe_symptoms"`
	CntHospitalizations  string `json:"cnt_hospitalizations"`
	CntConfirmations     string `json:"cnt_confirmations"`
	Mmddhh               string `json:"mmddhh"`
	RateDeaths           string `json:"rate_deaths"`
}
