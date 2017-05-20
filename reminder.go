package main

import "strings"

const REMIND_ME = "remind me to"
const ALARMING_TIME_PREFIX = "on"

func remindMe(msg string) string {
	var content string
	time := ""
	response := "Will remind you to "
	if len(REMIND_ME) > len(msg) {
		return ""
	}
	afterPrefix := strings.TrimPrefix(msg, REMIND_ME+" ")

	// TODO: Handle M"on"day
	alarmingSuffixIndex := strings.LastIndex(afterPrefix, ALARMING_TIME_PREFIX)

	if alarmingSuffixIndex > -1 {
		content = afterPrefix[:alarmingSuffixIndex-1]
		time = extractTime(afterPrefix[alarmingSuffixIndex+len(ALARMING_TIME_PREFIX):])
		// TODO: Create reminder on specified time
		response = response + content + ": on " + time
	} else {
		// TODO: Ask for the time
		content = afterPrefix
		response = response + content
	}

	return response
}

func extractTime(msg string) string {
	return strings.Trim(msg, " ")
}
