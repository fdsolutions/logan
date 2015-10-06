package usage_test

import (
	"strings"

	//"github.com/kr/pretty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/usage"
)

var expectedUsage = `Usage:  logan  [options]  <intent>  [<param>...]

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.

Arguments
  <intent>   The intent describing the action being performed.
             Intent is formed of '<verb>:<target>:<context>'.
             Eg: create:file:windows

  <param>    The argument passed as an action parameter.
             You can pass multiple parameters separated by space.
             By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
             Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'
`

var _ = Describe("LoganUsage", func() {
	It("Should render logan usage", func() {
		expected := strings.TrimSpace(expectedUsage)
		Expect(LoganUsage()).To(Equal(expected))
	})
})
