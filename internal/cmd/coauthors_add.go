package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-git-mob/internal/revParse"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"net/mail"
	"strings"
)

type CoauthorsAddOptions struct {
	*printers.PrinterOptions
	AllCoAuthorsByInitials map[string]authors.Author
	Author                 authors.Author
	AuthorInitials         string
}

func NewCoauthorsAddOptions(s printers.IOStreams) *CoauthorsAddOptions {
	return &CoauthorsAddOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdCoauthorsAdd(s printers.IOStreams) *cobra.Command {
	o := NewCoauthorsAddOptions(s)
	var cmd = &cobra.Command{
		Use:   "add <initials> <name> <email_address>",
		Short: "add a new co-authors to the list of co-authors available to git-mob",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *CoauthorsAddOptions) Complete(cmd *cobra.Command, args []string) error {
	o.AuthorInitials = args[0]
	o.Author = authors.Author{
		Name:  args[1],
		Email: args[2],
	}

	if allCoAuthorsByInitials, err := gitConfig.ReadAllCoAuthorsFromFile(); err != nil {
		return err
	} else {
		o.AllCoAuthorsByInitials = allCoAuthorsByInitials
	}

	return nil
}

// Validate the options
func (o *CoauthorsAddOptions) Validate() error {
	if !revParse.InsideWorkTree() {
		return fmt.Errorf("not a git repository")
	}

	if _, err := mail.ParseAddress(o.Author.Email); err != nil {
		return fmt.Errorf("invalid email address: %v", err)
	}

	var foundExisting = ""
	var foundAs = ""
	for ii, a := range o.AllCoAuthorsByInitials {
		if strings.EqualFold(o.Author.Email, a.Email) {
			foundExisting = a.Email
			foundAs = ii
		} else if strings.EqualFold(o.AuthorInitials, ii) {
			foundExisting = ii
			foundAs = a.Email
		}
	}

	if foundExisting != "" {
		return fmt.Errorf("coauthor '%s' already exists in '~/.git-coauthors' as '%s'", foundExisting, foundAs)
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *CoauthorsAddOptions) Run() error {
	//fmt.Fprintf(o.Out, "%#v\n", o.Author)
	o.AllCoAuthorsByInitials[o.AuthorInitials] = o.Author

	return authors.WriteCoauthorsContent(authors.CoAuthorsFileContent{
		CoAuthorsByInitial: o.AllCoAuthorsByInitials,
	})
}
