
# azure-plugins

TravisCI: [![TravisCI Build Status](https://travis-ci.org/sreejita-biswas/azure-plugins.svg?branch=develop)](https://travis-ci.org/sreejita-biswas/azure-plugins)

## Functionality

 - check-azurerm-core-usage
 - check-azurerm-cores-d-usage
 - check-azurerm-cores-ds-usage
 - check-azurerm-cores-dsv2-usage
 - check-azurerm-cores-dv2-usage
 - check-azurerm-cores-f-usage
 - check-azurerm-cores-fs-usage
 - check-azurerm-network-interfaces-usage
 - check-azurerm-network-security-groups-usage
 - check-azurerm-public-ip-addresses-usage
 - check-azurerm-route-tables-usage
 - check-azurerm-service-bus-subscription
 - check-azurerm-service-bus-topic-size
 - check-azurerm-static-public-ip-addresses-usage
 - check-azurerm-virtual-machines-usage
 - check-azurerm-virtual-network-gateway-connected
 - check-azurerm-virtual-network-gateway-failover-connected
 - check-azurerm-virtual-networks-usage
 - check-azurerm-monitor-metric
 - metric-azurerm-service-bus-subscription-message-count
 - metric-azurerm-virtual-network-gateway-usage

## Files

* /bin/check-azurerm-core-usage
* /bin/check-azurerm-cores-d-usage
* /bin/check-azurerm-cores-ds-usage
* /bin/check-azurerm-cores-dsv2-usage
* /bin/check-azurerm-cores-dv2-usage
* /bin/check-azurerm-cores-f-usage
* /bin/check-azurerm-cores-fs-usage
* /bin/check-azurerm-network-interfaces-usage
* /bin/check-azurerm-network-security-groups-usage
* /bin/check-azurerm-public-ip-addresses-usage
* /bin/check-azurerm-route-tables-usage
* /bin/check-azurerm-service-bus-subscription
* /bin/check-azurerm-service-bus-topic-size
* /bin/check-azurerm-static-public-ip-addresses-usage
* /bin/check-azurerm-virtual-machines-usage
* /bin/check-azurerm-virtual-network-gateway-connected
* /bin/check-azurerm-virtual-network-gateway-failover-connected
* /bin/check-azurerm-virtual-networks-usage
* /bin/metric-azurerm-service-bus-subscription-message-count
* /bin/metric-azurerm-virtual-network-gateway-usage
* /bin/check-azurerm-monitor-metric


## Usage

**check-azurerm-core-usage**
```
./check-azurerm-core-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-core-usage --auth_type=1  
                           --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth 
                           --location="westeurope" --warning=80 --critical=90

./check-azurerm-core-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --location="eastus2" --warning=80 --critical=90
                           
Sample Output : OK: Current usage:1 of 4 Cores         
```

**check-azurerm-cores-d-usage**
```
./check-azurerm-cores-d-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-d-usage --auth_type=1 
                              --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                              --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-d-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                              --client_id="00000000-0000-0000-0000-000000000000"
                              --secret="00000000-0000-0000-0000-000000000000"
                              --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                              --location="eastus2" --warning=80 --critical=90
                              
Sample Output : CRITICAL: Current usage:4 of 4 D Family Cores   
```

**check-azurerm-cores-ds-usage**
```
./check-azurerm-cores-ds-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-ds-usage --auth_type=1 
                               --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                               --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-ds-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                               --client_id="00000000-0000-0000-0000-000000000000"
                               --secret="00000000-0000-0000-0000-000000000000"
                               --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                               --location="eastus2" --warning=80 --critical=90
                               
Sample Output : OK: Current usage:0 of 4 DS Family Cores                              
```

**check-azurerm-cores-dsv2-usage**
```
./check-azurerm-cores-dsv2-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dsv2-usage --auth_type=1 
                                 --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                 --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dsv2-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90
                                 
Sample Output : OK: Current usage:0 of 4 DS Family Cores                                     
```

**check-azurerm-cores-dv2-usage**
```
./check-azurerm-cores-dv2-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dv2-usage --auth_type=1 
                                --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dv2-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                --client_id="00000000-0000-0000-0000-000000000000"
                                --secret="00000000-0000-0000-0000-000000000000"
                                --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                --location="eastus2" --warning=80 --critical=90
                                
Sample Output : WARNING: Current usage:85 of 100 DSv2 Family Cores                      
```

**check-azurerm-cores-f-usage**
```
./check-azurerm-cores-f-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-f-usage --auth_type=1 
                              --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                              --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-f-usage.rb --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90

Sample Output : OK: Current usage:0 of 4 F Family Cores
```

**check-azurerm-cores-fs-usage**
```
./check-azurerm-cores-fs-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-fs-usage --auth_type=1 
                               --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                               --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-fs-usage   --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90
 
Sample Output : OK: Current usage:1 of 6 FS Family Cores
```

**check-azurerm-load-balancers-usage**
```
./check-azurerm-load-balancers-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-load-balancers-usage --auth_type=1 
                                     --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                     --location="westeurope" --warning=80 --critical=90

./check-azurerm-load-balancers-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                     --client_id="00000000-0000-0000-0000-000000000000"
                                     --secret="00000000-0000-0000-0000-000000000000"
                                     --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                     --location="eastus2" --warning=80 --critical=90
                                     
Sample Output : OK: Current usage:20 of 1000 Load Balancers                                   
```

**check-azurerm-network-interfaces-usage**
```
./check-azurerm-network-interfaces-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-interfaces-usage --auth_type=1 
                                         --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                         --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-interfaces-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                         --client_id="00000000-0000-0000-0000-000000000000"
                                         --secret="00000000-0000-0000-0000-000000000000"
                                         --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                         --location="eastus2" --warning=80 --critical=90

Sample Output : OK: Current usage:57 of 65536 Network Interfaces
```

**check-azurerm-network-security-groups-usage**
```
./check-azurerm-network-security-groups-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-security-groups-usage --auth_type=1 
                                              --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                              --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-security-groups-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                              --client_id="00000000-0000-0000-0000-000000000000"
                                              --secret="00000000-0000-0000-0000-000000000000"
                                              --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                              --location="eastus2" --warning=80 --critical=90
                                              
Sample Output : OK: Current usage:100 of 5000 Network Security Groups                                            
```

**check-azurerm-public-ip-addresses-usage**
```
./check-azurerm-public-ip-addresses-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-public-ip-addresses-usage --auth_type=1 
                                          --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                          --location="westeurope" --warning=80 --critical=90

./check-azurerm-public-ip-addresses-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                          --client_id="00000000-0000-0000-0000-000000000000"
                                          --secret="00000000-0000-0000-0000-000000000000"
                                          --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                          --location="eastus2" --warning=80 --critical=90
                                          
Sample Output : OK: Current usage:100 of 5000 Network Security Groups                                         
```

**check-azurerm-static-public-ip-addresses-usage**
```
./check-azurerm-static-public-ip-addresses-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-static-public-ip-addresses-usage --auth_type=1 
                                                 --azure_auth_location=/Users/sreejita.biswas/.azure/credential.auth
                                                 --location="westeurope" --warning=80 --critical=90

./check-azurerm-static-public-ip-addresses-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                                 --client_id="00000000-0000-0000-0000-000000000000"
                                                 --secret="00000000-0000-0000-0000-000000000000"
                                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                                 --location="eastus2" --warning=80 --critical=90
                                                 
Sample Output : OK: Current usage:1 of 1000 Public IP Addresses                                                  
```

**check-azurerm-route-tables-usage**
```
./check-azurerm-route-tables-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-route-tables-usage --auth_type=1 --location="westeurope" --warning=80 --critical=90

./check-azurerm-route-tables-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                   --client_id="00000000-0000-0000-0000-000000000000"
                                   --secret="00000000-0000-0000-0000-000000000000"
                                   --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                   --location="eastus2" --warning=80 --critical=90
                                   
Sample Output : OK: Current usage:23 of 200 Route Tables                                        
```

**check-azurerm-service-bus-subscription**
```
./check-azurerm-service-bus-subscription
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"
                           
./check-azurerm-service-bus-subscription
                           --auth_type=1
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"

./check-azurerm-service-bus-subscription
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscriptionName="subscriptionName"
                           
Sample Output : CRITICAL:Subscription 'FreeTrial' not found in topic 'hello'                        
```

**check-azurerm-service-bus-topic-size**
```
./check-azurerm-service-bus-topic-size
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --warning=60
                           --critical=80
                           
./check-azurerm-service-bus-topic-size
                           --auth_type=1
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --warning=60
                           --critical=80

./check-azurerm-service-bus-topic-size
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --warning=60
                           --critical=80
                           
Sample Output : OK: Current size of topic 'test-topic' is NaN                         
```

**check-azurerm-virtual-machines-usage**
```
./check-azurerm-virtual-machines-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-virtual-machines-usage --auth_type=1 --location="westeurope" --warning=80 --critical=90

./check-azurerm-virtual-machines-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                       --client_id="00000000-0000-0000-0000-000000000000"
                                       --secret="00000000-0000-0000-0000-000000000000"
                                       --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                       --location="eastus2" --warning=80 --critical=90
   
Sample Output : OK: Current usage:1 of 10000 VMs       
```

**check-azurerm-virtual-network-gateway-connected**
```
./check-azurerm-virtual-network-gateway-connected --resource_group="resourcegroup" --name="gatewayname"

./check-azurerm-virtual-network-gateway-connected --auth_type=1 --resource_group="resourcegroup" --name="gatewayname"

./check-azurerm-virtual-network-gateway-connected
                               --tenant_id="00000000-0000-0000-0000-000000000000"
                               --client_id="00000000-0000-0000-0000-000000000000"
                               --secret="00000000-0000-0000-0000-000000000000"
                               --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                               --resource_group="resourcegroup" --name="gatewayname"
                               
Sample Output : OK: Current usage:1 of 10000 VMs                              
```

**check-azurerm-virtual-network-gateway-failover-connected**
```
./check-azurerm-virtual-network-gateway-failover-connected
                           --resource_group="resourcegroup"
                           --primary_name="primaryname"
                           --secondary_name="secondaryname"
                           
./check-azurerm-virtual-network-gateway-failover-connected
                           --auth_type=1
                           --resource_group="resourcegroup"
                           --primary_name="primaryname"
                           --secondary_name="secondaryname"

./check-azurerm-virtual-network-gateway-failover-connected
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup" --name="gatewayname"
```

**check-azurerm-virtual-networks-usage**
```
./check-azurerm-virtual-networks-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-virtual-networks-usage  --auth_type=1 --location="westeurope" --warning=80 --critical=90

./check-azurerm-virtual-networks-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                       --client_id="00000000-0000-0000-0000-000000000000"
                                       --secret="00000000-0000-0000-0000-000000000000"
                                       --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                       --location="eastus2" --warning=80 --critical=90
                                       
Sample Output : OK: Current usage:2 of 1000 Virtual Networks                                  
```

**metric-azurerm-service-bus-subscription-message-count**
```
./metric-azurerm-service-bus-subscription-message-count
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"
                           
 ./metric-azurerm-service-bus-subscription-message-count
                           --auth_type=1
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"

./metric-azurerm-service-bus-subscription-message-count.rb
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"

./metric-azurerm-service-bus-subscription-message-count.rb
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --subscription_name="subscriptionName"
                           --coustom_scheme="foo"
                           
Sample Output : test.testing-azure.test-azure-bus.test-topic.sub.message_count-0 - timestamp ::2019-01-15T16:50:22+05:30    
```

**Authentication Setup**

```
Two types of Authentication supported -
1. Credential based authentication : Either provide tenant_id,client_id,secret and location as the command line arguments or set three environment variables AZURE_TENANT_ID, AZURE_CLIENT_ID,AZURE_CLIENT_SECRET and AZURE_AUTH_LOCATION as below

export AZURE_TENANT_ID=00000000-0000-0000-0000-000000000000
export AZURE_CLIENT_ID=00000000-0000-0000-0000-000000000000
export AZURE_CLIENT_SECRET=00000000-0000-0000-0000-000000000000
export AZURE_AUTH_LOCATION=eastus2

2. File based authentication : sample file format is mentioned below. Provide auth_type (as 1) and azure_auth_location(the file location for e.g. /User/xyz/.azure/credential.auth) as the command line arguments.

{
"tenantId":"00000000-0000-0000-0000-000000000000",
"clientId":"00000000-0000-0000-0000-000000000000",
"clientSecret":"00000000-0000-0000-0000-000000000000",
"activeDirectoryEndpointUrl":"https://www.xyz.com",
"resourceManagerEndpointUrl":"https://www.xyz.com"
}
```

**Azure Account Setup**

```
- Step 1: Create an Azure Account
- Step 2: Sign In using username and paasword.
- Step 3: Select Azure Active Directory
- Step 4: Click on App Registration
- Step 5: Enter App Name and Sign In Url. Click on Create. The generated Application Id is the client id.
- Step 6: Go to Settings and create one key. 
- Step 7: The genarted key value will be the client secret.
- Step 8: Go to All Services and click on subscription name. 
- Step 9: Select Access Control(IAM) and click on "Add role assignment"
- Step10: Select role as Owner/Contributor and select the create App name in the last text field.
- Step 11: We are all set to execute the plugins with the azure credential details.

```

