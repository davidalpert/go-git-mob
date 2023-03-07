package authors

import (
	"fmt"
	"strings"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a Author) String() string {
	return fmt.Sprintf("%s <%s>", a.Name, a.Email)
}

func (a Author) CoauthorTag() string {
	return fmt.Sprintf("Co-Authored-By: %s <%s>", a.Name, a.Email)
}

func (a Author) CoauthorTagBytes() []byte {
	return []byte(a.CoauthorTag())
}

func (a Author) InitialsFromName() string {
	//return name.split(' ').map(word => word[0].toLowerCase()).join('');
	words := strings.Split(a.Name, " ")
	initials := make([]string, len(words))
	for i, w := range words {
		initials[i] = strings.ToLower(w[0:1])
	}

	return strings.Join(initials, "")
}

// MustParseOne parses an author string into an Author and panics if parsing fails
func MustParseOne(s string) Author {
	if a, err := ParseOne(s); err != nil {
		panic(err)
	} else {
		return a
	}
}

// ParseOne parses an author string into an Author
func ParseOne(s string) (Author, error) {
	res := reAuthorString.FindAllStringSubmatch(s, 1)

	if len(res) > 0 {
		return Author{
			Name:  res[0][1],
			Email: res[0][2],
		}, nil
	}

	return Author{}, fmt.Errorf("failed to parse co-author from config option: '%s'", s)
}
