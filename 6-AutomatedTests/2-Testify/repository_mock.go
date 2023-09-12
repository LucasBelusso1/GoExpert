package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (mock TaxRepositoryMock) SaveTax(tax float64) error {
	args := mock.Called(tax)
	return args.Error(0)
}
