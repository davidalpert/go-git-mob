package msg

import (
	"bytes"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"regexp"
	"sort"
	"strings"
)

var (
	empty = []byte("")
	space = []byte(" ")
	nl    = []byte("\n")
)

// AppendCoauthorMarkup appends Co-Authored-By markup to a commit message
func AppendCoauthorMarkup(newCoauthors []authors.Author, msgBytes []byte) ([]byte, error) {
	re := regexp.MustCompile(`(?im)^co-authored-by: ([^<]+?)\s+<([^>]+)>`)
	existingCoauthorTags := re.FindAllStringSubmatch(string(msgBytes), -1)
	coauthorsByEmail := make(map[string]authors.Author, 0)
	coauthorEmails := make([]string, 0)
	for _, capture := range existingCoauthorTags {
		a := authors.Author{
			Name:  capture[1],
			Email: capture[2],
		}
		if _, found := coauthorsByEmail[a.Email]; !found {
			coauthorsByEmail[a.Email] = a
			coauthorEmails = append(coauthorEmails, a.Email)
		}
	}
	cleanedB := bytes.TrimSpace(re.ReplaceAll(msgBytes, empty))

	// add in new ones
	for _, a := range newCoauthors {
		if _, found := coauthorsByEmail[a.Email]; !found {
			coauthorsByEmail[a.Email] = a
			coauthorEmails = append(coauthorEmails, a.Email)
		}
	}

	sort.Strings(coauthorEmails)

	coAuthorBytes := make([]byte, 0)
	for _, e := range coauthorEmails {
		coAuthorBytes = append(coAuthorBytes, bytes.Join([][]byte{
			nl, coauthorsByEmail[e].CoauthorTagBytes(),
		}, empty)...)
	}
	coauthorsB := bytes.TrimSpace(coAuthorBytes)

	updated := make([]byte, 0)
	if commentPos := strings.Index(string(cleanedB), "# "); commentPos > -1 {
		gitMessage := bytes.TrimSpace(cleanedB[0:commentPos])
		gitComments := cleanedB[commentPos:]
		updated = append(updated, bytes.Join([][]byte{
			gitMessage, nl,
			nl,
			coauthorsB, nl,
			nl,
			gitComments, nl,
		}, empty)...)
	} else if len(coauthorsB) == 0 {
		return msgBytes, nil
	} else {
		updated = append(updated, bytes.Join([][]byte{
			cleanedB, nl,
			nl,
			coauthorsB, nl,
			nl,
		}, empty)...)
	}
	return updated, nil
}
