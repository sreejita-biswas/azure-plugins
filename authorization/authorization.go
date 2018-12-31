package authorization

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
	"github.com/sreejita-biswas/azure-plugins/constants"
)

func newSessionFromFile() (autorest.Authorizer, bool) {
	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't initialize authorizer"))
		return nil, false
	}
	return authorizer, true
}

func newSessionFromCredentials() (autorest.Authorizer, bool) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Can't initialize authorizer"))
		return nil, false
	}
	return authorizer, true
}

func Authorize() (autorest.Authorizer, bool) {
	success := validateAuthenticationParameters()
	if !success {
		return nil, false
	}
	if authType == constants.ClientBasedAuthentication {
		return newSessionFromCredentials()
	}
	if authType == constants.FileBasedAuthentication {
		return newSessionFromFile()
	}
	fmt.Println("Authentication type : 0 - Client Based , 1 - File Based Authentication, values other than 0 or 1  - Invalid auth type")
	return nil, false
}
