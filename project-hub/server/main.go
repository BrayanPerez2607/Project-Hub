package main

import (
	"fmt"
	"log"
	"net/http"
	"project-hub/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Software on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
