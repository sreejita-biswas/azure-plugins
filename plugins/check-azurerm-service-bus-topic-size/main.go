package main

/*
#
# check-azurerm-service-bus-topic-size
#
# DESCRIPTION:
#   This plugin checks a given Service Bus Topic percentage used with warning/critical thresholds
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   Linux
#
# DEPENDENCIES:
#   gem: azure_mgmt_service_bus
#   gem: sensu-plugin
#
# USAGE:
#
#   ./check-azurerm-service-bus-topic-size
#                             --resource_group="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --warning=60
#                             --critical=80
#
#   ./check-azurerm-service-bus-topic-size
#                             --tenant_id="00000000-0000-0000-0000-000000000000"
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
#   Andy Royle
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
	checkServiceBusTopicSize()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-service-bus-topic-size",
		Short: "The Sensu Go Azure handler for service bus topic size management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	utils.GetServiceBusPluginSpecificParameters(cmd)

	_ = cmd.MarkFlagRequired("resource_group")
	_ = cmd.MarkFlagRequired("namespace")
	_ = cmd.MarkFlagRequired("topic")

	return cmd
}

func checkServiceBusTopicSize() {
	var maxSizeInBytes int32 = 0
	var currentSize int64 = 0
	var percentageused float64
	var message string
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetTopicClient(utils.GetSubscription(), authorizer)
	result, err := client.Get(context.Background(), utils.GetResourceGroup(), utils.GetNamespace(), utils.GetTopic())

	if err != nil {
		fmt.Println("Error while calling Azure Service Bus Topic Api,", err)
		return
	}

	if result.MaxSizeInMegabytes == nil {
		return
	}

	maxSizeInBytes = *result.MaxSizeInMegabytes * int32(1024) * int32(1024)
	currentSize = *result.SizeInBytes
	percentageused = (float64(currentSize) / (float64(maxSizeInBytes))) * (float64(100))
	message = fmt.Sprintf("Current size of topic '%s' is %v", utils.GetTopic(), percentageused)

	if percentageused >= utils.GetCritical() {
		fmt.Println("CRITICAL:", message)
	} else if percentageused >= utils.GetWarning() {
		fmt.Println("WARNING:", message)
	} else {
		fmt.Println("OK:", message)
	}
}
