package cli

import (
	"fmt"
	"regexp"

	docopt "github.com/docopt/docopt-go"
)

const (
	argsRegex = `(?:(\w*)=?'(\S*)?')*`

	// UsageStr is logan's help text
	UsageStr = `A Command line tool helps to organise our scripts.

Usage:
  logan [options] <goal> [<parameter>...]
  logan -h | --help
  logan --version

Arguments
  <goal>        The goal is expressed as a composition of this 3 items '<intent>:<target>:<context>'
                - <intent>  : Define the action verb we want to perform.
                              Eg: 'create'
                - <target>  : Define the target that we have the intention to
                              operate on.
                              Eg: 'file'
                - <context> : Define the context in which the action is performed.
                              Eg: 'windows'
                Eg: create:file:windows

  <parameter>   Additional parameter we want to pass to our goal.
                You can set multiple parameters with space as a separator.
                By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
                Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the action in sudo mode.
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
	errInvalidExprCode ErrorCode = "[Parser] Invalid command expression"
	errInvalidArgsExpr ErrorCode = "[Parser] Invalid arguments expression"
)

// ErrorCode is an error code type for parsing errors
type ErrorCode string

// ParserError is an error that results when a parsing fails
type ParserError struct {
	code ErrorCode
	expr string
}

func (e *ParserError) Error() string {
	return fmt.Sprintf("%s : expr=<%v>", e.code, e.expr)
}

// Argv is a CLI arguments holder as key value pairs struct
type Argv map[string]interface{}

// Parser is a set of parsing operations
type Parser interface {
	ParseExpr(expr string) Argv
}

// ParserImp is a concrete parser
type ParserImp struct {
	regex *regexp.Regexp
}

// NewParser is the constructor
func NewParser() *ParserImp {
	return FromRegexExpr(argsRegex)
}

// FromRegexExpr return a parser using the given regex
func FromRegexExpr(expr string) *ParserImp {
	return &ParserImp{regexp.MustCompile(expr)}
}

// ParseExpr parses a given CLI expression and returns
// arg variables enter by the user and error if any
func (p *ParserImp) ParseExpr(expr string) (Argv, ParserError) {

	return nil, ParserError{errInvalidExprCode, expr}
}
