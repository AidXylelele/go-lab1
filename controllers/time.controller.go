package controllers

import (
	"fmt"
	"go-lab1/services"
	"io"
	"net/http"
)

func TimeController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request\n")
	fmt.Printf(r.Method)
	data := services.GetCurrentTime()
	io.WriteString(w, string(data))
}