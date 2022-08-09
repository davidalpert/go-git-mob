package authors

type AuthorWithInitials struct {
	Initials string `json:"initials"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
