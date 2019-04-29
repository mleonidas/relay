package cmd

import (
	"github.com/puppetlabs/nebula/pkg/cmd/login"
	"github.com/puppetlabs/nebula/pkg/cmd/version"
	"github.com/puppetlabs/nebula/pkg/cmd/workflow"
	"github.com/puppetlabs/nebula/pkg/config/runtimefactory"
	"github.com/puppetlabs/nebula/pkg/loader"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	r, err := runtimefactory.NewRuntimeFactory()
	if err != nil {
		return nil, err
	}

	c := &cobra.Command{
		Use:   "nebula",
		Short: "Nebula workflow management cli",
		// don't show usage text for every error
		SilenceUsage: true,
		// we want to be able to handle our own errors for display; this allows us to use
		// the CLI display mechanism for errawr.
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fp, err := cmd.Flags().GetString("filepath")
			if err != nil {
				return err
			}

			if fp != "" {
				r.SetPlanLoader(loader.NewFilepathLoader(fp))
			}

			return nil
		},
	}

	c.PersistentFlags().StringP("filepath", "f", "", "optional path to a workflow.yaml")

	c.AddCommand(login.NewCommand(r))
	c.AddCommand(workflow.NewCommand(r))
	c.AddCommand(version.NewCommand(r))

	return c, nil
}
