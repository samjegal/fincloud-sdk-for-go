package vpcapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
)

// NatGatewayClientAPI contains the set of methods on the NatGatewayClient type.
type NatGatewayClientAPI interface {
	Create(ctx context.Context, vpcNo string, zoneCode string, natGatewayName string, natGatewayDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, natGatewayInstanceNo string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, natGatewayInstanceNo string) (result autorest.Response, err error)
	GetList(ctx context.Context, natGatewayInstanceNoListN string, publicIP string, vpcName string, natGatewayName string, natGatewayInstanceStatusCode vpc.NatGatewayInstanceStatusCode, pageNo string, pageSize string) (result autorest.Response, err error)
}

var _ NatGatewayClientAPI = (*vpc.NatGatewayClient)(nil)

// NetworkACLClientAPI contains the set of methods on the NetworkACLClient type.
type NetworkACLClientAPI interface {
	AddInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result vpc.NetworkACLInboundRuleResponse, err error)
	AddOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result vpc.NetworkACLOutboundRuleResponse, err error)
	Create(ctx context.Context, vpcNo string, networkACLName string, networkACLDescription string) (result vpc.NetworkACLResponse, err error)
	Delete(ctx context.Context, networkACLNo string) (result vpc.NetworkACLResponse, err error)
	GetDetail(ctx context.Context, networkACLNo string) (result vpc.NetworkACLDetailResponse, err error)
	GetList(ctx context.Context, networkACLName string, networkACLStatusCode vpc.NetworkACLStatusCode, networkACLNoListN string, pageNo string, pageSize string, vpcNo string) (result vpc.NetworkACLListResponse, err error)
	GetRuleList(ctx context.Context, networkACLNo string, networkACLRuleTypeCode vpc.NetworkACLRuleTypeCode) (result vpc.NetworkACLRuleListResponse, err error)
	RemoveInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, networkACLRuleListNportRange string) (result vpc.NetworkACLInboundRuleResponse, err error)
	RemoveOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, networkACLRuleListNportRange string) (result vpc.NetworkACLOutboundRuleResponse, err error)
	SetSubnet(ctx context.Context, networkACLNo string, subnetNo string) (result vpc.SubnetNetworkACLResponse, err error)
}

var _ NetworkACLClientAPI = (*vpc.NetworkACLClient)(nil)

// RouteTableClientAPI contains the set of methods on the RouteTableClient type.
type RouteTableClientAPI interface {
	Create(ctx context.Context, vpcNo string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, routeTableDescription string) (result vpc.RouteTableResponse, err error)
	Delete(ctx context.Context, routeTableNo string) (result vpc.RouteTableResponse, err error)
	GetDetail(ctx context.Context, routeTableNo string) (result vpc.RouteTableDetailResponse, err error)
	GetList(ctx context.Context, routeTableNoListN string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder vpc.SortingOrder, vpcNo string) (result vpc.RouteTableListResponse, err error)
}

var _ RouteTableClientAPI = (*vpc.RouteTableClient)(nil)

// RouteClientAPI contains the set of methods on the RouteClient type.
type RouteClientAPI interface {
	Add(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (result vpc.RouteResponse, err error)
	GetList(ctx context.Context, routeTableNo string, vpcNo string) (result vpc.RouteListResponse, err error)
	Remove(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (result vpc.RouteResponse, err error)
}

var _ RouteClientAPI = (*vpc.RouteClient)(nil)

// RouteTableSubnetClientAPI contains the set of methods on the RouteTableSubnetClient type.
type RouteTableSubnetClientAPI interface {
	Add(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (result vpc.RouteTableSubnetResponse, err error)
	GetList(ctx context.Context, routeTableNo string) (result vpc.RouteTableSubnetListResponse, err error)
	Remove(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (result vpc.RouteTableSubnetResponse, err error)
}

var _ RouteTableSubnetClientAPI = (*vpc.RouteTableSubnetClient)(nil)

// SubnetClientAPI contains the set of methods on the SubnetClient type.
type SubnetClientAPI interface {
	Create(ctx context.Context, zoneCode string, vpcNo string, subnet string, networkACLNo string, subnetTypeCode vpc.SubnetTypeCode, subnetName string, usageTypeCode vpc.UsageTypeCode) (result vpc.SubnetResponse, err error)
	Delete(ctx context.Context, subnetNo string) (result vpc.SubnetResponse, err error)
	GetDetail(ctx context.Context, subnetNo string) (result vpc.SubnetDetailResponse, err error)
	GetList(ctx context.Context, subnetNoListN string, subnetName string, subnet string, subnetTypeCode vpc.SubnetTypeCode, usageTypeCode vpc.UsageTypeCode, networkACLNo string, pageNo string, pageSize string, subnetStatusCode vpc.SubnetStatusCode, vpcNo string, zoneCode string) (result vpc.SubnetListResponse, err error)
}

var _ SubnetClientAPI = (*vpc.SubnetClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, ipv4CidrBlock string, vpcName string) (result vpc.Response, err error)
	Delete(ctx context.Context, vpcNo string) (result vpc.Response, err error)
	GetDetail(ctx context.Context, vpcNo string) (result vpc.DetailResponse, err error)
	GetList(ctx context.Context, vpcStatusCode vpc.StatusCode, vpcName string, vpcNoListN string) (result vpc.ListResponse, err error)
}

var _ ClientAPI = (*vpc.Client)(nil)

// PeeringClientAPI contains the set of methods on the PeeringClient type.
type PeeringClientAPI interface {
	AcceptOrReject(ctx context.Context, vpcPeeringInstanceNo string, isAccept string) (result vpc.PeeringInstanceResponse, err error)
	Create(ctx context.Context, sourceVpcNo string, targetVpcNo string, vpcPeeringName string, targetVpcName string, targetVpcLoginID string, vpcPeeringDescription string) (result vpc.PeeringInstanceResponse, err error)
	Delete(ctx context.Context, vpcPeeringInstanceNo string) (result vpc.PeeringInstanceResponse, err error)
	GetDetail(ctx context.Context, vpcPeeringInstanceNo string) (result autorest.Response, err error)
	GetList(ctx context.Context, vpcPeeringInstanceNoListN string, sourceVpcName string, vpcPeeringName string, targetVpcName string, vpcPeeringInstanceStatusCode vpc.PeeringInstanceStatusCode, pageNo string, pageSize string, sortedBy vpc.SortedBy, sortingOrder vpc.SortingOrder) (result autorest.Response, err error)
}

var _ PeeringClientAPI = (*vpc.PeeringClient)(nil)
