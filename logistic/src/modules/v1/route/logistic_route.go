package route

import (
	"github.com/depri11/lolipad/logistic/src/middleware"
	"github.com/depri11/lolipad/logistic/src/modules/v1/handler"
	"github.com/depri11/lolipad/logistic/src/modules/v1/repository"
	"github.com/depri11/lolipad/logistic/src/modules/v1/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AuthRouter(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/logistic").Subrouter()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	route.HandleFunc("/all", middleware.Do(handler.GetAllData, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/", middleware.Do(handler.GetData, middleware.CheckAuth)).Methods("GET")
}
