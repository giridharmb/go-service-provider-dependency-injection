package service

import (
	"log"
	"service-provider/entities"
)

type IWeatherQuery interface {
	GetWeatherV2() (entities.Weather, error)
}

type WeatherST struct {
	iData IWeatherQuery
}

func NewWeatherST(wq IWeatherQuery) *WeatherST {
	return &WeatherST{iData: wq}
}

func (w *WeatherST) FetchWeatherDetailsV2() (entities.Weather, error) {
	var weather entities.Weather
	var err error
	weather, err = w.iData.GetWeatherV2()
	if err != nil {
		log.Printf("error : FetchWeatherDetailsV2 : %v", err.Error())
		return weather, err
	}
	return weather, nil
}
