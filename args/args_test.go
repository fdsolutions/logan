package args_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/args"
	"github.com/fdsolutions/logan/errors"
	"github.com/fdsolutions/logan/fixtures"
)

var paramParser ParamParser

var _ = Describe("args", func() {

	Context("With no input or multiple space as input", func() {
		It("Should fail and return an error ErrInvalidInput ", func() {
			examples := []string{
				"",
				" ",
				"     ",
			}
			for _, exp := range examples {
				var _, err = ParseInputWithParser(exp, nil)
				Expect(err.Code()).To(Equal(errors.ErrInvalidInput))
			}
		})
	})

	Context("With no given param parser", func() {
		It("Should parse user input using the default parser", func() {
			for _, input := range fixtures.UserInputExamples {
				var got, _ = ParseInputWithParser(input.In, nil)
				var expected = input.Expected
				Expect(got).To(Equal(expected))
			}
		})
	})

	Context("With given param parser", func() {
		It("should parse user inputs", func() {
			paramParser = NewParamParser()
			for _, input := range fixtures.UserInputExamples {
				var got, _ = ParseInputWithParser(input.In, paramParser)
				var expected = input.Expected
				Expect(got).To(Equal(expected))
			}
		})
	})

})
