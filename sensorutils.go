package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

type sensorpayload struct {
	DeviceId   string `json:"device_id"`
	Temp       int    `json:"temp"`
	Timestamp  int64  `json:"timestamp"`
	DeviceType string `json:"device_type"`
}

// sends a random temperature with timestamp when it gets a POST request from server
// running at 8081
func mockSensorStream() {

	getTemp := func() int {
		return rand.Intn(100) // returnign random int
	}

	r := mux.NewRouter()
	r.HandleFunc("/sensor/1", func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Sensor 1 data requested")
		payload := sensorpayload{
			DeviceId:   "1",
			Temp:       getTemp(),
			Timestamp:  time.Now().Unix(),
			DeviceType: "sensor",
		}

		x := json.NewEncoder(w)

		err := x.Encode(payload)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		return

	})

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		return
	}

}
