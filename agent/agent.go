package agent

import (
	// "fmt"
	"regexp"
	//"strconv"
	"errors"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Agent is the one which handle logan action
type Agent struct {
	Options map[string]interface{}
}

// NewAgent create a new logan agent
func NewAgent() *Agent {
	return &Agent{make(map[string]interface{})}
}

// Action is a combination of <intent:target:context>
// and <parameters> if any
type Action struct {
	Intent     string
	Target     string
	Context    string
	Parameters map[string]string
}

// NewAction create a new logan action
func NewAction(intent string, target string, context string) *Action {
	return &Action{intent, target, context, make(map[string]string)}
}

// SetParams assign params to the action
func (ac *Action) SetParams(params map[string]string) {
	ac.Parameters = params
}

// type ParsingError struct {
//   Cause string
//   Expr string
// }
// type ParamsParsingError ParsingError
// type ActionParsingError ParsingError

// func (e *ParsingError) Error() string {
//   return e.Cause + ": Failed to parse <" + e.Expr + ">"
// }

const (
	// LoganSemVer is the version number of logan.
	// it follows SEMVER convention
	LoganSemVer = "0.1.0"

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

// Tuts about regex : https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown

// var r = regexp.MustCompile(`(?:(\w*)='?(\S*)'?)*`)
// fmt.Println("var %v", r.FindAllStringSubmatch("VER=1.0 DETAIL=NULL DATE=12/12/2015", -1))
// => var %v [[VER=1.0 VER 1.0] [DETAIL=NULL DETAIL NULL] [DATE=12/12/2015 DATE 12/12/2015]]

var paramsRegex = regexp.MustCompile(`(?:(\w*)=?'(\S*)?')*`)

// Errors
// const (

// )

// ParseParams parses params as strign and returns Map respresenting
// params as a key, value pairs collection
// It return nil if the parsing fails
//   Ex: ("VER=1.0 DETAIL=NULL DATE=12/12/2015") => map[VER:1.0 DETAIL:NULL DATE:12/12/2015]
func (a *Agent) ParseParams(in string) (map[string]string, error) {
	paramsArr := paramsRegex.FindAllStringSubmatch(in, -1)
	if in == "" || len(paramsArr) < 1 {
		return nil, errors.New("Invalid params <" + in + ">")
	}
	paramsMap := arr2Map(paramsArr)

	return paramsMap, nil
}

// arr2Map is an helper function taht tranform an array of array to a map
func arr2Map(arr [][]string) map[string]string {
	paramsMap := make(map[string]string)
	for _, innerArr := range arr {
		// Build the param map
		key := innerArr[1]
		val := innerArr[2]
		paramsMap[key] = val
	}
	return paramsMap
}

func (a *Agent) buildAction(goal string, paramsArr []string) (action Action, e error) {
	intent, target, ctx, partMissing := getGoalParts(goal)
	if partMissing {
		e = errors.New("<intent> and <target> params are required. They must not be empty")
	}
	action = *NewAction(intent, target, ctx)
	params := strings.Join(paramsArr, " ")   // Join all parameter string
	actionParams, _ := a.ParseParams(params) // No error handling
	action.SetParams(actionParams)

	return
}

// ParseAction parses a given action command as a string and
// returns an action object representing the action string
// TODO: Review the casting 'cause it seems not right
func (a *Agent) ParseAction(cmd string) (action Action, e error) {
	var argv = strings.Split(cmd, " ")
	args, _ := docopt.Parse(UsageStr, argv, true, LoganSemVer, false)
	goal, _ := args[availableArgsNames["GOAL"]]
	paramsArr, _ := args[availableArgsNames["PARAMS"]]

	return a.buildAction(goal.(string), paramsArr.([]string))
}

// Return intent, target, context and a boolean isPartMissing.
// isPartMissing ensure that the intent and the target are present.
// When count mismatch we fill the missing value with 'nil'
func getGoalParts(goal string) (intent string, target string, ctx string, partMissing bool) {
	parts := strings.SplitN(goal, ":", 3)
	switch len(parts) {
	case 1: // We have at list the intent
		intent, target, ctx, partMissing = parts[0], "", "", true
	case 2:
		intent, target, ctx, partMissing = parts[0], parts[1], "", false
	default:
		intent, target, ctx, partMissing = parts[0], parts[1], parts[2], false
	}
	return
}

// TODO: parse func =>
// if !a.isInitialize() { // Agent not initialized
//   return nil, errors.New(`Agent must intialize first.
//     Use method 'a.initialize(options)' before a.ParseAction`)
// }
