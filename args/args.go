package args

import (
	"strings"

	docopt "github.com/docopt/docopt-go"

	errors "github.com/fdsolutions/logan/errors"
	helper "github.com/fdsolutions/logan/helper"
	usage "github.com/fdsolutions/logan/usage"
	version "github.com/fdsolutions/logan/version"
)

const (
	argsParamsRegexPattern = `(?:(\w*)='?([^']*)'?)*`
)

var (
	availableFlagNames = map[string]string{
		"HELP":    "--help",
		"VERSION": "--version",
		"SUDO":    "--sudo",
	}

	defaultParamParser = NewParamParser()
)

// Arg holds CLI argument elements
type Arg struct {
	Name   string
	Flags  map[string]bool
	Params map[string]string
}

// ArgFromInput returns argument elements from user command input
func ParseInputWithParser(input string, pp ParamParser) (arg Arg, e errors.LoganError) {
	if pp == nil {
		pp = GetDefaultParamParser()
	}

	argv := getArgvFromInput(input)
	parsedArgs, err := docopt.Parse(usage.LoganUsage(), argv, true, version.LoganVersion, true, false)
	if err != nil || parsedArgs == nil {
		e = errors.New(errors.ErrInvalidInput)
		return
	}

	arg = parseArgElementsWithParser(parsedArgs, pp)
	return
}

func getArgvFromInput(input string) (argv []string) {
	// Sanitize the given  input
	in := strings.TrimSpace(input)
	if in == "" {
		return
	}
	return strings.Split(in, " ")
}

// GetDefaultParamParser returns the default param parser for args
func GetDefaultParamParser() ParamParser {
	return defaultParamParser
}

func parseArgElementsWithParser(args map[string]interface{}, pp ParamParser) Arg {
	name := parseName(args)
	flags := parseFlags(args)
	params, _ := parseParamsWithParser(args, pp) // No error handling if fails

	return Arg{name, flags, params}
}

func parseName(args map[string]interface{}) string {
	name, _ := args[usage.CommandArgName]
	return name.(string)
}

func parseFlags(args map[string]interface{}) (flags map[string]bool) {
	flags = map[string]bool{}
	for _, flagName := range availableFlagNames {
		flags[flagName] = args[flagName].(bool)
	}
	return
}

// parseParamsWithParser  parses arguments and retrives argument's parameters as a key/value pairs
func parseParamsWithParser(args map[string]interface{}, pp ParamParser) (map[string]string, errors.LoganError) {
	argsParamList, _ := args[usage.CommandArgParamsName].([]string)
	params, ok := parseParamListWithParser(argsParamList, pp)
	if !ok {
		return nil, errors.New(errors.ErrInvalidParams)
	}
	return helper.ArrayToMap(params), nil
}

func parseParamListWithParser(paramList []string, pp ParamParser) ([][]string, bool) {
	if len(paramList) < 1 {
		return nil, false
	}
	inlineParamList := strings.Join(paramList, " ")
	parsedParams := pp.Parse(inlineParamList)
	if parsedParams == nil {
		return nil, false
	}
	return parsedParams, true
}
