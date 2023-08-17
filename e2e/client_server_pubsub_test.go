// Integration integration_test
package e2e_test

import (
	"github.com/stretchr/testify/assert"
	"mini-redis-go/e2e/internal"
	"mini-redis-go/internal/test_utils"
	"os"
	"testing"
)

func TestSubscribeAndPublish(t *testing.T) {
	tempFolder := test_utils.CreateTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := internal.StartServer(tempFolder)
	defer s.Stop()

	topic := "t1"
	client1 := internal.ConnectToServer(internal.GetClientConfigTest())
	client2 := internal.ConnectToServer(internal.GetClientConfigTest())

	subscriber1 := client1.Subscribe(topic)
	assert.Equal(t, "Subscribed\n", internal.Read(t, client1))

	subscriber2 := client2.Subscribe(topic)
	assert.Equal(t, "Subscribed\n", internal.Read(t, client2))

	// subscriber1 publishes
	internal.Publish(t, subscriber1, "Hello I'm client1")
	assert.Equal(t, "Hello I'm client1\n", internal.NextMessage(t, subscriber2))

	// subscriber2 publishes
	internal.Publish(t, subscriber2, "Hello there")
	assert.Equal(t, "Hello there\n", internal.NextMessage(t, subscriber1))

	// subscriber1 unsubscribes
	internal.Unsubscribe(t, subscriber1)
}
