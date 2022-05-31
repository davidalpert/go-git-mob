package cfg

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CoAuthorsFileContent struct {
	CoAuthorsByInitial map[string]Author `json:"coauthors"`
}
