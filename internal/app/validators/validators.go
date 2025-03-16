package validators

import (
	"errors"
	"strconv"

	"github.com/phedde/luhn-algorithm"
)

var ErrInvalidOrder = errors.New("invalid order")

func OrderIsValid(number string) error {
	numberInt, err := strconv.ParseInt(string(number), 10, 64)
	if err != nil || !luhn.IsValid(numberInt) {
		return ErrInvalidOrder
	}
	return nil
}
