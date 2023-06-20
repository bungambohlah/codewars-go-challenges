package rectintorectangles_test

import (
	rectintorectangles "codewars/solutions/rect_into_rectangles"
	"sort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RectIntoRectangles Suite")
}

var _ = Describe("RectIntoRectangles", func() {
	It("Should handle 8 basic tests", func() {
		example_inputs := [...][2]int{{13, 5}, {22, 6}, {8, 5}, {20, 8}, {4, 3}, {7, 5}, {6, 4}, {15, 4}}
		example_tests := [...][]string{
			{"(10*5)", "(13*5)", "(2*1)", "(3*2)", "(5*3)", "(8*5)"},
			{"(10*6)", "(12*6)", "(12*6)", "(16*6)", "(18*6)", "(22*6)", "(4*2)", "(6*4)"},
			{"(2*1)", "(3*2)", "(5*3)", "(8*5)"},
			{"(12*8)", "(16*8)", "(20*8)", "(8*4)"},
			{"(2*1)", "(2*1)", "(3*1)", "(4*3)"},
			{"(2*1)", "(3*2)", "(4*2)", "(5*2)", "(7*5)"},
			{"(4*2)", "(6*4)"},
			{"(11*4)", "(12*4)", "(15*4)", "(2*1)", "(2*1)", "(3*1)", "(4*3)", "(7*4)", "(8*4)", "(8*4)"},
		}

		for i, v := range example_inputs {
			user := rectintorectangles.RectIntoRects(v[0], v[1])
			sort.Strings(user)
			Expect(user).To(Equal(example_tests[i]))
		}
	})
})
