package cmd

import (
	"fmt"
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
		Short: "Show version information",
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
		s := fmt.Sprintf("%s %s - %s", o.VersionDetails.AppName, o.VersionDetails.Version, o.VersionDetails.GitCommit)
		if o.VersionDetails.GitDirty {
			s = fmt.Sprintf("%s [dirty]", s)
		}
		o.WriteStringln(s)
	} else {
		if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
			o.OutputFormat = utils.StringPointer("json")
		}

		s, _, err := o.PrinterOptions.FormatOutput(o.VersionDetails)
		utils.ExitIfErr(err)
		o.WriteStringln(s)
	}

	return nil
}
