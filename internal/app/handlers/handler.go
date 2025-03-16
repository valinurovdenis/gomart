package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/valinurovdenis/gomart/internal/app/accrualorder"
	"github.com/valinurovdenis/gomart/internal/app/orderstorage"
	"github.com/valinurovdenis/gomart/internal/app/service"
	"github.com/valinurovdenis/gomart/internal/app/validators"
	"github.com/valinurovdenis/gomart/internal/app/withdrawstorage"
)

type ApiHandler struct {
	Service service.OrderService
}

func (h *ApiHandler) AddUserOrder(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	number, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.AddUserOrder(r.Context(), login, string(number))

	if errors.Is(err, orderstorage.ErrOrderExists) {
		http.Error(w, err.Error(), http.StatusConflict)
	} else if errors.Is(err, orderstorage.ErrAlreadySent) {
		w.WriteHeader(http.StatusOK)
	} else if errors.Is(err, validators.ErrInvalidOrder) || errors.Is(err, accrualorder.ErrNoSuchOrder) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}
}

func (h *ApiHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	orders, err := h.Service.GetUserOrders(r.Context(), login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *ApiHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	userBalance, err := h.Service.GetUserBalance(r.Context(), login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userBalance)
}

func (h *ApiHandler) WithdrawOrder(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	var withdraw withdrawstorage.UserWithdraw
	if err := json.NewDecoder(r.Body).Decode(&withdraw); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	withdraw.Login = login

	err := h.Service.AddUserWithdraw(r.Context(), withdraw)

	if errors.Is(err, service.ErrNotEnoughBalance) {
		http.Error(w, err.Error(), http.StatusPaymentRequired)
		return
	} else if errors.Is(err, validators.ErrInvalidOrder) {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ApiHandler) GetWithdrawals(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	withdrawals, err := h.Service.GetUserWithdrawals(r.Context(), login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(withdrawals)
}

func NewApiHandler(service service.OrderService) *ApiHandler {
	return &ApiHandler{Service: service}
}
