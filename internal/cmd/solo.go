package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/gitMobCommands"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"strings"
)

type SoloOptions struct {
	*printers.PrinterOptions
}

func NewSoloOptions(s printers.IOStreams) *SoloOptions {
	return &SoloOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
	}
}

func NewCmdSolo(s printers.IOStreams) *cobra.Command {
	o := NewSoloOptions(s)
	var cmd = &cobra.Command{
		Use:   "solo",
		Short: "return to solo work (remove co-authors)",
		Args:  cobra.MinimumNArgs(0), // positional args are allowed (by spec) but ignored
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
func (o *SoloOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *SoloOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *SoloOptions) Run() error {
	if err := gitMobCommands.ResetMob(); err != nil {
		return err
	}

	me, err := gitMobCommands.GetGitAuthor()
	if err != nil {
		return err
	}

	meTag := fmt.Sprintf("%s <%s>", me.Name, me.Email)

	_, err = fmt.Fprintln(o.Out, strings.Join(append([]string{meTag}), "\n"))
	return err
}
