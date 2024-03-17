package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/danielgz405/Resev/handlers"
	"github.com/danielgz405/Resev/middleware"
	"github.com/danielgz405/Resev/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URI := os.Getenv("DB_URI")
	DB_URI_2 := os.Getenv("DB_URI_2")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      ":" + PORT,
		JWTSecret: JWT_SECRET,
		DbURI:     DB_URI,
		DB_URI_2:  DB_URI_2,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/welcome", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//Auth
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	//user
	r.HandleFunc("/user/delete", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/user/update", handlers.UpdateUserHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/user/profile/{userId}", handlers.ProfileHandler(s)).Methods(http.MethodGet)

	//Table
	r.HandleFunc("/table/create", handlers.CreateTableHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/table/{id}", handlers.GetTableByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/table/updated/{id}", handlers.UpdateTableHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/table/delete/{id}", handlers.DeleteTableHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/table/list/{limit}/{page}", handlers.ListTablesByPageHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/table/list/all", handlers.ListTablesHandler(s)).Methods(http.MethodGet)

	//Role
	r.HandleFunc("/role/create", handlers.CreateRoleHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/role/{id}", handlers.GetRoleByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/role/updated/{id}", handlers.UpdateRoleHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/role/delete/{id}", handlers.DeleteRoleHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/role/list/{limit}/{page}", handlers.ListRolesByPageHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/role/list/all", handlers.ListRolesHandler(s)).Methods(http.MethodGet)

	//Plate
	r.HandleFunc("/plate/create", handlers.CreatePlateHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/plate/{id}", handlers.GetPlateByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/plate/updated/{id}", handlers.UpdatePlateHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/plate/delete/{id}", handlers.DeletePlateHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/plate/list/{limit}/{page}", handlers.ListPlatesByPageHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/plate/list/all", handlers.ListPlatesHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/plate/list/ids", handlers.GetPlatesByIdsHandler(s)).Methods(http.MethodPatch)

	//Order
	r.HandleFunc("/order/create", handlers.CreateOrderHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/order/{id}", handlers.GetOrderByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/order/updated/{id}", handlers.UpdateOrderHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/order/delete/{id}", handlers.DeleteOrderHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/order/list/{limit}/{page}", handlers.ListOrdersByPageHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/order/list/all", handlers.ListOrdersHandler(s)).Methods(http.MethodGet)

	//Booking
	r.HandleFunc("/booking/create", handlers.CreateBookingHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/booking/{id}", handlers.GetBookingByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/booking/updated/{id}", handlers.UpdateBookingHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/booking/delete/{id}", handlers.DeleteBookingHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/booking/list/{limit}/{page}", handlers.ListBookingsByPageHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/booking/list/all", handlers.ListBookingsHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/booking/list/ids", handlers.GetBookingsByIdsHandler(s)).Methods(http.MethodPatch)
}
