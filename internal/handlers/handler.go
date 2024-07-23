package handlers

import (
	"Hisobot/internal/service"
	"Hisobot/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler) UpdateFirstTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload models.RegionTable_1

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateFirstTable(uint(id), payload.StudentsCount, payload.ItRoomsCount, payload.FurnitureCount,
		payload.LocalNet, payload.ElectroLibCount, payload.MaterialsCount, payload.ItTeachersCount,
		payload.WithHighEduCount, payload.WithSecondEduCount, payload.UnqualifiedCount,
		payload.CompleteCourseCount, payload.LibraryCount, uint(regionID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
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

func (h *Handler) GetRegions(w http.ResponseWriter, r *http.Request) {
	regions, err := h.Service.GetAllRegions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regions)
}
