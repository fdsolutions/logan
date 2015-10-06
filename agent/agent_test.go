package agent_test

import (
	//"fmt"
	//"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fdsolutions/logan/action"
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

	var (
		ag             *Agent
		repoSample     metadata.Repository
		defaultFactory action.Factory

		usersInputExamples = map[int]string{
			WithNothing:         "",
			WithOnlyAction:      "show:help",
			WithActionAndParams: "install:pkg:ubuntu PKG_NAME='apache'",
		}

		_ = action.NewFactory()
	)

	BeforeEach(func() {
		store, _ := metadata.NewFileStore(fixtures.ExistingPath)
		repoSample = metadata.NewRepositoryFromStore(store)
		defaultFactory = action.NewFactory()
		ag, _ = FromFactoryAndRepos(defaultFactory, nil)
	})

	Describe(".FromFactoryAndRepos", func() {
		Context("With no given factory with whatever repos", func() {
			It("should fail with the error ErrMissingActionFactory.", func() {
				_, err := FromFactoryAndRepos(nil, nil)
				Expect(err.Code()).To(Equal(errors.ErrMissingActionFactory))
			})
		})

		Context("With a given factory", func() {
			Context("And no given repos list", func() {
				It("Should return a new agent with no repos", func() {
					ag, _ = FromFactoryAndRepos(defaultFactory, nil)
					Expect(ag).NotTo(BeNil())
					Expect(ag.GetMetadataRepos()).To(BeNil())
				})
			})
			Context("And an empty list of repos", func() {
				It("Should return a new agent with an empty list of repos", func() {
					ag, _ = FromFactoryAndRepos(defaultFactory, []metadata.Repository{})
					Expect(ag).NotTo(BeNil())
					Expect(ag.GetMetadataRepos()).To(BeEmpty())
				})
			})
			Context("And a given list of repos", func() {
				It("Should return a new agent with some repos ", func() {
					r := []metadata.Repository{repoSample}
					ag, _ = FromFactoryAndRepos(defaultFactory, r)
					Expect(ag).NotTo(BeNil())
					Expect(ag.GetMetadataRepos()).NotTo(BeEmpty())
				})
			})
		})
	})

	Describe(".ParseUserInput", func() {

		Context("With an invalid input", func() {
			It("Should fail with the error ErrInvalidUserInput", func() {
				s := ag.ParseUserInput(usersInputExamples[WithNothing])
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrInvalidUserInput))
			})
		})
		Context("With a valid input", func() {
			It("Should parse user's input as an action with no parameters", func() {
				s := ag.ParseUserInput(usersInputExamples[WithOnlyAction])
				Expect(s.GetCode()).To(Equal(StatusSuccess))
			})
		})
	})

	Describe(".LookupActionInRepos", func() {
		Context("With empty goal and no repos provided", func() {
			It("Should fail with an ErrInvalidGoal error.", func() {
				s := ag.LookupActionInRepos("", nil)
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrInvalidGoal))
			})
		})
		Context("With a goal that has no action related in any provided repos", func() {
			It("Should return an ErrActionNotFound error.", func() {
				s := ag.LookupActionInRepos(fixtures.UnkownActionGoal, []metadata.Repository{repoSample})
				Expect(s.GetCode()).To(Equal(StatusFail))
				Expect(s.GetErrorStackCodes()).To(ContainElement(errors.ErrActionNotFound))
			})
		})
		Context("With a goal that the related action is present in one of the provided repos", func() {
			It("Should return the action related to that given goal.", func() {
				s := ag.LookupActionInRepos(fixtures.ExistingGoal, []metadata.Repository{repoSample})
				expected := defaultFactory.MakeActionFromMetadata(metadata.Entry{
					"copy:file:unix", // goal
					"copy",           // intent
					"file",           // target
					"unix",           // context
					"/usr/bin/cp <SOURCE_FILE> <DESTINATION_FILE>",  // path
					[]string{"<SOURCE_FILE>", "<DESTINATION_FILE>"}, // required_params
				})

				Expect(s.GetCode()).NotTo(Equal(StatusFail))

				if a, ok := s.GetValue().(action.LoganAction); ok {
					Expect(a).To(BeEquivalentTo(expected))
				} else {
					Fail("Wrong action value returned")
				}
			})
		})
	})

	Describe(".PerformAction", func() {
		Context("with an invalid action (no action)", func() {
			It("Should do nothing", func() {
				_ = ag.PerformAction(nil)
				Expect(ag.Output).To(BeEmpty())
			})
		})
		Context("with a given action", func() {
			// Context("When the execution of the command linked to that action failed", func() {
			// 	It("Should fail with error ErrActionPerfomingFailed and report failure (output).", func() {

			// 	})
			// })
			// Context("When the execution of the command linked to that action failed", func() {
			// 	XIt("Should succeed and keep the related output.", func() {

			// 	})
			// })

			It("Should execute the command linked to that action and keep the result", func() {
				act := defaultFactory.MakeActionFromMetadata(metadata.Entry{
					"copy:file:unix", // goal
					"copy",           // intent
					"file",           // target
					"unix",           // context
					"/usr/bin/cp <SOURCE_FILE> <DESTINATION_FILE>",  // path
					[]string{"<SOURCE_FILE>", "<DESTINATION_FILE>"}, // required_params
				})
				_ = ag.PerformAction(act)

			})
		})
	})
})
