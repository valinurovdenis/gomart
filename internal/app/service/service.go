package service

import (
	"context"
	"errors"

	"github.com/valinurovdenis/gomart/internal/app/accrualorder"
	"github.com/valinurovdenis/gomart/internal/app/currencybalance"
	"github.com/valinurovdenis/gomart/internal/app/orderstorage"
	"github.com/valinurovdenis/gomart/internal/app/userstorage"
	"github.com/valinurovdenis/gomart/internal/app/validators"
	"github.com/valinurovdenis/gomart/internal/app/withdrawstorage"
)

type OrderService struct {
	AccrualOrderService accrualorder.AccrualOrderService
	OrderServiceStorage ServiceStorage
}

func (s *OrderService) AddUserOrder(context context.Context, login string, number string) error {
	if err := validators.OrderIsValid(number); err != nil {
		return err
	}
	order, err := s.AccrualOrderService.GetOrder(context, number)
	if err != nil {
		return err
	}
	var balance currencybalance.CurrencyBalance
	balance.SetFloat(order.Accrual)
	userOrder := orderstorage.UserOrder{Login: login, Number: number, Status: order.Status, Balance: balance}
	err = s.OrderServiceStorage.AddUserOrder(context, userOrder)
	if err == nil && !orderstorage.IsFinal(userOrder.Status) {
		err = s.AccrualOrderService.EnqueueOrderUpdate(context, login, number)
	}
	return err
}

func (s *OrderService) GetUserOrders(context context.Context, login string) ([]orderstorage.UserOrder, error) {
	return s.OrderServiceStorage.GetUserOrders(context, login)
}

func (s *OrderService) GetUserBalance(context context.Context, login string) (userstorage.UserBalance, error) {
	return s.OrderServiceStorage.GetBalance(context, login)
}

var ErrNotEnoughBalance = errors.New("not enough balance for withdraw")

func (s *OrderService) AddUserWithdraw(context context.Context, withdraw withdrawstorage.UserWithdraw) error {
	if err := validators.OrderIsValid(withdraw.Number); err != nil {
		return err
	}
	userBalance, err := s.OrderServiceStorage.GetBalance(context, withdraw.Login)
	if err != nil {
		return err
	}
	if userBalance.Current.Less(withdraw.Withdraw) {
		return ErrNotEnoughBalance
	}
	err = s.OrderServiceStorage.AddUserWithdraw(context, withdraw)
	return err
}

func (s *OrderService) GetUserWithdrawals(context context.Context, login string) ([]withdrawstorage.UserWithdraw, error) {
	return s.OrderServiceStorage.GetUserWithdrawals(context, login)
}

func NewOrderService(serviceStorage ServiceStorage, accrualOrderService accrualorder.AccrualOrderService) *OrderService {

	ret := &OrderService{
		OrderServiceStorage: serviceStorage,
		AccrualOrderService: accrualOrderService,
	}
	return ret
}
