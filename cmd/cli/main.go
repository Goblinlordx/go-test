package main

import (
	"os"
	"path/filepath"

	accounting_cli "github.com/goblinlordx/go-test/accounting/cli"
	"github.com/goblinlordx/go-test/cli"
	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{
		Use:     filepath.Base(os.Args[0]),
		PreRunE: cli.DefaultHelp,
	}

	root.AddCommand(accounting_cli.CLIRoot())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
