package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/command"
)

// Regex tester : https://regex-golang.appspot.com/assets/html/index.html

var (
	defaultFlags = map[string]bool{"--help": false, "--version": false, "--sudo": false}

	userInputExamples = []struct {
		in       string
		expected Arg
	}{
		{"--sudo start:server",
			Arg{"start:server",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				nil}},
		{"show:help",
			Arg{"show:help",
				defaultFlags,
				nil}},
		{"install:pkg:ubuntu PKG_NAME='apache'",
			Arg{"install:pkg:ubuntu",
				defaultFlags,
				map[string]string{"PKG_NAME": "apache"}}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Arg{"connect:database:mysql",
				defaultFlags,
				map[string]string{"DATABASE_NAME": "mysqldb"}}},
		{"--sudo connect:database:mysql DATABASE_NAME='mysqldb' USER='root' PASSWORD='root' VERSION=1.0.1",
			Arg{"connect:database:mysql",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				map[string]string{"DATABASE_NAME": "mysqldb",
					"USER":     "root",
					"PASSWORD": "root",
					"VERSION":  "1.0.1"}}},
	}
)

var _ = Describe("Args", func() {

	It("should parse user inputs", func() {
		for _, input := range userInputExamples {
			var got, _ = ArgFromInput(input.in)
			var expected = input.expected
			Expect(got).To(Equal(expected))
		}
	})

})
