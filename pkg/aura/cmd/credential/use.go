package credential

import (
	"errors"

	"github.com/neo4j/cli/pkg/clictx"
	"github.com/spf13/cobra"
)

func NewUseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "use",
		Short: "Sets the default credential to be used",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, ok := clictx.Config(cmd.Context())

			if !ok {
				return errors.New("error fetching configuration values")
			}

			err := config.Aura.SetDefaultCredential(args[0])
			if err != nil {
				return err
			}

			err = config.Write()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
