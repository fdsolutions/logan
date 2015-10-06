package usage

const (
	loganUsageTpl string = `Usage:
  logan [options] {{.ArgName}}  [{{.ArgParamName}}...]

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.

Arguments
  {{.ArgName}}      The name of the action to execute expressed as a composition of '<intent>:<target>:<context>'
                Eg: create:file:windows

  {{.ArgParamName}}      Parameters we want to pass to the action. You can add multiple parameters with space separated.
                By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
                Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'
`
)
