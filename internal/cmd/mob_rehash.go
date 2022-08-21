package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
	"path"
)

type MobRehashOptions struct {
	*printers.PrinterOptions
}

func NewMobRehashOptions(s printers.IOStreams) *MobRehashOptions {
	return &MobRehashOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
	}
}

func NewCmdMobRehash(s printers.IOStreams) *cobra.Command {
	o := NewMobRehashOptions(s)
	var cmd = &cobra.Command{
		Use:     "rehash",
		Short:   "creates helper git plugin scripts",
		Aliases: []string{"install", "explode"},
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
func (o *MobRehashOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *MobRehashOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobRehashOptions) Run() error {
	e, err := os.Executable()
	if err != nil {
		return err
	}
	eDir := path.Dir(e)

	for plugin, cmd := range ShimMap {
		p := fmt.Sprintf("%s", path.Join(eDir, plugin))
		c := fmt.Sprintf(`
#!/bin/sh
%s "$@"
`, cmd)

		fmt.Println("writing:", p)
		if err := os.WriteFile(p, []byte(c), 0755); err != nil {
			return fmt.Errorf("writing '%s': %v", p, err)
		}
	}

	return nil
}
