package usage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/usage"
)

var expectedUsage = `Usage:
  logan [options] <action>  [<param>...]

Options:
  -h --help     Show this screen.
  --version     Show version.
  -s, --sudo    Run the coammand in sudo mode.

Arguments
  <action>      The name of the action to execute expressed as a composition of '<intent>:<target>:<context>'
                Eg: create:file:windows

  <param>      Parameters we want to pass to the action. You can add multiple parameters with space separated.
                By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
                Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'
`

var _ = Describe("LoganUsage", func() {
	It("Should render logan usage", func() {
		Expect(LoganUsage()).To(Equal(expectedUsage))
	})
})
