package decodethemorsecode_test

import (
	decodethemorsecode "codewars/solutions/decode_the_morse_code"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DecodeTheMorseCode Suite")
}

var _ = Describe("Test Example", func() {
	It("Example from description", func() {
		Expect(decodethemorsecode.DecodeMorse(".... . -.--   .--- ..- -.. .")).To(Equal("HEY JUDE"))
	})
})
