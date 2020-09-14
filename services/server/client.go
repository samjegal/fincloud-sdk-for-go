// Package server implements the Azure ARM Server service API version 1.0.0.
//
// Server Client
package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Server
	DefaultBaseURI = "https://fin-ncloud.apigw.fin-ntruss.com/vserver/v2"
)

// BaseClient is the base client for Server.
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
