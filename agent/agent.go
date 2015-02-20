package agent

import "regexp"

type Agent struct{}

type Action struct {
	intent     string
	target     string
	context    string
	parameters map[string]string
}

const VERSION = "0.1.0"

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

var params_regex = regexp.MustCompile(`(\w*)='?(\S*)'?`)

func buildParams(attr_values ...string) (params map[string]string) {
	params = make(map[string]string)

	if len(attr_values) > 0 {
		for _, attr_val := range attr_values {
			g := params_regex.FindStringSubmatch(attr_val)

			if len(g) > 2 {
				params[g[1]] = g[2]
			}
		}
	}
	return
}
