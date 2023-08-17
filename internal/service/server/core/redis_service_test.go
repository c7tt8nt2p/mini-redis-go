package core

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"mini-redis-go/internal/mock"
	"testing"
)

func NewRedisServiceTestInstance(cacheReader *mock.MockCacheReaderService, cacheWriter *mock.MockCacheWriterService) *redisService {
	return &redisService{
		cacheReaderService: cacheReader,
		cacheWriterService: cacheWriter,
		db: &MyDb{
			cache: map[string][]byte{},
		},
	}
}

func TestNewRedisService(t *testing.T) {
	service := NewRedisService()

	assert.NotNil(t, service)
}

func TestRedisService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	key := "k1"
	value := []byte("v1")

	service := NewRedisServiceTestInstance(mock.NewMockCacheReaderService(ctrl), mock.NewMockCacheWriterService(ctrl))
	service.db.cache[key] = value

	assert.Equal(t, value, service.Get(key))
}

func TestRedisService_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	key := "k2"
	value := []byte("v2")

	service := NewRedisServiceTestInstance(mock.NewMockCacheReaderService(ctrl), mock.NewMockCacheWriterService(ctrl))
	service.Set(key, value)

	assert.Equal(t, value, service.db.cache[key])
}

func TestRedisService_ExistsByKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	key := "k3"
	value := []byte("v3")

	service := NewRedisServiceTestInstance(mock.NewMockCacheReaderService(ctrl), mock.NewMockCacheWriterService(ctrl))
	service.Set(key, value)

	assert.True(t, service.ExistsByKey(key))
}

func TestRedisService_ReadCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	cacheReader := mock.NewMockCacheReaderService(ctrl)
	service := NewRedisServiceTestInstance(cacheReader, mock.NewMockCacheWriterService(ctrl))
	tempFolder := "/temp"
	cacheReader.EXPECT().ReadFromFile(tempFolder).Return(map[string][]byte{"cache1": []byte("cacheValue1")})

	service.ReadCache(tempFolder)

	assert.Equal(t, 1, len(service.db.cache))
	assert.Equal(t, []byte("cacheValue1"), service.db.cache["cache1"])
}

func TestRedisService_WriteCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	cacheWriter := mock.NewMockCacheWriterService(ctrl)
	service := NewRedisServiceTestInstance(mock.NewMockCacheReaderService(ctrl), cacheWriter)
	tempFolder := "/temp"
	key := "cache1"
	value := []byte("cacheValue1")

	cacheWriter.EXPECT().WriteToFile(tempFolder, key, value).Times(1).Return(nil)

	response := service.WriteCache(tempFolder, key, value)

	assert.Nil(t, response)
}
