package usage

const (
	loganUsageTpl string = `A Command line tool helps to organise our scripts.

Usage:
	logan [options] {{.ArgName}}  [{{.ArgParamName}}...]
	logan -h | --help
	logan --version

Arguments
	{{.ArgName}}:   The name of the command expressed as a composition of this 3 items '<intent>:<target>:<context>'
	    	- <intent>  : Define the action we want to perform as a verb.
	    	              Eg: 'create'
	    	- <target>  : Define the target that we have the intention to
	    	              operate on.
	    	              Eg: 'file'
	    	- <context> : Define the context in which the action is performed.
	    	              Eg: 'windows'
	    		Eg: create:file:windows

	{{.ArgParamName}}:  Additional parameters we want to pass wth the command.
	        You can add multiple parameters with space separated.
	        By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
	        	Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
	-h --help     Show this screen.
	--version     Show version.
	-s, --sudo    Run the coammand in sudo mode.
`
)
