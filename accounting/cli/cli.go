package accounting_cli

import (
	"context"
	"fmt"

	"github.com/goblinlordx/go-test/accounting"
	"github.com/goblinlordx/go-test/app"
	"github.com/goblinlordx/go-test/cli"
	"github.com/goblinlordx/go-test/db"
	"github.com/spf13/cobra"
)

func CLIRoot() *cobra.Command {
	// asdf := cobra.Command{
	// 	Use: "asdf",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("asdf")
	// 	},
	// }

	root := cobra.Command{
		Use:     "accounting",
		PreRunE: cli.DefaultHelp,
	}

	root.AddCommand(CLIAccountCommand())

	return &root
}

func CLIAccountCommand() *cobra.Command {
	ac := cobra.Command{
		Use:     "account",
		PreRunE: cli.DefaultHelp,
	}

	createCommand := cobra.Command{
		Use:  "create",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			app := app.InitializeApp()
			ctx := context.Background()
			db.WithDbCtx(ctx, func(ctx context.Context) error {

				id, err := app.AccountCommander.Create(ctx, accounting.AccountCreateInput{
					Name:           args[0],
					CurrencySymbol: args[1],
				})

				fmt.Println("Created Account ID: ", id)

				if err != nil {
					panic(err)
				}
				return nil
			})
		},
	}

	ac.AddCommand(&createCommand)

	return &ac
}
