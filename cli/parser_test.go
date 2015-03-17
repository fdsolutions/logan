package cli_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/cli"
)

// Regex tester : https://regex-golang.appspot.com/assets/html/index.html

var (
	defaultFlags = map[string]bool{"--help": false, "--version": false, "--sudo": false}

	cmdInputs = []struct {
		in       string
		expected Argv
	}{
		{"--sudo start:server",
			Argv{"start:server",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				nil}},
		{"show:help",
			Argv{"show:help",
				defaultFlags,
				nil}},
		{"install:pkg:ubuntu PKG_NAME='apache'",
			Argv{"install:pkg:ubuntu",
				defaultFlags,
				map[string]string{"PKG_NAME": "apache"}}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Argv{"connect:database:mysql",
				defaultFlags,
				map[string]string{"DATABASE_NAME": "mysqldb"}}},
		{"--sudo connect:database:mysql DATABASE_NAME='mysqldb' USER='root' PASSWORD='root' VERSION=1.0.1",
			Argv{"connect:database:mysql",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				map[string]string{"DATABASE_NAME": "mysqldb",
					"USER":     "root",
					"PASSWORD": "root",
					"VERSION":  "1.0.1"}}},
	}
)

var _ = Describe("Parser", func() {
	var parser *ParserImp = NewParser()

	It("should parse CLI inputs", func() {
		for _, cmd := range cmdInputs {
			var got, _ = parser.ParseUserInput(cmd.in)
			var expected = cmd.expected
			Expect(got).To(Equal(expected))
		}
	})

})
