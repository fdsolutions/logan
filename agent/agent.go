package agent

import (
	"fmt"
	"regexp"
	"strings"

	docopt "github.com/docopt/docopt-go"
)

type Agent struct{}

type Action struct {
	intent     string
	target     string
	context    string
	parameters map[string]string
}

const LOGAN_VERSION = "0.1.0"

const USAGE = `A Command line tool helps to organise our scripts.

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

// Tuts about regex : https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown

// var r = regexp.MustCompile(`(?:(\w*)='?(\S*)'?)*`)
// fmt.Println("var %v", r.FindAllStringSubmatch("VER=1.0 DETAIL=NULL DATE=12/12/2015", -1))
// => var %v [[VER=1.0 VER 1.0] [DETAIL=NULL DETAIL NULL] [DATE=12/12/2015 DATE 12/12/2015]]

var params_regex = regexp.MustCompile(`(\w*)='?(\S*)'?`)

// Parse params as strign and returns Map respresenting
// params as a key, value pairs collection
func (this *Agent) ParseParams(params string) map[string]string {
	return nil
}

// Parse a given action string and
// returns an action object representing the action string
func (this *Agent) ParseAction(action string) (Action, error) {
	var argv = strings.Split(action, " ")
	args, _ := docopt.Parse(USAGE, argv, true, LOGAN_VERSION, false)
	fmt.Printf("args: %v\n", args)
	return Action{}, nil
}
