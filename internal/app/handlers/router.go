package handlers

import (
	"github.com/go-chi/chi"
	"github.com/valinurovdenis/gomart/internal/app/auth"
	"github.com/valinurovdenis/gomart/internal/app/gzip"
	"github.com/valinurovdenis/gomart/internal/app/logger"
)

func MartRouter(handler ApiHandler, auth auth.JwtAuthenticator) chi.Router {
	r := chi.NewRouter()
	r.Use(logger.RequestLogger)
	r.Use(gzip.GzipMiddleware)

	r.Post("/api/user/register", auth.Register)
	r.Post("/api/user/login", auth.Login)

	r.Route("/", func(r chi.Router) {
		r.Use(auth.Authenticate)
		r.Post("/api/user/orders", handler.AddUserOrder)
		r.Get("/api/user/orders", handler.GetUserOrders)
		r.Get("/api/user/balance", handler.GetUserBalance)
		r.Post("/api/user/balance/withdraw", handler.WithdrawOrder)
		r.Get("/api/user/withdrawals", handler.GetWithdrawals)
	})

	return r
}
