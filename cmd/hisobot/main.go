package main

import (
	"Hisobot/configs"
	"Hisobot/internal/handlers"
	"Hisobot/internal/repository"
	"Hisobot/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouters(handler *handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get/{id}", handler.GetReportWithRegion).Methods("GET")

	return r
}

func Run() error {
	err := configs.InitDatabase()
	if err != nil {
		panic(err)
	}

	Repository := repository.NewRepository(configs.DB)
	Service := service.Service{Repository}
	Handler := handlers.Handler{&Service}

	router := InitRouters(&Handler)
	config, err := configs.InitConfigs()

	if err != nil {
		return err
	}

	address := config.Server.Host + ":" + config.Server.Port
	err = http.ListenAndServe(address, router)

	if err != nil {
		log.Println("listen and serve error", err)
		return err
	}
	return nil
}

func main() {
	Run()
}
