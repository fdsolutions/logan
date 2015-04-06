package usage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUsage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usage Suite")
}
