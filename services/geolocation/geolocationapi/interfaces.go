package geolocationapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/geolocation"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Get(ctx context.Context, IP string, enc string, ext string, responseFormatType string) (result geolocation.ResultParameter, err error)
}

var _ ClientAPI = (*geolocation.Client)(nil)

// QuotaClientAPI contains the set of methods on the QuotaClient type.
type QuotaClientAPI interface {
	Get(ctx context.Context) (result geolocation.QuotaParameter, err error)
	Set(ctx context.Context, quota string) (result autorest.Response, err error)
}

var _ QuotaClientAPI = (*geolocation.QuotaClient)(nil)
