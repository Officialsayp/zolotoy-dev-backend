package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/service"
)

type createOrderRequest struct {
	Product string `json:"product"`
}

type createOrderResponse struct {
	Product string `json:"product"`
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	idInt, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}
	if idInt < 1 {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	detailsStr := r.URL.Query().Get("details")
	details, err2 := strconv.ParseBool(detailsStr)
	if err2 != nil && detailsStr != "" {
		http.Error(w, "Некорректный параметр details", http.StatusBadRequest)
		return
	}

	switch details {
	case true:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "order id: %d, details: %v", idInt, details)
		return
	default:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "order id: %d, details:", idInt)
		return
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func createOrderHandler(orderService *service.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createOrderRequest
		errReq := json.NewDecoder(r.Body).Decode(&req)
		if errReq != nil {
			http.Error(w, "incorrect Body", http.StatusBadRequest)
			return
		}
		product := strings.TrimSpace(req.Product)

		if product == "" {
			http.Error(
				w,
				"product is required",
				http.StatusBadRequest,
			)
			return
		}

		err := orderService.CreateOrder(product)
		if err != nil {
			if errors.Is(
				err,
				service.ErrProductUnavailable,
			) {
				http.Error(
					w,
					"product cannot be ordered",
					http.StatusBadRequest,
				)
				return
			}

			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createOrderResponse{
			Product: product,
		})
		return
	}
}

func main() {
	orderService := &service.OrderService{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", getOrderHandler)
	mux.Handle("GET /health", http.HandlerFunc(healthHandler))
	mux.HandleFunc("POST /orders", createOrderHandler(orderService))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
