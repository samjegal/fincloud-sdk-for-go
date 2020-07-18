package wmsapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/samjegal/fincloud-sdk-for-go/services/wms"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Detail(ctx context.Context, scriptID string) (result wms.SetObject, err error)
	List(ctx context.Context) (result wms.SetObject, err error)
	Result(ctx context.Context, scriptID string, timestamp string) (result wms.SetObject, err error)
	Start(ctx context.Context, scriptID string) (result wms.ReturnParameter, err error)
	Stop(ctx context.Context, scriptID string) (result wms.ReturnParameter, err error)
}

var _ ClientAPI = (*wms.Client)(nil)
