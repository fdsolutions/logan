package usage

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

const (
	ActionIntentToken string = "<intent>"
	ActionParamToken  string = "<param>"
)

type usageInfo struct {
	Intent string
	Param  string
}

func LoganUsage() string {
	b := new(bytes.Buffer)

	info := usageInfo{
		ActionIntentToken,
		ActionParamToken,
	}

	t := template.Must(template.New("logan usage").Parse(strings.TrimSpace(loganUsageTpl)))
	err := t.Execute(b, info)
	if err != nil {
		panic(fmt.Sprintf("Fail to compile the logan usage template: %v", err))
	}
	return b.String()
}
