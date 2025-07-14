package warmup

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SimpleArraySum", func() {
	It("should handle positive numbers", func() {
		input := []int32{1, 2, 3, 4, 5}
		expected := int32(15)
		Expect(SimpleArraySum(input)).To(Equal(expected))
	})
	It("should handle negative numbers", func() {
		input := []int32{-1, -2, -3, -4, -5}
		expected := int32(-15)
		Expect(SimpleArraySum(input)).To(Equal(expected))
	})
})
