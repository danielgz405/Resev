package handlers

import (
	"encoding/json"
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
		utils.DatabaseConnection_3(s)
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
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		userID, err := primitive.ObjectIDFromHex(req.UserId)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		orderID, err := primitive.ObjectIDFromHex(req.OrderId)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		layout := "Mon Jan 02 2006"
		// dates should have the same format "Fri Apr 05 2024"
		date, err := time.Parse(layout, req.Date)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		if !utils.IsTimeFormat(req.Hour) {
			responses.BadRequest(w, "Invalid request body"+"req.Hour  hour is invalid.\n")
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
		utils.DatabaseConnection_3(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		booking, err := repository.GetBookingById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting booking")
			return
		}

		table, err := repository.GetTableById(r.Context(), booking.TableId.Hex())
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		profile, err := repository.GetUserById(r.Context(), booking.UserId.Hex())
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		order, err := repository.GetOrderById(r.Context(), booking.OrderId.Hex())
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		responseBooking := responses.BookingResponse{
			Id:          booking.Id.Hex(),
			Table:       *table,
			User:        *profile,
			Order:       *order,
			Name:        booking.Name,
			Description: booking.Description,
			Date:        booking.Date,
			Hour:        booking.Hour,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseBooking)
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
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		layout := "Mon Jan 02 2006"
		date, err := time.Parse(layout, req.Date)
		if err != nil {
			responses.BadRequest(w, "Invalid request body"+err.Error())
			return
		}
		if !utils.IsTimeFormat(req.Hour) {
			responses.BadRequest(w, "Invalid request body"+"req.Hour  hour is invalid.\n")
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
		utils.DatabaseConnection_3(s)
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
		utils.DatabaseConnection_3(s)
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

		responseBookings := []responses.BookingResponse{}
		for _, booking := range bookings {
			table, err := repository.GetTableById(r.Context(), booking.TableId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			profile, err := repository.GetUserById(r.Context(), booking.UserId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			order, err := repository.GetOrderById(r.Context(), booking.OrderId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			response := responses.BookingResponse{
				Id:          booking.Id.Hex(),
				Table:       *table,
				User:        *profile,
				Order:       *order,
				Name:        booking.Name,
				Description: booking.Description,
				Date:        booking.Date,
				Hour:        booking.Hour,
			}

			responseBookings = append(responseBookings, response)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.BookingsResponse{
			Booking:  responseBookings,
			Quantity: quantity,
		})
	}
}

func ListBookingsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_3(s)
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
		utils.DatabaseConnection_3(s)
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

		responseBookings := []responses.BookingResponse{}
		for _, booking := range bookings {
			table, err := repository.GetTableById(r.Context(), booking.TableId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			profile, err := repository.GetUserById(r.Context(), booking.UserId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			order, err := repository.GetOrderById(r.Context(), booking.OrderId.Hex())
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			response := responses.BookingResponse{
				Id:          booking.Id.Hex(),
				Table:       *table,
				User:        *profile,
				Order:       *order,
				Name:        booking.Name,
				Description: booking.Description,
				Date:        booking.Date,
				Hour:        booking.Hour,
			}

			responseBookings = append(responseBookings, response)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseBookings)
	}
}
