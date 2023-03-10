package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitCommands"
	"github.com/davidalpert/go-git-mob/internal/revParse"
	"github.com/davidalpert/go-printers/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type CoauthorsSuggestOptions struct {
	*printers.PrinterOptions
	IncludeAll                 bool
	CurrentCoAuthorsByInitials map[string]authors.Author
}

func NewCoauthorsSuggestOptions(s printers.IOStreams) *CoauthorsSuggestOptions {
	return &CoauthorsSuggestOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter().WithDefaultOutput("text"),
		IncludeAll:     false,
	}
}

func NewCmdCoauthorsSuggest(s printers.IOStreams) *cobra.Command {
	o := NewCoauthorsSuggestOptions(s)
	var cmd = &cobra.Command{
		Use:   "suggest",
		Short: "suggest some co-authors to add based on existing committers to your current repo",
		Args:  cobra.NoArgs,
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

	cmd.Flags().BoolVarP(&o.IncludeAll, "all", "a", false, "include all (committers in your coauthors file and committers who are missing")

	return cmd
}

// Complete the options
func (o *CoauthorsSuggestOptions) Complete(cmd *cobra.Command, args []string) error {
	if allCoAuthorsByInitials, err := completeExistingCoauthorsByInitial(o.PrinterOptions); err != nil {
		return err
	} else {
		o.CurrentCoAuthorsByInitials = allCoAuthorsByInitials
	}
	return nil
}

// Validate the options
func (o *CoauthorsSuggestOptions) Validate() error {
	if !revParse.InsideWorkTree() {
		return fmt.Errorf("not a git repository; suggesting co-authors requires a commit history")
	}

	return o.PrinterOptions.Validate()
}

func (o *CoauthorsSuggestOptions) lookupExistingCoauthorInitialsByEmail(e string) (string, bool) {
	for ii, a := range o.CurrentCoAuthorsByInitials {
		if strings.EqualFold(e, a.Email) {
			return ii, true
		}
	}
	return "", false
}

// Run the command
func (o *CoauthorsSuggestOptions) Run() error {
	aa, err := gitCommands.ShortLogAuthorSummary()
	if err != nil || len(aa) == 0 {
		return fmt.Errorf("unable to find existing authors in this repository")
	}

	foundInitials := make([]string, 0)
	suggestedInitials := make([]string, 0)
	anonymousSuggestions := make([]string, 0)
	if o.IncludeAll {
		for ii, _ := range aa {
			suggestedInitials = append(suggestedInitials, ii)
		}
	} else {
		for ii, _ := range aa {
			if existingInitials, found := o.lookupExistingCoauthorInitialsByEmail(aa[ii].Email); found {
				foundInitials = append(foundInitials, existingInitials)
			} else if aa[ii].LooksAnonymous() {
				anonymousSuggestions = append(anonymousSuggestions, ii)
			} else {
				suggestedInitials = append(suggestedInitials, ii)
			}
		}
	}

	sort.Strings(foundInitials)
	sort.Strings(suggestedInitials)
	sort.Strings(anonymousSuggestions)

	if o.FormatCategory() == "text" {
		if len(foundInitials) > 0 {
			_, _ = fmt.Fprintf(o.Out, "The following authors from your coauthors file have contributed to this repository:\n\n")
			for _, ii := range foundInitials {
				a := o.CurrentCoAuthorsByInitials[ii]
				_, _ = fmt.Fprintf(o.Out, "- %s \"%s\" %s\n", ii, a.Name, a.Email)
			}
			_, _ = fmt.Fprintf(o.Out, "\n")
		}

		if len(suggestedInitials) == 0 {
			_, _ = fmt.Fprintf(o.Out, ":tada: You already know all the coauthors who have contributed to this repository!\n")
		} else {
			_, _ = fmt.Fprintf(o.Out, "Here are some suggestions for coauthors based on existing authors of this repository:\n\n")
			for _, ii := range suggestedInitials {
				a := aa[ii]
				_, _ = fmt.Fprintf(o.Out, "git add-coauthor %s \"%s\" %s\n", ii, a.Name, a.Email)
			}
			if len(anonymousSuggestions) > 0 {
				_, _ = fmt.Fprintf(o.Out, "\n")
			}
			for _, ii := range anonymousSuggestions {
				a := aa[ii]
				_, _ = fmt.Fprintf(o.Out, "git add-coauthor %s \"%s\" %s\n", ii, a.Name, a.Email)
			}
			_, _ = fmt.Fprintf(o.Out, "\nPaste any line above.\n")
		}
		return nil
	}

	output := make([]authors.AuthorWithInitials, len(suggestedInitials))
	for i, ii := range suggestedInitials {
		a := aa[ii]
		output[i] = authors.AuthorWithInitials{
			Initials: ii,
			Name:     a.Name,
			Email:    a.Email,
		}
	}

	return o.WithTableWriter("suggested co-authors", func(table *tablewriter.Table) {
		table.SetHeader([]string{"Initials", "Name", "Email"})
		for _, ii := range suggestedInitials {
			a := aa[ii]
			table.Append([]string{ii, a.Name, a.Email})
		}

	}).WriteOutput(output)
}
