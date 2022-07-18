package cmd

import (
	"strings"

	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-printers/v1"
	"github.com/davidalpert/go-git-mob/internal/msg"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/spf13/cobra"
)

type PrintOptions struct {
	*printers.PrinterOptions
	printers.IOStreams
	VersionDetails *version.DetailStruct
	InitialsOnly   bool
}

func NewPrintOptions(ioStreams printers.IOStreams) *PrintOptions {
	return &PrintOptions{
		IOStreams:      ioStreams,
		PrinterOptions: printers.NewPrinterOptions().WithDefaultOutput("text"),
		VersionDetails: &version.Detail,
		InitialsOnly:   false,
	}
}

func NewCmdPrint(ioStreams printers.IOStreams) *cobra.Command {
	o := NewPrintOptions(ioStreams)
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

	o.PrinterOptions.AddPrinterFlags(cmd.Flags())

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
	aa, err := cfg.GetCoAuthors()
	if err != nil {
		return err
	}

	if o.InitialsOnly {
		return o.printInitialsOnly(aa)
	}

	return o.printFullMarkup(aa)
}

func (o *PrintOptions) printInitialsOnly(aa []authors.Author) error {
	aaByInitial, err := cfg.ReadAllCoAuthorsFromFile()
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

	o.WriteStringln(strings.Join(parts, " "))

	return nil
}

func (o *PrintOptions) printFullMarkup(aa []authors.Author) error {
	if strings.EqualFold(*o.OutputFormat, "text") {
		return o.WriteStringln(msg.FormatCoAuthorList(aa))
		return nil
	}

	if o.FormatCategory() == "table" {
		o.OutputFormat = printers.StringPointer("json")
	}

	if s, _, err := o.PrinterOptions.FormatOutput(aa); err != nil {
		return err
	} else {
		o.WriteStringln(s)
	}

	return nil
}
