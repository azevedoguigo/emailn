package internalerros

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("internal server error")
var ErrNotFound error = errors.New("not found")

func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}

	return err
}
