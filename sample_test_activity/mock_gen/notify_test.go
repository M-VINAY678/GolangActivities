package mockgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUserServiceNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"send success", "Hello", "Hello"},
		{"send failure", "Hello", "Hello!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockNotifier := NewMockNotifier(ctrl)
			service := &UserService{Notifier: mockNotifier}

			// Expectation: Send("Hello") should be called once and return nil
			mockNotifier.EXPECT().Send(tt.expected).Return(nil).Times(1) //Default is once if not specified for no of times.

			err := service.Notify(tt.input)
			assert.NoError(t, err)
		})
	}

}
