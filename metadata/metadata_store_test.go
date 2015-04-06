package metadata_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fdsolutions/logan/errors"
	. "github.com/fdsolutions/logan/metadata"
)

var (
	UnexistingPath      string = filepath.Join("..", "fixtures", "nofile.metas")
	ExistingPath        string = filepath.Join("..", "fixtures", "command_examples.metas")
	EmptyFilePath       string = filepath.Join("..", "fixtures", "empty.metas")
	UnsupportedFilePath string = filepath.Join("..", "fixtures", "unsupported_yaml.metas")
)

var _ = Describe("MetadataStore", func() {

	var (
		store, emptyStore *FileStore
	)

	BeforeEach(func() {
		store, _ = NewFileStore(ExistingPath)
		emptyStore, _ = NewFileStore(EmptyFilePath)
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
				_, err := NewFileStore(UnexistingPath)
				Expect(err).To(MatchError(errors.New(ErrInvalidFilePath)))
			})
		})

		Context("Given an existing filepath", func() {
			It("should get back an intance of the store", func() {
				store, _ = NewFileStore(ExistingPath)
				Expect(store.Filepath()).To(Equal(ExistingPath))
			})
		})
	})

	Describe("FileStore", func() {
		Describe("#QueryAll", func() {
			Context("With an a empty source file", func() {
				It("should return an empty metadata collection", func() {
					entries, _ := emptyStore.QueryAll()
					Expect(entries).To(BeEmpty())
				})
			})

			Context("With a source file containing metadata in a non-JOSN format", func() {
				It("should return an ErrUnsupportedFileFormat", func() {
					store, _ := NewFileStore(UnsupportedFilePath)
					_, err := store.QueryAll()
					Expect(err).To(MatchError(errors.New(ErrUnsupportedFileFormat)))
				})
			})

			Context("With a file containing metadata in JSON format", func() {
				// TODO : Make the choice of element random in asserion 'containElement'
				It("should return all metadata entries from store file", func() {
					entries, _ := store.QueryAll()
					expEntry := NewFromGoal(TestGoalWithContext)
					expEntry.SetPath("/usr/bin/cp <SOURCE_FILE> <DESTINATION_FILE>")
					expEntry.SetRequiredParams([]string{
						"<SOURCE_FILE>",
						"<DESTINATION_FILE>",
					})
					Expect(entries).To(ContainElement(*expEntry))
				})
			})

			Context("When call twice or more", func() {
				It("Should not load data again", func() {
					store.QueryAll()
					store.QueryAll()
					Expect(store.HasDataAlreadyLoaded()).To(BeTrue())
				})
			})
		})

		Describe("#Query", func() {
			Context("With no predicate provided", func() {
				It("should always return an empty collection from an empty store", func() {
					entries := emptyStore.Query(nil)
					Expect(entries).To(BeEmpty())
				})

				It("should return all entries from an none empty store", func() {
					entries := store.Query(nil)
					Expect(entries).NotTo(BeEmpty())
				})
			})

			Context("With a predicate and a goal name provided ", func() {
				It("should always return an empty collection form an empty store", func() {
					entries := emptyStore.Query(PredicateForGoal(TestGoal))
					Expect(entries).To(BeEmpty())
				})

				It("should return a collection with only one entry", func() {
					entries := store.Query(PredicateForGoal(TestGoal))
					expEntry := NewFromGoal(TestGoal)
					Expect(entries).To(ContainElement(*expEntry))
				})

				PIt("Must return an ErrMetadataConflict if many metadata are found for that goal name.")
			})
		})
	})
})
