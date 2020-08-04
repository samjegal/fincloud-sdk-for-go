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
        return []NatGatewayInstanceStatusCode{INIT,RUN,TERMTING}
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
        return []NetworkACLRuleTypeCode{INBND,OTBND}
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
        return []NetworkACLStatusCode{NetworkACLStatusCodeINIT,NetworkACLStatusCodeRUN,NetworkACLStatusCodeSET,NetworkACLStatusCodeTERMTING}
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
        return []PeeringInstanceStatusCode{PeeringInstanceStatusCodeINIT,PeeringInstanceStatusCodeRUN,PeeringInstanceStatusCodeTERMTING}
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
        return []ProtocolTypeCode{ICMP,TCP,UDP}
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
        return []RuleActionCode{ALLOW,DROP}
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
        return []SortedBy{SourceVpcName,TargetVpcName,VpcPeeringName}
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
        return []SortingOrder{ASC,DESC}
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
        return []StatusCode{StatusCodeCREATING,StatusCodeINIT,StatusCodeRUN,StatusCodeTERMTING}
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
        return []SubnetStatusCode{SubnetStatusCodeCREATING,SubnetStatusCodeINIT,SubnetStatusCodeRUN,SubnetStatusCodeTERMTING}
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
        return []SubnetTypeCode{PRIVATE,PUBLIC}
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
        return []SupportedSubnetTypeCode{SupportedSubnetTypeCodePRIVATE,SupportedSubnetTypeCodePUBLIC}
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
        return []TargetTypeCode{NATGW,VGW,VPCPEERING}
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
        return []UsageTypeCode{BM,GEN,LOADB}
    }

            // DetailResponse ...
            type DetailResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcList *[]List `json:"vpcList,omitempty"`
            }

            // List ...
            type List struct {
            VpcNo *string `json:"vpcNo,omitempty"`
            VpcName *string `json:"vpcName,omitempty"`
            Ipv4CidrBlock *string `json:"ipv4CidrBlock,omitempty"`
            VpcStatus *Status `json:"vpcStatus,omitempty"`
            RegionCode *string `json:"regionCode,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            }

            // ListResponse ...
            type ListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcList *[]List `json:"vpcList,omitempty"`
            }

            // NatGatewayInstanceDetailResponse ...
            type NatGatewayInstanceDetailResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
            }

            // NatGatewayInstanceList ...
            type NatGatewayInstanceList struct {
            VpcName *string `json:"vpcName,omitempty"`
            NatGatewayInstanceNo *string `json:"natGatewayInstanceNo,omitempty"`
            NatGatewayName *string `json:"natGatewayName,omitempty"`
            PublicIP *string `json:"publicIp,omitempty"`
            NatGatewayInstanceStatus *NatGatewayInstanceStatus `json:"natGatewayInstanceStatus,omitempty"`
            NatGatewayInstanceStatusName *string `json:"natGatewayInstanceStatusName,omitempty"`
            NatGatewayInstanceOperation interface{} `json:"natGatewayInstanceOperation,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`
            }

            // NatGatewayInstanceListResponse ...
            type NatGatewayInstanceListResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
            }

            // NatGatewayInstanceResponse ...
            type NatGatewayInstanceResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
            }

            // NatGatewayInstanceStatus ...
            type NatGatewayInstanceStatus struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // NetworkACLDetailResponse ...
            type NetworkACLDetailResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
            }

            // NetworkACLInboundRuleResponse ...
            type NetworkACLInboundRuleResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
            }

            // NetworkACLList ...
            type NetworkACLList struct {
            NetworkACLNo *string `json:"networkAclNo,omitempty"`
            NetworkACLName *string `json:"networkAclName,omitempty"`
            VpcNo *string `json:"vpcNo,omitempty"`
            NetworkACLStatus *NetworkACLStatus `json:"networkAclStatus,omitempty"`
            NetworkACLDescription *string `json:"networkAclDescription,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            IsDefault *bool `json:"isDefault,omitempty"`
            }

            // NetworkACLListResponse ...
            type NetworkACLListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
            }

            // NetworkACLOutboundRuleResponse ...
            type NetworkACLOutboundRuleResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
            }

            // NetworkACLResponse ...
            type NetworkACLResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
            }

            // NetworkACLRuleList ...
            type NetworkACLRuleList struct {
            NetworkACLNo *string `json:"networkAclNo,omitempty"`
            Priority *int32 `json:"priority,omitempty"`
            ProtocolType *ProtocolType `json:"protocolType,omitempty"`
            PortRange *string `json:"portRange,omitempty"`
            RuleAction *RuleAction `json:"ruleAction,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            IPBlock *string `json:"ipBlock,omitempty"`
            NetworkACLRuleType *NetworkACLRuleType `json:"networkAclRuleType,omitempty"`
            NetworkACLRuleDescription *string `json:"networkAclRuleDescription,omitempty"`
            }

            // NetworkACLRuleListResponse ...
            type NetworkACLRuleListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
            }

            // NetworkACLRuleType ...
            type NetworkACLRuleType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // NetworkACLStatus ...
            type NetworkACLStatus struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // PeeringInstanceDetailResponse ...
            type PeeringInstanceDetailResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
            }

            // PeeringInstanceList ...
            type PeeringInstanceList struct {
            VpcPeeringInstanceNo *string `json:"vpcPeeringInstanceNo,omitempty"`
            VpcPeeringName *string `json:"vpcPeeringName,omitempty"`
            RegionCode *string `json:"regionCode,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            VpcPeeringInstanceStatus *Status `json:"vpcPeeringInstanceStatus,omitempty"`
            VpcPeeringInstanceStatusName *string `json:"vpcPeeringInstanceStatusName,omitempty"`
            VpcPeeringInstanceOperation interface{} `json:"vpcPeeringInstanceOperation,omitempty"`
            SourceVpcNo *string `json:"sourceVpcNo,omitempty"`
            SourceVpcName *string `json:"sourceVpcName,omitempty"`
            SourceVpcIpv4CidrBlock *string `json:"sourceVpcIpv4CidrBlock,omitempty"`
            SourceVpcLoginID *string `json:"sourceVpcLoginId,omitempty"`
            TargetVpcNo *string `json:"targetVpcNo,omitempty"`
            TargetVpcName *string `json:"targetVpcName,omitempty"`
            TargetVpcIpv4CidrBlock *string `json:"targetVpcIpv4CidrBlock,omitempty"`
            TargetVpcLoginID *string `json:"targetVpcLoginId,omitempty"`
            VpcPeeringDescription *string `json:"vpcPeeringDescription,omitempty"`
            HasReverseVpcPeering *bool `json:"hasReverseVpcPeering,omitempty"`
            IsBetweenAccounts *bool `json:"isBetweenAccounts,omitempty"`
            ReverseVpcPeeringInstanceNo *string `json:"reverseVpcPeeringInstanceNo,omitempty"`
            }

            // PeeringInstanceListResponse ...
            type PeeringInstanceListResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
            }

            // PeeringInstanceResponse ...
            type PeeringInstanceResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
            }

            // PeeringInstanceStatus ...
            type PeeringInstanceStatus struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // PeeringResponse ...
            type PeeringResponse struct {
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
            }

            // ProtocolType ...
            type ProtocolType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // Response ...
            type Response struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            VpcList *[]List `json:"vpcList,omitempty"`
            }

            // RouteList ...
            type RouteList struct {
            DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`
            TargetName *string `json:"targetName,omitempty"`
            RouteTableNo *string `json:"routeTableNo,omitempty"`
            TargetType *TargetType `json:"targetType,omitempty"`
            TargetNo *string `json:"targetNo,omitempty"`
            IsDefault *bool `json:"isDefault,omitempty"`
            }

            // RouteListResponse ...
            type RouteListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            RouteList *[]RouteList `json:"routeList,omitempty"`
            }

            // RouteResponse ...
            type RouteResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            RouteList *[]RouteList `json:"routeList,omitempty"`
            }

            // RouteTableDetailResponse ...
            type RouteTableDetailResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
            }

            // RouteTableList ...
            type RouteTableList struct {
            RouteTableNo *string `json:"routeTableNo,omitempty"`
            RouteTableName *string `json:"routeTableName,omitempty"`
            RegionCode *string `json:"regionCode,omitempty"`
            VpcNo *string `json:"vpcNo,omitempty"`
            SupportedSubnetType *SupportedSubnetType `json:"supportedSubnetType,omitempty"`
            IsDefault *bool `json:"isDefault,omitempty"`
            RouteTableStatus *RouteTableStatus `json:"routeTableStatus,omitempty"`
            RouteTableDescription *string `json:"routeTableDescription,omitempty"`
            }

            // RouteTableListResponse ...
            type RouteTableListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
            }

            // RouteTableResponse ...
            type RouteTableResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
            }

            // RouteTableStatus ...
            type RouteTableStatus struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // RouteTableSubnetListResponse ...
            type RouteTableSubnetListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            SubnetList *[]SubnetList `json:"subnetList,omitempty"`
            }

            // RouteTableSubnetResponse ...
            type RouteTableSubnetResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            SubnetList *[]SubnetList `json:"subnetList,omitempty"`
            }

            // RuleAction ...
            type RuleAction struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // Status ...
            type Status struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // SubnetDetailResponse ...
            type SubnetDetailResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            SubnetList *[]SubnetList `json:"subnetList,omitempty"`
            }

            // SubnetList ...
            type SubnetList struct {
            SubnetNo *string `json:"subnetNo,omitempty"`
            VpcNo *string `json:"vpcNo,omitempty"`
            ZoneCode *string `json:"zoneCode,omitempty"`
            SubnetName *string `json:"subnetName,omitempty"`
            Subnet *string `json:"subnet,omitempty"`
            SubnetStatus *SubnetStatus `json:"subnetStatus,omitempty"`
            SubnetType *SubnetType `json:"subnetType,omitempty"`
            CreateDate *string `json:"createDate,omitempty"`
            UsageType *UsageType `json:"usageType,omitempty"`
            NetworkACLNo *string `json:"networkAclNo,omitempty"`
            }

            // SubnetListResponse ...
            type SubnetListResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            SubnetList *[]SubnetList `json:"subnetList,omitempty"`
            }

            // SubnetNetworkACLResponse ...
            type SubnetNetworkACLResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
            }

            // SubnetResponse ...
            type SubnetResponse struct {
            autorest.Response `json:"-"`
            ReturnCode *string `json:"returnCode,omitempty"`
            ReturnMessage *string `json:"returnMessage,omitempty"`
            TotalRows *int32 `json:"totalRows,omitempty"`
            SubnetList *[]SubnetList `json:"subnetList,omitempty"`
            }

            // SubnetStatus ...
            type SubnetStatus struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // SubnetType ...
            type SubnetType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // SupportedSubnetType ...
            type SupportedSubnetType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // TargetType ...
            type TargetType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

            // UsageType ...
            type UsageType struct {
            Code *string `json:"code,omitempty"`
            CodeName *string `json:"codeName,omitempty"`
            }

