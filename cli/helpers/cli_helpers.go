package helpers

import (
	log "github.com/fdsolutions/logan/log"
)

func fieldify(input string) map[string]string {
	return map[string]string{
		"inputCmd": input,
	}
}

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
