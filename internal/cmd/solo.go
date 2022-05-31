package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/spf13/cobra"
	"strings"
)

type SoloOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
}

func NewSoloOptions(ioStreams utils.IOStreams) *SoloOptions {
	return &SoloOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
	}
}

func NewCmdSolo(ioStreams utils.IOStreams) *cobra.Command {
	o := NewSoloOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:   "solo",
		Short: "return to solo work (remove co-authors)",
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
func (o *SoloOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *SoloOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *SoloOptions) Run() error {
	if err := cfg.ResetMob(); err != nil {
		return err
	}

	me, err := cfg.GetUser()
	if err != nil {
		return err
	}

	meTag := fmt.Sprintf("%s <%s>", me.Name, me.Email)

	o.WriteStringln(strings.Join(append([]string{meTag}), "\n"))

	return nil
}
