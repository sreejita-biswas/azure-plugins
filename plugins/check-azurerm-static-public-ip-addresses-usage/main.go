package main

/*
#
# check-azurerm-static-public-ip-addresses-usage
#
# DESCRIPTION:
#   This plugin checks the number of Static Public IP Addresses allocated and available in a Region in Azure.
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
#   ./check-azurerm-static-public-ip-addresses-usage --location="westeurope" --warning=80 --critical=90
#
#   ./check-azurerm-static-public-ip-addresses-usage
#								--tenant_id="00000000-0000-0000-0000-000000000000"
#                             --client_id="00000000-0000-0000-0000-000000000000"
#                             --secret="00000000-0000-0000-0000-000000000000"
#                             --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                             --resourceGroup="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --warning=60
#                             --critical=80
#
#   ./check-azurerm-static-public-ip-addresses-usage
#							  --tenant_id="00000000-0000-0000-0000-000000000000"
#                             --client_id="00000000-0000-0000-0000-000000000000"
#                             --secret="00000000-0000-0000-0000-000000000000"
#                             --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                             --resourceGroup="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --warning=60
#                             --critical=80
#
# NOTES:
#
# LICENSE:
#   Tom Harvey
#   Released under the same terms as Sensu (the MIT license); see LICENSE
#   for details.
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
	checkStaticPublicIpAddressesUsage()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-static-public-ip-addresses-usage",
		Short: "The Sensu Go Azure handler for static public ip addresses management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	return cmd
}

func checkStaticPublicIpAddressesUsage() {
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetNetworkClient(utils.GetSubscription(), authorizer)
	utils.PrintNetworkUsageStatistic(client, "StaticPublicIPAddresses")
}
