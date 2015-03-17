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
		{"start:server",
			Action{"start", "server", "", nil}},
		// {"help",
		// 	Action{"help", "", "", nil}},
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

	paramsExps = []struct {
		in       string
		expected map[string]string
	}{
		{"", nil},
		{"DATABASE_NAME='mysqldb'",
			map[string]string{"DATABASE_NAME": "mysqldb"},
		},
		{"DATABASE_NAME='mysqldb' USER='root' PASSWORD='root'",
			map[string]string{
				"DATABASE_NAME": "mysqldb",
				"USER":          "root",
				"PASSWORD":      "root",
			},
		},
	}

	agent = &Agent{}
)

func assertFail(t *testing.T, msg string, got interface{}, expected interface{}) {
	t.Errorf(msg+"\nGot : %v \nExpected : %v",
		got,
		expected)
}

func TestParseParams(t *testing.T) {
	for _, param := range paramsExps {
		var got, _ = agent.ParseParams(param.in)
		var expected = param.expected
		if !reflect.DeepEqual(got, expected) {
			assertFail(t,
				"It Fails parsing paramters <"+param.in+">",
				got,
				expected)
		}
	}
}

func TestParseAction(t *testing.T) {
	for _, action := range actionExps {
		var got, _ = agent.ParseAction(action.in)
		var expected = action.expected
		if !reflect.DeepEqual(got, expected) {
			assertFail(t,
				"It Fails parsing action <"+action.in+">",
				got,
				expected)
		}
	}
}
