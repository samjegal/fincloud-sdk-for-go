// Package kubernetes implements the Azure ARM Kubernetes service API version 2.0.0.
//
// Kubernetes Client
package kubernetes

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Kubernetes
	DefaultBaseURI = "https://nks.apigw.fin-ntruss.com/nks/v2"
)

// BaseClient is the base client for Kubernetes.
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

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}
