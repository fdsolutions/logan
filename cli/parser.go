package cli

import (
	"fmt"
	"regexp"
	"strings"

	docopt "github.com/docopt/docopt-go"

	"github.com/fdsolutions/logan/version"
)

const (
	argsRegexPattern = `(?:(\w*)=?'(\S*)?')*`

	// UsageStr is logan's help text
	Usage = `A Command line tool helps to organise our scripts.

Usage:
  logan [options] <command_name> [<command_arg>...]
  logan -h | --help
  logan --version

Arguments
  <command_name>  	is the name of the command expressed as a composition of this 3 items '<intent>:<target>:<context>'
                	- <intent>  : Define the action we want to perform as a verb.
                	              Eg: 'create'
                	- <target>  : Define the target that we have the intention to
                	              operate on.
                	              Eg: 'file'
                	- <context> : Define the context in which the action is performed.
                	              Eg: 'windows'
                	Eg: create:file:windows

  <command_arg>   	Additional arguments we want to pass wth the command.
                	You can add multiple arguments with space separated.
                	By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
                	Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.
`
)

// CLI options
var (
	availableOptionNames = map[string]string{
		"HELP":    "--help",
		"VERSION": "--version",
		"SUDO":    "--sudo",
	}

	availableArgsNames = map[string]string{
		"GOAL":   "<goal>",
		"PARAMS": "<parameter>",
	}
)

var (
	errInvalidCommandInputCode ErrorCode = "[Parser] Invalid command input"
)

// ErrorCode is an error code type for parsing errors
type ErrorCode string

// ParserError is an error that results when a parsing fails
type ParserError struct {
	code  ErrorCode
	input string
}

func NewError(code ErrorCode, input string) *ParserError {
	return &ParserError{code, input}
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("%s : inputCmd=<%v>", e.code, e.input)
}

// ParserImp is a concrete parser
type ParserImp struct {
	regex *regexp.Regexp
}

// NewParser is the constructor
func NewParser() *ParserImp {
	return FromRegexPattern(argsRegexPattern)
}

// FromRegexPattern return a parser using the given regex pattern
func FromRegexPattern(pattern string) *ParserImp {
	return &ParserImp{regexp.MustCompile(pattern)}
}

// Argv holds CLI arguments as key/value pairs
type Argv map[string]interface{}

// Parser is a set of parsing operations
type Parser interface {
	ParseUserInput(cmd string) Argv
}

// ParseUserInput parses user command input
func (p *ParserImp) ParseUserInput(cmd string) (Argv, error) {
	cmdParts := strings.Split(cmd, " ")

	args, err := docopt.Parse(Usage, cmdParts, true, version.LoganVersion, false)
	if err != nil {
		return nil, NewError(errInvalidCommandInputCode, cmd)
	}
	return args, nil
}
