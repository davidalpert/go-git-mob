package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdCoauthors(s printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "coauthors",
		Short:   "co-author related subcommands",
		Aliases: []string{"co-authors", "co-author", "coauthor", "c"},
		Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdCoauthorsSuggest(s))
	cmd.AddCommand(NewCmdCoauthorsAdd(s))

	return cmd
}
