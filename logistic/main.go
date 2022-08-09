package main

import (
	"net/http"
	"os"

	"github.com/depri11/lolipad/logistic/src/routes"
)

func main() {
	if mainRoute, err := routes.SetupRouter(); err == nil {
		var addrs string = "0.0.0.0:3001"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}

		if err := http.ListenAndServe(addrs, mainRoute); err != nil {
			panic(err)
		}
	}

}
