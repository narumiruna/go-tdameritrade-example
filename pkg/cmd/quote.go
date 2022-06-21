package cmd

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/narumiruna/go-tdameritrade-example/pkg/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var quoteCmd = &cobra.Command{
	Use: "get quote",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil {
			log.WithError(err).Error("failed to get symbol")
		}

		ctx := context.Background()

		client, err := client.New(ctx)
		if err != nil {
			log.WithError(err).Error("failed to create client")
		}

		quotes, _, err := client.Quotes.GetQuotes(ctx, symbol)
		if err != nil {
			log.Fatal(err)
		}

		quote, ok := (*quotes)[symbol]
		if !ok {
			log.WithError(err).Fatal("no quote found")
		}

		log.Info(spew.Sdump(quote))
	},
}

func init() {
	quoteCmd.Flags().StringP("symbol", "s", "", "symbol")

	RootCmd.AddCommand(quoteCmd)
}
