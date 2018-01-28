package main

import (
	"github.com/spf13/cobra"
)

type priceCmd struct{}

func registerPriceCmds(cli *Cli) {
	cli.priceCmd = &priceCmd{}

	priceCmd := &cobra.Command{
		Use:   "price",
		Short: "Coin price related commands",
		Long:  "Commands to list coin price accross exchanges",
		RunE:  cli.usageRunner(),
	}

	listPriceCmd := &cobra.Command{
		Use:     "list",
		Short:   "list price",
		Long:    "Command to list coin price accross exchanges",
		Example: `coin price list coin_name`,
		RunE:    cli.runner(priceListRun),
	}

	priceCmd.AddCommand(listPriceCmd)
	cli.rootCmd.AddCommand(priceCmd)

}
