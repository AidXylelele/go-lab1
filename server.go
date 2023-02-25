package main

import (
	"go-lab1/controllers"
	"net/http"
)

func main() {
 http.HandleFunc("/time", controllers.TimeController)
 http.ListenAndServe(":8795", nil)
}