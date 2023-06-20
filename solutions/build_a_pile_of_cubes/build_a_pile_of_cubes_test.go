package buildapileofcubes_test

import (
	buildapileofcubes "codewars/solutions/build_a_pile_of_cubes"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuildAPileOfCubes Suite")
}

func dotest(n int, exp int) {
	var ans = buildapileofcubes.FindNb(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(4183059834009, 2022)
		dotest(24723578342962, -1)
	})
})
