package warmup

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompareTheTriplets", func() {
	It("should handle the case when A wins", func() {
		input1 := []int32{17, 28, 30}
		input2 := []int32{99, 16, 8}
		output := []int32{2, 1}
		Expect(CompareTheTriplets(input1, input2)).To(Equal(output))
	})
	It("should handle the case when B wins", func() {
		input1 := []int32{17, 9, 30}
		input2 := []int32{99, 16, 8}
		output := []int32{1, 2}
		Expect(CompareTheTriplets(input1, input2)).To(Equal(output))
	})
	It("should handle the case when it's a draw", func() {
		input1 := []int32{17, 16, 30}
		input2 := []int32{99, 16, 8}
		output := []int32{1, 1}
		Expect(CompareTheTriplets(input1, input2)).To(Equal(output))
	})
})
