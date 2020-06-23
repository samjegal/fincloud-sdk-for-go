package networkapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/network"
)

// VirtualPrivateCloudClientAPI contains the set of methods on the VirtualPrivateCloudClient type.
type VirtualPrivateCloudClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters network.VirtualPrivateCloudParameter) (result network.VirtualPrivateCloudListParameter, err error)
	Delete(ctx context.Context, vpcNo string, parameters network.VirtualPrivateCloudParameter) (result autorest.Response, err error)
	List(ctx context.Context) (result network.VirtualPrivateCloudListParameter, err error)
}

var _ VirtualPrivateCloudClientAPI = (*network.VirtualPrivateCloudClient)(nil)

// ACLClientAPI contains the set of methods on the ACLClient type.
type ACLClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters network.ACLRequestParameter) (result network.ACLResponseParameter, err error)
	Delete(ctx context.Context, networkACLNo string) (result autorest.Response, err error)
	Get(ctx context.Context, networkACLNo string) (result network.ACLDetailParameter, err error)
	List(ctx context.Context) (result network.ACLResponseParameter, err error)
}

var _ ACLClientAPI = (*network.ACLClient)(nil)

// InboundRuleClientAPI contains the set of methods on the InboundRuleClient type.
type InboundRuleClientAPI interface {
	Get(ctx context.Context, networkACLNo string) (result network.ACLRuleContentParameter, err error)
}

var _ InboundRuleClientAPI = (*network.InboundRuleClient)(nil)

// OutboundRuleClientAPI contains the set of methods on the OutboundRuleClient type.
type OutboundRuleClientAPI interface {
	Get(ctx context.Context, networkACLNo string) (result network.ACLRuleContentParameter, err error)
}

var _ OutboundRuleClientAPI = (*network.OutboundRuleClient)(nil)

// RuleClientAPI contains the set of methods on the RuleClient type.
type RuleClientAPI interface {
	CreateOrUpdate(ctx context.Context, networkACLNo string, parameters network.ACLRuleListParameter) (result autorest.Response, err error)
}

var _ RuleClientAPI = (*network.RuleClient)(nil)

// SubnetClientAPI contains the set of methods on the SubnetClient type.
type SubnetClientAPI interface {
	Create(ctx context.Context, parameters network.SubnetParameter) (result network.ErrorMessageParameter, err error)
	Delete(ctx context.Context, subnetNo string) (result network.SubnetSearchListParameter, err error)
	List(ctx context.Context, parameters network.SubnetSearchParameter) (result network.SubnetSearchListParameter, err error)
}

var _ SubnetClientAPI = (*network.SubnetClient)(nil)

// NatGatewayClientAPI contains the set of methods on the NatGatewayClient type.
type NatGatewayClientAPI interface {
	Create(ctx context.Context, parameters network.NatGatewayParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, natGatewayNo string, parameters interface{}) (result autorest.Response, err error)
	List(ctx context.Context, parameters network.NatGatewaySearchParameter) (result network.NatGatewaySearchListParameter, err error)
}

var _ NatGatewayClientAPI = (*network.NatGatewayClient)(nil)

// RouteTableClientAPI contains the set of methods on the RouteTableClient type.
type RouteTableClientAPI interface {
	Create(ctx context.Context, parameters network.RouteTableParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, routeTableNo string) (result autorest.Response, err error)
	Get(ctx context.Context) (result network.RouteTableRuleSearchParameter, err error)
	GetEndpoints(ctx context.Context) (result network.RouteTableEndpointsParameter, err error)
	GetEndpointTypes(ctx context.Context) (result network.RouteTableEndpointTypesParameter, err error)
	IsSetted(ctx context.Context) (result network.RouteTableIsSettedParameter, err error)
	List(ctx context.Context, parameters network.RouteTableSearchParameter) (result network.RouteTableSearchListParameter, err error)
	Update(ctx context.Context, routeTableNo string, parameters network.RouteTableRuleParameter) (result autorest.Response, err error)
}

var _ RouteTableClientAPI = (*network.RouteTableClient)(nil)

// RouteTableSubnetClientAPI contains the set of methods on the RouteTableSubnetClient type.
type RouteTableSubnetClientAPI interface {
	List(ctx context.Context, routeTableNo string) (result network.RouteTableSubnetListParameter, err error)
	Update(ctx context.Context, routeTableNo string, parameters network.RouteTableSubnetParameter) (result autorest.Response, err error)
}

var _ RouteTableSubnetClientAPI = (*network.RouteTableSubnetClient)(nil)

// RouteTableDescriptionClientAPI contains the set of methods on the RouteTableDescriptionClient type.
type RouteTableDescriptionClientAPI interface {
	Update(ctx context.Context, routeTableNo string, parameters network.RouteTableDescriptionParameter) (result autorest.Response, err error)
}

var _ RouteTableDescriptionClientAPI = (*network.RouteTableDescriptionClient)(nil)

// LoadBalancerClientAPI contains the set of methods on the LoadBalancerClient type.
type LoadBalancerClientAPI interface {
	Search(ctx context.Context, parameters network.LoadBalancerSearchParameter) (result network.LoadBalancerSearchListParameter, err error)
}

var _ LoadBalancerClientAPI = (*network.LoadBalancerClient)(nil)
