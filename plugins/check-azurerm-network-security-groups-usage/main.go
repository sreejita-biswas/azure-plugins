package main

/*
#
# check-azurerm-network-security-groups-usage
#
# DESCRIPTION:
#   This plugin checks the number of Network Security Groups allocated and available in a Region in Azure.
#   Warning and Critical Percentage thresholds may be set as needed.
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   Linux
#
#
# USAGE:
#   ./check-azurerm-network-security-groups-usage --location="westeurope" --warning=80 -critical=90
#
#   ./check-azurerm-network-security-groups-usage --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 -subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="eastus2" --warning=80 --critical=90
#
#   ./check-azurerm-network-security-groups-usage --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="westeurope"
#                                 --warning=80
#                                 --critical=90
# NOTES:
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
	checkNetworkSecurityGroupUsage()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-network-security-groups-usage",
		Short: "The Sensu Go Azure handler for network security group management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	return cmd
}

func checkNetworkSecurityGroupUsage() {
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetNetworkClient(utils.GetSubscription(), authorizer)
	utils.PrintNetworkUsageStatistic(client, "NetworkSecurityGroups")
}
