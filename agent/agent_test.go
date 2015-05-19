package agent_test

import (
	//"fmt"
	//"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/agent"
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/fixtures"
	"github.com/fdsolutions/logan/metadata"
)

var _ = Describe("Agent", func() {

	const (
		WithNothing = iota
		WithOnlyAction
		WithActionAndParams
	)

	const UnkownActionGoal = "fail:test"

	var (
		agent      *Agent
		repoSample metadata.Repository

		usersInputExamples = map[int]string{
			WithNothing:         "",
			WithOnlyAction:      "show:help",
			WithActionAndParams: "install:pkg:ubuntu PKG_NAME='apache'",
		}
	)

	BeforeEach(func() {
		agent = FromFactoryAndRepos(nil, nil)
		store, _ := metadata.NewFileStore(fixtures.ExistingPath)
		repoSample = metadata.NewRepositoryFromStore(store)
	})

	Describe(".ParseUserInput", func() {
		Context("With an invalid input", func() {
			It("Should fail with the error ErrInvalidUserInput", func() {
				s := agent.ParseUserInput(usersInputExamples[WithNothing])
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrInvalidUserInput))
			})
		})
		Context("With a valid input", func() {
			It("Should parse user's input as an action with no parameters", func() {
				s := agent.ParseUserInput(usersInputExamples[WithOnlyAction])
				Expect(s.GetCode()).To(Equal(StatusSuccess))
			})
		})
	})

	Describe(".LookupActionInRepos", func() {
		Context("With empty goal and no repos provided", func() {
			It("Should fail with an ErrInvalidGoal error.", func() {
				s := agent.LookupActionInRepos("", nil)
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrInvalidGoal))
			})
		})
		Context("With a goal that has no action related in any provided repos", func() {
			It("Should return an ErrActionNotFound error.", func() {
				s := agent.LookupActionInRepos(UnkownActionGoal, []metadata.Repository{repoSample})
				Expect(s.GetCode()).To(Equal(StatusFail))
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrActionNotFound))
			})
		})
		Context("With a goal that hthe related action is present in one of the provided repos", func() {
			XIt("Should return the action related to that given goal.", func() {

			})
		})
	})
})
