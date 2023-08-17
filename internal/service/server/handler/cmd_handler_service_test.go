package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mini-redis-go/internal/mock"
	"testing"
)

func NewCmdHandlerServiceInstance(redisService *mock.MockIRedis, brokerService *mock.MockIBroker) *CmdHandlerService {
	return &CmdHandlerService{
		redisService:  redisService,
		brokerService: brokerService,
	}
}

func TestCmdHandlerService_NewBrokerService(t *testing.T) {
	service := NewCmdHandlerService()

	assert.NotNil(t, service)
}

func TestCmdHandlerService_PingCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Write([]byte("PONG\n")).Times(1).Return(0, nil)

	service := NewCmdHandlerServiceInstance(mock.NewMockIRedis(ctrl), mock.NewMockIBroker(ctrl))
	err := service.PingCmdHandler(conn)

	assert.Nil(t, err)
}

func TestCmdHandlerService_SetCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	cacheFolder := "/tempFolder"
	message := "set a b"

	redisService := mock.NewMockIRedis(ctrl)
	redisService.EXPECT().WriteCache(cacheFolder, "a", []byte{1, 98}).Times(1)
	redisService.EXPECT().Set("a", []byte{98}).Times(1)

	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Write([]byte("Set ok"+"\n")).Times(1).Return(0, nil)

	service := NewCmdHandlerServiceInstance(redisService, mock.NewMockIBroker(ctrl))
	response := service.SetCmdHandler(conn, cacheFolder, message)

	assert.Nil(t, response)
}

func TestCmdHandlerService_SetCmdHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	cacheFolder := "/tempFolder"
	message := "set a b"
	err := errors.New("mock error")

	redisService := mock.NewMockIRedis(ctrl)
	redisService.EXPECT().WriteCache(cacheFolder, "a", []byte{1, 98}).Times(1).Return(err)

	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Write([]byte("Set failed"+"\n")).Times(1).Return(0, nil)

	service := NewCmdHandlerServiceInstance(redisService, mock.NewMockIBroker(ctrl))
	response := service.SetCmdHandler(conn, cacheFolder, message)

	assert.Equal(t, err, response)
}

func TestCmdHandlerService_GetCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	message := "get a"

	redisService := mock.NewMockIRedis(ctrl)
	redisService.EXPECT().Get("a").Times(1).Return([]byte("b"))

	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Write(append([]byte("b"), []byte("\n")...)).Times(1).Return(0, nil)

	service := NewCmdHandlerServiceInstance(redisService, mock.NewMockIBroker(ctrl))
	response := service.GetCmdHandler(conn, message)

	assert.Nil(t, response)
}

func TestCmdHandlerService_SubscribeCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	message := "subscribe t1"

	conn := mock.NewMockConn(ctrl)
	conn.EXPECT().Write([]byte("Subscribed\n")).Times(1).Return(0, nil)

	brokerService := mock.NewMockIBroker(ctrl)
	brokerService.EXPECT().Subscribe(conn, "t1").Times(1)

	service := NewCmdHandlerServiceInstance(mock.NewMockIRedis(ctrl), brokerService)
	response := service.SubscribeCmdHandler(conn, message)

	assert.Nil(t, response)
}

func TestCmdHandlerService_OtherCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	message := "hello there"
	conn.EXPECT().Write([]byte(message)).Times(1).Return(0, nil)

	service := NewCmdHandlerServiceInstance(mock.NewMockIRedis(ctrl), mock.NewMockIBroker(ctrl))
	response := service.OtherCmdHandler(conn, message)

	assert.Nil(t, response)
}

func TestCmdHandlerService_UnsubscribeCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	topic := "t2"

	brokerService := mock.NewMockIBroker(ctrl)
	brokerService.EXPECT().GetTopicFromConnection(conn).Times(1).Return(topic, true)
	brokerService.EXPECT().Unsubscribe(conn).Times(1)

	service := NewCmdHandlerServiceInstance(mock.NewMockIRedis(ctrl), brokerService)
	service.UnsubscribeCmdHandler(conn)
}

func TestCmdHandlerService_PublishCmdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := mock.NewMockConn(ctrl)
	topic := "t3"
	message := "hello everyone"

	brokerService := mock.NewMockIBroker(ctrl)
	brokerService.EXPECT().GetTopicFromConnection(conn).Times(1).Return(topic, true)
	brokerService.EXPECT().Publish(conn, topic, message).Times(1)

	service := NewCmdHandlerServiceInstance(mock.NewMockIRedis(ctrl), brokerService)
	service.PublishCmdHandler(conn, message)
}
