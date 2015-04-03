package metadata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/metadata"
)

var _ = Describe("MetadataRepository", func() {
	var (
		r             Repository
		s, emptyStore Store
	)

	BeforeEach(func() {
		s, _ = NewFileStore(ExistingPath)
		emptyStore, _ = NewFileStore(EmptyFilePath)
		r = NewRepositoryFromStore(s)
	})

	Describe("#FindAll", func() {
		It("should return an no entries from a empty store", func() {
			r.SetStore(emptyStore)
			entries := r.FindAll()
			Expect(entries).To(BeEmpty())
		})

		It("should return all metadata entries from a non-empty store", func() {
			entries := r.FindAll()
			Expect(entries).NotTo(BeEmpty())
		})
	})

	Describe("#FindByGoal", func() {
		It("Should no metatdata entry for the given from an empty store", func() {
			r.SetStore(emptyStore)
			entry := r.FindByGoal(TestGoal)
			Expect(entry).To(BeZero())
		})

		It("Should the metatdata entry of the given goal.", func() {
			entry := r.FindByGoal(TestGoal)
			expected := NewFromGoal(TestGoal)
			Expect(entry).To(Equal(*expected))
		})
	})

	PDescribe("#FindByContext", func() {

	})
})
