package metadata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/logan/metadata"
)

var _ = Describe("MetadataRepository", func() {
	var (
		r Repository
		s Store
	)

	BeforeEach(func() {
		s, _ = NewFileStore(ExistingPath)
		r = NewRepositoryFromStore(s)

	})

	PDescribe("#FindAll", func() {
		It("should return all metadata entries.", func() {
			entries := r.FindAll()
			Expect(entries).NotTo(BeEmpty())
		})
	})
	PDescribe("#FindByGoal", func() {

	})
	PDescribe("#FindByContext", func() {

	})
})
