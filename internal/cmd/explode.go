package cmd

import (
	"fmt"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
	"path"
)

type ExplodeOptions struct {
	*printers.PrinterOptions
}

func NewExplodeOptions(s printers.IOStreams) *ExplodeOptions {
	return &ExplodeOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
	}
}

func NewCmdExplode(s printers.IOStreams) *cobra.Command {
	o := NewExplodeOptions(s)
	var cmd = &cobra.Command{
		Use:     "explode",
		Short:   "creates helper git plugin scripts",
		Aliases: []string{"install", "rehash"},
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
func (o *ExplodeOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// Validate the options
func (o *ExplodeOptions) Validate() error {
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *ExplodeOptions) Run() error {
	e, err := os.Executable()
	if err != nil {
		return err
	}
	eDir := path.Dir(e)

	for plugin, cmd := range GitMobPluginMap {
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

var GitMobPluginMap = map[string]string{
	"git-mob-print":         "git-mob print",
	"git-mob-version":       "git-mob version",
	"git-solo":              "git-mob solo",
	"git-suggest-coauthors": "git-mob coauthors suggest",
}
