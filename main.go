package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func (w *Weather) checkStatus() (resWater string, resWind string) {
	switch {
	case w.Water < 5:
		resWater = "aman"
	case w.Water >= 6 && w.Water <= 8:
		resWater = "siaga"
	default:
		resWater = "tidak terdefinisi"
	}

	switch {
	case w.Wind < 5:
		resWind = "aman"
	case w.Wind >= 6 && w.Wind <= 8:
		resWind = "siaga"
	default:
		resWind = "tidak terdefinisi"
	}
	return resWater, resWind
}

func generateJSON(weather Weather) {

	res, err := json.Marshal(weather)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s\n", res)
	resWater, resWind := weather.checkStatus()
	fmt.Printf("status air : %s\n", resWater)
	fmt.Printf("status angin : %s\n", resWind)
}

func main() {
	weather := Weather{}

	for {
		time.Sleep(2 * time.Second)

		weather.Water = rand.Intn(15)
		weather.Wind = rand.Intn(15)
		generateJSON(weather)
	}
}
