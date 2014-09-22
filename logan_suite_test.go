package logan_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLogan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logan Suite")
}
