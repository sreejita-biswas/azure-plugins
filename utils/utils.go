package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/authorization"
	"github.com/sreejita-biswas/azure-plugins/constants"
)

var (
	location     string
	warning      float64
	critical     float64
	subscription string
)

func GetPluginSpecificParameters(cmd *cobra.Command) {
	authorization.GetParameterValues(cmd)
	cmd.Flags().StringVar(&location, "location", "", "Azure Location (e.g. westeurope/eastus2)")
	cmd.Flags().Float64Var(&warning, "warning", 80, "Warning Percentage threshold for filter")
	cmd.Flags().Float64Var(&critical, "critical", 90, "Critical Percentage threshold for filter")
	cmd.Flags().StringVar(&subscription, "subscription", "", "ARM Subscription ID")
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
