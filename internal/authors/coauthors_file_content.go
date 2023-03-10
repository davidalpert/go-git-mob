package authors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CoAuthorsFileContent struct {
	CoAuthorsByInitial map[string]Author `json:"coauthors"`
}

func ReadCoAuthorsContent() (CoAuthorsFileContent, error) {
	return ReadCoAuthorsContentFromFilePath(CoAuthorsFilePath)
}

func ReadCoAuthorsContentFromFilePath(filePath string) (CoAuthorsFileContent, error) {
	if err := EnsureCoauthorsFileExists(filePath); err != nil {
		return CoAuthorsFileContent{}, nil
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		return CoAuthorsFileContent{}, err
	}
	return ReadCoAuthorsContentFromBytes(b)
}

func ReadCoAuthorsContentFromBytes(b []byte) (CoAuthorsFileContent, error) {
	var c CoAuthorsFileContent
	if err := checkJsonBytesForDuplicateCoauthors(b); err != nil {
		return c, err
	}
	err := json.Unmarshal(b, &c)
	return c, err
}

func (f CoAuthorsFileContent) LookupByInitials(initials ...string) []Author {
	parts := make([]Author, 0)
	for _, i := range initials {
		for initial, author := range f.CoAuthorsByInitial {
			if strings.EqualFold(i, initial) {
				parts = append(parts, author)
				continue
			}
		}
	}

	return parts
}

func (f CoAuthorsFileContent) LookupByEmail(email ...string) []Author {
	parts := make([]Author, 0)
	for _, e := range email {
		for _, author := range f.CoAuthorsByInitial {
			if strings.EqualFold(e, author.Email) {
				parts = append(parts, author)
				continue
			}
		}
	}

	return parts
}

func (f CoAuthorsFileContent) FindAndFormatAsList(initials ...string) []string {
	aa := f.LookupByInitials(initials...)
	result := make([]string, len(aa))

	for i, a := range aa {
		result[i] = a.String()
	}

	return result
}

func (f CoAuthorsFileContent) FindInitialsFromCoAuthorStrings(ss ...string) []string {
	result := make([]string, 0)

	for _, s := range ss {
		a := MustParseOne(s)
		for k, v := range f.CoAuthorsByInitial {
			if strings.EqualFold(a.Email, v.Email) {
				result = append(result, k)
			}
		}
	}

	return result
}

func EnsureCoauthorsFileExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cc := CoAuthorsFileContent{
			CoAuthorsByInitial: make(map[string]Author, 0),
		}
		return WriteCoauthorsContent(cc)
	}
	return nil
}

func WriteCoauthorsContent(cc CoAuthorsFileContent) error {
	return WriteCoauthorsContentToFilePath(CoAuthorsFilePath, cc)
}

func WriteCoauthorsContentToFilePath(path string, cc CoAuthorsFileContent) error {
	b, err := json.MarshalIndent(cc, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, os.ModePerm)
}

func checkJsonBytesForDuplicateCoauthors(b []byte) error {
	return check(json.NewDecoder(bytes.NewReader(b)), nil)
}

func check(d *json.Decoder, path []string) error {
	// Get the next token
	t, err := d.Token()
	if err != nil {
		return err
	}

	// Is it a delimiter?
	delim, ok := t.(json.Delim)
	// No, nothing more to check.
	if !ok {
		// scalar type, nothing to do
		return nil
	}

	switch delim {
	case '{':
		for d.More() {

			// Get field key.
			t, err := d.Token()
			if err != nil {
				return err
			}
			key := t.(string)

			if key == "coauthors" {
				return checkCoauthorsObjectForDuplicatesByInitials(d, append(path, key))
			}

			// Check value.
			if err := check(d, append(path, key)); err != nil {
				return err
			}
		}
		// consume trailing }
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := check(d, append(path, strconv.Itoa(i))); err != nil {
				return err
			}
			i++
		}
		// consume trailing ]
		if _, err := d.Token(); err != nil {
			return err
		}

	}
	return nil
}

type DuplicateInitialsError struct {
	DuplicateAuthorsByInitial map[string][]Author
	Message                   string
}

func (e DuplicateInitialsError) Error() string {
	return e.Message
}

func NewDuplicateInitialsError(duplicates map[string][]Author) DuplicateInitialsError {
	return DuplicateInitialsError{
		DuplicateAuthorsByInitial: duplicates,
		Message:                   "duplicate coauthor initials found",
	}
}

func checkCoauthorsObjectForDuplicatesByInitials(d *json.Decoder, path []string) error {
	parsedCoauthors := make(map[string][]Author, 0)
	duplicates := make(map[string][]Author, 0)

	// consume the opening '{'
	t, err := d.Token()
	if err != nil {
		return err
	}
	delim, ok := t.(json.Delim)
	if !(ok && delim == '{') {
		return fmt.Errorf("expected an opening { after \"coauthors\"")
	}

	for d.More() {
		// get an initials key
		t, err = d.Token()
		if err != nil {
			return err
		}
		key := t.(string)

		if _, ok = parsedCoauthors[key]; !ok {
			parsedCoauthors[key] = make([]Author, 0)
		}

		var a Author
		if err2 := d.Decode(&a); err2 != nil {
			return err2
		}
		//fmt.Printf("found: %#v: %#v\n", key, a)
		parsedCoauthors[key] = append(parsedCoauthors[key], a)
	}

	// consume the closing '}'
	t, err = d.Token()
	if err != nil {
		return err
	}
	delim, ok = t.(json.Delim)
	if !(ok && delim == '}') {
		return fmt.Errorf("expected an closing } at the end of  \"coauthors\"")
	}

	var foundDuplicates = false
	for k, aa := range parsedCoauthors {
		if len(aa) > 1 {
			foundDuplicates = true
			duplicates[k] = aa
		}
	}

	if foundDuplicates {
		return NewDuplicateInitialsError(duplicates)
	}
	return nil
}
