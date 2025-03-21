package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var rdb *redis.Client

func main() {
	ctx := context.Background()

	log.Println("started")
	r := mux.NewRouter()
	rdb = initdb()

	registerhandlers(ctx, r)

	go func() { http.ListenAndServe(":8080", r) }()
	go func() {
		mockSensorStream()
	}()
	tick := time.NewTicker(2 * time.Second)

	go sendRequestToSensor(ctx, tick)

	// start  main server at 8080
	http.ListenAndServe(":8080", r)
	return

}

// register diplay handlers
func registerhandlers(ctx context.Context, r *mux.Router) {

	r.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {

		sensorID := r.URL.Query().Get("device_id")
		startTime := r.URL.Query().Get("start")
		endTime := r.URL.Query().Get("end")

		if sensorID == "" || startTime == "" || endTime == "" {
			http.Error(w, "Missing query parameters", http.StatusBadRequest)
			return
		}
		startTimeInt, _ := strconv.Atoi(startTime)
		endTimeInt, _ := strconv.Atoi(endTime)

		data, err := rdb.TSRange(ctx, sensorID, startTimeInt, endTimeInt).Result()

		if err != nil {
			log.Fatalf("Error querying time series data: %v", err)
		}

		fmt.Println(data)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		}

	}).Methods("GET")

}

func sendRequestToSensor(ctx context.Context, tick *time.Ticker) {
	client := http.Client{Timeout: 3 * time.Second}

	for {
		select {
		case <-tick.C:
			//fmt.Println("send request to sensor")

			request, err := http.NewRequest("GET", "http://localhost:8081/sensor/1", nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			// Execute request using http.Client
			response, err := client.Do(request)
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				continue
			}
			res := &sensorpayload{}
			err = json.Unmarshal(body, res)
			if err != nil {
				log.Println("Error unmarshalling JSON:", err)
				continue
			}
			fmt.Println("Sensor data:", res)

			if ok := validate(res); !ok {
				log.Println("validation failed for payload, continue")
				continue
			}
			_, err = rdb.TSAdd(ctx, res.DeviceId, res.Timestamp, float64(res.Temp)).Result()
			if err != nil {
				log.Println("Error adding to Redis:", err)
			}

		}
	}
}

func validate(res *sensorpayload) bool {

	if res.DeviceId != "1" {
		log.Println("device id invalid")
		return false
	}
	if res.Temp > 100 {
		log.Println("temp invalid")

		return false
	}
	return true

}
