package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/revParse"
	"github.com/davidalpert/go-printers/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
)

type CoauthorsSuggestOptions struct {
	*printers.PrinterOptions
	printers.IOStreams
}

func NewCoauthorsSuggestOptions(ioStreams printers.IOStreams) *CoauthorsSuggestOptions {
	return &CoauthorsSuggestOptions{
		IOStreams:      ioStreams,
		PrinterOptions: printers.NewPrinterOptions().WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdCoauthorsSuggest(ioStreams printers.IOStreams) *cobra.Command {
	o := NewCoauthorsSuggestOptions(ioStreams)
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

	o.PrinterOptions.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *CoauthorsSuggestOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *CoauthorsSuggestOptions) Validate() error {
	if !revParse.InsideWorkTree() {
		return fmt.Errorf("not a git repository")
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *CoauthorsSuggestOptions) Run() error {
	aa, err := cfg.ShortLogAuthorSummary()
	if err != nil || len(aa) == 0 {
		return fmt.Errorf("unable to find existing authors in this repository")
	}

	initials := make([]string, len(aa))
	i := 0
	for ii, _ := range aa {
		initials[i] = ii
		i++
	}

	sort.Strings(initials)

	if o.FormatCategory() == "text" {
		if len(initials) > 0 {
			fmt.Printf("Here are some suggestions for coauthors based on existing authors of this repository:\n\n")
			for _, ii := range initials {
				a := aa[ii]
				fmt.Printf("git mob add-coauthor %s %s %s\n", ii, a.Name, a.Email)
			}
			fmt.Println("\nPaste any line above.")
		}
		return nil
	}

	output := make([]authors.AuthorWithInitials, len(initials))
	for i, ii := range initials {
		a := aa[ii]
		output[i] = authors.AuthorWithInitials{
			Initials: ii,
			Name:     a.Name,
			Email:    a.Email,
		}
	}

	return o.WithTableWriter("suggested co-authors", func(table *tablewriter.Table) {
		table.SetHeader([]string{"Initials", "Name", "Email"})
		for _, ii := range initials {
			a := aa[ii]
			table.Append([]string{ii, a.Name, a.Email})
		}

	}).WriteOutput(output)
}
