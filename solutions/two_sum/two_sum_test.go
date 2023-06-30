package twosum_test

import (
	twosum "codewars/solutions/two_sum"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwoSum Suite")
}

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Expect(twosum.TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
		Expect(twosum.TwoSum([]int{1234, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
		Expect(twosum.TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
	})
})
