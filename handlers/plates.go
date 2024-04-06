package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/danielgz405/Resev/utils"
	"github.com/gorilla/mux"
)

type InsertPlateRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Details  string  `json:"details"`
	Image    string  `json:"image"`
	ImageRef string  `json:"image_ref"`
}

type GetPlatesByIdsRequest struct {
	Ids []string `json:"ids"`
}

func CreatePlateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertPlateRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		createPlate := models.InsertPlate{
			Name:     req.Name,
			Price:    req.Price,
			Details:  req.Details,
			Image:    req.Image,
			ImageRef: req.ImageRef,
		}
		plate, err := repository.InsertPlate(r.Context(), &createPlate)
		if err != nil {
			responses.BadRequest(w, "Error creating plate")
			return
		}

		err = repository.AuditOperation(r.Context(), plate.Id.Hex(), "plate", "insert")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plate)
	}
}

func GetPlateByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		plate, err := repository.GetPlateById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting plate")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plate)
	}
}

func UpdatePlateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := InsertPlateRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		updatePlate := models.UpdatePlate{
			Name:     req.Name,
			Price:    req.Price,
			Details:  req.Details,
			Image:    req.Image,
			ImageRef: req.ImageRef,
		}

		err = repository.AuditOperation(r.Context(), params["id"], "plate", "update")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		plate, err := repository.UpdatePlate(r.Context(), &updatePlate, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating plate")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plate)
	}
}

func DeletePlateHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		err := repository.AuditOperation(r.Context(), params["id"], "plate", "delete")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = repository.DeletePlate(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting plate")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListPlatesByPageHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		limit, err := strconv.Atoi(params["limit"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}
		page, err := strconv.Atoi(params["page"])
		if err != nil {
			responses.BadRequest(w, "Bad request")
			return
		}
		plates, quantity, err := repository.ListPlatesByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting plates")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.PlateResponse{
			Plate:    plates,
			Quantity: quantity,
		})
	}
}

func ListPlatesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		plates, err := repository.ListPlates(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting plates")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plates)
	}
}

func GetPlatesByIdsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := GetPlatesByIdsRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		plates, err := repository.GetPlatesByIds(r.Context(), req.Ids)
		if err != nil {
			responses.BadRequest(w, "Error getting plates")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(plates)
	}
}
