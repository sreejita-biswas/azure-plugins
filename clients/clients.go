package clients

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-10-01/network"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/Azure/go-autorest/autorest"
)

func GetVirtualNetworkClient(subscription string, authorizer autorest.Authorizer) network.VirtualNetworkGatewayConnectionsClient {
	client := network.NewVirtualNetworkGatewayConnectionsClient(subscription)
	client.Authorizer = authorizer
	return client
}

func GetUsageClient(subscription string, authorizer autorest.Authorizer) compute.UsageClient {
	client := compute.NewUsageClient(subscription)
	client.Authorizer = authorizer
	return client
}

func GetSubscriptionClient(subscription string, authorizer autorest.Authorizer) servicebus.SubscriptionsClient {
	client := servicebus.NewSubscriptionsClient(subscription)
	client.Authorizer = authorizer
	return client
}

func GetTopicClient(subscription string, authorizer autorest.Authorizer) servicebus.TopicsClient {
	client := servicebus.NewTopicsClient(subscription)
	client.Authorizer = authorizer
	return client
}

func GetNetworkClient(subscription string, authorizer autorest.Authorizer) network.UsagesClient {
	client := network.NewUsagesClient(subscription)
	client.Authorizer = authorizer
	return client
}
