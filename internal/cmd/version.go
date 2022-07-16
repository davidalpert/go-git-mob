package cmd

import (
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/version"
	"github.com/spf13/cobra"
	"strings"
)

type VersionOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
	VersionDetails *version.DetailStruct
}

func NewVersionOptions(ioStreams utils.IOStreams) *VersionOptions {
	return &VersionOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
		VersionDetails: &version.Detail,
	}
}

func NewCmdVersion(ioStreams utils.IOStreams) *cobra.Command {
	o := NewVersionOptions(ioStreams)
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

	o.PrinterOptions.AddPrinterFlags(cmd)

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
	if strings.EqualFold(*o.OutputFormat, "text") {
		if s, _, err := o.FormatOutput(o.VersionDetails); err != nil {
			return err
		} else {
			return o.WriteStringf("%s %s\n", o.VersionDetails.AppName, s)
		}
	}
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.OutputFormat = utils.StringPointer("json")
	}

	return o.IOStreams.WriteOutput(o.VersionDetails, o.PrinterOptions)
}
