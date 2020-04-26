// Package outboundmailer implements the Azure ARM Outboundmailer service API version 0.0.1.
//
// Cloud Outbound Mailer Client
package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Outboundmailer
	DefaultBaseURI = "https://mail.apigw.fin-ntruss.com/api/v1"
)

// BaseClient is the base client for Outboundmailer.
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
