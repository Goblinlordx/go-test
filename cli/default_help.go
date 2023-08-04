package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func DefaultHelp(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
	return nil
}
