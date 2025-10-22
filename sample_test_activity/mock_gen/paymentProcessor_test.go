package mockgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestPaymentProcessor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name  string
		total float64
	}{
		{"send success", 10},
		{"send failure", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProcessor := NewMockPaymentProcessor(ctrl)
			order := &Order{Processor: mockProcessor, Total: tt.total}

			// Expectation: Send("Hello") should be called once and return nil
			mockProcessor.EXPECT().Charge(tt.total).Return(nil)

			err := order.Checkout()
			assert.NoError(t, err)
		})
	}
}
