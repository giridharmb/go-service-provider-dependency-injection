package infrastructure

import (
	"encoding/json"
	"net/http"
	"service-provider/entities"
	"service-provider/providers"
	"service-provider/service"
)

func FetchWeather(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var weatherData entities.WeatherDataHTTP
	err := json.NewDecoder(req.Body).Decode(&weatherData)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	apiDataProvider := providers.NewAPIData(weatherData.APIKey)
	weatherService := service.NewWeatherUsecase(apiDataProvider)
	weather, err := weatherService.FetchWeatherData(weatherData.City)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(res).Encode(ErrorResponse{Message: "Could Not Fetch Weather"})
		return
	}
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(weather)
}
