package usage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/usage"
)

var expectedUsage = `A Command line tool helps to organise our scripts.

Usage:
	logan [options] NAME  [PARAM...]
	logan -h | --help
	logan --version

Arguments
	NAME:   The name of the command expressed as a composition of this 3 items '<intent>:<target>:<context>'
	    	- <intent>  : Define the action we want to perform as a verb.
	    	              Eg: 'create'
	    	- <target>  : Define the target that we have the intention to
	    	              operate on.
	    	              Eg: 'file'
	    	- <context> : Define the context in which the action is performed.
	    	              Eg: 'windows'
	    		Eg: create:file:windows

	PARAM:  Additional parameters we want to pass wth the command.
	        You can add multiple parameters with space separated.
	        By convention, we use UPPERCASE_VAR_NAME='<var_value' ...
	        	Eg: FILE_NAME='sample.txt' OWNER='fdsolutions'

Options:
	-h --help     Show this screen.
	--version     Show version.
	-s, --sudo    Run the coammand in sudo mode.
`

var _ = Describe("LoganUsage", func() {
	It("Should render logan usage", func() {
		Expect(LoganUsage()).To(Equal(expectedUsage))
	})
})
