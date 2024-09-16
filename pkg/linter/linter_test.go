package linter_test

import (
	"github.com/illume/supply-chain-linters/pkg/linter"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linter", func() {
	Context("when running the linter", func() {
		It("should not return an error for a valid path", func() {
			err := linter.Run("./")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
