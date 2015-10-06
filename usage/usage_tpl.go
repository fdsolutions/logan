package usage

const (
	loganUsageTpl string = `Usage:  logan  [options]  {{.Intent}}  [{{.Param}}...]

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.

Arguments
  {{.Intent}}   The intent describing the action being performed.
             Intent is formed of '<verb>:<target>:<context>'.
             Eg: create:file:windows

  {{.Param}}    The argument passed as an action parameter.
             You can pass multiple parameters separated by space.
             By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
             Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'
`
)
