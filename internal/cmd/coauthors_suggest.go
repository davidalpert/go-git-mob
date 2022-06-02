package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
)

type CoauthorsSuggestOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
}

func NewCoauthorsSuggestOptions(ioStreams utils.IOStreams) *CoauthorsSuggestOptions {
	return &CoauthorsSuggestOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultTableWriter().WithDefaultOutput("text"),
	}
}

func NewCmdCoauthorsSuggest(ioStreams utils.IOStreams) *cobra.Command {
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

	o.PrinterOptions.AddPrinterFlags(cmd)

	return cmd
}

// Complete the options
func (o *CoauthorsSuggestOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *CoauthorsSuggestOptions) Validate() error {
	if !cfg.InsideWorkTree() {
		return fmt.Errorf("not a git repository")
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *CoauthorsSuggestOptions) Run() error {
	aa, err := cfg.ShortLogAuthorSummary()
	if err != nil {
		return err
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
			fmt.Println("Here are some suggestions for coauthors based on existing authors of this repository:")
			for _, ii := range initials {
				a := aa[ii]
				fmt.Printf("git mob add-coauthor %s %s %s\n", ii, a.Name, a.Email)
			}
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
