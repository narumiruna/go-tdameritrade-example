package cmd

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/narumiruna/go-tdameritrade-example/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zricethezav/go-tdameritrade"
)

var accountCmd = &cobra.Command{
	Use: "get accounts",
	Run: func(cmd *cobra.Command, args []string) {
		position, err := cmd.Flags().GetBool("position")
		if err != nil {
			log.WithError(err).Error("failed to get <position> flag")
		}

		orders, err := cmd.Flags().GetBool("orders")
		if err != nil {
			log.WithError(err).Error("failed to get <orders> flag")
		}

		ctx := context.Background()

		client, err := client.New(ctx)
		if err != nil {
			log.WithError(err).Error("failed to create client")
		}

		accounts, _, err := client.Account.GetAccounts(ctx, &tdameritrade.AccountOptions{Position: position, Orders: orders})
		if err != nil {
			log.WithError(err).Error("failed to get accounts")
		}

		log.Info(spew.Sdump(accounts))
	},
}

func init() {
	accountCmd.Flags().BoolP("position", "p", false, "position")
	accountCmd.Flags().BoolP("orders", "o", false, "orders")

	RootCmd.AddCommand(accountCmd)
}
