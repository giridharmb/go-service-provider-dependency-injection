package main

import (
	"flag"
	"log"
	"service-provider/infrastructure"
	"service-provider/providers"
	"service-provider/service"
	"service-provider/utils"
)

func main() {

	operationPtr := flag.String("operation", "cli", "Run As 'cli' or 'api'")

	cityPtr := flag.String("city", "London", "City To Be Queried")
	apiKeyPtr := flag.String("apikey", "73ch4kdvw5fu3hxur7ahqq6f3nr9259x", "OpenWeather API Key")

	flag.Parse()

	city := *cityPtr
	apiKey := *apiKeyPtr
	operation := *operationPtr

	switch operation {
	case "cli":
		if city == "" {
			log.Printf("please provide valid 'city' !")
			return
		}

		if apiKey == "" {
			log.Printf("please provide valid 'apikey' !")
			return
		}

		apiDataProvider := providers.NewAPIData(apiKey)
		weatherService := service.NewWeatherUsecase(apiDataProvider)
		weatherData, err := weatherService.FetchWeatherData(city)
		if err != nil {
			log.Printf("error : %v", err.Error())
			return
		}
		log.Printf("WeatherUsecase >")
		utils.PrettyPrintData(weatherData)
	case "api":
		//weatherController := infrastructure.NewWeatherController()
		httpRouter := infrastructure.NewMuxRouter()
		httpRouter.Post("/weather", infrastructure.FetchWeather)
		httpRouter.ServeEndpoint(":8181")
	default:
		log.Printf("Please provide valid operation !")
		return
	}
}
