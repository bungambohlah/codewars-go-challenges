package eureka_test

import (
	eureka "codewars/solutions/eureka"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Eureka Suite")
}

func dotest(a, b uint64, expected []uint64) {
	actual := eureka.SumDigPow(a, b)
	if len(expected) == 0 {
		Expect(actual).To(BeEmpty(), "With a = %d, b = %d", a, b)
	} else {
		Expect(actual).To(Equal(expected), "With a = %d, b = %d", a, b)
	}
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest(1, 10, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		dotest(1, 100, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89})
		dotest(10, 89, []uint64{89})
		dotest(10, 100, []uint64{89})
		dotest(90, 100, nil)
		dotest(89, 135, []uint64{89, 135})
	})
})
