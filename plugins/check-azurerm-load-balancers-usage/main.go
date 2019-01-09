package main

/*
# check-azurerm-load-balancers-usage
#
# DESCRIPTION:
#   This plugin checks the number of Load Balancers allocated and available in a Region in Azure.
#   Warning and Critical Percentage thresholds may be set as needed.
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   MAC OS
#
# USAGE:
#   ./check-azurerm-load-balancers-usage --location="westeurope" --warning=80 -critical=90
#
#   ./check-azurerm-load-balancers-usage --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 -subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="eastus2" --warning=80 --critical=90
#
#   ./check-azurerm-load-balancers-usage --tenant_id="00000000-0000-0000-0000-000000000000"
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
	checkLoadBalancerUsage()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-load-balancers-usage",
		Short: "The Sensu Go Azure handler for load balancers management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	return cmd
}

func checkLoadBalancerUsage() {
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetNetworkClient(utils.GetSubscription(), authorizer)
	utils.PrintNetworkUsageStatistic(client, "LoadBalancers")
}
