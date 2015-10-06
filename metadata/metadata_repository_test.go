package metadata_test

import (
	//"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fdsolutions/logan/fixtures"
	. "github.com/fdsolutions/logan/metadata"
)

var _ = Describe("MetadataRepository", func() {
	var (
		r             Repository
		s, emptyStore Store
	)

	BeforeEach(func() {
		s, _ = NewFileStore(fixtures.ExistingPath)
		emptyStore, _ = NewFileStore(fixtures.EmptyFilePath)
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
		It("Should return no metatdata entry for the given goal from an empty store", func() {
			r.SetStore(emptyStore)
			entry, _ := r.FindByGoal(TestGoal)
			Expect(entry).To(BeZero())
		})

		It("Should return the metatdata entry of the given goal.", func() {
			entry, _ := r.FindByGoal(TestGoal)
			expected := FromGoal(TestGoal)
			Expect(entry).To(Equal(*expected))
		})
	})

	Describe("#FindByContext", func() {
		Context("With no given context", func() {
			Context("From empty store", func() {
				It("Should return no metadata", func() {
					r.SetStore(emptyStore)
					entries := r.FindByContext(TestNoContext)
					Expect(entries).To(BeEmpty())
				})
			})

			Context("From a store with data", func() {
				It("Should return all metatdata entries with with no context set", func() {
					entries := r.FindByContext(TestNoContext)
					Expect(entries).NotTo(BeEmpty())
					for _, entry := range entries {
						Expect(TestGoalsWithNoContext).To(ContainElement(entry.Goal))
					}
				})
			})
		})

		Context("With a given context", func() {
			It("Sould return no metadata when the context appears no where", func() {
				entries := r.FindByContext(TestUnkownContext)
				Expect(entries).To(BeEmpty())
			})

			It("Sould return all metadata entries related to the given context", func() {
				entries := r.FindByContext(TestContext)
				Expect(entries).NotTo(BeEmpty())
				for _, entry := range entries {
					Expect(TestGoalsWithContext).To(ContainElement(entry.Goal))
				}
			})
		})
	})
})
