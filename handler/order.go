package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Abiji-2020/go-curd/model"
	"github.com/Abiji-2020/go-curd/repository/order"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Order struct {
	Repo *order.RedisRepo
}

// Utility function for centralized error handling
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Order) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerId uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	now := time.Now().UTC()
	order := model.Order{
		OrderId:    rand.Uint64(),
		CustomerId: body.CustomerId,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	if err := h.Repo.Insert(r.Context(), order); err != nil {
		log.Printf("Failed to insert order: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (h *Order) List(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}
	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid cursor value")
		return
	}

	const size = 50
	res, err := h.Repo.FindAll(r.Context(), order.FindAllPage{
		Offset: uint(cursor),
		Size:   size,
	})
	if err != nil {
		log.Printf("Failed to find orders: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	response := struct {
		Items []model.Order `json:"items"`
		Next  uint64        `json:"next,omitempty"`
	}{
		Items: res.Orders,
		Next:  res.Cursor,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write response: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func (h *Order) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	o, err := h.Repo.FindById(r.Context(), orderId)
	if errors.Is(err, order.ErrNotExist) {
		writeError(w, http.StatusNotFound, "Order not found")
		return
	} else if err != nil {
		log.Printf("Failed to find order by ID: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to retrieve order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Printf("Failed to write response: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func (h *Order) UpdateBYID(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	re := regexp.MustCompile(`/orders/(\d+)`) // Matches numbers after /orders/
	match := re.FindStringSubmatch(r.URL.Path)
	idParam := match[1]

	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Invalid order ID %v", err))
		return
	}

	theOrder, err := h.Repo.FindById(r.Context(), orderId)
	if errors.Is(err, order.ErrNotExist) {
		writeError(w, http.StatusNotFound, "Order not found")
		return
	} else if err != nil {
		log.Printf("Failed to find order by ID: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to retrieve order")
		return
	}

	const completedStatus = "completed"
	const shippedStatus = "shipped"
	now := time.Now().UTC()
	switch body.Status {
	case shippedStatus:
		if theOrder.ShippedAt != nil {
			writeError(w, http.StatusBadRequest, "Order already shipped")
			return
		}
		theOrder.ShippedAt = &now
	case completedStatus:
		if theOrder.CompletedAt != nil || theOrder.ShippedAt == nil {
			writeError(w, http.StatusBadRequest, "Order cannot be completed")
			return
		}
		theOrder.CompletedAt = &now
	default:
		writeError(w, http.StatusBadRequest, "Invalid status value")
		return
	}

	if err := h.Repo.Update(r.Context(), theOrder); err != nil {
		log.Printf("Failed to update order: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to update order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(theOrder); err != nil {
		log.Printf("Failed to write response: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
func (h *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	const base = 10
	const bitSize = 64
	orderId, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	if err := h.Repo.DeleteById(r.Context(), orderId); errors.Is(err, order.ErrNotExist) {
		writeError(w, http.StatusNotFound, "Order not found")
		return
	} else if err != nil {
		log.Printf("Failed to delete order: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to delete order")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"message": "Order deleted successfully"}`)); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
