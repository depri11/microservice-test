package route

import (
	"github.com/depri11/lolipad/auth/src/modules/v1/handler"
	"github.com/depri11/lolipad/auth/src/modules/v1/repository"
	"github.com/depri11/lolipad/auth/src/modules/v1/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AuthRouter(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/auth").Subrouter()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	route.HandleFunc("/register", handler.Register).Methods("POST")
	route.HandleFunc("/", handler.Login).Methods("POST")
}
