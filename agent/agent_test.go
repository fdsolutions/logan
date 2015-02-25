package agent

import (
	"reflect"
	"testing"
)

// type Action struct {
// 	intent     string
// 	target     string
// 	context    string
// 	parameters map[string]string
// }

// Regex tester : https://regex-golang.appspot.com/assets/html/index.html

var (
	actionExps = []struct {
		in       string
		expected Action
	}{
		{"help",
			Action{"help", "", "", nil}},
		{"show:help",
			Action{"show", "help", "", nil}},
		{"install:pkg:ubuntu",
			Action{"install", "pkg", "ubuntu", nil}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Action{"connect", "database", "mysql", map[string]string{
				"DATABASE_NAME": "mysqldb"},
			},
		},
	}

	agent = &Agent{}
)

func TestParseAction(t *testing.T) {
	for _, action := range actionExps {
		var got, _ = agent.ParseAction(action.in)
		var expected = action.expected
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Parsing fails got<%v> != expected<%v>",
				got,
				expected)
		}
	}
}
