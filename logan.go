package logan

type LoganAgent struct {}

type LoganAction struct {
	intent string
	target string
	context string
	parameters map[string]string
}


var Usage = `Command line scripts organiser.

Usage:
  logan [options] <goal> [<parameter>...]
  logan -h | --help
  logan --version

Arguments
  <goal>        The part of the action made up with <intent>:<target>:<context>
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
                       Don't use it a lot 'cause it decreases the meaning of the action

Options:
  -h --help   Show this screen.
  --version   Show version.
  -s, --sudo  Run the action in sudo mode.
`

func NewAgent() *LoganAgent {
	return &LoganAgent{}
}

// Perform an action and returns and output string and an error if any
// appears in the function execution.
// @TODO: Handle multiple actions and not only one
func (agent *LoganAgent)Do(a ...interface{}) (output string, e error) {
	if (len(a) > 1) { // Only one action is allow at the moment

	}

	return agent.DoAction("", "", "", []string{})
}

func (agent *LoganAgent)DoAction(intent string, target string, ctx string, parameters []string) (output string, e error) {
	return Usage,nil
}
