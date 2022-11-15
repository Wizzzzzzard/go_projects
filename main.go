package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

const ()

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Zaphod",
		LastName:  "Beeblebrox",
	}

	fmt.Println(u)

	fmt.Println("Hello World")
	port := 3000
	_, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, nil
}
