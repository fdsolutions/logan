package command

import (
	"regexp"
	"strings"

	docopt "github.com/docopt/docopt-go"

	errors "github.com/fdsolutions/logan/errors"
	helpers "github.com/fdsolutions/logan/helpers"
	usage "github.com/fdsolutions/logan/usage"
	version "github.com/fdsolutions/logan/version"
)

const (
	argsParamsRegexPattern = `(?:(\w*)='?([^']*)'?)*`

	ErrInvalidInput  errors.ErrorCode = "Invalid input."
	ErrInvalidParams errors.ErrorCode = "Invalid params : No params retreived from user input."
)

var (
	availableFlagNames = map[string]string{
		"HELP":    "--help",
		"VERSION": "--version",
		"SUDO":    "--sudo",
	}
)

// Args holds CLI argument elements
type Args struct {
	Name   string
	Flags  map[string]bool
	Params map[string]string
}

// ArgFromInput returns argument elements from user command input
func ArgsFromInput(input string) (args Args, e error) {
	argv := strings.Split(input, " ")

	parsedArgs, err := docopt.Parse(usage.LoganUsage(), argv, true, version.LoganVersion, false)
	if err != nil {
		e = errors.New(ErrInvalidInput)
		return
	}
	args = extractArgsElementsFrom(parsedArgs)
	return
}

func extractArgsElementsFrom(args map[string]interface{}) Args {
	name := extractNameFrom(args)
	flags := extractFlagsFrom(args)
	params, _ := extractParamsFrom(args) // No error handling if fails

	return Args{name, flags, params}
}

func extractNameFrom(args map[string]interface{}) string {
	name, _ := args[usage.CommandArgName]
	return name.(string)
}

func extractFlagsFrom(args map[string]interface{}) (flags map[string]bool) {
	flags = map[string]bool{}
	for _, flagName := range availableFlagNames {
		flags[flagName] = args[flagName].(bool)
	}
	return
}

func extractParamsFrom(args map[string]interface{}) (map[string]string, error) {
	argsParamList, _ := args[usage.CommandArgParamsName].([]string)
	return makeParamsFromList(argsParamList)
}

// makeParamsFromList transforms list of arg's params to a map
//  	["DATABASE_NAME='mysqldb'", "USER='root'"] => {"DATABASE_NAME": "mysqldb", "USER": "root"}
func makeParamsFromList(paramList []string) (map[string]string, error) {
	params, ok := scanParamList(paramList)
	if !ok {
		return nil, errors.New(ErrInvalidParams)
	}
	return helpers.ArrayToMap(params), nil
}

func scanParamList(paramList []string) ([][]string, bool) {
	if len(paramList) < 1 {
		return nil, false
	}
	inlineParamList := strings.Join(paramList, " ")
	paramParser := Parser(ArgsParamParser{})
	parsedParams := paramParser.Parse(inlineParamList)
	if pp, ok := parsedParams.([][]string); !ok {
		return nil, false
	} else {
		return pp, true
	}
}

// ArgsParamParser is used to parse arg parameters.
// It implements the parser.Parser interface to define its own way to parse input.
type ArgsParamParser struct {
	Parser
}

// Parse define the way arg parameters must be parsed
func (a ArgsParamParser) Parse(input string) (params interface{}) {
	r := regexp.MustCompile(argsParamsRegexPattern)
	params = r.FindAllStringSubmatch(input, -1)
	return
}
