package main

/*
#
# check-azurerm-service-bus-subscription
#
# DESCRIPTION:
#   This plugin asserts that a given Service Bus Subscription exists
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
#   ./check-azurerm-service-bus-subscription
#                             --resource_group "resourcegroup"
#                             --namespace "namespace"
#                             --topic "topic"
#                             --subscription_name "subscriptionName"
#
#   ./check-azurerm-service-bus-subscription.rb
#                             --tenant_id="00000000-0000-0000-0000-000000000000"
#                             --client_id="00000000-0000-0000-0000-000000000000"
#                             --secret="00000000-0000-0000-0000-000000000000"
#                             --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
#                             --location="westeurope"
#                             --resource_group "resourcegroup"
#                             --namespace "namespace"
#                             --topic "topic"
#                             --subscription_name "subscriptionName"
#
# NOTES:
#
# LICENSE:
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
	checkServiceBusSubscription()
	return nil
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-azurerm-service-bus-subscription",
		Short: "The Sensu Go Azure handler for service bus subscription management",
		RunE:  run,
	}

	utils.GetPluginSpecificParameters(cmd)
	utils.GetServiceBusPluginSpecificParameters(cmd)

	_ = cmd.MarkFlagRequired("resource_group")
	_ = cmd.MarkFlagRequired("namespace")
	_ = cmd.MarkFlagRequired("topic")
	_ = cmd.MarkFlagRequired("subscription_name")

	return cmd
}

func checkServiceBusSubscription() {
	var message string
	authorizer, success := authorization.Authorize()
	if !success {
		return
	}
	success = utils.IsValidPluginSpecificParameters()
	if !success {
		return
	}
	client := clients.GetSubscriptionClient(utils.GetSubscription(), authorizer)
	result, _ := utils.GetServiceBusresult(client)

	// if err != nil {
	// 	fmt.Println("Error while calling Azure Api,", err)
	// 	return
	// }

	if result == nil || result.SBSubscriptionProperties == nil {
		message = fmt.Sprintf("CRITICAL:Subscription '%s' not found in topic '%s'", utils.GetSubscriptionName(), utils.GetTopic())
	} else {
		message = fmt.Sprintf("OK:Subscription '%s' was found in topic '%s'", utils.GetSubscriptionName(), utils.GetTopic())
	}
	fmt.Println(message)
}
