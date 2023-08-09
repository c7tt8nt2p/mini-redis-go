package server

import (
	"regexp"
	"strings"
)

type NonSubscriptionCmdType uint

const (
	exitCmd NonSubscriptionCmdType = iota
	pingCmd
	setCmd
	getCmd
	subscribeCmd
	otherCmd
)

type SubscriptionCmdType uint

const (
	unsubscribeCmd SubscriptionCmdType = iota
	publishCmd
)

const unsubscribeCmdRegex = `^unsubscribe$`
const setCmdRegex = `^set ([a-zA-Z0-9]+) ([a-zA-Z0-9]+)$`
const getCmdRegex = `^get ([a-zA-Z0-9]+)$`
const subscribeCmdRegex = `^subscribe ([a-zA-Z0-9]+)$`

func parseNonSubscriptionCommand(s string) NonSubscriptionCmdType {
	text := strings.ToLower(strings.TrimSpace(s))

	if isExit(text) {
		return exitCmd
	} else if isPing(text) {
		return pingCmd
	} else if isSetCmd(text) {
		return setCmd
	} else if isGetCmd(text) {
		return getCmd
	} else if isSubscribeCmd(text) {
		return subscribeCmd
	} else {
		return otherCmd
	}
}

func parseSubscriptionCommand(s string) SubscriptionCmdType {
	text := strings.ToLower(strings.TrimSpace(s))

	if isUnsubscribeCmd(text) {
		return unsubscribeCmd
	} else {
		return publishCmd
	}
}

func isExit(s string) bool {
	return s == "exit"
}

func isUnsubscribeCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(unsubscribeCmdRegex, message)
	return matches
}

func isPing(s string) bool {
	return strings.TrimSpace(s) == "ping"
}

func isSetCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(setCmdRegex, message)
	return matches
}

func extractSetCmd(s string) (string, string) {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(setCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 3 {
		return rs[1], rs[2]
	}
	return "", ""
}

func isGetCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(getCmdRegex, message)
	return matches
}

func extractGetCmd(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(getCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}

func isSubscribeCmd(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(subscribeCmdRegex, message)
	return matches
}

func extractSubscribeCmd(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(subscribeCmdRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}
