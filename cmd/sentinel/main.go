package main

import (
	"fmt"

	"github.com/2pizzzza/sentinetAgent/internal/config"
)

func main() {

	config, err := config.New("config/config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Println(*config)
}
