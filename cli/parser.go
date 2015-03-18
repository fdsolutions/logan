package cli

import (
	"fmt"
	"regexp"
	"strings"

	docopt "github.com/docopt/docopt-go"

	helper "github.com/fdsolutions/logan/cli/helpers"
	"github.com/fdsolutions/logan/version"
)

const (
	paramsRegexPattern = `(?:(\w*)='?([^']*)'?)*`

	// Usage is logan's help text
	Usage = `A Command line tool helps to organise our scripts.

Usage:
  logan [options] <name> [<param>...]
  logan -h | --help
  logan --version

Arguments
  <name>  	The name of the command expressed as a composition of this 3 items '<intent>:<target>:<context>'
        	- <intent>  : Define the action we want to perform as a verb.
        	              Eg: 'create'
        	- <target>  : Define the target that we have the intention to
        	              operate on.
        	              Eg: 'file'
        	- <context> : Define the context in which the action is performed.
        	              Eg: 'windows'
        		Eg: create:file:windows

  <param>   Additional parameters we want to pass wth the command.
             You can add multiple parameters with space separated.
             By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
             	Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.
`

	commandNameKey   = "<name>"
	commandParamsKey = "<param>"
)

// CLI options
var (
	availableFlagNames = map[string]string{
		"HELP":    "--help",
		"VERSION": "--version",
		"SUDO":    "--sudo",
	}
)

var (
	errInvalidCommandInputCode ErrorCode = "Invalid command input"
	errInvalidParamsCode       ErrorCode = "Invalid params - No matched chunks found!"
)

// ErrorCode is an error code type for parsing errors
type ErrorCode string

// ParserError is an error that results when a parsing fails
type ParserError struct {
	code  ErrorCode
	input string
}

// NewError create parser error for a specific code
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

// NewParser creates a parser using the default params regex pattern
// To create one with your own regex pattern, use func FromParamsRegexPattern(...)
func NewParser() *ParserImp {
	return FromParamsRegexPattern(paramsRegexPattern)
}

// FromParamsRegexPattern returns a parser using the given regex pattern
func FromParamsRegexPattern(pattern string) *ParserImp {
	return &ParserImp{regexp.MustCompile(pattern)}
}

// Argv holds CLI argument information
type Argv struct {
	Name   string
	Flags  map[string]bool
	Params map[string]string
}

// Parser is a set of parsing operations
type Parser interface {
	ParseUserInput(cmd string) Argv
}

// ParseUserInput parses user command input
func (p *ParserImp) ParseUserInput(cmd string) (argv Argv, e error) {
	cmdParts := strings.Split(cmd, " ")

	userInputArgs, err := docopt.Parse(Usage, cmdParts, true, version.LoganVersion, false)
	if err != nil {
		e = NewError(errInvalidCommandInputCode, cmd)
		return
	}
	return p.extractArgvFromArgs(userInputArgs), nil
}

func (p *ParserImp) extractArgvFromArgs(args map[string]interface{}) Argv {
	name := p.extractNameFromArgs(args)
	flags := p.extractFlagsFromArgs(args)
	params, _ := p.extractParamsFromArgs(args) // No error handling if fails

	return Argv{name, flags, params}
}

func (p *ParserImp) extractNameFromArgs(args map[string]interface{}) string {
	name, _ := args[commandNameKey]
	return name.(string)
}

func (p *ParserImp) extractFlagsFromArgs(args map[string]interface{}) (flags map[string]bool) {
	flags = map[string]bool{}
	for _, flagName := range availableFlagNames {
		flags[flagName] = args[flagName].(bool)
	}
	return
}

func (p *ParserImp) extractParamsFromArgs(args map[string]interface{}) (map[string]string, error) {
	chunkOfParamTexts, _ := args[commandParamsKey].([]string)
	return p.buildParamsFromChunks(chunkOfParamTexts)
}

// buildParamsFromChunksUsingRegex collects all chunks of param in
// a structured way.
//  	["DATABASE_NAME='mysqldb'", "USER='root'"] => {"DATABASE_NAME": "mysqldb", "USER": "root"}
func (p *ParserImp) buildParamsFromChunks(chunks []string) (params map[string]string, e error) {
	matchedChunks, got := p.getMatchedChunks(chunks)
	if !got {
		return nil, NewError(errInvalidParamsCode, strings.Join(chunks, " "))
	}
	params = helper.ArrayToMap(matchedChunks)
	return
}

func (p *ParserImp) getMatchedChunks(chunks []string) (matchedChunks [][]string, got bool) {
	lineOfParamTexts := strings.Join(chunks, " ")
	// get chunks of param text matching paramsRegexPattern
	if len(chunks) < 1 {
		return nil, false
	}
	matchedChunks = p.regex.FindAllStringSubmatch(lineOfParamTexts, -1)
	return matchedChunks, true
}
