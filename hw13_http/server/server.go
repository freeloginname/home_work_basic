package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/freeloginname/home_work_basic/hw13_http/internal"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := internal.User{
		ID:    1,
		Name:  "John Doe",
		Email: "xk0e5@example.com",
	}

	json.NewEncoder(w).Encode(user)

	log.Printf("Handled GET request for /v1/get_user")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser internal.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)

	log.Printf("Handled POST request for /v1/create_user with user: %+v", newUser)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	order := internal.Order{
		ID:          1,
		UserID:      "1",
		OrderDate:   "2022-01-01",
		TotalAmount: 100.0,
	}

	json.NewEncoder(w).Encode(order)

	log.Printf("Handled GET request for /v1/get_order")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newOrder internal.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrder)

	log.Printf("Handled POST request for /v1/create_order with order: %+v", newOrder)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	product := internal.Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10.0,
	}

	json.NewEncoder(w).Encode(product)

	log.Printf("Handled GET request for /v1/get_product")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newProduct internal.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)

	log.Printf("Handled POST request for /v1/create_product with product: %+v", newProduct)
}

func Server(ip *string, port *string) {
	// для получения через флаги
	// ip := flag.String("ip", "127.0.0.1", "IP address")
	// port := flag.String("port", "8080", "Port number")
	// flag.Parse()

	fmt.Printf("Starting server on %s:%s\n", *ip, *port)
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/v1/get_user", GetUser)
	serverMux.HandleFunc("/v1/create_user", CreateUser)
	serverMux.HandleFunc("/v1/get_order", GetOrder)
	serverMux.HandleFunc("/v1/create_order", CreateOrder)
	serverMux.HandleFunc("/v1/get_product", GetProduct)
	serverMux.HandleFunc("/v1/create_product", CreateProduct)

	server := &http.Server{
		Addr:         net.JoinHostPort(*ip, *port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      serverMux,
	}

	server.ListenAndServe()
}
