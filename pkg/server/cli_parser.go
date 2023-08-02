package server

import (
	"regexp"
	"strings"
)

func isExit(s string) bool {
	message := strings.ToLower(strings.TrimSpace(s))
	return message == "exit"
}

func isPing(s string) bool {
	message := strings.ToLower(strings.TrimSpace(s))
	return message == "ping"
}

func isSetCli(s string) (bool, string, string) {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(`^SET ([a-zA-Z0-9]+) ([a-zA-Z0-9]+)$`)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 3 {
		return true, rs[1], rs[2]
	}
	return false, "", ""
}

func isGetCli(s string) (bool, string) {
	message := strings.TrimSpace(s)
	rgx := regexp.MustCompile(`^GET ([a-zA-Z0-9]+)$`)
	rs := rgx.FindStringSubmatch(message)

	if len(rs) == 2 {
		return true, rs[1]
	}
	return false, ""
}
