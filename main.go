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
}
