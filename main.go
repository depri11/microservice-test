package main

import (
	"fmt"

	"github.com/depri11/lolipad/src/configs"
)

func main() {
	db, err := configs.SetupDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connect")
}
