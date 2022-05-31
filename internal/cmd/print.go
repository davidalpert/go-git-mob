package cmd

import (
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/msg"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/spf13/cobra"
	"strings"
)

type PrintOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
	VersionDetails *version.DetailStruct
	InitialsOnly   bool
}

func NewPrintOptions(ioStreams utils.IOStreams) *PrintOptions {
	return &PrintOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
		VersionDetails: &version.Detail,
		InitialsOnly:   false,
	}
}

func NewCmdPrint(ioStreams utils.IOStreams) *cobra.Command {
	o := NewPrintOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:   "print",
		Short: "show co-authors",
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

	cmd.Flags().BoolVarP(&o.InitialsOnly, "initials-only", "i", false, "show initials only")

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

func (o *PrintOptions) printInitialsOnly(aa []cfg.Author) error {
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

func (o *PrintOptions) printFullMarkup(aa []cfg.Author) error {
	if strings.EqualFold(*o.OutputFormat, "text") {
		return o.WriteStringln(msg.FormatCoAuthorList(aa))
		return nil
	}

	if o.FormatCategory() == "table" {
		o.OutputFormat = utils.StringPointer("json")
	}

	if s, _, err := o.PrinterOptions.FormatOutput(aa); err != nil {
		return err
	} else {
		o.WriteStringln(s)
	}

	return nil
}
