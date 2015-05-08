package agent_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/agent"
	"github.com/fdsolutions/logan/errors"
)

var _ = Describe("Agent", func() {

	const (
		WithNothing = iota
		WithOnlyAction
		WithActionAndParams
	)

	var (
		agent *Agent

		usersInputExamples = map[int]string{
			WithNothing:         "",
			WithOnlyAction:      "show:help",
			WithActionAndParams: "install:pkg:ubuntu PKG_NAME='apache'",
		}
	)

	BeforeEach(func() {
		agent = FromFactoryAndRepos(nil, nil)
	})

	Describe(".PerformActionFromInput", func() {
		Context("With an invalid input", func() {
			It("Should fail with the error ErrInvalidUserInput", func() {
				status := agent.PerformActionFromInput(usersInputExamples[WithNothing])
				errValue := status.GetValue()
				if _, ok := errValue.(errors.ErrorCode); !ok {
					Fail(fmt.Sprintf("The status value must be of type errors.ErrorCode but %v [%v]", reflect.TypeOf(errValue), errValue))
				}
				Expect(status.GetErrorStackCodes()).To(ContainElement(errors.ErrInvalidUserInput))
			})
		})
		Context("With a valid input", func() {
			It("Should parse user's input as an action with no parameters", func() {
				status := agent.PerformActionFromInput(usersInputExamples[WithOnlyAction])
				Expect(status.GetCode()).To(Equal(StatusSuccess))
			})
		})
	})
})
