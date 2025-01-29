package openapi

import (
	"applicationDesignTest/internal/entity"
	"applicationDesignTest/internal/helpers"
	"applicationDesignTest/internal/tools/logger"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (h *handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if err := func() error {
		var newOrder entity.Order
		if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
			return errors.Join(errors.New("failed to decode request"), err)
		}

		available, err := h.AvailabilityRepository.GetAvailable()
		if err != nil {
			return err
		}

		availableMap := available.ToDaysMap()
		for _, dayToBook := range helpers.DaysBetween(newOrder.From, newOrder.To) {
			if availableDay, ok := availableMap[dayToBook.Format(time.DateOnly)]; ok {
				if availableDay.Quota < 1 {
					http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
					return fmt.Errorf("Hotel room is not available for selected date:\n%v\n%v", newOrder, availableDay.Date)
				}
				availableDay.Quota--
				if err := h.AvailabilityRepository.Save(&availableDay); err != nil {
					return errors.Join(errors.New("failed to save availability"), err)
				}
			}
		}

		if err := h.OrdersRepository.Create(&newOrder); err != nil {
			return errors.Join(errors.New("failed to create order"), err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newOrder); err != nil {
			return errors.Join(errors.New("failed to encode answer"), err)
		}

		logger.LogInfo("Order successfully created: %v", newOrder)
		return nil
	}(); err != nil {
		logger.LogErrorf("failed to process request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
