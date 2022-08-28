package main

import (
	"flag"
	"log"
	"service-provider/infrastructure"
	"service-provider/providers"
	"service-provider/service"
	"service-provider/utils"
)

func checkForEmptyStr(data string, logstr string) bool {
	if data == "" {
		log.Printf("'%v' cannot be empty string !", logstr)
		return false
	}
	return true
}

func main() {

	operationPtr := flag.String("operation", "cli", "Run As 'cli' or 'api'")
	versionPtr := flag.String("version", "v1", "version of CLI , ex: 'v1' or 'v2' etc.")

	cityPtr := flag.String("city", "London", "City To Be Queried")
	apiKeyPtr := flag.String("apikey", "73ch4kdvw5fu3hxur7ahqq6f3nr9259x", "OpenWeather API Key")

	flag.Parse()

	city := *cityPtr
	apiKey := *apiKeyPtr
	operation := *operationPtr
	version := *versionPtr

	switch operation {

	case "cli":

		if checkForEmptyStr(city, "city") == false {
			return
		}

		if checkForEmptyStr(apiKey, "apikey") == false {
			return
		}

		switch version {

		case "v1":
			apiDataProvider := providers.NewAPIData(apiKey)
			weatherService := service.NewWeatherUsecase(apiDataProvider)
			weatherData, err := weatherService.FetchWeatherData(city)
			if err != nil {
				log.Printf("error : %v", err.Error())
				return
			}
			log.Printf("WeatherUsecase (1) >")
			utils.PrettyPrintData(weatherData)

		case "v2":
			wqDetails := providers.NewWeatherQueryDetails(apiKey, city)
			weather := service.NewWeatherST(wqDetails)
			weatherData, err := weather.FetchWeatherDetailsV2()
			if err != nil {
				log.Printf("error : %v", err.Error())
				return
			}
			log.Printf("WeatherUsecase (2) >")
			utils.PrettyPrintData(weatherData)

		default:
			log.Printf("please provide valid version of CLI !")
			return
		}

	case "api":

		httpRouter := infrastructure.NewMuxRouter()
		httpRouter.Post("/weather", infrastructure.FetchWeather)
		httpRouter.ServeEndpoint(":8181")
	default:
		log.Printf("Please provide valid operation !")
		return
	}
}
