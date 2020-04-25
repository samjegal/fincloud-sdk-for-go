package sslvpnapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/sslvpn"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	List(ctx context.Context) (result sslvpn.Parameter, err error)
	UpdateSpec(ctx context.Context, parameters sslvpn.LimitUserCountParameter) (result autorest.Response, err error)
}

var _ ClientAPI = (*sslvpn.Client)(nil)
