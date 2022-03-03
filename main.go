package main

import (
	"fmt"
	"snykAPIClient/internal/config"
)

func main() {
	fileName := "config.yaml"
	fmt.Println(config.Config(fileName))
}
