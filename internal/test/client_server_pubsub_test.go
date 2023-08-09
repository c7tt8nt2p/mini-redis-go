// Integration test
package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/test/utils"
	"os"
	"testing"
)

func TestSubscribeAndPublish(t *testing.T) {
	tempFolder := utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := utils.StartServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	topic := "t1"
	client1 := utils.ConnectToServer(config.ConnectionHost, config.ConnectionPort)
	client2 := utils.ConnectToServer(config.ConnectionHost, config.ConnectionPort)

	subscriber1 := client1.Subscribe(topic)
	assert.Equal(t, "Subscribed\n", utils.Read(t, client1))

	subscriber2 := client2.Subscribe(topic)
	assert.Equal(t, "Subscribed\n", utils.Read(t, client2))

	assert.Equal(t,
		fmt.Sprintf("%s has joined us.", client2.GetConnection().LocalAddr()),
		utils.NextMessage(t, subscriber1))

	// subscriber1 publishes
	utils.Publish(t, subscriber1, "Hello I'm client1")
	assert.Equal(t, "Hello I'm client1\n", utils.NextMessage(t, subscriber2))

	// subscriber2 publishes
	utils.Publish(t, subscriber2, "Hello there")
	assert.Equal(t, "Hello there\n", utils.NextMessage(t, subscriber1))

	// subscriber1 unsubscribes
	utils.Unsubscribe(t, subscriber1)
	assert.Equal(t,
		fmt.Sprintf("%s has left.", client1.GetConnection().LocalAddr()),
		utils.NextMessage(t, subscriber2))

	s.Stop()
}
