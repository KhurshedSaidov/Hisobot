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

	var payload models.RegionTable_1

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateFirstTable(payload.StudentsCount, payload.ItRoomsCount, payload.FurnitureCount,
		payload.LocalNet, payload.ElectroLibCount, payload.MaterialsCount, payload.ItTeachersCount,
		payload.WithHighEduCount, payload.WithSecondEduCount, payload.UnqualifiedCount,
		payload.CompleteCourseCount, payload.LibraryCount, uint(regionID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateSecondTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload models.RegionTable_2

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateSecondTable(payload.SchoolsCount, payload.ComputersCount,
		payload.LearningComputersCount, payload.RepairComputerCount, payload.BrokeComputerCount,
		payload.DecommissionedComputersCount, payload.PurchasedComputersCount, payload.LicenseComputersCount,
		payload.PrintersCount, payload.BrokePrintersCount, payload.ScannersCount, payload.ConnectedToInternetCount,
		payload.ConnectionType, payload.ElectronicBoardsCount, payload.ProjectorsCount, payload.WebsiteCount,
		uint(regionID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateThirdTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var payload models.RegionTable_3

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateThirdTable(payload.ComputersBuyPlan, payload.ComputersBuyPlan6Month,
		payload.BoardsBuyPlan6Month, payload.BoardsBuy6Month, payload.BoardsBuyPercent, payload.VideoProjectorCount,
		payload.PrintersBuy, payload.IctFinancing, uint(regionID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateFoundations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionID, err := strconv.ParseUint(vars["region_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var payload models.Foundations

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateFoundations(uint(id), payload.Foundation, payload.ComputersCount,
		payload.BoardsCount, payload.ProjectorsCount, payload.PrinterCount, uint(regionID)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload models.Regions

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateInfo(uint(ID), payload.School, payload.NameSurname, payload.Specialization,
		payload.WhereGraduate, payload.WhenGraduate, payload.WorkExperience, payload.BirthYear, payload.Education,
		payload.PhoneNumber); err != nil {
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

func (h *Handler) GetAllSums(w http.ResponseWriter, r *http.Request) {
	sums, err := h.Service.GetAllSums()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sums)
}

func (h *Handler) CreateInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionId, err := strconv.ParseUint(vars["region_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var info models.Regions
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateInfo(info.School, info.NameSurname, info.Specialization, info.WhereGraduate,
		info.WhenGraduate, info.WorkExperience, info.BirthYear, info.Education, info.PhoneNumber, uint(regionId)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Информация добавлена!"))
}

func (h *Handler) DeleteInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Неверно указан идентификатор id"))
		return
	}

	err = h.Service.DeleteInfo(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetRegionsByRegionID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionIDStr := vars["region_id"]
	regionID, err := strconv.Atoi(regionIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid region ID"))
		return
	}

	regions, err := h.Service.GetRegionsByRegionID(uint(regionID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving regions"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regions)
}
