// Package sens implements the Azure ARM Sens service API version 2019-10-23T06:20:43Z.
//
// SENS Client
package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Sens
	DefaultBaseURI = "https://sens.apigw.fin-ntruss.com"
)

// BaseClient is the base client for Sens.
type BaseClient struct {
	autorest.Client
	BaseURI string

	AccessKey string
	Secretkey string

	APIGatewayAPIKey string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

func NewWithKey(accessKey string, secretKey string) BaseClient {
	return NewWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

func NewWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) BaseClient {
	return BaseClient{
		Client:    autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:   baseURI,
		AccessKey: accessKey,
		Secretkey: secretKey,
	}
}
