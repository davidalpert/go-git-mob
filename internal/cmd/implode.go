package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
	"path"
)

type ImplodeOptions struct {
	*printers.PrinterOptions
}

func NewImplodeOptions(s printers.IOStreams) *ImplodeOptions {
	return &ImplodeOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
	}
}

func NewCmdImplode(s printers.IOStreams) *cobra.Command {
	o := NewImplodeOptions(s)
	var cmd = &cobra.Command{
		Use:     "implode",
		Short:   "uninstall git-mob (removes helper git plugin scripts and deletes the git-mob binary)",
		Aliases: []string{"uninstall"},
		Args:    cobra.NoArgs,
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
func (o *ImplodeOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *ImplodeOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *ImplodeOptions) Run() error {
	e, err := os.Executable()
	if err != nil {
		return err
	}
	eDir := path.Dir(e)

	for plugin, _ := range ShimMap {
		p := fmt.Sprintf("%s", path.Join(eDir, plugin))

		if _, err = os.Stat(p); err == nil {
			fmt.Fprintln(o.Out, "removing:", p)
			if err := os.Remove(p); err != nil {
				return fmt.Errorf("removing '%s': %v", p, err)
			}
		}
	}

	fmt.Fprintln(o.Out, "removing:", e)
	return os.Remove(e)
}
