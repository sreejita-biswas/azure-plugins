package authorization

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sreejita-biswas/azure-plugins/constants"
)

var (
	authLocation string
	authType     int
	clientID     string
	tenantID     string
	clientSecret string
)

func GetParameterValues(cmd *cobra.Command) {
	cmd.Flags().IntVar(&authType, "auth_type", 0, "authentication type : 0 - Client Based , 1 - File Based Authentication, values other than 0 or 1  - In valid auth type")
	cmd.Flags().StringVar(&authLocation, "azure_auth_location", "", "path to the auth credential file we created before")
	cmd.Flags().StringVar(&clientID, "client_id", "", "ARM Client ID")
	cmd.Flags().StringVar(&tenantID, "tenant_id", "", "ARM Tenant ID")
	cmd.Flags().StringVar(&clientSecret, "secret", "", "ARM Client Secret")
}

func isFileBasedAuthentication() (bool, error) {
	if authType != constants.FileBasedAuthentication {
		return false, nil
	}
	if len(authLocation) > 0 {
		os.Setenv("AZURE_AUTH_LOCATION", authLocation)
		return true, nil
	}
	location := os.Getenv("AZURE_AUTH_LOCATION")
	if len(location) > 0 {
		fmt.Println("WARNING:Azure credential location path not provided, but AZURE_AUTH_LOCATION environment variable was set")
		return true, nil
	}

	fmt.Println("set AZURE_AUTH_LOCATION enviroment variable or provide path to the auth credential file, e.g. /Users/xyz/.azure/credential.auth")
	return true, errors.New("aure location path not yet provided")
}

func isClientBasedAuthentication() (bool, error) {
	if authType != constants.ClientBasedAuthentication {
		return false, nil
	}
	if len(clientID) > 0 && len(tenantID) > 0 && len(clientSecret) > 0 {
		os.Setenv("AZURE_TENANT_ID", tenantID)
		os.Setenv("AZURE_CLIENT_ID", clientID)
		os.Setenv("AZURE_CLIENT_SECRET", clientSecret)
		return true, nil
	}
	if len(os.Getenv("AZURE_TENANT_ID")) > 0 && len(os.Getenv("AZURE_CLIENT_ID")) > 0 && len(os.Getenv("AZURE_CLIENT_SECRET")) > 0 {
		fmt.Println("WARNING:reading tenant id, client id and client secret from environment variables")
		return true, nil
	}

	fmt.Println("set AZURE_TENANT_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET enviroment variables or provide tenant_id, client_id and secret")
	return true, errors.New("tenant_id amd/or client_id and/or client secret not yet provided")
}

func validateAuthenticationParameters() bool {
	fileBasedAuthentication, err := isFileBasedAuthentication()
	if fileBasedAuthentication && err != nil {
		return false
	}
	if fileBasedAuthentication && err == nil {
		return true
	}
	clientBasedAuthentication, err := isClientBasedAuthentication()
	if clientBasedAuthentication && err != nil {
		return false
	}
	if clientBasedAuthentication && err == nil {
		return true
	}
	fmt.Println("Authentication type : 0 - Client Based , 1 - File Based Authentication, values other than 0 or 1  - Invalid auth type")
	return false
}
