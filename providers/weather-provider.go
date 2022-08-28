package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service-provider/entities"
)

const (
	endpoint                = "https://api.openweathermap.org/data/2.5"
	pathFormatWeatherByCity = "/weather?q=%s&appid=%s&units=metric"
)

type WeatherApi struct {
	ApiKey string
}

func NewAPIData(apiKey string) *WeatherApi {
	return &WeatherApi{
		ApiKey: apiKey,
	}
}

func (data *WeatherApi) GetWeather(city string) (entities.Weather, error) {
	// compose the url. note that it's not the best way to add query params.
	path := fmt.Sprintf(pathFormatWeatherByCity, city, data.ApiKey)
	log.Printf("@ provider : path : %v", path)

	completeURL := endpoint + path
	log.Printf("@ provider : completeURL : %v", completeURL)

	res, err := http.Get(completeURL)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeather failed http GET: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// read the response body and encode it into the respose struct
	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeather failed reading body: %s", err)
	}

	var weatherResponse entities.WeatherResponse
	if err = json.Unmarshal(bodyRaw, &weatherResponse); err != nil {
		return entities.Weather{}, fmt.Errorf("providers.GetWeather failed encoding body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return entities.Weather{}, fmt.Errorf("providers.GetWeather got error from OpenWeather: %s", weatherResponse.Message)
	}

	// return the external response converted into an entity
	return weatherResponse.ToWeather(), nil
}
