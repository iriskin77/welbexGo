package models

type Location struct {
	Id         int     `json:"id"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Zip        int     `json:"zip"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	Created_at string  `json:"created_at"`
}