package agent

import (
	"fmt"
	"regexp"
	//"strconv"
	"errors"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

// Agent is the one which handle logan action
type Agent struct{}

// Action is a combination of <intent:target:context>
// and <parameters> if any
type Action struct {
	intent     string
	target     string
	context    string
	parameters map[string]string
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
	// LoganVersion is the version number of logan.
	// it follows SEMVER convention
	LoganSemVer = "0.1.0"

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

  <paramater>   Additional parameter we want to pass to our goal.
                You can set multiple parameters with space as a separator.
                By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
                Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the action in sudo mode.
`
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

// ParseAction parses a given action string and
// returns an action object representing the action string
func (a *Agent) ParseAction(action string) (Action, error) {
	var argv = strings.Split(action, " ")
	args, _ := docopt.Parse(UsageStr, argv, true, LoganSemVer, false)
	fmt.Printf("args: %v\n", args)
	return Action{}, nil
}
