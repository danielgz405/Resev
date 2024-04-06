package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/danielgz405/Resev/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role_id  string `bson:"role_id" json:"role_id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
}

type UpdateUserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Role_id  string   `bson:"role_id" json:"role_id"`
	Image    string   `json:"image"`
	ImageRef string   `json:"image_ref"`
	Bookings []string `json:"bookings"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)
		w.Header().Set("Content-Type", "application/json")
		var req = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		client, err := repository.GetRoleByName(r.Context(), "client")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		createUser := models.InsertUser{
			Email:    req.Email,
			Phone:    req.Phone,
			Role_id:  client.Id.Hex(),
			Password: string(hashedPassword),
			Name:     req.Name,
		}

		profile, err := repository.InsertUser(r.Context(), &createUser)
		if err != nil {
			responses.BadRequest(w, "Error creating user")
			return
		}

		err = repository.AuditOperation(r.Context(), profile.Id.Hex(), "users", "insert")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		role, err := repository.GetRoleById(r.Context(), profile.Role_id)
		if err != nil {
			responses.BadRequest(w, "Error getting role")
			return
		}

		responseProfile := responses.UserResponse{
			Id:       profile.Id.Hex(),
			Name:     profile.Name,
			Email:    profile.Email,
			Phone:    profile.Phone,
			Role:     *role,
			Image:    profile.Image,
			ImageRef: profile.ImageRef,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseProfile)
	}
}

func CreateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)
		w.Header().Set("Content-Type", "application/json")
		var req = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		_, err = primitive.ObjectIDFromHex(req.Role_id)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		createUser := models.InsertUser{
			Email:    req.Email,
			Phone:    req.Phone,
			Role_id:  req.Role_id,
			Password: string(hashedPassword),
			Name:     req.Name,
		}
		profile, err := repository.InsertUser(r.Context(), &createUser)
		if err != nil {
			responses.BadRequest(w, "Error creating user")
			return
		}

		err = repository.AuditOperation(r.Context(), profile.Id.Hex(), "users", "insert")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		role, err := repository.GetRoleById(r.Context(), req.Role_id)
		if err != nil {
			responses.BadRequest(w, "Error getting role")
			return
		}

		responseProfile := responses.UserResponse{
			Id:       profile.Id.Hex(),
			Name:     profile.Name,
			Email:    profile.Email,
			Phone:    profile.Phone,
			Role:     *role,
			Image:    profile.Image,
			ImageRef: profile.ImageRef,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseProfile)
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)
		var req = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		user, err := repository.GetUserByEmail(r.Context(), req.Email)
		if err != nil {
			responses.BadRequest(w, "Error getting user")
			return
		}
		if user == nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		var profile = models.Profile{
			Id:       user.Id,
			Name:     user.Name,
			Phone:    user.Phone,
			Email:    user.Email,
			Role_id:  user.Role_id,
			Image:    user.Image,
			ImageRef: user.ImageRef,
			Bookings: user.Bookings,
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		role, err := repository.GetRoleById(r.Context(), user.Role_id)
		if err != nil {
			responses.BadRequest(w, "Error getting role")
			return
		}

		booking := []models.Booking{}

		if len(profile.Bookings) > 0 {
			booking, err = repository.GetBookingsByIds(r.Context(), profile.Bookings)
			if err != nil {
				responses.BadRequest(w, "Error getting booking"+err.Error())
				return
			}
		}

		responseProfile := responses.UserResponse{
			Id:       profile.Id.Hex(),
			Name:     profile.Name,
			Email:    profile.Email,
			Phone:    profile.Phone,
			Role:     *role,
			Image:    profile.Image,
			ImageRef: profile.ImageRef,
			Bookings: booking,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseProfile)
	}
}

func ProfileHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		profile, err := repository.GetUserById(r.Context(), params["userId"])
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		role, err := repository.GetRoleById(r.Context(), profile.Role_id)
		if err != nil {
			responses.BadRequest(w, "Error getting role"+err.Error())
			return
		}

		booking := []models.Booking{}

		if len(profile.Bookings) > 0 {
			booking, err = repository.GetBookingsByIds(r.Context(), profile.Bookings)
			if err != nil {
				responses.BadRequest(w, "Error getting booking"+err.Error())
				return
			}
		}

		responseProfile := responses.UserResponse{
			Id:       profile.Id.Hex(),
			Name:     profile.Name,
			Email:    profile.Email,
			Phone:    profile.Phone,
			Role:     *role,
			Image:    profile.Image,
			ImageRef: profile.ImageRef,
			Bookings: booking,
		}

		// Handle request
		json.NewEncoder(w).Encode(responseProfile)
	}
}

func UpdateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)

		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		var req = UpdateUserRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		data := models.UpdateUser{
			Id:       params["id"],
			Name:     req.Name,
			Email:    req.Email,
			Phone:    req.Phone,
			Role_id:  req.Role_id,
			Image:    req.Image,
			ImageRef: req.ImageRef,
			Bookings: req.Bookings,
		}

		err = repository.AuditOperation(r.Context(), params["id"], "users", "updated")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		updatedUser, err := repository.UpdateUser(r.Context(), data)
		if err != nil {
			responses.BadRequest(w, "Error updating user")
			return
		}

		role, err := repository.GetRoleById(r.Context(), updatedUser.Role_id)
		if err != nil {
			responses.BadRequest(w, "Error getting role"+err.Error())
			return
		}

		booking := []models.Booking{}

		if len(updatedUser.Bookings) > 0 {
			booking, err = repository.GetBookingsByIds(r.Context(), updatedUser.Bookings)
			if err != nil {
				responses.BadRequest(w, "Error getting booking"+err.Error())
				return
			}
		}

		responseProfile := responses.UserResponse{
			Id:       updatedUser.Id.Hex(),
			Name:     updatedUser.Name,
			Email:    updatedUser.Email,
			Phone:    updatedUser.Phone,
			Role:     *role,
			Image:    updatedUser.Image,
			ImageRef: updatedUser.ImageRef,
			Bookings: booking,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseProfile)
	}
}

func DeleteUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//conexion
		utils.DatabaseConnection(s)
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		err := repository.AuditOperation(r.Context(), params["id"], "users", "delete")
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		err = repository.DeleteUser(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting user")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
