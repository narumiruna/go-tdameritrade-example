package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-tdameritrade-example",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.WithError(err).Error("failed to execute")
	}
}
