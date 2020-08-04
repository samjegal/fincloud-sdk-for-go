package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/server"

// AccessControlGroupRuleTypeCode enumerates the values for access control group rule type code.
type AccessControlGroupRuleTypeCode string

const (
	// INBND ...
	INBND AccessControlGroupRuleTypeCode = "INBND"
	// OTBND ...
	OTBND AccessControlGroupRuleTypeCode = "OTBND"
)

// PossibleAccessControlGroupRuleTypeCodeValues returns an array of possible values for the AccessControlGroupRuleTypeCode const type.
func PossibleAccessControlGroupRuleTypeCodeValues() []AccessControlGroupRuleTypeCode {
	return []AccessControlGroupRuleTypeCode{INBND, OTBND}
}

// AccessControlGroupStatusCode enumerates the values for access control group status code.
type AccessControlGroupStatusCode string

const (
	// INIT ...
	INIT AccessControlGroupStatusCode = "INIT"
	// RUN ...
	RUN AccessControlGroupStatusCode = "RUN"
	// SET ...
	SET AccessControlGroupStatusCode = "SET"
	// TERMTING ...
	TERMTING AccessControlGroupStatusCode = "TERMTING"
)

// PossibleAccessControlGroupStatusCodeValues returns an array of possible values for the AccessControlGroupStatusCode const type.
func PossibleAccessControlGroupStatusCodeValues() []AccessControlGroupStatusCode {
	return []AccessControlGroupStatusCode{INIT, RUN, SET, TERMTING}
}

// BlockStorageDiskDetailTypeCode enumerates the values for block storage disk detail type code.
type BlockStorageDiskDetailTypeCode string

const (
	// HDD ...
	HDD BlockStorageDiskDetailTypeCode = "HDD"
	// SDD ...
	SDD BlockStorageDiskDetailTypeCode = "SDD"
)

// PossibleBlockStorageDiskDetailTypeCodeValues returns an array of possible values for the BlockStorageDiskDetailTypeCode const type.
func PossibleBlockStorageDiskDetailTypeCodeValues() []BlockStorageDiskDetailTypeCode {
	return []BlockStorageDiskDetailTypeCode{HDD, SDD}
}

// BlockStorageDiskTypeCode enumerates the values for block storage disk type code.
type BlockStorageDiskTypeCode string

const (
	// NET ...
	NET BlockStorageDiskTypeCode = "NET"
)

// PossibleBlockStorageDiskTypeCodeValues returns an array of possible values for the BlockStorageDiskTypeCode const type.
func PossibleBlockStorageDiskTypeCodeValues() []BlockStorageDiskTypeCode {
	return []BlockStorageDiskTypeCode{NET}
}

// BlockStorageInstanceStatusCode enumerates the values for block storage instance status code.
type BlockStorageInstanceStatusCode string

const (
	// BlockStorageInstanceStatusCodeATTAC ...
	BlockStorageInstanceStatusCodeATTAC BlockStorageInstanceStatusCode = "ATTAC"
	// BlockStorageInstanceStatusCodeCREATE ...
	BlockStorageInstanceStatusCodeCREATE BlockStorageInstanceStatusCode = "CREATE"
	// BlockStorageInstanceStatusCodeINIT ...
	BlockStorageInstanceStatusCodeINIT BlockStorageInstanceStatusCode = "INIT"
)

// PossibleBlockStorageInstanceStatusCodeValues returns an array of possible values for the BlockStorageInstanceStatusCode const type.
func PossibleBlockStorageInstanceStatusCodeValues() []BlockStorageInstanceStatusCode {
	return []BlockStorageInstanceStatusCode{BlockStorageInstanceStatusCodeATTAC, BlockStorageInstanceStatusCodeCREATE, BlockStorageInstanceStatusCodeINIT}
}

// BlockStorageSnapshotInstanceStatusCode enumerates the values for block storage snapshot instance status
// code.
type BlockStorageSnapshotInstanceStatusCode string

const (
	// BlockStorageSnapshotInstanceStatusCodeATTAC ...
	BlockStorageSnapshotInstanceStatusCodeATTAC BlockStorageSnapshotInstanceStatusCode = "ATTAC"
	// BlockStorageSnapshotInstanceStatusCodeCREAT ...
	BlockStorageSnapshotInstanceStatusCodeCREAT BlockStorageSnapshotInstanceStatusCode = "CREAT"
	// BlockStorageSnapshotInstanceStatusCodeINIT ...
	BlockStorageSnapshotInstanceStatusCodeINIT BlockStorageSnapshotInstanceStatusCode = "INIT"
)

// PossibleBlockStorageSnapshotInstanceStatusCodeValues returns an array of possible values for the BlockStorageSnapshotInstanceStatusCode const type.
func PossibleBlockStorageSnapshotInstanceStatusCodeValues() []BlockStorageSnapshotInstanceStatusCode {
	return []BlockStorageSnapshotInstanceStatusCode{BlockStorageSnapshotInstanceStatusCodeATTAC, BlockStorageSnapshotInstanceStatusCodeCREAT, BlockStorageSnapshotInstanceStatusCodeINIT}
}

// BlockStorageTypeCode enumerates the values for block storage type code.
type BlockStorageTypeCode string

const (
	// BASIC ...
	BASIC BlockStorageTypeCode = "BASIC"
	// NSI ...
	NSI BlockStorageTypeCode = "NSI"
	// SNAP ...
	SNAP BlockStorageTypeCode = "SNAP"
	// SVRBS ...
	SVRBS BlockStorageTypeCode = "SVRBS"
)

// PossibleBlockStorageTypeCodeValues returns an array of possible values for the BlockStorageTypeCode const type.
func PossibleBlockStorageTypeCodeValues() []BlockStorageTypeCode {
	return []BlockStorageTypeCode{BASIC, NSI, SNAP, SVRBS}
}

// FeeSystemTypeCode enumerates the values for fee system type code.
type FeeSystemTypeCode string

const (
	// FXSUM ...
	FXSUM FeeSystemTypeCode = "FXSUM"
	// MTRAT ...
	MTRAT FeeSystemTypeCode = "MTRAT"
)

// PossibleFeeSystemTypeCodeValues returns an array of possible values for the FeeSystemTypeCode const type.
func PossibleFeeSystemTypeCodeValues() []FeeSystemTypeCode {
	return []FeeSystemTypeCode{FXSUM, MTRAT}
}

// MemberServerImageInstanceStatusCode enumerates the values for member server image instance status code.
type MemberServerImageInstanceStatusCode string

const (
	// MemberServerImageInstanceStatusCodeCREATE ...
	MemberServerImageInstanceStatusCodeCREATE MemberServerImageInstanceStatusCode = "CREATE"
	// MemberServerImageInstanceStatusCodeINIT ...
	MemberServerImageInstanceStatusCodeINIT MemberServerImageInstanceStatusCode = "INIT"
)

// PossibleMemberServerImageInstanceStatusCodeValues returns an array of possible values for the MemberServerImageInstanceStatusCode const type.
func PossibleMemberServerImageInstanceStatusCodeValues() []MemberServerImageInstanceStatusCode {
	return []MemberServerImageInstanceStatusCode{MemberServerImageInstanceStatusCodeCREATE, MemberServerImageInstanceStatusCodeINIT}
}

// NetworkInterfaceStatusCode enumerates the values for network interface status code.
type NetworkInterfaceStatusCode string

const (
	// NetworkInterfaceStatusCodeNOTUSED ...
	NetworkInterfaceStatusCodeNOTUSED NetworkInterfaceStatusCode = "NOTUSED"
	// NetworkInterfaceStatusCodeSET ...
	NetworkInterfaceStatusCodeSET NetworkInterfaceStatusCode = "SET"
	// NetworkInterfaceStatusCodeUNSET ...
	NetworkInterfaceStatusCodeUNSET NetworkInterfaceStatusCode = "UNSET"
	// NetworkInterfaceStatusCodeUSED ...
	NetworkInterfaceStatusCodeUSED NetworkInterfaceStatusCode = "USED"
)

// PossibleNetworkInterfaceStatusCodeValues returns an array of possible values for the NetworkInterfaceStatusCode const type.
func PossibleNetworkInterfaceStatusCodeValues() []NetworkInterfaceStatusCode {
	return []NetworkInterfaceStatusCode{NetworkInterfaceStatusCodeNOTUSED, NetworkInterfaceStatusCodeSET, NetworkInterfaceStatusCodeUNSET, NetworkInterfaceStatusCodeUSED}
}

// OsTypeCode enumerates the values for os type code.
type OsTypeCode string

const (
	// LNX ...
	LNX OsTypeCode = "LNX"
	// WND ...
	WND OsTypeCode = "WND"
)

// PossibleOsTypeCodeValues returns an array of possible values for the OsTypeCode const type.
func PossibleOsTypeCodeValues() []OsTypeCode {
	return []OsTypeCode{LNX, WND}
}

// PlatformTypeCode enumerates the values for platform type code.
type PlatformTypeCode string

const (
	// LNX32 ...
	LNX32 PlatformTypeCode = "LNX32"
	// LNX64 ...
	LNX64 PlatformTypeCode = "LNX64"
	// UBD64 ...
	UBD64 PlatformTypeCode = "UBD64"
	// UBS64 ...
	UBS64 PlatformTypeCode = "UBS64"
	// WND32 ...
	WND32 PlatformTypeCode = "WND32"
	// WND64 ...
	WND64 PlatformTypeCode = "WND64"
)

// PossiblePlatformTypeCodeValues returns an array of possible values for the PlatformTypeCode const type.
func PossiblePlatformTypeCodeValues() []PlatformTypeCode {
	return []PlatformTypeCode{LNX32, LNX64, UBD64, UBS64, WND32, WND64}
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

// BaseBlockStorageDiskDetailType ...
type BaseBlockStorageDiskDetailType struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

// BaseBlockStorageDiskType ...
type BaseBlockStorageDiskType struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceList ...
type InstanceList struct {
	ServerInstanceNo               *string                         `json:"serverInstanceNo,omitempty"`
	ServerName                     *string                         `json:"serverName,omitempty"`
	ServerDescription              *string                         `json:"serverDescription,omitempty"`
	CPUCount                       *int32                          `json:"cpuCount,omitempty"`
	MemorySize                     *float64                        `json:"memorySize,omitempty"`
	PlatformType                   *PlatformType                   `json:"platformType,omitempty"`
	LoginKeyName                   *string                         `json:"loginKeyName,omitempty"`
	PublicIPInstanceNo             *string                         `json:"publicIpInstanceNo,omitempty"`
	PublicIP                       *string                         `json:"publicIp,omitempty"`
	ServerInstanceStatus           *InstanceStatus                 `json:"serverInstanceStatus,omitempty"`
	ServerInstanceOperation        *InstanceOperation              `json:"serverInstanceOperation,omitempty"`
	ServerInstanceStatusName       *string                         `json:"serverInstanceStatusName,omitempty"`
	CreateDate                     *string                         `json:"createDate,omitempty"`
	Uptime                         *string                         `json:"uptime,omitempty"`
	ServerImageProductCode         *string                         `json:"serverImageProductCode,omitempty"`
	ServerProductCode              *string                         `json:"serverProductCode,omitempty"`
	IsProtectServerTermination     *bool                           `json:"isProtectServerTermination,omitempty"`
	ZoneCode                       *string                         `json:"zoneCode,omitempty"`
	RegionCode                     *string                         `json:"regionCode,omitempty"`
	VpcNo                          *string                         `json:"vpcNo,omitempty"`
	SubnetNo                       *string                         `json:"subnetNo,omitempty"`
	NetworkInterfaceNoList         *NetworkInterfaceNoList         `json:"networkInterfaceNoList,omitempty"`
	InitScriptNo                   *string                         `json:"initScriptNo,omitempty"`
	ServerInstanceType             *InstanceType                   `json:"serverInstanceType,omitempty"`
	BaseBlockStorageDiskType       *BaseBlockStorageDiskType       `json:"baseBlockStorageDiskType,omitempty"`
	BaseBlockStorageDiskDetailType *BaseBlockStorageDiskDetailType `json:"baseBlockStorageDiskDetailType,omitempty"`
	PlacementGroupNo               *string                         `json:"placementGroupNo,omitempty"`
	PlacementGroupName             *string                         `json:"placementGroupName,omitempty"`
}

// InstanceListResponse ...
type InstanceListResponse struct {
	autorest.Response  `json:"-"`
	ReturnCode         *string         `json:"returnCode,omitempty"`
	ReturnMessage      *string         `json:"returnMessage,omitempty"`
	TotalRows          *int32          `json:"totalRows,omitempty"`
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// InstanceOperation ...
type InstanceOperation struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceStatus ...
type InstanceStatus struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceType ...
type InstanceType struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkInterfaceNoList ...
type NetworkInterfaceNoList struct {
	NetworkInterfaceNo *string `json:"networkInterfaceNo,omitempty"`
}

// PlatformType ...
type PlatformType struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}
