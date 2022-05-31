package authors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authors Suite")
}
