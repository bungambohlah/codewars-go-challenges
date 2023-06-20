package firstnonrepeatingcharacter_test

import (
	firstNonRepeatingCharacter "codewars/solutions/first_non_repeating_character"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCodewarsGoChallenges(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FirstNonRepeating Suite")
}

var _ = Describe("Basic Tests", func() {
	It("should handle simple tests", func() {
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("a")).To(Equal("a"))
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("stress")).To(Equal("t"))
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("moonmen")).To(Equal("e"))
	})
	It("should handle empty strings", func() {
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("")).To(Equal(""))
	})
	It("should handle all repeating strings", func() {
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("abba")).To(Equal(""))
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("aa")).To(Equal(""))
	})
	It("should handle odd characters", func() {
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("~><#~><")).To(Equal("#"))
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("hello world, eh?")).To(Equal("w"))
	})
	It("should handle letter cases", func() {
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("sTreSS")).To(Equal("T"))
		Expect(firstNonRepeatingCharacter.FirstNonRepeating("Go hang a salami, I'm a lasagna hog!")).To(Equal(","))
	})
})
