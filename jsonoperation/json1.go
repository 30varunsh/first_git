package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}
type Info struct {
	Description string `json:"desc"`
}

func main() {
	fmt.Println("Hello World")
	jsonString := `{"name":"Battery sensor","capacity":40,"time":"2019-01-21T19:07:28Z","info":{
		"desc":"a sensor reading"
	}}` //nested json
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
