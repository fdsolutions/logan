package agent

import (
	"testing"
)

// type Action struct {
// 	intent     string
// 	target     string
// 	context    string
// 	parameters map[string]string
// }

// Regex tester : https://regex-golang.appspot.com/assets/html/index.html

type actionExp struct {
	in       string
	expected Action
}

var (
	emptyParams = make(map[string]string)

	actionExps = []actionExp{
		{"help",
			Action{"help", "", "", emptyParams}},
		{"show:help",
			Action{"show", "help", "", emptyParams}},
		{"install:pkg:ubuntu",
			Action{"install", "pkg", "ubuntu", emptyParams}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Action{"connect", "database", "mysql", map[string]string{
				"DATABASE_NAME": "mysqldb"},
			},
		},
	}

	agent = &Agent{}
)

func TestParse(t *testing.T) {
	for _, action := range actionExps {

	}
}
