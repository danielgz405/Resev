package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/danielgz405/Resev/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateRoleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertRoleRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		if strings.Contains(strings.ToLower(req.Name), strings.ToLower("client")) || strings.Contains(strings.ToLower(req.Name), strings.ToLower("admin")) {
			responses.BadRequest(w, "The name you have entered is reserved, please enter another one.")
			return
		}
		createRole := models.InsertRole{
			Id:          primitive.NewObjectID(),
			Name:        req.Name,
			Description: req.Description,
		}
		role, err := repository.InsertRole(r.Context(), &createRole)
		if err != nil {
			responses.BadRequest(w, "Error creating role")
			return
		}

		err = repository.AuditOperation(r.Context(), role.Id.Hex(), "role", "insert")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(role)
	}
}

func GetRoleByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		role, err := repository.GetRoleById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting role")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(role)
	}
}

func UpdateRoleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := InsertRoleRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		updateRole := models.UpdateRole{
			Name:        req.Name,
			Description: req.Description,
		}

		err = repository.AuditOperation(r.Context(), params["id"], "role", "update")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		role, err := repository.UpdateRole(r.Context(), &updateRole, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating role")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(role)
	}
}

func DeleteRoleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		err := repository.AuditOperation(r.Context(), params["id"], "role", "delete")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = repository.DeleteRole(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting role")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListRolesByPageHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
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
		roles, quantity, err := repository.ListRolesByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting roles")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.RoleResponse{
			Role:     roles,
			Quantity: quantity,
		})
	}
}

func ListRolesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		roles, err := repository.ListRoles(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting roles")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(roles)
	}
}
