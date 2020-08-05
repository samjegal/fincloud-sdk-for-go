package vpc

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/vpc"

// NatGatewayInstanceStatusCode enumerates the values for nat gateway instance status code.
type NatGatewayInstanceStatusCode string

const (
	// INIT ...
	INIT NatGatewayInstanceStatusCode = "INIT"
	// RUN ...
	RUN NatGatewayInstanceStatusCode = "RUN"
	// TERMTING ...
	TERMTING NatGatewayInstanceStatusCode = "TERMTING"
)

// PossibleNatGatewayInstanceStatusCodeValues returns an array of possible values for the NatGatewayInstanceStatusCode const type.
func PossibleNatGatewayInstanceStatusCodeValues() []NatGatewayInstanceStatusCode {
	return []NatGatewayInstanceStatusCode{INIT, RUN, TERMTING}
}

// NetworkACLRuleTypeCode enumerates the values for network acl rule type code.
type NetworkACLRuleTypeCode string

const (
	// INBND ...
	INBND NetworkACLRuleTypeCode = "INBND"
	// OTBND ...
	OTBND NetworkACLRuleTypeCode = "OTBND"
)

// PossibleNetworkACLRuleTypeCodeValues returns an array of possible values for the NetworkACLRuleTypeCode const type.
func PossibleNetworkACLRuleTypeCodeValues() []NetworkACLRuleTypeCode {
	return []NetworkACLRuleTypeCode{INBND, OTBND}
}

// NetworkACLStatusCode enumerates the values for network acl status code.
type NetworkACLStatusCode string

const (
	// NetworkACLStatusCodeINIT ...
	NetworkACLStatusCodeINIT NetworkACLStatusCode = "INIT"
	// NetworkACLStatusCodeRUN ...
	NetworkACLStatusCodeRUN NetworkACLStatusCode = "RUN"
	// NetworkACLStatusCodeSET ...
	NetworkACLStatusCodeSET NetworkACLStatusCode = "SET"
	// NetworkACLStatusCodeTERMTING ...
	NetworkACLStatusCodeTERMTING NetworkACLStatusCode = "TERMTING"
)

// PossibleNetworkACLStatusCodeValues returns an array of possible values for the NetworkACLStatusCode const type.
func PossibleNetworkACLStatusCodeValues() []NetworkACLStatusCode {
	return []NetworkACLStatusCode{NetworkACLStatusCodeINIT, NetworkACLStatusCodeRUN, NetworkACLStatusCodeSET, NetworkACLStatusCodeTERMTING}
}

// PeeringInstanceStatusCode enumerates the values for peering instance status code.
type PeeringInstanceStatusCode string

const (
	// PeeringInstanceStatusCodeINIT ...
	PeeringInstanceStatusCodeINIT PeeringInstanceStatusCode = "INIT"
	// PeeringInstanceStatusCodeRUN ...
	PeeringInstanceStatusCodeRUN PeeringInstanceStatusCode = "RUN"
	// PeeringInstanceStatusCodeTERMTING ...
	PeeringInstanceStatusCodeTERMTING PeeringInstanceStatusCode = "TERMTING"
)

// PossiblePeeringInstanceStatusCodeValues returns an array of possible values for the PeeringInstanceStatusCode const type.
func PossiblePeeringInstanceStatusCodeValues() []PeeringInstanceStatusCode {
	return []PeeringInstanceStatusCode{PeeringInstanceStatusCodeINIT, PeeringInstanceStatusCodeRUN, PeeringInstanceStatusCodeTERMTING}
}

// ProtocolTypeCode enumerates the values for protocol type code.
type ProtocolTypeCode string

const (
	// ICMP ...
	ICMP ProtocolTypeCode = "ICMP"
	// TCP ...
	TCP ProtocolTypeCode = "TCP"
	// UDP ...
	UDP ProtocolTypeCode = "UDP"
)

// PossibleProtocolTypeCodeValues returns an array of possible values for the ProtocolTypeCode const type.
func PossibleProtocolTypeCodeValues() []ProtocolTypeCode {
	return []ProtocolTypeCode{ICMP, TCP, UDP}
}

// RuleActionCode enumerates the values for rule action code.
type RuleActionCode string

const (
	// ALLOW ...
	ALLOW RuleActionCode = "ALLOW"
	// DROP ...
	DROP RuleActionCode = "DROP"
)

// PossibleRuleActionCodeValues returns an array of possible values for the RuleActionCode const type.
func PossibleRuleActionCodeValues() []RuleActionCode {
	return []RuleActionCode{ALLOW, DROP}
}

// SortedBy enumerates the values for sorted by.
type SortedBy string

const (
	// SourceVpcName ...
	SourceVpcName SortedBy = "sourceVpcName"
	// TargetVpcName ...
	TargetVpcName SortedBy = "targetVpcName"
	// VpcPeeringName ...
	VpcPeeringName SortedBy = "vpcPeeringName"
)

// PossibleSortedByValues returns an array of possible values for the SortedBy const type.
func PossibleSortedByValues() []SortedBy {
	return []SortedBy{SourceVpcName, TargetVpcName, VpcPeeringName}
}

// SortingOrder enumerates the values for sorting order.
type SortingOrder string

const (
	// ASC ...
	ASC SortingOrder = "ASC"
	// DESC ...
	DESC SortingOrder = "DESC"
)

// PossibleSortingOrderValues returns an array of possible values for the SortingOrder const type.
func PossibleSortingOrderValues() []SortingOrder {
	return []SortingOrder{ASC, DESC}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// StatusCodeCREATING ...
	StatusCodeCREATING StatusCode = "CREATING"
	// StatusCodeINIT ...
	StatusCodeINIT StatusCode = "INIT"
	// StatusCodeRUN ...
	StatusCodeRUN StatusCode = "RUN"
	// StatusCodeTERMTING ...
	StatusCodeTERMTING StatusCode = "TERMTING"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{StatusCodeCREATING, StatusCodeINIT, StatusCodeRUN, StatusCodeTERMTING}
}

// SubnetStatusCode enumerates the values for subnet status code.
type SubnetStatusCode string

const (
	// SubnetStatusCodeCREATING ...
	SubnetStatusCodeCREATING SubnetStatusCode = "CREATING"
	// SubnetStatusCodeINIT ...
	SubnetStatusCodeINIT SubnetStatusCode = "INIT"
	// SubnetStatusCodeRUN ...
	SubnetStatusCodeRUN SubnetStatusCode = "RUN"
	// SubnetStatusCodeTERMTING ...
	SubnetStatusCodeTERMTING SubnetStatusCode = "TERMTING"
)

// PossibleSubnetStatusCodeValues returns an array of possible values for the SubnetStatusCode const type.
func PossibleSubnetStatusCodeValues() []SubnetStatusCode {
	return []SubnetStatusCode{SubnetStatusCodeCREATING, SubnetStatusCodeINIT, SubnetStatusCodeRUN, SubnetStatusCodeTERMTING}
}

// SubnetTypeCode enumerates the values for subnet type code.
type SubnetTypeCode string

const (
	// PRIVATE ...
	PRIVATE SubnetTypeCode = "PRIVATE"
	// PUBLIC ...
	PUBLIC SubnetTypeCode = "PUBLIC"
)

// PossibleSubnetTypeCodeValues returns an array of possible values for the SubnetTypeCode const type.
func PossibleSubnetTypeCodeValues() []SubnetTypeCode {
	return []SubnetTypeCode{PRIVATE, PUBLIC}
}

// SupportedSubnetTypeCode enumerates the values for supported subnet type code.
type SupportedSubnetTypeCode string

const (
	// SupportedSubnetTypeCodePRIVATE ...
	SupportedSubnetTypeCodePRIVATE SupportedSubnetTypeCode = "PRIVATE"
	// SupportedSubnetTypeCodePUBLIC ...
	SupportedSubnetTypeCodePUBLIC SupportedSubnetTypeCode = "PUBLIC"
)

// PossibleSupportedSubnetTypeCodeValues returns an array of possible values for the SupportedSubnetTypeCode const type.
func PossibleSupportedSubnetTypeCodeValues() []SupportedSubnetTypeCode {
	return []SupportedSubnetTypeCode{SupportedSubnetTypeCodePRIVATE, SupportedSubnetTypeCodePUBLIC}
}

// TargetTypeCode enumerates the values for target type code.
type TargetTypeCode string

const (
	// NATGW ...
	NATGW TargetTypeCode = "NATGW"
	// VGW ...
	VGW TargetTypeCode = "VGW"
	// VPCPEERING ...
	VPCPEERING TargetTypeCode = "VPCPEERING"
)

// PossibleTargetTypeCodeValues returns an array of possible values for the TargetTypeCode const type.
func PossibleTargetTypeCodeValues() []TargetTypeCode {
	return []TargetTypeCode{NATGW, VGW, VPCPEERING}
}

// UsageTypeCode enumerates the values for usage type code.
type UsageTypeCode string

const (
	// BM ...
	BM UsageTypeCode = "BM"
	// GEN ...
	GEN UsageTypeCode = "GEN"
	// LOADB ...
	LOADB UsageTypeCode = "LOADB"
)

// PossibleUsageTypeCodeValues returns an array of possible values for the UsageTypeCode const type.
func PossibleUsageTypeCodeValues() []UsageTypeCode {
	return []UsageTypeCode{BM, GEN, LOADB}
}

// DetailResponse ...
type DetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// List ...
type List struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// Ipv4CidrBlock - IP 주소 CIDR
	Ipv4CidrBlock *string `json:"ipv4CidrBlock,omitempty"`
	// VpcStatus - VPC 상태
	VpcStatus *Status `json:"vpcStatus,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
}

// ListResponse ...
type ListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// NatGatewayInstanceDetailResponse ...
type NatGatewayInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceList ...
type NatGatewayInstanceList struct {
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// NatGatewayInstanceNo - NAT 게이트웨이 인스턴스 번호
	NatGatewayInstanceNo *string `json:"natGatewayInstanceNo,omitempty"`
	// NatGatewayName - NAT 게이트웨이 이름
	NatGatewayName *string `json:"natGatewayName,omitempty"`
	// PublicIP - 공인 IP 주소
	PublicIP *string `json:"publicIp,omitempty"`
	// NatGatewayInstanceStatus - NAT 게이트웨이 인스턴스 상태
	NatGatewayInstanceStatus *NatGatewayInstanceStatus `json:"natGatewayInstanceStatus,omitempty"`
	// NatGatewayInstanceStatusName - NAT 게이트웨이 인스턴스 상태 이름
	NatGatewayInstanceStatusName *string     `json:"natGatewayInstanceStatusName,omitempty"`
	NatGatewayInstanceOperation  interface{} `json:"natGatewayInstanceOperation,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// NatGatewayDescription - NAT 게이트웨이 설명
	NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`
}

// NatGatewayInstanceListResponse ...
type NatGatewayInstanceListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceResponse ...
type NatGatewayInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceStatus ...
type NatGatewayInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkACLDetailResponse ...
type NetworkACLDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLInboundRuleResponse ...
type NetworkACLInboundRuleResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLList ...
type NetworkACLList struct {
	// NetworkACLNo - 네트워크 ACL 번호
	NetworkACLNo *string `json:"networkAclNo,omitempty"`
	// NetworkACLName - 네트워크 ACL 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// NetworkACLStatus - 네트워크 ACL 상태
	NetworkACLStatus *NetworkACLStatus `json:"networkAclStatus,omitempty"`
	// NetworkACLDescription - 네트워크 ACL 설명
	NetworkACLDescription *string `json:"networkAclDescription,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// IsDefault - 네트워크 ACL 기본 여부
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NetworkACLListResponse ...
type NetworkACLListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLOutboundRuleResponse ...
type NetworkACLOutboundRuleResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLResponse ...
type NetworkACLResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLRuleList ...
type NetworkACLRuleList struct {
	// NetworkACLNo - 네트워크 ACL 번호
	NetworkACLNo *string `json:"networkAclNo,omitempty"`
	// Priority - 우선순위
	Priority *int32 `json:"priority,omitempty"`
	// ProtocolType - 프로토콜 타입
	ProtocolType *ProtocolType `json:"protocolType,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// RuleAction - 룰 액션
	RuleAction *RuleAction `json:"ruleAction,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// IPBlock - 블럭 여부
	IPBlock *string `json:"ipBlock,omitempty"`
	// NetworkACLRuleType - 네트워크 ACL 룰 타입
	NetworkACLRuleType *NetworkACLRuleType `json:"networkAclRuleType,omitempty"`
	// NetworkACLRuleDescription - 네트워크 ACL 룰 설명
	NetworkACLRuleDescription *string `json:"networkAclRuleDescription,omitempty"`
}

// NetworkACLRuleListResponse ...
type NetworkACLRuleListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLRuleType ...
type NetworkACLRuleType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkACLStatus ...
type NetworkACLStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PeeringInstanceDetailResponse ...
type PeeringInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceList ...
type PeeringInstanceList struct {
	// VpcPeeringInstanceNo - VPC 피어링 인스턴스 번호
	VpcPeeringInstanceNo *string `json:"vpcPeeringInstanceNo,omitempty"`
	// VpcPeeringName - VPC 피어링 이름
	VpcPeeringName *string `json:"vpcPeeringName,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// VpcPeeringInstanceStatus - VPC 피어링 인스턴스 상태
	VpcPeeringInstanceStatus *Status `json:"vpcPeeringInstanceStatus,omitempty"`
	// VpcPeeringInstanceStatusName - VPC 피어링 인스턴스 상태 이름
	VpcPeeringInstanceStatusName *string     `json:"vpcPeeringInstanceStatusName,omitempty"`
	VpcPeeringInstanceOperation  interface{} `json:"vpcPeeringInstanceOperation,omitempty"`
	// SourceVpcNo - 소스 VPC 번호
	SourceVpcNo *string `json:"sourceVpcNo,omitempty"`
	// SourceVpcName - 소스 VPC 이름
	SourceVpcName *string `json:"sourceVpcName,omitempty"`
	// SourceVpcIpv4CidrBlock - 소스 VPC IP 주소 CIDR
	SourceVpcIpv4CidrBlock *string `json:"sourceVpcIpv4CidrBlock,omitempty"`
	// SourceVpcLoginID - 소스 VPC 로그인 아이디
	SourceVpcLoginID *string `json:"sourceVpcLoginId,omitempty"`
	// TargetVpcNo - 타겟 VPC 번호
	TargetVpcNo *string `json:"targetVpcNo,omitempty"`
	// TargetVpcName - 타겟 VPC 이름
	TargetVpcName *string `json:"targetVpcName,omitempty"`
	// TargetVpcIpv4CidrBlock - 타겟 VPC IP 주소 CIDR
	TargetVpcIpv4CidrBlock *string `json:"targetVpcIpv4CidrBlock,omitempty"`
	// TargetVpcLoginID - 타겟 VPC 로그인 ID
	TargetVpcLoginID *string `json:"targetVpcLoginId,omitempty"`
	// VpcPeeringDescription - VPC 피어링 설명
	VpcPeeringDescription *string `json:"vpcPeeringDescription,omitempty"`
	// HasReverseVpcPeering - 리버스 VPC 피어링 여부
	HasReverseVpcPeering *bool `json:"hasReverseVpcPeering,omitempty"`
	// IsBetweenAccounts - 계정간 피어링 여부
	IsBetweenAccounts *bool `json:"isBetweenAccounts,omitempty"`
	// ReverseVpcPeeringInstanceNo - 리버스 VPC 피어링 인스턴스 번호
	ReverseVpcPeeringInstanceNo *string `json:"reverseVpcPeeringInstanceNo,omitempty"`
}

// PeeringInstanceListResponse ...
type PeeringInstanceListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceResponse ...
type PeeringInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceStatus ...
type PeeringInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PeeringResponse ...
type PeeringResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// ProtocolType ...
type ProtocolType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// Response ...
type Response struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// RouteList ...
type RouteList struct {
	// DestinationCidrBlock - 목적이 IP 주소 CIDR
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`
	// TargetName - 타겟 이름
	TargetName *string `json:"targetName,omitempty"`
	// RouteTableNo - 라우트 테이블 번호
	RouteTableNo *string `json:"routeTableNo,omitempty"`
	// TargetType - 타겟 타입
	TargetType *TargetType `json:"targetType,omitempty"`
	// TargetNo - 타겟 번호
	TargetNo *string `json:"targetNo,omitempty"`
	// IsDefault - 기본 라우트 여부
	IsDefault *bool `json:"isDefault,omitempty"`
}

// RouteListResponse ...
type RouteListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteList - 라우트 리스트
	RouteList *[]RouteList `json:"routeList,omitempty"`
}

// RouteResponse ...
type RouteResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteList - 라우트 리스트
	RouteList *[]RouteList `json:"routeList,omitempty"`
}

// RouteTableDetailResponse ...
type RouteTableDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableList ...
type RouteTableList struct {
	// RouteTableNo - 라우트 테이블 번호
	RouteTableNo *string `json:"routeTableNo,omitempty"`
	// RouteTableName - 라우트 테이블 이름
	RouteTableName *string `json:"routeTableName,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SupportedSubnetType - 지원 서브넷 타입
	SupportedSubnetType *SupportedSubnetType `json:"supportedSubnetType,omitempty"`
	// IsDefault - 기본 라우트 여부
	IsDefault *bool `json:"isDefault,omitempty"`
	// RouteTableStatus - 라우트 테이블 상태
	RouteTableStatus *RouteTableStatus `json:"routeTableStatus,omitempty"`
	// RouteTableDescription - 라우트 테이블 설명
	RouteTableDescription *string `json:"routeTableDescription,omitempty"`
}

// RouteTableListResponse ...
type RouteTableListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableResponse ...
type RouteTableResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableStatus ...
type RouteTableStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// RouteTableSubnetListResponse ...
type RouteTableSubnetListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// RouteTableSubnetResponse ...
type RouteTableSubnetResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// RuleAction ...
type RuleAction struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// Status ...
type Status struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SubnetDetailResponse ...
type SubnetDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetList ...
type SubnetList struct {
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// ZoneCode - ZONE 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// Subnet - 서브넷 CIDR
	Subnet *string `json:"subnet,omitempty"`
	// SubnetStatus - 서브넷 상태
	SubnetStatus *SubnetStatus `json:"subnetStatus,omitempty"`
	// SubnetType - 서브넷 타입
	SubnetType *SubnetType `json:"subnetType,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// UsageType - 사용 타입
	UsageType *UsageType `json:"usageType,omitempty"`
	// NetworkACLNo - 네트워크 ACL 번호
	NetworkACLNo *string `json:"networkAclNo,omitempty"`
}

// SubnetListResponse ...
type SubnetListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetNetworkACLResponse ...
type SubnetNetworkACLResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// SubnetResponse ...
type SubnetResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetStatus ...
type SubnetStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SubnetType ...
type SubnetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SupportedSubnetType ...
type SupportedSubnetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// TargetType ...
type TargetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// UsageType ...
type UsageType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}
