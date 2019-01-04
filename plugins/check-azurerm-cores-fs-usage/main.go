package main

/*
#
# check-azurerm-cores-f-usage
#
# DESCRIPTION:
#   This plugin checks the number of FS Family CPU Cores allocated and available in a Region in Azure.
#   Warning and Critical Percentage thresholds may be set as needed.
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   MAC OS
#
#
# USAGE:
#   ./check-azurerm-cores-fs-usage --location="westeurope" --warning=80 -critical=90
#
#   ./check-azurerm-cores-fs-usage --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 -subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="eastus2" --warning=80 --critical=90
#
#   ./check-azurerm-cores-fs-usage --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="westeurope"
#                                 --warning=80
#                                 --critical=90
#
# NOTES:
#
# LICENSE:
#   TODO
#
*/

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/authorization"
	"github.com/sreejita-biswas/azure-plugins/clients"
	"github.com/sreejita-biswas/azure-plugins/utils"
)

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}
	checkCoreFsUsage()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-core-fs-usage",
		Short: "The Sensu Go Azure handler for usage management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	return cmd
}

func checkCoreFsUsage() {
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetUsageClient(utils.GetSubscription(), authorizer)
	utils.PrintUsageStatistic(client, "standardFSFamily")
}
