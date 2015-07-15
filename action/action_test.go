package action_test

import (
	. "github.com/fdsolutions/logan/action"

	//. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Action", func() {

	BeforeEach(func() {
		act := NewAction()
	})

	Describe(".BuildCommand", func() {
		Context("With no parameters", func() {
			It("Should fail when action metadata doesn't require paramters", func() {

			})
		})
	})
})
