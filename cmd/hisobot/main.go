package main

import (
	"Hisobot/configs"
	"Hisobot/internal/handlers"
	"Hisobot/internal/repository"
	"Hisobot/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func InitRouters(handler *handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/get", handler.GetRegions).Methods("GET")
	r.HandleFunc("/avg", handler.GetAllSums).Methods("GET")
	r.HandleFunc("/get/{id}", handler.GetReportWithRegion).Methods("GET")
	r.HandleFunc("/table1/{id}", handler.UpdateFirstTable).Methods("PUT")
	r.HandleFunc("/table2/{id}", handler.UpdateSecondTable).Methods("PUT")
	r.HandleFunc("/table3/{id}", handler.UpdateThirdTable).Methods("PUT")
	r.HandleFunc("/table4/{region_id}/{id}", handler.UpdateFoundations).Methods("PUT")
	r.HandleFunc("/info/{region_id}", handler.CreateInfoHandler).Methods("POST")
	r.HandleFunc("/info/{region_id}", handler.GetRegionsByRegionID).Methods("GET")
	r.HandleFunc("/info/{id}", handler.UpdateInfoHandler).Methods("PUT")
	r.HandleFunc("/info/{id}", handler.DeleteInfoHandler).Methods("DELETE")

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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	address := config.Server.Host + ":" + config.Server.Port
	err = http.ListenAndServe(address, handler)

	if err != nil {
		log.Println("listen and serve error", err)
		return err
	}
	return nil
}

func main() {
	Run()
}
