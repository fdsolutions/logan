package logan_test

import (
	. "github.com/fdsolutions/logan"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logan", func() {

	Describe("Logan CLI parse arguments", func() {
		Context("With no arguments", func() {
			It("Should display command line usage text", func(){
					var Agent = NewAgent()
					Expect(Agent).ShouldNot(BeNil())
			})
		})
	})


})
