package main

/*
#
# check-azurerm-virtual-network-gateway-failover-connected
#
# DESCRIPTION:
#   This plugin checks that at least one of the specified Virtual Network Gateways is connected.
#   Firing a Critical alert if both are not in the Connected state.
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   Linux
#
# DEPENDENCIES:
#   gem: azure_mgmt_network
#   gem: sensu-plugin
#
# USAGE:
#   ./check-azurerm-virtual-network-gateway-failover-connected --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="westeurope" --resource_group="resourceGroup" --primary_name="name"
#                                 --secondary_name="name2"
# NOTES:
#
# LICENSE:
#   Tom Harvey
#   Released under the same terms as Sensu (the MIT license); see LICENSE
#   for details.
#
*/

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/authorization"
	"github.com/sreejita-biswas/azure-plugins/clients"
	"github.com/sreejita-biswas/azure-plugins/utils"
)

var primaryName string
var secondaryName string
var resourceGroupName string

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
	checkVirtualNetworkGatewayUsage()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-virtual-network-gateway-failover-connected",
		Short: "The Sensu Go Azure handler for virtual network gateway management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	cmd.Flags().StringVar(&resourceGroupName, "resource_group", "", "Azure Resource Group Name")
	cmd.Flags().StringVar(&primaryName, "primary_name", "", "Azure Virtual Network Connection Gateway Primary Name")
	cmd.Flags().StringVar(&secondaryName, "seondary_name", "", "Azure Virtual Network Connection Gateway Secondary Name")
	cmd.MarkFlagRequired("primary_name")
	cmd.MarkFlagRequired("seondary_name")
	cmd.MarkFlagRequired("resource_group")
	return cmd
}

func checkVirtualNetworkGatewayUsage() {
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetVirtualNetworkClient(utils.GetSubscription(), authorizer)
	primaryResult, err := client.Get(context.Background(), resourceGroupName, primaryName)

	if err != nil {
		fmt.Errorf("Error while calling Azure API,", err)
		return
	}

	secondaryResult, err := client.Get(context.Background(), resourceGroupName, secondaryName)

	if err != nil {
		fmt.Errorf("Error while calling Azure API,", err)
		return
	}

	primaryConnectionStatus := primaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.ConnectionStatus
	primaryInbound := *primaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.IngressBytesTransferred
	primaryOutbound := *primaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.ExpressRouteGatewayBypass

	message := fmt.Sprintf("Primary:State is '%s'. Usage is %d in / %d out", primaryConnectionStatus, primaryInbound, primaryOutbound)

	secondaryConnectionStatus := secondaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.ConnectionStatus
	secondaryprimaryInbound := *secondaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.IngressBytesTransferred
	secondaryOutbound := *secondaryResult.VirtualNetworkGatewayConnectionPropertiesFormat.ExpressRouteGatewayBypass

	message = fmt.Sprintf("%s\nSecondary:State is '%s'. Usage is %d in / %d out", message, secondaryConnectionStatus, secondaryprimaryInbound, secondaryOutbound)

	if primaryConnectionStatus == "Connected" || secondaryConnectionStatus == "Connected" {
		fmt.Println("OK:", message)
	} else {
		fmt.Println("CRITICAL:", message)
	}

}
