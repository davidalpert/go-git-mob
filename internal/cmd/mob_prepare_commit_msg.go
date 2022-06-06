package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/davidalpert/go-git-mob/internal/msg"
	"github.com/spf13/cobra"
	"os"
)

type MobPrepareCommitMsgOptions struct {
	*utils.PrinterOptions
	utils.IOStreams

	// 1-3 positional args provided by git
	CommitMessageFile string
	Source            msg.Source // optional
	CommitObject      string     // optional (required when Source is CommitSource)

	RawArgs []string
}

func NewMobPrepareCommitMsgOptions(ioStreams utils.IOStreams) *MobPrepareCommitMsgOptions {
	return &MobPrepareCommitMsgOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
	}
}

func NewCmdMobPrepareCommitMsg(ioStreams utils.IOStreams) *cobra.Command {
	o := NewMobPrepareCommitMsgOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:     "prepare-commit-msg",
		Short:   "edits a message file to include current co-authors",
		Aliases: []string{},
		Args:    cobra.RangeArgs(1, 3),
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
func (o *MobPrepareCommitMsgOptions) Complete(cmd *cobra.Command, args []string) error {
	o.CommitMessageFile = args[0]
	if len(args) > 1 {
		o.Source = msg.CommitMsgSourceFromString(args[1])
	} else {
		o.Source = msg.EmptySource
	}
	if len(args) > 2 {
		o.CommitObject = args[2]
	}
	o.RawArgs = args
	return nil
}

// Validate the options
func (o *MobPrepareCommitMsgOptions) Validate() error {
	if o.Source == msg.UnknownSource {
		return fmt.Errorf("'%s' is not a recognized message source", o.RawArgs[1])
	}
	if o.Source == msg.CommitSource && o.CommitObject == "" {
		return fmt.Errorf("must provide a commit SHA with a message source of '%s'", msg.CommitSource.String())
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobPrepareCommitMsgOptions) Run() error {
	fileBytes, err := os.ReadFile(o.CommitMessageFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("opening '%s': %v", o.CommitMessageFile, err)
	}

	aa, err := cfg.GetCoAuthors()
	if err != nil {
		return fmt.Errorf("reading co-authors: %v", err)
	}

	if len(aa) == 0 {
		return nil // nothing to do
	}

	updated, err := msg.AppendCoauthorMarkup(aa, fileBytes)
	if err != nil {
		return err
	}

	return os.WriteFile(o.CommitMessageFile, updated, os.ModePerm)
}
