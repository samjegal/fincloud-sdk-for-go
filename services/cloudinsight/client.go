// Package cloudinsight implements the Azure ARM Cloudinsight service API version 0.0.1.
//
// Cloud Insight Client
package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Cloudinsight
	DefaultBaseURI = "https://cw.apigw.fin-ntruss.com"
)

// BaseClient is the base client for Cloudinsight.
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

func NewWithKey(accessKey string, secretKey string, apiGatewayKey string) BaseClient {
	return NewWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey, apiGatewayKey)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

func NewWithBaseURIWithKey(baseURI string, accessKey string, secretKey string, apiGatewayKey string) BaseClient {
	return BaseClient{
		Client:           autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:          baseURI,
		AccessKey:        accessKey,
		Secretkey:        secretKey,
		APIGatewayAPIKey: apiGatewayKey,
	}
}
