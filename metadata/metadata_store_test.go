package metadata_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fdsolutions/logan/errors"
	. "github.com/fdsolutions/logan/metadata"
)

var _ = Describe("MetadataStore", func() {
	var (
		unexistingPath      string = filepath.Join("..", "fixtures", "nofile.metas")
		existingPath        string = filepath.Join("..", "fixtures", "command_examples.metas")
		emptyFilePath       string = filepath.Join("..", "fixtures", "empty.metas")
		unsupportedFilePath string = filepath.Join("..", "fixtures", "unsupported_yaml.metas")
		store, emptyStore   *FileStore
	)

	BeforeEach(func() {
		store, _ = NewFileStore(existingPath)
		emptyStore, _ = NewFileStore(emptyFilePath)
	})

	Describe(".NewFileStore", func() {
		Context("Given an empty filename", func() {
			It("should return an ErrInvalidFilePath", func() {
				_, err := NewFileStore("")
				Expect(err).To(MatchError(errors.New(ErrInvalidFilePath)))
			})
		})

		Context("Given an unexisting filepath", func() {
			It("should return an ErrFileDontExistAtPath", func() {
				_, err := NewFileStore(unexistingPath)
				Expect(err).To(MatchError(errors.New(ErrInvalidFilePath)))
			})
		})

		Context("Given an existing filepath", func() {
			It("should get back an intance of the store", func() {
				store, _ = NewFileStore(existingPath)
				Expect(store.Filepath()).To(Equal(existingPath))
			})
		})
	})

	Describe("FileStore", func() {
		Describe("#QueryAll", func() {
			Context("With an a empty source file", func() {
				It("should return an empty metadata collection", func() {
					metas, _ := emptyStore.QueryAll()
					Expect(metas).To(BeNil())
				})
			})
			Context("With a source file containing metadata in a non-JOSN format", func() {
				It("should return an ErrUnsupportedFileFormat", func() {
					store, _ := NewFileStore(unsupportedFilePath)
					_, err := store.QueryAll()
					Expect(err).To(MatchError(errors.New(ErrUnsupportedFileFormat)))
				})
			})
			Context("With a file containing metadata in JSON format", func() {
				// TODO : Make the choice of element random in asserion 'containElement'
				It("should return all metadata entries from store file", func() {
					entries, _ := store.QueryAll()
					expEntry := NewEntry()
					expEntry.Goal = "copy:file:unix"
					expEntry.Path = "/usr/bin/cp <SOURCE_FILE> <DESTINATION_FILE>"
					expEntry.RequiredParams = []string{
						"<SOURCE_FILE>",
						"<DESTINATION_FILE>",
					}
					Expect(entries).To(ContainElement(*expEntry))
				})
			})
		})

		Describe("#Query", func() {
			Context("With no predicate", func() {
				It("should query all entries", func() {
					entries := store.Query(nil)
					Expect(entries).ToNot(BeEmpty())
				})
			})

			Context("Entries of a given goal name", func() {
				It("should get all entries of that goal name", func() {
					goalName := "show:version"
					entries := store.Query(func(ent Entry) bool {
						return (ent.Goal == goalName)
					})
					expEntry := NewEntry()
					expEntry.Goal = "show:version"
					expEntry.Target = "version"
					Expect(entries).To(ContainElement(*expEntry))
				})
			})
		})
	})
})
