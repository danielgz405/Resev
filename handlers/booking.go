package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/danielgz405/Resev/models"
	"github.com/danielgz405/Resev/repository"
	"github.com/danielgz405/Resev/responses"
	"github.com/danielgz405/Resev/server"
	"github.com/danielgz405/Resev/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertBookingRequest struct {
	TableId     string `json:"table_id"`
	UserId      string `json:"user_id"`
	OrderId     string `json:"order_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Hour        string `json:"hour"`
}

type GetBookingsByIdsRequest struct {
	Ids []string `json:"ids"`
}

func CreateBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertBookingRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		tableID, err := primitive.ObjectIDFromHex(req.TableId)
		if err != nil {
			fmt.Println("Invalid request body", err)
			return
		}
		userID, err := primitive.ObjectIDFromHex(req.UserId)
		if err != nil {
			fmt.Println("Invalid request body", err)
			return
		}
		orderID, err := primitive.ObjectIDFromHex(req.OrderId)
		if err != nil {
			fmt.Println("Invalid request body", err)
			return
		}
		layout := "Mon Jan 02 2006"
		date, err := time.Parse(layout, req.Date)
		if err != nil {
			fmt.Println("Invalid request body")
			return
		}
		if !utils.IsTimeFormat(req.Hour) {
			fmt.Printf("Invalid request body: %s  hour is invalid.\n", req.Hour)
			return
		}
		createBooking := models.InsertBooking{
			TableId:     tableID,
			UserId:      userID,
			OrderId:     orderID,
			Name:        req.Name,
			Description: req.Description,
			Date:        date,
			Hour:        req.Hour,
		}
		booking, err := repository.InsertBooking(r.Context(), &createBooking)
		if err != nil {
			responses.BadRequest(w, "Error creating booking")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	}
}

func GetBookingByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		booking, err := repository.GetBookingById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting booking")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	}
}

func UpdateBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := InsertBookingRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		tableID, err := primitive.ObjectIDFromHex(req.TableId)
		if err != nil {
			fmt.Println("Invalid request body", err)
			return
		}
		layout := "Mon Jan 02 2006"
		date, err := time.Parse(layout, req.Date)
		if err != nil {
			fmt.Println("Invalid request body")
			return
		}
		if !utils.IsTimeFormat(req.Hour) {
			fmt.Printf("Invalid request body: %s  hour is invalid.\n", req.Hour)
			return
		}
		updateBooking := models.UpdateBooking{
			TableId:     tableID,
			Name:        req.Name,
			Description: req.Description,
			Date:        date,
			Hour:        req.Hour,
		}
		booking, err := repository.UpdateBooking(r.Context(), &updateBooking, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating booking")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booking)
	}
}

func DeleteBookingHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err := repository.DeleteBooking(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting booking")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListBookingsByPageHandler(s server.Server) http.HandlerFunc {
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
		bookings, quantity, err := repository.ListBookingsByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting bookings")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.BookingResponse{
			Booking:  bookings,
			Quantity: quantity,
		})
	}
}

func ListBookingsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		bookings, err := repository.ListBookings(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting bookings")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
	}
}

func GetBookingsByIdsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_2(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := GetBookingsByIdsRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		bookings, err := repository.GetBookingsByIds(r.Context(), req.Ids)
		if err != nil {
			responses.BadRequest(w, "Error getting bookings")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookings)
	}
}
