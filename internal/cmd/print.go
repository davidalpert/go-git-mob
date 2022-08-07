package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/gitMobCommands"
	"strings"

	"github.com/davidalpert/go-git-mob/internal/gitConfig"
	"github.com/davidalpert/go-git-mob/internal/gitMessage"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type PrintOptions struct {
	*printers.PrinterOptions
	printers.IOStreams
	VersionDetails *version.DetailStruct
	InitialsOnly   bool
}

func NewPrintOptions(s printers.IOStreams) *PrintOptions {
	return &PrintOptions{
		IOStreams:      s,
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		VersionDetails: &version.Detail,
		InitialsOnly:   false,
	}
}

func NewCmdPrint(s printers.IOStreams) *cobra.Command {
	o := NewPrintOptions(s)
	var cmd = &cobra.Command{
		Use:   "print",
		Short: "show current co-authors",
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

	cmd.Flags().BoolVarP(&o.InitialsOnly, "initials", "i", false, "show initials only")

	return cmd
}

// Complete the options
func (o *PrintOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *PrintOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *PrintOptions) Run() error {
	aa, err := gitMobCommands.GetCoAuthors()
	if err != nil {
		return err
	}

	if o.InitialsOnly {
		return o.printInitialsOnly(aa)
	}

	return o.printFullMarkup(aa)
}

func (o *PrintOptions) printInitialsOnly(aa []authors.Author) error {
	aaByInitial, err := gitConfig.ReadAllCoAuthorsFromFile()
	if err != nil {
		return err
	}

	parts := make([]string, len(aa))
	for i, a := range aa {
		for initial, author := range aaByInitial {
			if strings.EqualFold(a.Email, author.Email) {
				parts[i] = initial
				break
			}
		}
	}

	_, err = fmt.Fprintln(o.Out, strings.Join(parts, " "))
	return err
}

func (o *PrintOptions) printFullMarkup(aa []authors.Author) error {
	if strings.EqualFold(*o.OutputFormat, "text") {
		_, err := fmt.Fprintln(o.Out, gitMessage.FormatCoAuthorList(aa))
		return err
	}

	if o.FormatCategory() == "table" {
		o.PrinterOptions.WithDefaultOutput("json")
	}

	return o.WriteOutput(aa)
}
