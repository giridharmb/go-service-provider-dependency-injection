package service

import (
	"log"
	"service-provider/entities"
)

type WeatherRepository interface {
	GetWeather(city string) (entities.Weather, error)
}

type WeatherUsecase struct {
	repo WeatherRepository
}

func NewWeatherUsecase(r WeatherRepository) *WeatherUsecase {
	return &WeatherUsecase{repo: r}
}

func (uc *WeatherUsecase) FetchWeatherData(city string) (entities.Weather, error) {
	var weather entities.Weather
	var err error
	weather, err = uc.repo.GetWeather(city)
	if err != nil {
		log.Printf("error : FetchWeatherData : %v", err.Error())
		return weather, err
	}
	return weather, nil
}
