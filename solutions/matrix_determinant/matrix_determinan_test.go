package matrixdeterminant_test

// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

import (
	matrixdeterminant "codewars/solutions/matrix_determinant"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MatrixDeterminant Suite")
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(matrixdeterminant.Determinant([][]int{{1}})).To(Equal(1))
		Expect(matrixdeterminant.Determinant([][]int{{1, 3}, {2, 5}})).To(Equal(-1))
		Expect(matrixdeterminant.Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})).To(Equal(-20))
	})
})
