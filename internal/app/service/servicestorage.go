package service

import (
	"context"

	"github.com/valinurovdenis/gomart/internal/app/orderstorage"
	"github.com/valinurovdenis/gomart/internal/app/userstorage"
	"github.com/valinurovdenis/gomart/internal/app/withdrawstorage"
)

//go:generate mockery --name ServiceStorage
type ServiceStorage struct {
	UserBalanceStorage userstorage.BalanceStorage
	WithdrawStorage    withdrawstorage.WithdrawRepository
	OrderStorage       orderstorage.OrderStorage
}

func (s *ServiceStorage) GetUserOrders(context context.Context, login string) ([]orderstorage.UserOrder, error) {
	return s.OrderStorage.GetUserOrders(context, login)
}

func (s *ServiceStorage) AddUserOrder(context context.Context, order orderstorage.UserOrder) error {
	return s.OrderStorage.AddUserOrder(context, order)
}

func (s *ServiceStorage) GetBalance(context context.Context, login string) (userstorage.UserBalance, error) {
	return s.UserBalanceStorage.GetBalance(context, login)
}

func (s *ServiceStorage) AddUserWithdraw(ctx context.Context, order withdrawstorage.UserWithdraw) error {
	return s.WithdrawStorage.AddUserWithdraw(ctx, order)
}

func (s *ServiceStorage) GetUserWithdrawals(ctx context.Context, login string) ([]withdrawstorage.UserWithdraw, error) {
	return s.WithdrawStorage.GetUserWithdrawals(ctx, login)
}

func NewServiceStorage(userBalanceStorage userstorage.BalanceStorage,
	withdrawStorage withdrawstorage.WithdrawRepository,
	orderStorage orderstorage.OrderStorage) *ServiceStorage {

	ret := &ServiceStorage{
		UserBalanceStorage: userBalanceStorage,
		WithdrawStorage:    withdrawStorage,
		OrderStorage:       orderStorage,
	}
	return ret
}
