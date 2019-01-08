package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-10-01/network"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/authorization"
	"github.com/sreejita-biswas/azure-plugins/constants"
)

var (
	location     string
	warning      float64
	critical     float64
	subscription string

	resourceGroupName string
	namespace         string
	topic             string
	subscriptionName  string
)

func GetPluginSpecificParameters(cmd *cobra.Command) {
	authorization.GetParameterValues(cmd)
	cmd.Flags().StringVar(&location, "location", "", "Azure Location (e.g. westeurope/eastus2)")
	cmd.Flags().Float64Var(&warning, "warning", 80, "Warning Percentage threshold for filter")
	cmd.Flags().Float64Var(&critical, "critical", 90, "Critical Percentage threshold for filter")
	cmd.Flags().StringVar(&subscription, "subscription", "", "ARM Subscription ID")
}

func GetServiceBusPluginSpecificParameters(cmd *cobra.Command) {
	cmd.Flags().StringVar(&resourceGroupName, "resource_group", "", "Azure Resource Group Name")
	cmd.Flags().StringVar(&namespace, "namespace", "", "Azure Service Bus Namespace Name")
	cmd.Flags().StringVar(&topic, "topic", "", "Azure Service Bus Topic Name")
	cmd.Flags().StringVar(&subscriptionName, "subscription_name", "", "Azure Service Bus Subscription Name")
}

func IsValidPluginSpecificParameters() bool {
	if len(subscription) <= 0 && len(os.Getenv("AZURE_SUBSCRIPTION_ID")) <= 0 {
		fmt.Println("provide valid subscritpion id")
		return false
	}
	if len(location) <= 0 {
		fmt.Println("provide azure location")
		return false
	}
	return true
}

func GetSubscription() string {
	return subscription
}

func PrintUsageStatistic(client compute.UsageClient, statistic string) {
	var (
		currentUsage   int32
		allowance      int64
		message        string
		usageFound     = false
		percentageUsed float64
	)
	usageStatistics, err := client.List(context.Background(), location)
	if err != nil {
		fmt.Println("CRITICAL:Error while getting usage statictics details ,", err)
	}
	for _, usage := range usageStatistics.Values() {
		if usage.Name != nil && usage.Name.Value != nil && *usage.Name.Value == statistic {
			currentUsage = *usage.CurrentValue
			allowance = *usage.Limit
			usageFound = true
		}
	}
	if !usageFound {
		return
	}
	message = fmt.Sprintf("Current usage:%d of %d %s", currentUsage, allowance, constants.GetUnitByUsageStatisticType(statistic))
	percentageUsed = (float64(currentUsage) / float64(allowance)) * float64(100)
	if percentageUsed >= critical {
		fmt.Println("CRITICAL:", message)
	} else if percentageUsed >= warning {
		fmt.Println("WARNING:", message)
	} else {
		fmt.Println("OK:", message)
	}
}

func PrintNetworkUsageStatistic(client network.UsagesClient, statistic string) {
	var (
		currentUsage   int64
		allowance      int64
		message        string
		usageFound     = false
		percentageUsed float64
	)
	usageStatistics, err := client.List(context.Background(), location)
	if err != nil {
		fmt.Println("CRITICAL:Error while getting usage statictics details ,", err)
	}
	for _, usage := range usageStatistics.Values() {
		if usage.Name != nil && usage.Name.Value != nil && *usage.Name.Value == statistic {
			currentUsage = *usage.CurrentValue
			allowance = *usage.Limit
			usageFound = true
		}
	}
	if !usageFound {
		return
	}
	message = fmt.Sprintf("Current usage:%d of %d %s", currentUsage, allowance, constants.GetUnitByUsageStatisticType(statistic))
	percentageUsed = (float64(currentUsage) / float64(allowance)) * float64(100)
	if percentageUsed >= critical {
		fmt.Println("CRITICAL:", message)
	} else if percentageUsed >= warning {
		fmt.Println("WARNING:", message)
	} else {
		fmt.Println("OK:", message)
	}
}

func GetServiceBusresult(client servicebus.SubscriptionsClient) (*servicebus.SBSubscription, error) {
	result, err := client.Get(context.Background(), resourceGroupName, namespace, topic, subscriptionName)
	if err != nil {
		fmt.Errorf("Error while calling Azure Service Bus Api,", err)
		return &servicebus.SBSubscription{}, err
	}
	return &result, nil
}

func GetSubscriptionName() string {
	return subscriptionName
}

func GetTopic() string {
	return topic
}

func GetResourceGroup() string {
	return resourceGroupName
}

func GetNamespace() string {
	return namespace
}

func GetCritical() float64 {
	return critical
}
func GetWarning() float64 {
	return warning
}
