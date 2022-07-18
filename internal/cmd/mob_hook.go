package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdMobHooks(ioStreams printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "hooks",
		Aliases: []string{"hook"},
		Short:   "git hook utilities",
		Args:    cobra.NoArgs,
		//RunE: func(cmd *cobra.Command, args []string) error {
		//	if err := o.Complete(cmd, args); err != nil {
		//		return err
		//	}
		//	if err := o.Validate(); err != nil {
		//		return err
		//	}
		//	return o.Run()
		//},
	}

	//o.PrinterOptions.AddPrinterFlags(cmd.Flags())

	//cmd.Flags().BoolVarP(&o.ListOnly, "list", "l", false, "list which co-authors are available")

	cmd.AddCommand(NewCmdMobPrepareCommitMsg(ioStreams))

	return cmd
}
