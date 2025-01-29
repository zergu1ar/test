package openapi

import (
	"applicationDesignTest/internal/tools/logger"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *handlers) List(w http.ResponseWriter, r *http.Request) {
	if err := func() error {
		orders, err := h.OrdersRepository.List()
		if err != nil {
			return errors.Join(errors.New("failed to get orders"), err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(orders); err != nil {
			return errors.Join(errors.New("failed to encode answer"), err)
		}

		return nil
	}(); err != nil {
		logger.LogErrorf("failed to process request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
