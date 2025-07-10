package arrays

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseArray", func() {
	It("should reverse an array of integers", func() {
		input := []int32{1, 2, 3, 4, 5}
		expected := []int32{5, 4, 3, 2, 1}
		Expect(ReverseArray(input)).To(Equal(expected))
	})

	It("should return an empty array when input is empty", func() {
		input := []int32{}
		expected := []int32{}
		Expect(ReverseArray(input)).To(Equal(expected))
	})

	It("should handle a single-element array", func() {
		input := []int32{42}
		expected := []int32{42}
		Expect(ReverseArray(input)).To(Equal(expected))
	})

	It("should handle an array with duplicate elements", func() {
		input := []int32{1, 2, 2, 3}
		expected := []int32{3, 2, 2, 1}
		Expect(ReverseArray(input)).To(Equal(expected))
	})
})
