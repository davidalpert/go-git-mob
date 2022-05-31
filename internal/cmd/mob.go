package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/authors"
	"github.com/davidalpert/go-git-mob/internal/cfg"
	"github.com/davidalpert/go-git-mob/internal/cmd/utils"
	"github.com/spf13/cobra"
	"strings"
)

type MobOptions struct {
	*utils.PrinterOptions
	utils.IOStreams
	Initials []string
}

func NewMobOptions(ioStreams utils.IOStreams) *MobOptions {
	return &MobOptions{
		IOStreams:      ioStreams,
		PrinterOptions: utils.NewPrinterOptions().WithDefaultOutput("text"),
	}
}

func NewCmdMob(ioStreams utils.IOStreams) *cobra.Command {
	o := NewMobOptions(ioStreams)
	var cmd = &cobra.Command{
		Use:   "mob",
		Short: "configure co-authors",
		Args:  cobra.MinimumNArgs(1),
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
func (o *MobOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Initials = args
	return nil
}

// Validate the options
func (o *MobOptions) Validate() error {
	if len(o.Initials) < 1 {
		return fmt.Errorf("must supply at least one co-author")
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *MobOptions) Run() error {
	all, err := cfg.ReadAllCoAuthorsFromFile()
	if err != nil {
		return err
	}

	coauthors := make([]authors.Author, len(o.Initials))
	for i, initial := range o.Initials {
		for ii, a := range all {
			if strings.EqualFold(initial, ii) {
				coauthors[i] = a
				break
			}
		}
	}

	if err = cfg.ResetMob(); err != nil {
		return nil
	}
	if err = cfg.AddCoAuthors(coauthors...); err != nil {
		return err
	}

	me, err := cfg.GetUser()
	if err != nil {
		return err
	}

	parts := make([]string, len(o.Initials))
	meTag := fmt.Sprintf("%s <%s>", me.Name, me.Email)
	for i, initial := range o.Initials {
		for ii, a := range all {
			if strings.EqualFold(initial, ii) {
				parts[i] = fmt.Sprintf("%s <%s>", a.Name, a.Email)
				break
			}
		}
	}

	o.WriteStringln(strings.Join(append([]string{meTag}, parts...), "\n"))

	return nil
}
