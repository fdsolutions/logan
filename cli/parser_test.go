package cli_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/cli"
)

// Regex tester : https://regex-golang.appspot.com/assets/html/index.html

var (
	cliExprs = []struct {
		in       string
		expected Argv
	}{
		{"start:server",
			Argv{"name": "start:server"}},
		{"show:help",
			Argv{"name": "show:help"}},
		{"install:pkg:ubuntu",
			Argv{"name": "install:pkg:ubuntu"}},
		{"connect:database:mysql DATABASE_NAME='mysqldb'",
			Argv{"name": "connect:database:mysql",
				"args": map[string]string{
					"DATABASE_NAME": "mysqldb",
				},
			},
		},
		{"connect:database:mysql DATABASE_NAME='mysqldb' USER='root' PASSWORD='root' VERSION=1.0.1",
			Argv{"name": "connect:database:mysql",
				"args": map[string]string{
					"DATABASE_NAME": "mysqldb",
					"USER":          "root",
					"PASSWORD":      "root",
					"VERSION":       "1.0.1",
				},
			},
		},
	}
)

var _ = Describe("Parser", func() {
	var parser *ParserImp

	BeforeEach(func() {
		parser = NewParser()
	})

	It("should parse CLI expressions", func() {
		for _, cliExpr := range cliExprs {
			var got, _ = parser.ParseExpr(cliExpr.in)
			var expected = cliExpr.expected
			Expect(got).To(Equal(expected))
		}
	})

})
