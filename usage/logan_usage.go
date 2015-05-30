package usage

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

const (
	ActionTokenName       string = "<action>"
	ActionParamsTokenName string = "<param>"
)

type usageInfo struct {
	ArgName      string
	ArgParamName string
}

func LoganUsage() string {
	b := new(bytes.Buffer)

	info := usageInfo{
		ActionTokenName,
		ActionParamsTokenName,
	}

	t := template.Must(template.New("logan usage").Parse(strings.TrimSpace(loganUsageTpl)))
	err := t.Execute(b, info)
	if err != nil {
		panic(fmt.Sprintf("Fail to compile the logan usage template: %v", err))
	}
	return b.String()
}
