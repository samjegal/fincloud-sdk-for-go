package vpcapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
)

// NatGatewayClientAPI contains the set of methods on the NatGatewayClient type.
type NatGatewayClientAPI interface {
	Create(ctx context.Context, vpcNo string, zoneCode string, regionCode string, natGatewayName string, natGatewayDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, regionCode string, natGatewayInstanceNoListN string, publicIP string, vpcName string, natGatewayName string, natGatewayInstanceStatusCode vpc.NatGatewayInstanceStatusCode, pageNo string, pageSize string) (result autorest.Response, err error)
}

var _ NatGatewayClientAPI = (*vpc.NatGatewayClient)(nil)

// NetworkACLClientAPI contains the set of methods on the NetworkACLClient type.
type NetworkACLClientAPI interface {
	AddInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result vpc.NetworkACLInboundRuleResponse, err error)
	AddOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result vpc.NetworkACLOutboundRuleResponse, err error)
	Create(ctx context.Context, vpcNo string, regionCode string, networkACLName string, networkACLDescription string) (result vpc.NetworkACLResponse, err error)
	Delete(ctx context.Context, networkACLNo string, regionCode string) (result vpc.NetworkACLResponse, err error)
	GetDetail(ctx context.Context, networkACLNo string, regionCode string) (result vpc.NetworkACLDetailResponse, err error)
	GetList(ctx context.Context, regionCode string, networkACLName string, networkACLStatusCode vpc.NetworkACLStatusCode, networkACLNoListN string, pageNo string, pageSize string, vpcNo string) (result vpc.NetworkACLListResponse, err error)
	GetRuleList(ctx context.Context, networkACLNo string, regionCode string, networkACLRuleTypeCode vpc.NetworkACLRuleTypeCode) (result vpc.NetworkACLRuleListResponse, err error)
	RemoveInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result vpc.NetworkACLInboundRuleResponse, err error)
	RemoveOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result vpc.NetworkACLOutboundRuleResponse, err error)
	SetSubnet(ctx context.Context, networkACLNo string, subnetNo string, regionCode string) (result vpc.SubnetNetworkACLResponse, err error)
}

var _ NetworkACLClientAPI = (*vpc.NetworkACLClient)(nil)

// RouteTableClientAPI contains the set of methods on the RouteTableClient type.
type RouteTableClientAPI interface {
	Create(ctx context.Context, vpcNo string, regionCode string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, routeTableDescription string) (result vpc.RouteTableResponse, err error)
	Delete(ctx context.Context, routeTableNo string, regionCode string) (result vpc.RouteTableResponse, err error)
	GetDetail(ctx context.Context, routeTableNo string, regionCode string) (result vpc.RouteTableDetailResponse, err error)
	GetList(ctx context.Context, regionCode string, routeTableNoListN string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder vpc.SortingOrder, vpcNo string) (result vpc.RouteTableListResponse, err error)
}

var _ RouteTableClientAPI = (*vpc.RouteTableClient)(nil)

// RouteClientAPI contains the set of methods on the RouteClient type.
type RouteClientAPI interface {
	Add(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string, regionCode string) (result vpc.RouteResponse, err error)
	GetList(ctx context.Context, routeTableNo string, vpcNo string, regionCode string) (result vpc.RouteListResponse, err error)
	Remove(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string, regionCode string) (result vpc.RouteResponse, err error)
}

var _ RouteClientAPI = (*vpc.RouteClient)(nil)

// RouteTableSubnetClientAPI contains the set of methods on the RouteTableSubnetClient type.
type RouteTableSubnetClientAPI interface {
	Add(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string, regionCode string) (result vpc.RouteTableSubnetResponse, err error)
	GetList(ctx context.Context, routeTableNo string, regionCode string) (result vpc.RouteTableSubnetListResponse, err error)
	Remove(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string, regionCode string) (result vpc.RouteTableSubnetResponse, err error)
}

var _ RouteTableSubnetClientAPI = (*vpc.RouteTableSubnetClient)(nil)

// SubnetClientAPI contains the set of methods on the SubnetClient type.
type SubnetClientAPI interface {
	Create(ctx context.Context, zoneCode string, vpcNo string, subnet string, networkACLNo string, subnetTypeCode vpc.SubnetTypeCode, regionCode string, subnetName string, usageTypeCode vpc.UsageTypeCode) (result vpc.SubnetResponse, err error)
	Delete(ctx context.Context, subnetNo string, regionCode string) (result vpc.SubnetResponse, err error)
	GetDetail(ctx context.Context, subnetNo string, regionCode string) (result vpc.SubnetDetailResponse, err error)
	GetList(ctx context.Context, regionCode string, subnetNoListN string, subnetName string, subnet string, subnetTypeCode vpc.SubnetTypeCode, usageTypeCode vpc.UsageTypeCode, networkACLNo string, pageNo string, pageSize string, subnetStatusCode vpc.SubnetStatusCode, vpcNo string, zoneCode string) (result vpc.SubnetListResponse, err error)
}

var _ SubnetClientAPI = (*vpc.SubnetClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, ipv4CidrBlock string, regionCode string, vpcName string) (result vpc.Response, err error)
	Delete(ctx context.Context, vpcNo string, regionCode string) (result vpc.Response, err error)
	GetDetail(ctx context.Context, vpcNo string, regionCode string) (result vpc.DetailResponse, err error)
	GetList(ctx context.Context, regionCode string, vpcStatusCode vpc.StatusCode, vpcName string, vpcNoListN string) (result vpc.ListResponse, err error)
}

var _ ClientAPI = (*vpc.Client)(nil)

// PeeringClientAPI contains the set of methods on the PeeringClient type.
type PeeringClientAPI interface {
	AcceptOrReject(ctx context.Context, vpcPeeringInstanceNo string, isAccept string, regionCode string) (result vpc.PeeringInstanceResponse, err error)
	Create(ctx context.Context, sourceVpcNo string, targetVpcNo string, regionCode string, vpcPeeringName string, targetVpcName string, targetVpcLoginID string, vpcPeeringDescription string) (result vpc.PeeringInstanceResponse, err error)
	Delete(ctx context.Context, vpcPeeringInstanceNo string, regionCode string) (result vpc.PeeringInstanceResponse, err error)
	GetDetail(ctx context.Context, vpcPeeringInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, regionCode string, vpcPeeringInstanceNoListN string, sourceVpcName string, vpcPeeringName string, targetVpcName string, vpcPeeringInstanceStatusCode vpc.PeeringInstanceStatusCode, pageNo string, pageSize string, sortedBy vpc.SortedBy, sortingOrder vpc.SortingOrder) (result autorest.Response, err error)
}

var _ PeeringClientAPI = (*vpc.PeeringClient)(nil)
