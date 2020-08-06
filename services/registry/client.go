// Package registry implements the Azure ARM Registry service API version 2020-02-19T12:58:42Z.
//
// Container Registry Client
package registry

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Registry
	DefaultBaseURI = "https://ncr.apigw.fin-ntruss.com/ncr/api"
)

// BaseClient is the base client for Registry.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}
