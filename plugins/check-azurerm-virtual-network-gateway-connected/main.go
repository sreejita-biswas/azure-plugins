package main

/*
#
# check-azurerm-virtual-network-gateway-connected
#
# DESCRIPTION:
#   This plugin checks the specified Virtual Network Gateway is connected.
#   Firing a Critical alert if this is not.
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
#
#   ./check-azurerm-virtual-network-gateway-connected --tenant_id="00000000-0000-0000-0000-000000000000"
#                                 --client_id="00000000-0000-0000-0000-000000000000"
#                                 --secret="00000000-0000-0000-0000-000000000000"
#                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                                 --location="westeurope" --resource_group="resourceGroup" --name="name"
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

var gatewayName string
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
		Use:   "check-azurerm-virtual-network-gateway-connected",
		Short: "The Sensu Go Azure handler for virtual network gateway management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	cmd.Flags().StringVar(&resourceGroupName, "resource_group", "", "Azure Resource Group Name")
	cmd.Flags().StringVar(&gatewayName, "name", "", "Azure Virtual Network Connection Gateway Name")
	cmd.MarkFlagRequired("name")
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
	result, err := client.Get(context.Background(), resourceGroupName, gatewayName)

	if err != nil {
		fmt.Errorf("Error while calling Azure API,", err)
		return
	}

	connectionStatus := result.VirtualNetworkGatewayConnectionPropertiesFormat.ConnectionStatus
	inbound := *result.VirtualNetworkGatewayConnectionPropertiesFormat.IngressBytesTransferred
	outbound := *result.VirtualNetworkGatewayConnectionPropertiesFormat.ExpressRouteGatewayBypass

	message := fmt.Sprintf("State is '%s'. Usage is %d in / %d out", connectionStatus, inbound, outbound)
	if connectionStatus == "Connected" {
		fmt.Println("OK:", message)
	} else {
		fmt.Println("CRITICAL:", message)
	}

}
