package cmd

import (
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type VersionOptions struct {
	*printers.PrinterOptions
	VersionDetails *version.DetailStruct
}

func NewVersionOptions(s printers.IOStreams) *VersionOptions {
	return &VersionOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
		VersionDetails: &version.Detail,
	}
}

func NewCmdVersion(s printers.IOStreams) *cobra.Command {
	o := NewVersionOptions(s)
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "show version information",
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

	return cmd
}

// Complete the options
func (o *VersionOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *VersionOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *VersionOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	return o.WriteOutput(o.VersionDetails)
}
