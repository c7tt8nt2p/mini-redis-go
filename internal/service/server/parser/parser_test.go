package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNonSubscriptionCommand_ExitCmd(t *testing.T) {
	assert.Equal(t, ExitCmd, ParseNonSubscriptionCommand("exit"))
	assert.Equal(t, ExitCmd, ParseNonSubscriptionCommand("Exit"))
	assert.Equal(t, ExitCmd, ParseNonSubscriptionCommand(" ExiT "))
}

func TestParseNonSubscriptionCommand_PingCmd(t *testing.T) {
	assert.Equal(t, PingCmd, ParseNonSubscriptionCommand("ping"))
	assert.Equal(t, PingCmd, ParseNonSubscriptionCommand("Ping"))
	assert.Equal(t, PingCmd, ParseNonSubscriptionCommand(" PinG "))
}

func TestParseNonSubscriptionCommand_SetCmd(t *testing.T) {
	assert.Equal(t, SetCmd, ParseNonSubscriptionCommand("set a b"))
	assert.Equal(t, SetCmd, ParseNonSubscriptionCommand("Set a b"))
	assert.Equal(t, SetCmd, ParseNonSubscriptionCommand(" SeT a b "))
}

func TestParseNonSubscriptionCommand_GetCmd(t *testing.T) {
	assert.Equal(t, GetCmd, ParseNonSubscriptionCommand("get a"))
	assert.Equal(t, GetCmd, ParseNonSubscriptionCommand("Get a"))
	assert.Equal(t, GetCmd, ParseNonSubscriptionCommand(" GeT a "))
}

func TestParseNonSubscriptionCommand_SubscribeCmd(t *testing.T) {
	assert.Equal(t, SubscribeCmd, ParseNonSubscriptionCommand("subscribe t1"))
	assert.Equal(t, SubscribeCmd, ParseNonSubscriptionCommand("Subscribe t1"))
	assert.Equal(t, SubscribeCmd, ParseNonSubscriptionCommand(" SubscribE t1 "))
}

func TestParseNonSubscriptionCommand_OtherCmd(t *testing.T) {
	assert.Equal(t, OtherCmd, ParseNonSubscriptionCommand("Hello"))
}

func TestParseSubscriptionCommand_UnsubscribeCmd(t *testing.T) {
	assert.Equal(t, UnsubscribeCmd, ParseSubscriptionCommand("unsubscribe"))
	assert.Equal(t, UnsubscribeCmd, ParseSubscriptionCommand("Unsubscribe"))
	assert.Equal(t, UnsubscribeCmd, ParseSubscriptionCommand(" UnsubscribE "))
}

func TestParseNonSubscriptionCommand_PublishCmd(t *testing.T) {
	assert.Equal(t, PublishCmd, ParseSubscriptionCommand("Hello"))
	assert.Equal(t, PublishCmd, ParseSubscriptionCommand("set a b"))
	assert.Equal(t, PublishCmd, ParseSubscriptionCommand("get a"))
}

func TestExtractSetCmd(t *testing.T) {
	key, value := ExtractSetCmd("set a b")

	assert.Equal(t, "a", key)
	assert.Equal(t, "b", value)
}

func TestPExtractGetCmd(t *testing.T) {
	key := ExtractGetCmd("get aa")

	assert.Equal(t, "aa", key)
}

func TestExtractSubscribeCmd(t *testing.T) {
	topic := ExtractSubscribeCmd("subscribe tt")

	assert.Equal(t, "tt", topic)
}
