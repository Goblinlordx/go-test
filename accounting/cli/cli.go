package accounting_cli

import (
	"context"
	"fmt"
	"log"

	"github.com/goblinlordx/go-test/accounting"
	"github.com/goblinlordx/go-test/app"
	"github.com/goblinlordx/go-test/cli"
	"github.com/goblinlordx/go-test/db"
	"github.com/goblinlordx/go-test/db/modifier"
	"github.com/spf13/cobra"
)

func CLIRoot() *cobra.Command {
	root := cobra.Command{
		Use:     "accounting",
		PreRunE: cli.DefaultHelp,
	}

	root.AddCommand(cliAccountCommand())

	return &root
}

func cliAccountCreateCommand() *cobra.Command {
	return &cobra.Command{
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

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Created Account ID: ", id)

				return nil
			})
		},
	}
}

func cliAccountDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:  "delete",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			app := app.InitializeApp()
			ctx := context.Background()
			db.WithDbCtx(ctx, func(ctx context.Context) error {
				err := app.AccountCommander.Delete(ctx, accounting.AccountDeleteInput{
					Ids: []string{args[0]},
				})

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Delete Account ID: ", args[0])

				return nil
			})
		},
	}
}

func cliAccountQueryCommand() *cobra.Command {
	filters := accounting.AccountFilterInput{}
	cmd := cobra.Command{
		Use: "query",
		// Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			app := app.InitializeApp()
			ctx := context.Background()
			db.WithDbCtx(ctx, func(ctx context.Context) error {
				return modifier.WithModifierScope(ctx, func(ctx context.Context) error {
					ctx, err := app.AccountFilterer.Apply(ctx, filters)
					if err != nil {
						log.Fatal(err)
					}

					res, err := app.AccountRepository.Search(ctx)

					if err != nil {
						log.Fatal(err)
					}

					for _, v := range res {
						fmt.Println(v)
					}

					return nil
				})
			})
		},
	}

	cmd.Flags().StringVarP(&filters.Id, "id", "i", "", "test 2")
	cmd.Flags().StringVarP(&filters.Name, "name", "n", "", "test 2")
	cmd.Flags().StringVarP(&filters.CurrencySymbol, "currencySymbol", "c", "", "test 2")

	return &cmd
}

func cliAccountCommand() *cobra.Command {
	ac := cobra.Command{
		Use:     "account",
		PreRunE: cli.DefaultHelp,
	}

	ac.AddCommand(cliAccountCreateCommand())
	ac.AddCommand(cliAccountDeleteCommand())
	ac.AddCommand(cliAccountQueryCommand())

	return &ac
}
