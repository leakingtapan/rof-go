package rof

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("object factory", func() {
	It("should", func() {
		var a int

		err := Create(&a)
		Expect(err).To(BeNil())

		test(&a)
	})
})
