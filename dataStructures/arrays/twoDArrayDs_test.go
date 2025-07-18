package arrays

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HourglassSum", func() {
	It("should handle integer elements", func() {
		input := [][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		expected := int32(28)
		Expect(HourglassSum(input)).To(Equal(expected))
	})

	It("should handle all positive elements", func() {
		input := [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		expected := int32(19)
		Expect(HourglassSum(input)).To(Equal(expected))
	})

	It("should handle all negative elements", func() {
		input := [][]int32{
			{-1, -1, 0, -9, -2, -2},
			{-2, -1, -6, -8, -2, -5},
			{-1, -1, -1, -2, -3, -4},
			{-1, -9, -2, -4, -4, -5},
			{-7, -3, -3, -2, -9, -9},
			{-1, -3, -1, -2, -4, -5},
		}
		expected := int32(-6)
		Expect(HourglassSum(input)).To(Equal(expected))
	})

})
