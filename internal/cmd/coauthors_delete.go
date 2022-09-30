package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"strings"
)

type CoauthorsDeleteOptions struct {
	*printers.PrinterOptions
	AllCoAuthorsByInitials map[string]authors.Author
	AuthorInitials         string
}

func NewCoauthorsDeleteOptions(s printers.IOStreams) *CoauthorsDeleteOptions {
	return &CoauthorsDeleteOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdCoauthorsDelete(s printers.IOStreams) *cobra.Command {
	o := NewCoauthorsDeleteOptions(s)
	var cmd = &cobra.Command{
		Use:     "delete <initials>",
		Aliases: []string{"rm", "del"},
		Short:   "add a new co-authors to the list of co-authors available to git-mob",
		Args:    cobra.ExactArgs(1),
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
func (o *CoauthorsDeleteOptions) Complete(cmd *cobra.Command, args []string) error {
	o.AuthorInitials = args[0]

	if allCoAuthorsByInitials, err := gitConfig.ReadAllCoAuthorsFromFile(); err != nil {
		return err
	} else {
		o.AllCoAuthorsByInitials = allCoAuthorsByInitials
	}

	return nil
}

// Validate the options
func (o *CoauthorsDeleteOptions) Validate() error {
	var foundExisting = ""
	for ii, _ := range o.AllCoAuthorsByInitials {
		if strings.EqualFold(o.AuthorInitials, ii) {
			foundExisting = ii
		}
	}

	if foundExisting == "" {
		return fmt.Errorf("coauthor '%s' not found in '~/.git-coauthors'; git mob --list prints the available coauthors by initials", foundExisting)
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *CoauthorsDeleteOptions) Run() error {
	delete(o.AllCoAuthorsByInitials, o.AuthorInitials)

	return authors.WriteCoauthorsContent(authors.CoAuthorsFileContent{
		CoAuthorsByInitial: o.AllCoAuthorsByInitials,
	})
}
