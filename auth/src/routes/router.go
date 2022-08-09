package routes

import (
	"errors"
	"net/http"

	"github.com/depri11/lolipad/auth/src/configs"
	"github.com/depri11/lolipad/auth/src/modules/v1/route"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRouter() (http.Handler, error) {
	r := mux.NewRouter()
	db, err := configs.SetupDB()
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(r)

	route.AuthRouter(r, db)

	return c, nil
}
