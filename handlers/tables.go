package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/danielgz405/Resev/database"
	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/gorilla/mux"
)

type InsertTableRequest struct {
	Number      string `json:"number"`
	Observation string `json:"observation"`
}

func databaseConnection_2(s server.Server) {
	repo, err := database.NewMongoRepo(s.Config().DB_URI_2)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
}

func CreateTableHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		databaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertTableRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		createTable := models.InsertTable{
			Number:      req.Number,
			Observation: req.Observation,
		}
		table, err := repository.InsertTable(r.Context(), &createTable)
		if err != nil {
			responses.BadRequest(w, "Error creating table")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(table)
	}
}

func GetTableByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		databaseConnection_2(s)
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
		databaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err := repository.DeleteTable(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting table")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListTablesByPageHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		databaseConnection_2(s)
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
		databaseConnection_2(s)
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
