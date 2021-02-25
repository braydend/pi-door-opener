package main

import (
	"braydend/pi-door-opener/web"
	"log"
	"net/http"
)

func main() {
	web.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
