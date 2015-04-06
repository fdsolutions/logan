package metadata

import (
	"strings"
)

// SplitNameInGoalParts returns parts composing the command name
// A command name is made of 3 parts:  <intent>:<target>:<context>
// It fills the missing parts with empty value
func SplitInGoalParts(name string) (intent string, target string, ctx string) {
	parts := strings.SplitN(name, goalPartsSeparator, numberOfGoalParts)
	switch len(parts) {
	case 1: // We have at least the intent
		intent, target, ctx = parts[0], noPartValue, noPartValue
	case 2:
		intent, target, ctx = parts[0], parts[1], noPartValue
	default:
		intent, target, ctx = parts[0], parts[1], parts[2]
	}
	return
}
