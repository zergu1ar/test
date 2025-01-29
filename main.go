// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"applicationDesignTest/internal/gateways/openapi"
	"applicationDesignTest/internal/repository/availability"
	"applicationDesignTest/internal/repository/orders"
	"applicationDesignTest/internal/tools/logger"

	"errors"
	"net/http"
	"os"
)

func main() {
	handlers := openapi.NewHandlers(
		orders.NewRepository(&orders.Config{}),
		availability.NewRepository(&availability.Config{}),
	)
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handlers.CreateOrder)
	mux.HandleFunc("/list", handlers.List)

	logger.LogInfo("Server listening on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		logger.LogInfo("Server closed")
	} else if err != nil {
		logger.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
