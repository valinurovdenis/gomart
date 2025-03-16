package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valinurovdenis/gomart/internal/app/accrualorder"
	"github.com/valinurovdenis/gomart/internal/app/currencybalance"
	"github.com/valinurovdenis/gomart/internal/app/orderstorage"
	"github.com/valinurovdenis/gomart/internal/app/validators"
	"github.com/valinurovdenis/gomart/mocks"
)

func TestOrderService_AddUserOrder(t *testing.T) {
	ctx := context.Background()
	mockStorage := mocks.NewServiceStorage(t)
	mockService := mocks.NewAccrualOrderService(t)

	mockService.On("GetOrder", ctx, "79927398713").Return(accrualorder.AccrualOrder{}, accrualorder.ErrNoSuchOrder).Once()

	mockService.On("GetOrder", ctx, "79927398705").Return(accrualorder.AccrualOrder{}, accrualorder.ErrNoAnswer).Once()

	// invalid order
	invalidAccrualOrder := accrualorder.AccrualOrder{Order: "79927398721", Status: orderstorage.Invalid, Accrual: 5.}
	mockService.On("GetOrder", ctx, "79927398721").Return(invalidAccrualOrder, nil).Once()
	mockStorage.On("AddUserOrder", ctx, orderstorage.UserOrder{"a", "79927398721", orderstorage.Invalid, currencybalance.CurrencyBalance{500}, time.Time{}}).Return(nil).Once()

	// processed order
	processedAccrualOrder := accrualorder.AccrualOrder{Order: "79927398739", Status: orderstorage.Processed, Accrual: 5}
	mockService.On("GetOrder", ctx, "79927398739").Return(processedAccrualOrder, nil).Once()
	mockStorage.On("AddUserOrder", ctx, orderstorage.UserOrder{"a", "79927398739", orderstorage.Processed, currencybalance.CurrencyBalance{500}, time.Time{}}).Return(nil).Once()

	// processing order
	processingAccrualOrder := accrualorder.AccrualOrder{Order: "79927398747", Status: orderstorage.Processing, Accrual: 5}
	mockService.On("GetOrder", ctx, "79927398747").Return(processingAccrualOrder, nil).Once()
	mockStorage.On("AddUserOrder", ctx, orderstorage.UserOrder{"a", "79927398747", orderstorage.Processing, currencybalance.CurrencyBalance{500}, time.Time{}}).Return(nil).Once()
	mockService.On("EnqueueOrderUpdate", ctx, "a", "79927398747").Return(nil).Once()

	service := NewOrderService(mockStorage, mockService)
	tests := []struct {
		name   string
		login  string
		number string
		err    error
	}{
		{name: "invalid number", login: "a", number: "79927398712", err: validators.ErrInvalidOrder},
		{name: "no such order from accrual", login: "a", number: "79927398713", err: accrualorder.ErrNoSuchOrder},
		{name: "no answer from accrual", login: "a", number: "79927398705", err: accrualorder.ErrNoAnswer},
		{name: "invalid order from accrual", login: "a", number: "79927398721", err: nil},
		{name: "processed order from accrual", login: "a", number: "79927398739", err: nil},
		{name: "processing order from accrual", login: "a", number: "79927398747", err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.AddUserOrder(ctx, tt.login, tt.number)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err, "Ошибка не совпадает")
			}
		})
	}
}
