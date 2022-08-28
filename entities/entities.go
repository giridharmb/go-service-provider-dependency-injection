package entities

type Weather struct {
	Temp     float32
	Pressure float32
	MinTemp  float32
	MaxTemp  float32
}

type WeatherResponse struct {
	Message string
	Main    struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	}
}

func (r WeatherResponse) ToWeather() Weather {
	return Weather{
		Temp:     r.Main.Temp,
		Pressure: r.Main.Pressure,
		MinTemp:  r.Main.TempMin,
		MaxTemp:  r.Main.TempMax,
	}
}

type WeatherDataHTTP struct {
	APIKey string `json:"apikey"`
	City   string `json:"city"`
}
