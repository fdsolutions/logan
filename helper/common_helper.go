package helper

import (
	"fmt"
	log "github.com/fdsolutions/logan/log"
)

func I(msg string, arg ...interface{}) {
	log.Info(msg, fieldify(ToS(arg)))
}

func D(msg string, arg ...interface{}) {
	log.Debug(msg, fieldify(ToS(arg)))
}

func W(msg string, arg ...interface{}) {
	log.Warn(msg, fieldify(ToS(arg)))
}

func E(msg string, arg ...interface{}) {
	log.Error(msg, fieldify(ToS(arg)))
}

func ToS(arg ...interface{}) string {
	return fmt.Sprintf("%#v", arg)
}

func fieldify(input string) map[string]string {
	return map[string]string{
		"val": input,
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
