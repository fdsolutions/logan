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
	Defaults Options
}

// Options is a generic option type
type Options map[string]interface{}

// ActionCollection gathered 2 types of values.
// - v:  A simple map for basic implementation
// - ch: A channel implementation for concurrency
type ActionCollection struct {
	v  map[string]Action // For simple implementation
	ch <-chan Action
}

// Output is an abstraction to show output
type Output interface {
	Show()
}

// ActionProcessor is tha abstraction of an action processing flow
type ActionProcessor interface {
	ParseParams(params string) (ActionParams, error)
	ParseAction(actionCmd string) (Action, error)
	CheckAction(action Action) bool
	Configure(options Options)
	LoadActions(paths []string) ActionCollection
	// CacheActions(c ActionCollection) (bool, error)
	Lookup(a Action, c ActionCollection) (Action, bool)
	Perform(a Action) (Output, error)
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
	Parameters ActionParams
}

// ActionParams is an alias of type map[string]string
type ActionParams map[string]string

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

// Tuts about regex : https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown

// var r = regexp.MustCompile(`(?:(\w*)='?(\S*)'?)*`)
// fmt.Println("var %v", r.FindAllStringSubmatch("VER=1.0 DETAIL=NULL DATE=12/12/2015", -1))
// => var %v [[VER=1.0 VER 1.0] [DETAIL=NULL DETAIL NULL] [DATE=12/12/2015 DATE 12/12/2015]]

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

// arr2Map is an helper function that tranform an array of array to a map
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

// buildAction builds an action from its given goal and params
// Its returns the built action and an error if any
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

// ParseAction parses a action command given as a string and
// returns an action object and an error if something went wrong.
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
