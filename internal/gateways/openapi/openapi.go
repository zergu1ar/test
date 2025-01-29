package openapi

import (
	"applicationDesignTest/internal/repository/availability"
	"applicationDesignTest/internal/repository/orders"

	"net/http"
)

type Handlers interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	OrdersRepository       orders.Repository
	AvailabilityRepository availability.Repository
}

func NewHandlers(orders orders.Repository, availability availability.Repository) Handlers {
	return &handlers{
		OrdersRepository:       orders,
		AvailabilityRepository: availability,
	}
}
