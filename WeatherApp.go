package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	//"log"
	"net/http"
)

func showWeatherApp(w fyne.Window) {

	//a := app.New()
	//w := a.NewWindow("Weather App")

	//w.Resize(fyne.NewSize(500, 280))

	//Api part

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=delhi&APPID=c0ca10236320d334cd6a73764467b459")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}

	//UI

	img := canvas.NewImageFromFile("C:/Users/91817/Desktop/go-workspace/OS/5539060.png")
	img.FillMode = canvas.ImageFillOriginal

	//combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
	//		log.Println("Select set to", value)
	//})

	label1 := canvas.NewText("Weather Details", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label2 := canvas.NewText(fmt.Sprintf("Country %s", weather.Sys.Country), color.Black)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather.Wind.Speed), color.Black)
	label4 := canvas.NewText(fmt.Sprintf("Temperature %2f", weather.Main.Temp-273.15), color.Black)
	label5 := canvas.NewText(fmt.Sprintf("Humidity 	%#v", weather.Main.Humidity), color.Black)
	label6 := canvas.NewText(("Delhi"), color.Black)

	res1, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=c0ca10236320d334cd6a73764467b459")
	if err != nil {
		fmt.Println(err)
	}
	defer res1.Body.Close()

	body1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather1, err := UnmarshalWeather(body1)
	if err != nil {
		fmt.Println(err)
	}

	img1 := canvas.NewImageFromFile("C:/Users/91817/Desktop/go-workspace/OS/4743914.png")
	img1.FillMode = canvas.ImageFillOriginal

	label11 := canvas.NewText("Weather Details", color.White)
	label11.TextStyle = fyne.TextStyle{Bold: true}
	label12 := canvas.NewText(fmt.Sprintf("Country %s", weather1.Sys.Country), color.Black)
	label13 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather1.Wind.Speed), color.Black)
	label14 := canvas.NewText(fmt.Sprintf("Temperature  %2f", weather1.Main.Temp-273.15), color.Black)
	label15 := canvas.NewText(fmt.Sprintf("Humidity 	%#v", weather1.Main.Humidity), color.Black)
	label16 := canvas.NewText(("Mumbai"), color.Black)

	weatherContainer :=
		container.NewHSplit(
			container.NewVBox(
				label6,
				label1,
				img,
				label2,
				label3,
				label4,
				label5,
			),
			// 2nd Section
			container.NewVBox(
				label16,
				label11,
				img1,
				label12,
				label13,
				label14,
				label15,
			),
		)

	w.SetContent(container.NewBorder(nil, panelContent, nil, nil, weatherContainer))

	w.Show()
}

func UnmarshalWeather(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
