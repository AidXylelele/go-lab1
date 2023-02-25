package services

import (
	"encoding/json"
	"time"
)

type JSONTime struct {
	CURRENT_TIME string `json:"time"`
}

func GetCurrentTime() string {
	currentTime := JSONTime{CURRENT_TIME: time.Now().Format(time.RFC3339)}
	json_data, err := json.Marshal(currentTime)
	if err != nil {
		return "Something went wrong"
	}
	return string(json_data)
}
