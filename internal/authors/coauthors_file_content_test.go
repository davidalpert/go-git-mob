package authors_test

import (
	"github.com/davidalpert/go-git-mob/internal/authors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("authors", func() {
	Context("a valid coauthors file", func() {
		validJsonString := `
{
  "coauthors": {
    "jd": {
      "name": "Jane Doe",
      "email": "jane@findmypast.com"
    },
    "fb": {
      "name": "Frances Bar",
      "email": "frances-bar@findmypast.com"
    }
  }
}`

		It("read contents from .git-coauthors json", func() {
			_, err := authors.ReadCoAuthorsContentFromBytes([]byte(validJsonString))
			Expect(err).To(BeNil())
		})

		It(`find and format "jd" and "fb" to an array of co-authors`, func() {
			c, err := authors.ReadCoAuthorsContentFromBytes([]byte(validJsonString))
			Expect(err).To(BeNil())

			Expect(c.FindAndFormatAsList("jd", "fb")).To(Equal([]string{
				"Jane Doe <jane@findmypast.com>",
				"Frances Bar <frances-bar@findmypast.com>",
			}))
		})

		It(`find and format "jd" to an array of one co-author`, func() {
			c, err := authors.ReadCoAuthorsContentFromBytes([]byte(validJsonString))
			Expect(err).To(BeNil())

			Expect(c.FindAndFormatAsList("jd")).To(Equal([]string{
				"Jane Doe <jane@findmypast.com>",
			}))
		})

		It(`silently ignore when initials of author are not found`, func() {
			c, err := authors.ReadCoAuthorsContentFromBytes([]byte(validJsonString))
			Expect(err).To(BeNil())

			Expect(c.FindAndFormatAsList("jd", "xx")).To(Equal([]string{
				"Jane Doe <jane@findmypast.com>",
			}))
		})

		It(`find initials of co-authors`, func() {
			c, err := authors.ReadCoAuthorsContentFromBytes([]byte(validJsonString))
			Expect(err).To(BeNil())

			Expect(c.FindInitialsFromCoAuthorStrings(
				"Jane Doe <jane@findmypast.com>",
			)).To(Equal([]string{
				"jd",
			}))
		})
	})
})
