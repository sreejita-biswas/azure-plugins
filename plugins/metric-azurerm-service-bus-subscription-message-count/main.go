package main

/*
#
# metric-azurerm-service-bus-subscription-message-count
#
# DESCRIPTION:
#   This plugin exposes the Service Bus Subscription Message Counts as a Metric
#
# OUTPUT:
#   plain-text
#
# PLATFORMS:
#   Linux
#
#
# USAGE:
#
#   ./metric-azurerm-service-bus-subscription-message-count
#                             --resource_group="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --subscription_name="subscriptionName"
#
#   ./metric-azurerm-service-bus-subscription-message-count
#                             --tenant_id="00000000-0000-0000-0000-000000000000"
#                             --client_id="00000000-0000-0000-0000-000000000000"
#                             --secret="00000000-0000-0000-0000-000000000000"
#                             --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                             --resource_group="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --subscription_name="subscriptionName"
#
#   ./metric-azurerm-service-bus-subscription-message-count
#                             --tenant_id="00000000-0000-0000-0000-000000000000"
#                             --client_id="00000000-0000-0000-0000-000000000000"
#                             --secret="00000000-0000-0000-0000-000000000000"
#                             --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                              --resource_group="resourcegroup"
#                             --namespace="namespace"
#                             --topic="topic"
#                             --subscription_name="subscriptionName"
#                             --custom_scheme="foo"
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
	"time"

	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/authorization"
	"github.com/sreejita-biswas/azure-plugins/clients"
	"github.com/sreejita-biswas/azure-plugins/utils"
)

var customScheme string

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
	checkServiceBusSubscription()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metric-azurerm-service-bus-subscription-message-count",
		Short: "The Sensu Go Azure Metric handler for subscription message count management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	utils.GetServiceBusPluginSpecificParameters(cmd)
	cmd.Flags().StringVar(&customScheme, "coustom_scheme", "", "Metric naming scheme, text to prepend to .$parent.$child")

	_ = cmd.MarkFlagRequired("resource_group")
	_ = cmd.MarkFlagRequired("namespace")
	_ = cmd.MarkFlagRequired("topic")
	_ = cmd.MarkFlagRequired("subscription_name")

	return cmd
}

func checkServiceBusSubscription() {
	var message string
	var count int64 = 0
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetSubscriptionClient(utils.GetSubscription(), authorizer)
	result, err := utils.GetServiceBusresult(client)

	if err != nil {
		return
	}

	if result == nil {
		return
	}

	if result.MessageCount == nil {
		return
	}

	count = *result.MessageCount
	timestamp := time.Now()
	message = fmt.Sprintf("%s.%s.%s.%s.%s.%d", customScheme, utils.GetResourceGroup(), utils.GetNamespace(), utils.GetTopic(), utils.GetSubscriptionName(), "message_count")
	fmt.Println(message, "-", count, "timestamp :", timestamp.Format(time.RFC3339))
}
