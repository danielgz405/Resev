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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertOrderRequest struct {
	TableId  string   `json:"table_id"`
	UserId   string   `json:"user_id"`
	SubTotal float64  `json:"subtotal"`
	Iva      float64  `json:"iva"`
	Total    float64  `json:"total"`
	PlatesId []string `json:"plates_id"`
}

func CreateOrderHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_3(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		req := InsertOrderRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		PlatesId := []primitive.ObjectID{}
		for _, plateId := range req.PlatesId {
			OId, err := primitive.ObjectIDFromHex(plateId)
			if err != nil {
				responses.BadRequest(w, "Invalid request body")
				return
			}
			PlatesId = append(PlatesId, OId)
		}
		createOrder := models.InsertOrder{
			TableId:  req.TableId,
			UserId:   req.UserId,
			SubTotal: req.SubTotal,
			Iva:      req.Iva,
			Total:    req.Total,
			PlatesId: PlatesId,
		}
		order, err := repository.InsertOrder(r.Context(), &createOrder)
		if err != nil {
			responses.BadRequest(w, "Error creating order")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}
}

func GetOrderByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_3(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		order, err := repository.GetOrderById(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error getting order")
			return
		}

		platesIds := []string{}
		for _, plate := range order.PlatesId {
			platesIds = append(platesIds, plate.Hex())
		}

		table, err := repository.GetTableById(r.Context(), order.TableId)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		profile, err := repository.GetUserById(r.Context(), order.UserId)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		plates, err := repository.GetPlatesByIds(r.Context(), platesIds)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		responseOrder := responses.OrderResponse{
			Id:        order.Id.Hex(),
			Table:     *table,
			User:      *profile,
			Plates:    plates,
			SubTotal:  order.SubTotal,
			Iva:       order.Iva,
			Total:     order.Total,
			TimeStamp: order.TimeStamp,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseOrder)
	}
}

func UpdateOrderHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		req := InsertOrderRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		PlatesId := []primitive.ObjectID{}
		for _, plateId := range req.PlatesId {
			OId, err := primitive.ObjectIDFromHex(plateId)
			if err != nil {
				responses.BadRequest(w, "Invalid request body")
				return
			}
			PlatesId = append(PlatesId, OId)
		}
		updateOrder := models.UpdateOrder{
			TableId:  req.TableId,
			SubTotal: req.SubTotal,
			Total:    req.Total,
			PlatesId: PlatesId,
		}
		order, err := repository.UpdateOrder(r.Context(), &updateOrder, params["id"])
		if err != nil {
			responses.BadRequest(w, "Error updating order")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}
}

func DeleteOrderHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_3(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err := repository.DeleteOrder(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting order")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListOrdersByPageHandler(s server.Server) http.HandlerFunc {
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
		orders, quantity, err := repository.ListOrdersByPage(r.Context(), limit, page)
		if err != nil {
			responses.BadRequest(w, "Error getting orders")
			return
		}

		responsesOrders := []responses.OrderResponse{}

		for _, order := range orders {
			platesIds := []string{}
			for _, plate := range order.PlatesId {
				platesIds = append(platesIds, plate.Hex())
			}

			table, err := repository.GetTableById(r.Context(), order.TableId)
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			profile, err := repository.GetUserById(r.Context(), order.UserId)
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			plates, err := repository.GetPlatesByIds(r.Context(), platesIds)
			if err != nil {
				responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			response := responses.OrderResponse{
				Id:        order.Id.Hex(),
				Table:     *table,
				User:      *profile,
				Plates:    plates,
				SubTotal:  order.SubTotal,
				Iva:       order.Iva,
				Total:     order.Total,
				TimeStamp: order.TimeStamp,
			}

			responsesOrders = append(responsesOrders, response)

		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responses.OrdersResponse{
			Order:    responsesOrders,
			Quantity: quantity,
		})
	}
}

func ListOrdersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.DatabaseConnection_3(s)
		//Handle request
		w.Header().Set("Content-Type", "application/json")
		orders, err := repository.ListOrders(r.Context())
		if err != nil {
			responses.BadRequest(w, "Error getting orders")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(orders)
	}
}
