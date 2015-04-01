package helper

import (
	log "github.com/fdsolutions/logan/log"
)

func I(msg string, input string) {
	log.Info(msg, fieldify(input))
}

func D(msg string, input string) {
	log.Debug(msg, fieldify(input))
}

func W(msg string, input string) {
	log.Warn(msg, fieldify(input))
}

func E(msg string, input string) {
	log.Error(msg, fieldify(input))
}

func fieldify(input string) map[string]string {
	return map[string]string{
		"inputCmd": input,
	}
}

// ArrayToMap is an helper function that tranform an array of array to a map
func ArrayToMap(arr [][]string) map[string]string {
	m := make(map[string]string)
	for _, innerArr := range arr {
		// Build the param map
		key := innerArr[1]
		val := innerArr[2]
		m[key] = val
	}
	return m
}
