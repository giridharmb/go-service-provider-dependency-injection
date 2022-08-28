package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service-provider/entities"
)

type WeatherQueryDetails struct {
	ApiKey string
	City   string
}

func NewWeatherQueryDetails(apikey string, city string) *WeatherQueryDetails {
	return &WeatherQueryDetails{
		ApiKey: apikey,
		City:   city,
	}
}

func (wq *WeatherQueryDetails) GetWeatherV2() (entities.Weather, error) {
	// compose the url. note that it's not the best way to add query params.
	path := fmt.Sprintf(pathFormatWeatherByCity, wq.City, wq.ApiKey)
	log.Printf("@ provider : path : %v", path)

	completeURL := endpoint + path
	log.Printf("@ provider : completeURL : %v", completeURL)

	res, err := http.Get(completeURL)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeatherV2 failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// read the response body and encode it into the respose struct
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeatherV2 failed reading body: %s", err)
	}

	var weatherResponse entities.WeatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherResponse); err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeatherV2 failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return entities.Weather{}, fmt.Errorf("providers.GetWeatherV2 got error from OpenWeather: %s", weatherResponse.Message)
	}

	// return the external response converted into an entity
	return weatherResponse.ToWeather(), nil
}
