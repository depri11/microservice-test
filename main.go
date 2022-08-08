package main

import (
	"net/http"

	"github.com/depri11/lolipad/src/routers"
)

func main() {
	mainRoute, err := routers.SetupRouter()
	if err != nil {
		panic(err)
	}

	var addrs string = "0.0.0.0:3000"

	http.ListenAndServe(addrs, mainRoute)
}
