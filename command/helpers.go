package command

// GetCmdNameParts returns the tuple (intent, target, context) from
// a given command name.
func GetCmdNameParts(name string) (intent string, target string, ctx string) {
	parts := strings.SplitN(name, commandNamePartSeparator, numberOfCommandNameParts)
	switch len(parts) {
	case 1: // We have at least the intent
		intent, target, ctx = parts[0], "", ""
	case 2:
		intent, target, ctx = parts[0], parts[1], ""
	default:
		intent, target, ctx = parts[0], parts[1], parts[2]
	}
	return
}
