package romansnumeraldecoder_test

import (
	romansnumeral "codewars/solutions/romans_numeral_decoder"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RomansNumeralDecoder Suite")
}

var _ = Describe("test roman to decimal converter", func() {
	It("should give decimal number from roman", func() {
		Expect(romansnumeral.Decode("XXI")).To(Equal(21))

	})
	It("should give decimal number from roman", func() {
		Expect(romansnumeral.Decode("I")).To(Equal(1))
	})
	It("should give decimal number from roman", func() {
		Expect(romansnumeral.Decode("IV")).To(Equal(4))
	})
	It("should give decimal number from roman", func() {
		Expect(romansnumeral.Decode("MMVIII")).To(Equal(2008))
	})
	It("should give decimal number from roman", func() {
		Expect(romansnumeral.Decode("MDCLXVI")).To(Equal(1666))
	})
})
