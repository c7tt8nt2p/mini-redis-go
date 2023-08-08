package server

import (
	"regexp"
	"strings"
)

type CmdType uint

const (
	exitCmd CmdType = iota
	pingCmd
	setCmd
	getCmd
	subscribeCmd
	otherCmd
)

const setCliRegex = `^set ([a-zA-Z0-9]+) ([a-zA-Z0-9]+)$`
const getCliRegex = `^get ([a-zA-Z0-9]+)$`
const subscribeCliRegex = `^subscribe ([a-zA-Z0-9]+)$`

func parse(s string) CmdType {
	text := strings.ToLower(strings.TrimSpace(s))

	if isExit(text) {
		return exitCmd
	} else if isPing(text) {
		return pingCmd
	} else if isSetCli(text) {
		return setCmd
	} else if isGetCli(text) {
		return getCmd
	} else if isSubscribeCli(text) {
		return subscribeCmd
	} else {
		return otherCmd
	}
}

func isExit(s string) bool {
	return s == "exit"
}

func isPing(s string) bool {
	return s == "ping"
}

func isSetCli(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(setCliRegex, message)
	return matches
}

func extractSetCli(s string) (string, string) {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(setCliRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 3 {
		return rs[1], rs[2]
	}
	return "", ""
}

func isGetCli(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(getCliRegex, message)
	return matches
}

func extractGetCli(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(getCliRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}

func isSubscribeCli(s string) bool {
	message := strings.TrimSpace(s)
	matches, _ := regexp.MatchString(subscribeCliRegex, message)
	return matches
}

func extractSubscribeCli(s string) string {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(subscribeCliRegex)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return rs[1]
	}
	return ""
}
