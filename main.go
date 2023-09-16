package main

import (
	"fmt"
	"test/internal/pkg/app"
)

func main() {
	fmt.Println("Starting app")
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
