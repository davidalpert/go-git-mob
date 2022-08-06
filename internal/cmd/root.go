package cmd

import (
	"fmt"
	"github.com/davidalpert/go-git-mob/internal/diagnostics"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"os"
)

// cfgFile is an optional path to a configuration file used to initialize viper
var cfgFile string

// Execute builds the default root command and invokes it with os.Args
func Execute() {
	s := printers.DefaultOSStreams()
	// configure the logger here in the outer scope so that we can defer
	// any cleanup such as writing/flushing the stream
	logCleanupFn := diagnostics.ConfigureLogger(s)
	defer logCleanupFn()

	rootCmd := NewRootCmd(s)

	rootCmd.SetArgs(os.Args[1:]) // without program

	// look for matching subcommand
	var cmdFound bool
	for _, a := range rootCmd.Commands() {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			} else {
				for _, alias := range a.Aliases {
					if alias == b {
						cmdFound = true
						break
					}
				}
			}
		}
	}
	if cmdFound == false {
		// found no matching subcommand; run the default mob command
		args := append([]string{"mob"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(s.ErrOut, err)
		os.Exit(1)
	}
}

// NewRootCmd creates the 'root' command and configures it's nested children
func NewRootCmd(s printers.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "git",
		Short:         "A git plugin to help manage git coauthors.",
		Long:          "NOTE: this root command is not intended to be run by itself",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   false,
			DisableNoDescFlag:   false,
			DisableDescriptions: false,
			HiddenDefaultCmd:    true,
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
		//  RunE: func(cmd *cobra.Command, args []string) error { },
		Aliases: []string{},
	}

	// Register subcommands
	rootCmd.AddCommand(NewCmdMob(s))

	//rootCmd.PersistentFlags().BoolP("verbose", "vv", false, "enable verbose output")

	return rootCmd
}

func init() {
}
