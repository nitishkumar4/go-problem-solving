package warmup

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SolveMeFirst", func() {
	It("should handle positive numbers", func() {
		input1 := uint32(4)
		input2 := uint32(5)
		expected := uint32(9)
		Expect(SolveMeFirst(input1, input2)).To(Equal(expected))
	})
})
