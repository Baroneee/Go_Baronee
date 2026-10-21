package main

import (
	"fmt"
	"log"

	"maximilien-andile.com/errors/error-example/config"
)

func main() {
    confData, err := config.Load()
    if err != nil {
        log.Fatalf("Impossible to load application config because: %s", err)
    }
    fmt.Println(confData)
}