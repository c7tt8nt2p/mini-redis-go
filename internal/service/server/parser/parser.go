// Package parser is a utils to parse a String command and translates ot extracts it as needed
package parser

import (
	"regexp"
	"strings"
)

type NonSubscriptionCmdType uint

const (
	ExitCmd NonSubscriptionCmdType = iota
	PingCmd
	SetCmd
	GetCmd
	SubscribeCmd
	OtherCmd
)

type SubscriptionCmdType uint

const (
	UnsubscribeCmd SubscriptionCmdType = iota
	PublishCmd
)

const (
	UnsubscribeCmdRegex string = `^unsubscribe$`
	SetCmdRegex         string = `^set ([a-zA-Z0-9]+) ([a-zA-Z0-9]+)$`
	GetCmdRegex         string = `^get ([a-zA-Z0-9]+)$`
	SubscribeCmdRegex   string = `^subscribe ([a-zA-Z0-9]+)$`
)

func ParseNonSubscriptionCommand(s string) NonSubscriptionCmdType {
	text := strings.ToLower(strings.TrimSpace(s))

	if IsExit(text) {
		return ExitCmd
	} else if IsPing(text) {
		return PingCmd
	} else if IsSetCmd(text) {
		return SetCmd
	} else if IsGetCmd(text) {
		return GetCmd
	} else if IsSubscribeCmd(text) {
		return SubscribeCmd
	} else {
		return OtherCmd
	}
}

func ParseSubscriptionCommand(s string) SubscriptionCmdType {
	text := strings.ToLower(strings.TrimSpace(s))

	if IsUnsubscribeCmd(text) {
		return UnsubscribeCmd
	} else {
		return PublishCmd
	}
}

func IsExit(s string) bool {
	return s == "exit"
}

func IsUnsubscribeCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(UnsubscribeCmdRegex, message)
	return matches
}

func IsPing(s string) bool {
	return strings.TrimSpace(s) == "ping"
}

func IsSetCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(SetCmdRegex, message)
	return matches
}

func ExtractSetCmd(s string) (string, string) {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(SetCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 3 {
		return rs[1], rs[2]
	}
	return "", ""
}

func IsGetCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(GetCmdRegex, message)
	return matches
}

func ExtractGetCmd(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(GetCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}

func IsSubscribeCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(SubscribeCmdRegex, message)
	return matches
}

func ExtractSubscribeCmd(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(SubscribeCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}
