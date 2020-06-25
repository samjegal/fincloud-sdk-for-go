package loadbalancerapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/loadbalancer"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckName(ctx context.Context, loadBalancerName string) (result loadbalancer.CheckNameParameter, err error)
	Create(ctx context.Context, parameters interface{}) (result autorest.Response, err error)
	Delete(ctx context.Context, parameters interface{}) (result autorest.Response, err error)
	Search(ctx context.Context, parameters loadbalancer.SearchParameter) (result loadbalancer.SearchListParameter, err error)
	ServerInstance(ctx context.Context, vpcNo string, layerTypeCode string) (result loadbalancer.ServerInstanceListParameter, err error)
	Update(ctx context.Context, parameters interface{}) (result autorest.Response, err error)
	ZoneSubnet(ctx context.Context, vpcNo string) (result loadbalancer.ZoneSubnetParameter, err error)
}

var _ ClientAPI = (*loadbalancer.Client)(nil)

// ServerClientAPI contains the set of methods on the ServerClient type.
type ServerClientAPI interface {
	Update(ctx context.Context, parameters interface{}) (result autorest.Response, err error)
}

var _ ServerClientAPI = (*loadbalancer.ServerClient)(nil)

// ListenerClientAPI contains the set of methods on the ListenerClient type.
type ListenerClientAPI interface {
	Update(ctx context.Context, parameters interface{}) (result autorest.Response, err error)
}

var _ ListenerClientAPI = (*loadbalancer.ListenerClient)(nil)
