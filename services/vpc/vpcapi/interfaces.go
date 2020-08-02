package vpcapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
)

// NatGatewayClientAPI contains the set of methods on the NatGatewayClient type.
type NatGatewayClientAPI interface {
	Create(ctx context.Context, responseFormatType string, vpcNo string, zoneCode string, regionCode string, natGatewayName string, natGatewayDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, natGatewayInstanceNoListN string, publicIP string, vpcName string, natGatewayName string, natGatewayInstanceStatusCode vpc.NatGatewayInstanceStatusCode, pageNo string, pageSize string) (result autorest.Response, err error)
}

var _ NatGatewayClientAPI = (*vpc.NatGatewayClient)(nil)

// NetworkACLClientAPI contains the set of methods on the NetworkACLClient type.
type NetworkACLClientAPI interface {
	AddInboundRule(ctx context.Context, responseFormatType string, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result autorest.Response, err error)
	AddOutboundRule(ctx context.Context, responseFormatType string, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, vpcNo string, regionCode string, networkACLName string, networkACLDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, networkACLNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, networkACLNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, networkACLName string, networkACLStatusCode vpc.NetworkACLStatusCode, networkACLNoListN string, pageNo string, pageSize string, vpcNo string) (result autorest.Response, err error)
	GetRuleList(ctx context.Context, responseFormatType string, networkACLNo string, regionCode string, networkACLRuleTypeCode vpc.NetworkACLRuleTypeCode) (result autorest.Response, err error)
	RemoveInboundRule(ctx context.Context, responseFormatType string, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result autorest.Response, err error)
	RemoveOutboundRule(ctx context.Context, responseFormatType string, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode vpc.ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode vpc.RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result autorest.Response, err error)
	SetSubnet(ctx context.Context, responseFormatType string, networkACLNo string, subnetNo string, regionCode string) (result autorest.Response, err error)
}

var _ NetworkACLClientAPI = (*vpc.NetworkACLClient)(nil)

// RouteTableClientAPI contains the set of methods on the RouteTableClient type.
type RouteTableClientAPI interface {
	Create(ctx context.Context, responseFormatType string, vpcNo string, regionCode string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, routeTableDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, routeTableNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, routeTableNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, routeTableNoListN string, routeTableName string, supportedSubnetTypeCode vpc.SupportedSubnetTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder vpc.SortingOrder, vpcNo string) (result autorest.Response, err error)
}

var _ RouteTableClientAPI = (*vpc.RouteTableClient)(nil)

// RouteClientAPI contains the set of methods on the RouteClient type.
type RouteClientAPI interface {
	Add(ctx context.Context, responseFormatType string, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, routeTableNo string, vpcNo string, regionCode string) (result autorest.Response, err error)
	Remove(ctx context.Context, responseFormatType string, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode vpc.TargetTypeCode, routeListNtargetNo string, routeListNtargetName string, regionCode string) (result autorest.Response, err error)
}

var _ RouteClientAPI = (*vpc.RouteClient)(nil)

// RouteTableSubnetClientAPI contains the set of methods on the RouteTableSubnetClient type.
type RouteTableSubnetClientAPI interface {
	Add(ctx context.Context, responseFormatType string, vpcNo string, routeTableNo string, subnetNoListN string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, routeTableNo string, regionCode string) (result autorest.Response, err error)
	Remove(ctx context.Context, responseFormatType string, vpcNo string, routeTableNo string, subnetNoListN string, regionCode string) (result autorest.Response, err error)
}

var _ RouteTableSubnetClientAPI = (*vpc.RouteTableSubnetClient)(nil)

// SubnetClientAPI contains the set of methods on the SubnetClient type.
type SubnetClientAPI interface {
	Create(ctx context.Context, responseFormatType string, zoneCode string, vpcNo string, subnet string, networkACLNo string, subnetTypeCode vpc.SubnetTypeCode, regionCode string, subnetName string, usageTypeCode vpc.UsageTypeCode) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, subnetNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, subnetNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, subnetNoListN string, subnetName string, subnet string, subnetTypeCode vpc.SubnetTypeCode, usageTypeCode vpc.UsageTypeCode, networkACLNo string, pageNo string, pageSize string, subnetStatusCode vpc.SubnetStatusCode, vpcNo string, zoneCode string) (result autorest.Response, err error)
}

var _ SubnetClientAPI = (*vpc.SubnetClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	Create(ctx context.Context, ipv4CidrBlock string, responseFormatType string, regionCode string, vpcName string) (result autorest.Response, err error)
	Delete(ctx context.Context, vpcNo string, responseFormatType string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, vpcNo string, responseFormatType string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, vpcStatusCode vpc.StatusCode, vpcName string, vpcNoListN string) (result autorest.Response, err error)
}

var _ ClientAPI = (*vpc.Client)(nil)

// PeeringClientAPI contains the set of methods on the PeeringClient type.
type PeeringClientAPI interface {
	AcceptOrReject(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, isAccept string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, sourceVpcNo string, targetVpcNo string, regionCode string, vpcPeeringName string, targetVpcName string, targetVpcLoginID string, vpcPeeringDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, vpcPeeringInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, vpcPeeringInstanceNoListN string, sourceVpcName string, vpcPeeringName string, targetVpcName string, vpcPeeringInstanceStatusCode vpc.PeeringInstanceStatusCode, pageNo string, pageSize string, sortedBy vpc.SortedBy, sortingOrder vpc.SortingOrder) (result autorest.Response, err error)
}

var _ PeeringClientAPI = (*vpc.PeeringClient)(nil)
