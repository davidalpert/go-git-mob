package cmd

import (
	"fmt"
	"github.com/apex/log"
	"github.com/davidalpert/go-git-mob/internal/diagnostics"
	"github.com/davidalpert/go-git-mob/internal/gitMessage"
	"github.com/davidalpert/go-git-mob/internal/gitMobCommands"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
)

type MobPrepareCommitMsgOptions struct {
	*printers.PrinterOptions

	// 1-3 positional args provided by git
	CommitMessageFile string
	Source            gitMessage.Source // optional
	CommitObject      string            // optional (required when Source is CommitSource)

	RawArgs []string
}

func NewMobPrepareCommitMsgOptions(s printers.IOStreams) *MobPrepareCommitMsgOptions {
	return &MobPrepareCommitMsgOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("text"),
	}
}

func NewCmdMobPrepareCommitMsg(s printers.IOStreams) *cobra.Command {
	o := NewMobPrepareCommitMsgOptions(s)
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

	o.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *MobPrepareCommitMsgOptions) Complete(cmd *cobra.Command, args []string) error {
	o.CommitMessageFile = args[0]
	if len(args) > 1 {
		o.Source = gitMessage.CommitMsgSourceFromString(args[1])
	} else {
		o.Source = gitMessage.EmptySource
	}
	if len(args) > 2 {
		o.CommitObject = args[2]
	}
	o.RawArgs = args
	return nil
}

// Validate the options
func (o *MobPrepareCommitMsgOptions) Validate() error {
	if o.Source == gitMessage.UnknownSource {
		return fmt.Errorf("'%s' is not a recognized message source", o.RawArgs[1])
	}
	if o.Source == gitMessage.CommitSource && o.CommitObject == "" {
		return fmt.Errorf("must provide a commit SHA with a message source of '%s'", gitMessage.CommitSource.String())
	}
	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobPrepareCommitMsgOptions) Run() error {
	lg := diagnostics.Log.WithFields(log.Fields{
		"COMMIT_MSG_FILE": o.CommitMessageFile,
		"COMMIT_SOURCE":   o.Source,
		"COMMIT_SHA":      o.CommitObject,
	})

	fileBytes, err := os.ReadFile(o.CommitMessageFile)
	if os.IsNotExist(err) {
		wrapper := fmt.Errorf("opening '%s': %v", o.CommitMessageFile, err)
		lg.WithError(wrapper).Error("prepare-commit-msg")
		return wrapper
	}
	lg = lg.WithField("COMMIT_MSG", string(fileBytes))

	aa, err := gitMobCommands.GetCoAuthors()
	if err != nil {
		wrapper := fmt.Errorf("reading co-authors: %v", err)
		lg.WithError(wrapper).Error("prepare-commit-msg")
		return wrapper
	}

	if len(aa) == 0 {
		lg.Debug("prepare-commit-msg: no coauthors")
		return nil // nothing to do
	}
	lg = lg.WithField("co-authors", aa)

	updated, coauthoredTags, err := gitMessage.AppendCoauthorMarkup(aa, fileBytes)
	if err != nil {
		lg.WithError(err).Error("prepare-commit-msg")
		return err
	}

	lg = lg.WithField("COMMIT_MSG_AFTER", string(updated))
	lg.Debug("prepare-commit-msg")

	if o.FormatCategory() == "text" || o.FormatCategory() == "table" && len(coauthoredTags) > 0 {
		fmt.Fprintln(o.Out, string(coauthoredTags))
	}

	return os.WriteFile(o.CommitMessageFile, updated, os.ModePerm)
}
