# Covid-19 Data Api

This repository was made only for uploading api data to the database.

## Data reference

- 공공데이터 포털 [Data.go.kr](data.go.kr/index.do) 

## Data format

    {
        "response": {
            "resultMsg": "조회성공",
            "result": [{
                "rate_hospitalizations": "0.44",
                "rate_confirmations": "48.66",
                "cnt_severe_symptoms": "251",
                "cnt_deaths": "43",
                "rate_severe_symptoms": "0.49",
                "cnt_hospitalizations": "229",
                "cnt_confirmations": "25125",
                "mmddhh": "5.20.00시",
                "rate_deaths": "0.08"
            }],
            "resultCnt": "1",
            "resultCode": "1"
        }
    }

## Repository layout
    covidDataApi
    ├ api
    │   └ dataportalApi.go : connect to api by http GET
    ├ config
    │   └ dbConnection.go : connect to database
    ├ models
    │   └ covidModel.go : api data format
    ├ main.go
    ├ go.mod
    └ sample.env : database and api url, key setting

# Dev Version
- Go : 1.18.1
- GoLand : 2022.1