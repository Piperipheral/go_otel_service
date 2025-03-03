package main

import (
	"log"
	"net/http"

	"piperipheral.com/go_otel_service/internal/rolldice"
)

func main() {
	http.HandleFunc("/rolldice", rolldice.RollDice)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
