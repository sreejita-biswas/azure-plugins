package constants

type Client int

const (
	ClientBasedAuthentication          = 0
	FileBasedAuthentication            = 1
	UsageClient                 Client = 0
	VirtualNetwrokGatewayClient Client = 1
	SubscriptionClient          Client = 2
	TopicClient                 Client = 3
)

func GetUnitByUsageStatisticType(unit string) string {
	if unit == "cores" {
		return "Cores"
	} else if unit == "standardDFamily" {
		return "D Family Cores"
	} else if unit == "standardDSFamily" {
		return "DS Family Cores"
	} else if unit == "standardDSv2Family" {
		return "DSv2 Family Cores"
	} else if unit == "standardFFamily" {
		return "F Family Cores"
	} else if unit == "standardFSFamily" {
		return "FS Family Cores"
	} else if unit == "standardDv2Family" {
		return "Dv2 Family Cores"
	}
	return ""
}
