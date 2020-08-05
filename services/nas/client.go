// Package nas implements the Azure ARM Nas service API version 1.0.0.
//
// NAS Client
package nas

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Nas
	DefaultBaseURI = "https://fin-ncloud.apigw.fin-ntruss.com/vnas/v2"
)

// BaseClient is the base client for Nas.
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