package main

import (
	"database/sql"
	"errors"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/valinurovdenis/gomart/internal/app/accrualorder"
	"github.com/valinurovdenis/gomart/internal/app/auth"
	"github.com/valinurovdenis/gomart/internal/app/handlers"
	"github.com/valinurovdenis/gomart/internal/app/logger"
	"github.com/valinurovdenis/gomart/internal/app/orderstorage"
	"github.com/valinurovdenis/gomart/internal/app/service"
	"github.com/valinurovdenis/gomart/internal/app/userstorage"
	"github.com/valinurovdenis/gomart/internal/app/withdrawstorage"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	config := new(Config)
	parseFlags(config)
	config.updateFromEnv()

	if err := logger.Initialize(config.LogLevel); err != nil {
		return err
	}

	if config.DatabaseURI == "" {
		return errors.New("empty database config")
	}
	db, err := sql.Open("pgx", config.DatabaseURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userStorage := userstorage.NewDatabaseUserStorage(db)
	withdrawStorage := withdrawstorage.NewDatabaseWithdrawStorage(db)
	orderStorage := orderstorage.NewDatabaseOrderStorage(db)
	accrualSettings := accrualorder.AccrualServiceSettings{URL: config.AccrualSystemAddress, Timeout: config.AccrualTimeout, Delay: config.AccrualDelay, Retries: config.AccrualRetries}
	accrualOrderService := accrualorder.NewAccrualOrderQueue(db, 10, accrualSettings, orderStorage)
	auth := auth.NewAuthenticator(config.SecretKey, userStorage)
	serviceStorage := service.NewServiceStorage(userStorage, withdrawStorage, orderStorage)
	service := service.NewOrderService(serviceStorage, accrualOrderService)
	handler := handlers.NewApiHandler(*service)

	return http.ListenAndServe(config.RunAddress, handlers.MartRouter(*handler, *auth))
}
