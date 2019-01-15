# azure-plugins

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

./check-azurerm-core-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-d-usage**
```
./check-azurerm-cores-d-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-d-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                              --client_id="00000000-0000-0000-0000-000000000000"
                              --secret="00000000-0000-0000-0000-000000000000"
                              --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                              --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-ds-usage**
```
./check-azurerm-cores-ds-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-ds-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                               --client_id="00000000-0000-0000-0000-000000000000"
                               --secret="00000000-0000-0000-0000-000000000000"
                               --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                               --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-dsv2-usage**
```
./check-azurerm-cores-dsv2-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dsv2-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-dv2-usage**
```
./check-azurerm-cores-dv2-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-dv2-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                --client_id="00000000-0000-0000-0000-000000000000"
                                --secret="00000000-0000-0000-0000-000000000000"
                                --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-f-usage**
```
./check-azurerm-cores-f-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-f-usage.rb --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-cores-fs-usage**
```
./check-azurerm-cores-fs-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-cores-fs-usage   --tenant_id="00000000-0000-0000-0000-000000000000"
                                 --client_id="00000000-0000-0000-0000-000000000000"
                                 --secret="00000000-0000-0000-0000-000000000000"
                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                 --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-load-balancers-usage**
```
./check-azurerm-load-balancers-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-load-balancers-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                     --client_id="00000000-0000-0000-0000-000000000000"
                                     --secret="00000000-0000-0000-0000-000000000000"
                                     --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                     --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-network-interfaces-usage**
```
./check-azurerm-network-interfaces-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-interfaces-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                         --client_id="00000000-0000-0000-0000-000000000000"
                                         --secret="00000000-0000-0000-0000-000000000000"
                                         --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                         --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-network-security-groups-usage**
```
./check-azurerm-network-security-groups-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-network-security-groups-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                              --client_id="00000000-0000-0000-0000-000000000000"
                                              --secret="00000000-0000-0000-0000-000000000000"
                                              --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                              --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-public-ip-addresses-usage**
```
./check-azurerm-public-ip-addresses-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-public-ip-addresses-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                          --client_id="00000000-0000-0000-0000-000000000000"
                                          --secret="00000000-0000-0000-0000-000000000000"
                                          --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                          --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-static-public-ip-addresses-usage**
```
./check-azurerm-static-public-ip-addresses-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-static-public-ip-addresses-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                                 --client_id="00000000-0000-0000-0000-000000000000"
                                                 --secret="00000000-0000-0000-0000-000000000000"
                                                 --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                                 --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-route-tables-usage**
```
./check-azurerm-route-tables-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-route-tables-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                   --client_id="00000000-0000-0000-0000-000000000000"
                                   --secret="00000000-0000-0000-0000-000000000000"
                                   --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                   --location="eastus2" --warning=80 --critical=90
```

**check-azurerm-service-bus-subscription**
```
./check-azurerm-service-bus-subscription
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
                           --tenant_id="00000000-0000-0000-0000-000000000000"
                           --client_id="00000000-0000-0000-0000-000000000000"
                           --secret="00000000-0000-0000-0000-000000000000"
                           --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                           --resource_group="resourcegroup"
                           --namespace="namespace"
                           --topic="topic"
                           --warning=60
                           --critical=80
```

**check-azurerm-virtual-machines-usage**
```
./check-azurerm-virtual-machines-usage --location="westeurope" --warning=80 --critical=90

./check-azurerm-virtual-machines-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                       --client_id="00000000-0000-0000-0000-000000000000"
                                       --secret="00000000-0000-0000-0000-000000000000"
                                       --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                       --location="eastus2" --warning=80 --critical=90
```


**check-azurerm-virtual-network-gateway-connected**
```
./check-azurerm-virtual-network-gateway-connected --resource_group="resourcegroup" --name="gatewayname"

./check-azurerm-virtual-network-gateway-connected
                               --tenant_id="00000000-0000-0000-0000-000000000000"
                               --client_id="00000000-0000-0000-0000-000000000000"
                               --secret="00000000-0000-0000-0000-000000000000"
                               --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                               --resource_group="resourcegroup" --name="gatewayname"
```

**check-azurerm-virtual-network-gateway-failover-connected**
```
./check-azurerm-virtual-network-gateway-failover-connected
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

./check-azurerm-virtual-networks-usage --tenant_id="00000000-0000-0000-0000-000000000000"
                                       --client_id="00000000-0000-0000-0000-000000000000"
                                       --secret="00000000-0000-0000-0000-000000000000"
                                       --subscription="1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345678901234"
                                       --location="eastus2" --warning=80 --critical=90
```

**metric-azurerm-service-bus-subscription-message-count**
```
./metric-azurerm-service-bus-subscription-message-count
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
```
