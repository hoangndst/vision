package cmd

import (
	"os"

	"k8s.io/kubectl/pkg/util/i18n"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/kubectl/pkg/cmd/options"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/hoangndst/vision/cmd/server"
)

type VisionOptions struct {
	Arguments []string

	genericiooptions.IOStreams
}

// NewDefaultVisionCommand creates the `visionctl` command with default arguments
func NewDefaultVisionCommand() *cobra.Command {
	return NewDefaultVisionCommandWithArgs(VisionOptions{
		Arguments: os.Args,
		IOStreams: genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr},
	})
}

// NewDefaultVisionCommandWithArgs creates the `visionctl` command with arguments
func NewDefaultVisionCommandWithArgs(o VisionOptions) *cobra.Command {
	cmd := NewVisionCmd(o)

	if len(o.Arguments) > 1 {
		cmdPathPieces := o.Arguments[1:]
		if _, _, err := cmd.Find(cmdPathPieces); err == nil {
			// sub command exist
			return cmd
		}
	}

	return cmd
}

func NewVisionCmd(o VisionOptions) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "kusion",
		Short:         i18n.T(`Vision is a tool for all tech stuff services`),
		SilenceErrors: true,
		Run:           runHelp,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			if err := flushProfiling(); err != nil {
				return err
			}
			return nil
		},
	}

	// From this point and forward we get warnings on flags that contain "_" separators
	rootCmd.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

	flags := rootCmd.PersistentFlags()

	addProfilingFlags(flags)

	groups := templates.CommandGroups{
		templates.CommandGroup{
			Message: "Server Commands:",
			Commands: []*cobra.Command{
				server.NewCmdServer(),
			},
		},
	}
	groups.Add(rootCmd)

	filters := []string{"options"}

	templates.ActsAsRootCommand(rootCmd, filters, groups...)
	rootCmd.AddCommand(options.NewCmdOptions(o.IOStreams.Out))
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
