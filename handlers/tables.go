package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/danielgz405/Resev/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertTableRequest struct {
	Number      string `json:"number"`
	Observation string `json:"observation"`
	ImageBase64 string `json:"image_base64"`
}

func CreateTableHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertTableRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		imageBytes, err := base64.StdEncoding.DecodeString(req.ImageBase64)
		if err != nil {
			responses.BadRequest(w, "Error decoding base64 image")
			return
		}

		createTable := models.InsertTable{
			Number:      req.Number,
			Observation: req.Observation,
			ImageBase64: primitive.Binary{
				Subtype: 0,
				Data:    imageBytes,
			},
		}
		table, err := repository.InsertTable(r.Context(), &createTable)
		if err != nil {
			responses.BadRequest(w, "Error creating table")
			return
		}

		err = repository.AuditOperation(r.Context(), table.Id.Hex(), "table", "insert")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(table)
	}
}

func GetTableByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		table, err := repository.GetTableById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting table")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(table)
	}
}

func UpdateTableHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := InsertTableRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		updateTable := models.UpdateTable{
			Number:      req.Number,
			Observation: req.Observation,
		}

		err = repository.AuditOperation(r.Context(), params["id"], "table", "update")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		table, err := repository.UpdateTable(r.Context(), &updateTable, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating table")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(table)
	}
}

func DeleteTableHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		err := repository.AuditOperation(r.Context(), params["id"], "table", "delete")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = repository.DeleteTable(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting table")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListTablesByPageHandler(s server.Server) http.HandlerFunc {
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
		tables, quantity, err := repository.ListTablesByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting tables")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.TableResponse{
			Table:    tables,
			Quantity: quantity,
		})
	}
}

func ListTablesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		tables, err := repository.ListTables(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting tables")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tables)
	}
}
