package authors

import (
	"encoding/json"
	"os"
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
		b, err := json.Marshal(cc)
		if err != nil {
			return err
		}

		return os.WriteFile(path, b, os.ModePerm)
	}
	return nil
}
