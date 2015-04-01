package helper

import (
	"os"
)

// DoesFileExist checks whether or not a file exist at the given path
// TODO: Returns an error if any, this to enable error handling if needed
func DoesFileExist(path string) (doesExist bool) {
	doesExist = true
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		doesExist = false
	}
	return doesExist
}

// FileDoesntExist a handy function for the opposite of FileDoesExist
func FileDoesntExist(path string) bool {
	return !DoesFileExist(path)
}
