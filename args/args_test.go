package args_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/args"
)

var (
	defaultFlags = map[string]bool{"--help": false, "--version": false, "--sudo": false}

	paramParser ParamParser

	userInputExamples = []struct {
		in       string
		expected Args
	}{
		{"--sudo start:server",
			Args{"start:server",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				nil}},
		{"show:help",
			Args{"show:help",
				defaultFlags,
				nil}},
		{"install:pkg:ubuntu PKG_NAME='apache'",
			Args{"install:pkg:ubuntu",
				defaultFlags,
				map[string]string{"PKG_NAME": "apache"}}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Args{"connect:database:mysql",
				defaultFlags,
				map[string]string{"DATABASE_NAME": "mysqldb"}}},
		{"--sudo connect:database:mysql DATABASE_NAME='mysqldb' USER='root' PASSWORD='root' VERSION=1.0.1",
			Args{"connect:database:mysql",
				map[string]bool{"--help": false, "--version": false, "--sudo": true},
				map[string]string{"DATABASE_NAME": "mysqldb",
					"USER":     "root",
					"PASSWORD": "root",
					"VERSION":  "1.0.1"}}},
	}
)

var _ = Describe("args", func() {

	It("should parse user inputs", func() {
		paramParser = NewParamParser()
		for _, input := range userInputExamples {
			var got, _ = ParseInputWithParser(input.in, paramParser)
			var expected = input.expected
			Expect(got).To(Equal(expected))
		}
	})

})
