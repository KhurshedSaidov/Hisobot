package handlers

import (
	"Hisobot/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler) GetReportWithRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	region, err := h.Service.GetReportByRegionId(uint(regionId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(region)
}
