package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/availability/memory"
	"github.com/Officialsayp/zolotoy-dev-backend/services/order-service/internal/service"
)

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
		var req createOrderV1Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "incorrect body", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(req.BuyerID) == "" {
			http.Error(w, "buyer_id is required", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.PaymentMethod) == "" {
			http.Error(w, "payment_method is required", http.StatusBadRequest)
			return
		}

		if len(req.Items) == 0 {
			http.Error(w, "items are required", http.StatusBadRequest)
			return
		}

		if strings.TrimSpace(req.DeliveryAddress) == "" {
			http.Error(w, "delivery_address is required", http.StatusBadRequest)
			return
		}

		for _, item := range req.Items {
			if strings.TrimSpace(item.ProductID) == "" {
				http.Error(w, "product_id is required", http.StatusBadRequest)
				return
			}

			if item.Quantity <= 0 {
				http.Error(w, "quantity must be greater than zero", http.StatusBadRequest)
				return
			}
		}

		input := req.toServiceInput()

		order, err := orderService.CreateOrder(input)

		if err != nil {
			if errors.Is(err, service.ErrProductUnavailable) {
				http.Error(
					w,
					"product cannot be ordered",
					http.StatusBadRequest,
				)
				return
			}

			if errors.Is(err, service.ErrInvalidOrder) {
				http.Error(
					w,
					"invalid order",
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
		response := newCreateOrderResponse(order)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			return
		}

	}

}

func main() {
	availabilityChecker := &memory.AvailabilityChecker{}
	orderService := service.NewOrderService(availabilityChecker)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", getOrderHandler)
	mux.Handle("GET /health", http.HandlerFunc(healthHandler))
	mux.HandleFunc("POST /orders", createOrderHandler(orderService))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
