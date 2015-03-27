package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/command"
	errors "github.com/fdsolutions/logan/errors"
)

var _ = Describe("Metadata", func() {

	Describe("#NewFileMetadataStore", func() {
		Context("with empty filename", func() {
			It("should raise an ErrInvalidFilePath", func() {
				_, err := NewFileMetadataStore("")
				Expect(err).To(Equal(errors.New(ErrInvalidFilePath)))
			})
		})
	})

})
