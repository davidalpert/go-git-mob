package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"net/mail"
	"strings"
)

type CoauthorsEditOptions struct {
	*printers.PrinterOptions
	AllCoAuthorsByInitials map[string]authors.Author
	AuthorEmail            string
	AuthorInitials         string
	AuthorName             string
}

func NewCoauthorsEditOptions(s printers.IOStreams) *CoauthorsEditOptions {
	return &CoauthorsEditOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdCoauthorsEdit(s printers.IOStreams) *cobra.Command {
	o := NewCoauthorsEditOptions(s)
	var cmd = &cobra.Command{
		Use:   "edit <initials> [--name=<name>] [--email=<email_address>]",
		Short: "edit an existing co-author in the list of co-authors available to git-mob",
		Args:  cobra.ExactArgs(1),
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
	cmd.Flags().StringVarP(&o.AuthorName, "name", "n", "", "new name for the given coauthor")
	cmd.Flags().StringVarP(&o.AuthorEmail, "email", "e", "", "new email for the given coauthor")

	return cmd
}

// Complete the options
func (o *CoauthorsEditOptions) Complete(cmd *cobra.Command, args []string) error {
	o.AuthorInitials = args[0]

	if allCoAuthorsByInitials, err := gitConfig.ReadAllCoAuthorsFromFile(); err != nil {
		return err
	} else {
		o.AllCoAuthorsByInitials = allCoAuthorsByInitials
	}

	return nil
}

// Validate the options
func (o *CoauthorsEditOptions) Validate() error {
	if o.AuthorEmail != "" {
		if _, err := mail.ParseAddress(o.AuthorEmail); err != nil {
			return fmt.Errorf("invalid email address: %v", err)
		}
	}

	var foundExisting = ""
	for ii, _ := range o.AllCoAuthorsByInitials {
		if strings.EqualFold(o.AuthorInitials, ii) {
			foundExisting = ii
		}
	}

	if foundExisting == "" {
		return fmt.Errorf("coauthor '%s' not found in '~/.git-coauthors'", foundExisting)
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *CoauthorsEditOptions) Run() error {
	// extract current coauthor
	a := o.AllCoAuthorsByInitials[o.AuthorInitials]

	// update requested fields
	if o.AuthorName != "" {
		a.Name = o.AuthorName
	}
	if o.AuthorEmail != "" {
		a.Email = o.AuthorEmail
	}

	// save changes
	o.AllCoAuthorsByInitials[o.AuthorInitials] = a

	return authors.WriteCoauthorsContent(authors.CoAuthorsFileContent{
		CoAuthorsByInitial: o.AllCoAuthorsByInitials,
	})
}
