package main

type Sensor struct {
	Name string `json:"name"`
	Val1 int    `json:"val1"`
	Val2 int    `json:"val2"`
}

type Sensors []Sensor
