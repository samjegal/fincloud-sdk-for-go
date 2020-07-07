// Package geolocation implements the Azure ARM Geolocation service API version 0.0.1.
//
// GeoLocation Client
package geolocation

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Geolocation
	DefaultBaseURI = "https://geolocation.apigw.fin-ntruss.com/geolocation/v2"
)

// BaseClient is the base client for Geolocation.
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
