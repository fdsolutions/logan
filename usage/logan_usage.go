package usage

import (
	"bytes"
	"fmt"
	"text/template"
)

const (
	CommandArgName       string = "<name>"
	CommandArgParamsName string = "<param>"
)

type usageInfo struct {
	ArgName      string
	ArgParamName string
}

func LoganUsage() string {
	b := new(bytes.Buffer)

	info := usageInfo{
		CommandArgName,
		CommandArgParamsName,
	}

	t := template.Must(template.New("logan usage").Parse(loganUsageTpl))
	err := t.Execute(b, info)
	if err != nil {
		panic(fmt.Sprintf("Fail to compile the logan usage template: %v", err))
	}
	return b.String()
}
