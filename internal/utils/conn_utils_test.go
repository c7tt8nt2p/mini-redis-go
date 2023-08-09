package utils

import (
	"errors"
	"go.uber.org/mock/gomock"
	"mini-redis-go/internal/mock"
	"mini-redis-go/internal/test_utils"
	"testing"
)

func TestWriteToServerShouldBeWritten(t *testing.T) {
	ctrl := gomock.NewController(t)
	msg := "test_utils message"
	writer := mock.NewMockWriter(ctrl)
	writer.EXPECT().Write([]byte(msg)).Times(1)

	WriteToServer(writer, msg)
}

func TestWriteToServerShouldPanic(t *testing.T) {
	ctrl := gomock.NewController(t)
	msg := "test_utils message"
	writer := mock.NewMockWriter(ctrl)
	writer.EXPECT().Write([]byte(msg)).Return(0, errors.New("error test_utils")).Times(1)

	test_utils.ShouldPanicWithError(t, func() { WriteToServer(writer, msg) }, "Error sending message to server: error test_utils")
}
