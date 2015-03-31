package command_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/command"
	errors "github.com/fdsolutions/logan/errors"
)

var _ = Describe("metadata", func() {

	var (
		unexistingPath      string = filepath.Join("..", "fixtures", "nofile.metas")
		existingPath        string = filepath.Join("..", "fixtures", "command_examples.metas")
		emptyFilePath       string = filepath.Join("..", "fixtures", "empty.metas")
		unsupportedFilePath string = filepath.Join("..", "fixtures", "unsupported_yaml.metas")
		store, emptyStore   *FileMetadataStore
	)

	BeforeEach(func() {
		store, _ = NewFileMetadataStore(existingPath)
		emptyStore, _ = NewFileMetadataStore(emptyFilePath)
	})

	Describe(".NewFileMetadataStore", func() {
		Context("Given an empty filename", func() {
			It("should return an ErrInvalidFilePath", func() {
				_, err := NewFileMetadataStore("")
				Expect(err).To(MatchError(errors.New(ErrInvalidFilePath)))
			})
		})

		Context("Given an unexisting filepath", func() {
			It("should return an ErrFileDontExistAtPath", func() {
				_, err := NewFileMetadataStore(unexistingPath)
				Expect(err).To(MatchError(errors.New(ErrInvalidFilePath)))
			})
		})

		Context("Given an existing filepath", func() {
			It("should get back an intance of the store", func() {
				store, _ = NewFileMetadataStore(existingPath)
				Expect(store.Filepath()).To(Equal(existingPath))
			})
		})
	})

	Describe("FileMetadataStore", func() {
		Describe("#FindAll", func() {
			Context("With an a empty source file", func() {
				It("should return an empty metadata collection", func() {
					metas, _ := emptyStore.FindAll()
					Expect(metas).To(BeNil())
				})
			})
			Context("With a source file containing metadata in a non-JOSN format", func() {
				It("should return an ErrUnsupportedFileFormat", func() {
					store, _ := NewFileMetadataStore(unsupportedFilePath)
					_, err := store.FindAll()
					Expect(err).To(MatchError(errors.New(ErrUnsupportedFileFormat)))
				})
			})
			Context("With a file containing metadata in JSON format", func() {
				It("should return all metadata from store file", func() {
					metas, _ := store.FindAll()
					expectedMetadata := NewMetadata()
					expectedMetadata.Name = "copy:file:unix"
					expectedMetadata.Path = "/usr/bin/cp <SOURCE_FILE> <DESTINATION_FILE>"
					expectedMetadata.RequiredParams = []string{
						"<SOURCE_FILE>",
						"<DESTINATION_FILE>",
					}
					Expect(metas).To(ContainElement(*expectedMetadata))
				})
			})
		})
	})

})
