package cmd

// ShimMap stores the mapping of helper scripts to the equivalent git-mob subcommands; this
// mapping is used to auto-generate shim scripts and enable simpler syntax for calling these subcommands
// as if they were their own git plugins
var ShimMap = map[string]string{
	"git-add-coauthor":      "git-mob coauthors add",
	"git-delete-coauthor":   "git-mob coauthors delete",
	"git-edit-coauthor":     "git-mob coauthors edit",
	"git-mob-print":         "git-mob print",
	"git-mob-version":       "git-mob version",
	"git-solo":              "git-mob solo",
	"git-suggest-coauthors": "git-mob coauthors suggest",
}
