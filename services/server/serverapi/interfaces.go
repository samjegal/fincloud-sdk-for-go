package serverapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/server"
)

// ACGClientAPI contains the set of methods on the ACGClient type.
type ACGClientAPI interface {
	Create(ctx context.Context, vpcNo string, accessControlGroupName string, accessControlGroupDescription string) (result server.AccessControlGroupResponse, err error)
	Delete(ctx context.Context, vpcNo string, accessControlGroupNo string) (result server.AccessControlGroupResponse, err error)
	GetDetail(ctx context.Context, accessControlGroupNo string) (result server.AccessControlGroupDetailResponse, err error)
	GetList(ctx context.Context, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode server.AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (result server.AccessControlGroupListResponse, err error)
	GetRuleList(ctx context.Context, accessControlGroupNo string, accessControlGroupRuleTypeCode server.AccessControlGroupRuleTypeCode) (result server.AccessControlGroupRuleListResponse, err error)
}

var _ ACGClientAPI = (*server.ACGClient)(nil)

// ACGInboundClientAPI contains the set of methods on the ACGInboundClient type.
type ACGInboundClientAPI interface {
	AddRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result server.AccessControlGroupInboundRuleResponse, err error)
	RemoveRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result server.AccessControlGroupInboundRuleResponse, err error)
}

var _ ACGInboundClientAPI = (*server.ACGInboundClient)(nil)

// ACGOutboundClientAPI contains the set of methods on the ACGOutboundClient type.
type ACGOutboundClientAPI interface {
	AddRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result server.AccessControlGroupOutboundRuleResponse, err error)
	RemoveRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result server.AccessControlGroupOutboundRuleResponse, err error)
}

var _ ACGOutboundClientAPI = (*server.ACGOutboundClient)(nil)

// RegionClientAPI contains the set of methods on the RegionClient type.
type RegionClientAPI interface {
	GetList(ctx context.Context) (result server.RegionListResponse, err error)
}

var _ RegionClientAPI = (*server.RegionClient)(nil)

// ZoneClientAPI contains the set of methods on the ZoneClient type.
type ZoneClientAPI interface {
	GetList(ctx context.Context) (result autorest.Response, err error)
}

var _ ZoneClientAPI = (*server.ZoneClient)(nil)

// ImageProductClientAPI contains the set of methods on the ImageProductClient type.
type ImageProductClientAPI interface {
	GetList(ctx context.Context, blockStorageSize string, exclusionProductCode string, productCode string, platformTypeCodeListN server.PlatformTypeCode) (result server.ImageProductListResponse, err error)
}

var _ ImageProductClientAPI = (*server.ImageProductClient)(nil)

// ProductClientAPI contains the set of methods on the ProductClient type.
type ProductClientAPI interface {
	GetList(ctx context.Context, serverImageProductCode string, zoneCode string, exclusionProductCode string, productCode string, generationCode string) (result server.ProductListResponse, err error)
}

var _ ProductClientAPI = (*server.ProductClient)(nil)

// InitScriptClientAPI contains the set of methods on the InitScriptClient type.
type InitScriptClientAPI interface {
	Create(ctx context.Context, initScriptContent string, osTypeCode server.OsTypeCode, initScriptName string, initScriptDescription string) (result server.InitScriptResponse, err error)
	Delete(ctx context.Context, initScriptNoListN string) (result server.InitScriptResponse, err error)
	GetDetail(ctx context.Context, initScriptNo string) (result server.InitScriptDetailResponse, err error)
	GetList(ctx context.Context, osTypeCode server.OsTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder server.SortingOrder, initScriptName string, initScriptNoListN string) (result server.InitScriptListResponse, err error)
}

var _ InitScriptClientAPI = (*server.InitScriptClient)(nil)

// LoginKeyClientAPI contains the set of methods on the LoginKeyClient type.
type LoginKeyClientAPI interface {
	Create(ctx context.Context, keyName string) (result server.CreateLoginKeyResponse, err error)
	Delete(ctx context.Context, keyNameListN string) (result server.DeleteLoginKeysResponse, err error)
	GetList(ctx context.Context, pageNo string, pageSize string) (result server.LoginKeyListResponse, err error)
	Import(ctx context.Context, publicKey string, keyName string) (result server.LoginKeyResponse, err error)
}

var _ LoginKeyClientAPI = (*server.LoginKeyClient)(nil)

// NetworkInterfaceClientAPI contains the set of methods on the NetworkInterfaceClient type.
type NetworkInterfaceClientAPI interface {
	AddACG(ctx context.Context, networkInterfaceNo string, accessControlGroupNoListN string) (result server.NetworkInterfaceAccessControlGroupResponse, err error)
	Attach(ctx context.Context, subnetNo string, networkInterfaceNo string, serverInstanceNo string) (result server.NetworkInterfaceResponse, err error)
	Create(ctx context.Context, vpcNo string, subnetNo string, networkInterfaceName string, accessControlGroupNoListN string, serverInstanceNo string, IP string, networkInterfaceDescription string) (result server.NetworkInterfaceResponse, err error)
	Delete(ctx context.Context, networkInterfaceNo string) (result server.NetworkInterfaceResponse, err error)
	Detach(ctx context.Context, subnetNo string, networkInterfaceNo string, serverInstanceNo string) (result server.NetworkInterfaceResponse, err error)
	GetDetail(ctx context.Context, networkInterfaceNo string) (result server.NetworkInterfaceDetailResponse, err error)
	GetList(ctx context.Context, networkInterfaceNoListN string, IP string, networkInterfaceName string, serverInstanceName string, subnetName string, pageNo string, pageSize string, networkInterfaceStatusCode server.NetworkInterfaceStatusCode) (result server.NetworkInterfaceListResponse, err error)
	RemoveACG(ctx context.Context, networkInterfaceNo string, accessControlGroupNoListN string) (result server.NetworkInterfaceAccessControlGroupResponse, err error)
}

var _ NetworkInterfaceClientAPI = (*server.NetworkInterfaceClient)(nil)

// PlacementGroupClientAPI contains the set of methods on the PlacementGroupClient type.
type PlacementGroupClientAPI interface {
	Add(ctx context.Context, placementGroupNo string, serverInstanceNo string) (result server.PlacementGroupServerInstanceResponse, err error)
	Create(ctx context.Context, placementGroupName string, placementGroupTypeCode string) (result server.PlacementGroupResponse, err error)
	Delete(ctx context.Context, placementGroupNo string) (result server.PlacementGroupResponse, err error)
	GetDetail(ctx context.Context, placementGroupNo string) (result server.PlacementGroupDetailResponse, err error)
	GetList(ctx context.Context, placementGroupName string, placementGroupNoListN string) (result server.PlacementGroupListResponse, err error)
	Remove(ctx context.Context, placementGroupNo string, serverInstanceNo string) (result server.PlacementGroupServerInstanceResponse, err error)
}

var _ PlacementGroupClientAPI = (*server.PlacementGroupClient)(nil)

// PublicIPClientAPI contains the set of methods on the PublicIPClient type.
type PublicIPClientAPI interface {
	Associate(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result server.PublicIPWithServerInstanceResponse, err error)
	Create(ctx context.Context, serverInstanceNo string, publicIPDescription string) (result server.PublicIPInstanceResponse, err error)
	Delete(ctx context.Context, publicIPInstanceNo string) (result server.PublicIPInstanceResponse, err error)
	Disassociate(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result server.PublicIPFromServerInstanceResponse, err error)
	GetDetail(ctx context.Context, publicIPInstanceNo string) (result server.PublicIPInstanceDetailResponse, err error)
	GetList(ctx context.Context, publicIPInstanceNoListN string, publicIP string, privateIP string, isAssociated *bool) (result server.PublicIPInstanceListResponse, err error)
	GetTargetList(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result server.PublicIPTargetServerInstanceListResponse, err error)
}

var _ PublicIPClientAPI = (*server.PublicIPClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	ChangeSpec(ctx context.Context, serverInstanceNo string, serverProductCode string) (result server.InstanceSpecResponse, err error)
	Create(ctx context.Context, vpcNo string, subnetNo string, networkInterfaceListNnetworkInterfaceOrder string, networkInterfaceListNaccessControlGroupNoListN string, memberServerImageInstanceNo string, serverImageProductCode string, serverProductCode string, isEncryptedBaseBlockStorageVolume *bool, feeSystemTypeCode server.FeeSystemTypeCode, serverCreateCount string, serverCreateStartNo string, serverName string, networkInterfaceListNnetworkInterfaceNo string, networkInterfaceListNsubnetNo string, networkInterfaceListNip string, placementGroupNo string, isProtectServerTermination *bool, serverDescription string, initScriptNo string, loginKeyName string) (result server.InstancesResponse, err error)
	GetDetail(ctx context.Context, serverInstanceNo string) (result server.InstanceDetailResponse, err error)
	GetInstanceList(ctx context.Context, serverInstanceNoListN string, vpcNo string, pageNo string, serverInstanceStatusCode string, baseBlockStorageDiskTypeCode string, baseBlockStorageDiskDetailTypeCode string, serverName string, IP string, sortedBy string, sortingOrder string, placementGroupNoListN string) (result server.InstanceListResponse, err error)
	Reboot(ctx context.Context, serverInstanceNoListN string) (result server.InstancesResponse, err error)
	Start(ctx context.Context, serverInstanceNoListN string) (result server.InstancesResponse, err error)
	Stop(ctx context.Context, serverInstanceNoListN string) (result server.InstancesResponse, err error)
	Terminate(ctx context.Context, serverInstanceNoListN string) (result server.InstancesResponse, err error)
}

var _ ClientAPI = (*server.Client)(nil)

// RootPasswordClientAPI contains the set of methods on the RootPasswordClient type.
type RootPasswordClientAPI interface {
	Get(ctx context.Context, serverInstanceNo string, privateKey string) (result server.RootPassword, err error)
	GetList(ctx context.Context, rootPasswordServerInstanceListNserverInstanceNo string, rootPasswordServerInstanceListNprivateKey string) (result server.RootPasswordServerInstanceListResponse, err error)
}

var _ RootPasswordClientAPI = (*server.RootPasswordClient)(nil)

// ImageClientAPI contains the set of methods on the ImageClient type.
type ImageClientAPI interface {
	Create(ctx context.Context, serverInstanceNo string, memberServerImageName string, memberServerImageDescription string) (result server.MemberServerImageInstanceResponse, err error)
	Delete(ctx context.Context, memberServerImageInstanceNoListN string) (result server.MemberServerImageInstanceResponse, err error)
	GetDetail(ctx context.Context, memberServerImageInstanceNo string) (result server.MemberServerImageInstanceListResponse, err error)
	GetList(ctx context.Context, memberServerImageName string, memberServerImageInstanceStatusCode server.MemberServerImageInstanceStatusCode, memberServerImageInstanceNoListN string, platformTypeCodeListN server.PlatformTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder server.SortingOrder) (result server.MemberServerImageInstanceListResponse, err error)
}

var _ ImageClientAPI = (*server.ImageClient)(nil)

// SnapshotClientAPI contains the set of methods on the SnapshotClient type.
type SnapshotClientAPI interface {
	Create(ctx context.Context, originalBlockStorageInstanceNo string, blockStorageSnapshotName string, blockStorageSnapshotDescription string) (result server.BlockStorageSnapshotInstanceResponse, err error)
	Delete(ctx context.Context, blockStorageSnapshotInstanceNoListN string) (result server.BlockStorageSnapshotInstanceResponse, err error)
	GetDetail(ctx context.Context, blockStorageSnapshotInstanceNo string) (result server.BlockStorageSnapshotInstanceDetailResponse, err error)
	GetList(ctx context.Context, blockStorageSnapshotInstanceNoListN string, originalBlockStorageInstanceNoListN string, blockStorageSnapshotInstanceStatusCode server.BlockStorageSnapshotInstanceStatusCode, pageNo string, pageSize string, blockStorageSnapshotVolumeSize string, isEncryptedOriginalBlockStorageVolume *bool, blockStorageSnapshotName string, sortedBy string, sortingOrder server.SortingOrder) (result server.BlockStorageSnapshotInstanceListResponse, err error)
}

var _ SnapshotClientAPI = (*server.SnapshotClient)(nil)

// BlockStorageClientAPI contains the set of methods on the BlockStorageClient type.
type BlockStorageClientAPI interface {
	Attach(ctx context.Context, blockStorageInstanceNo string, serverInstanceNo string) (result server.BlockStorageInstanceResponse, err error)
	ChangeVolumeSize(ctx context.Context, blockStorageInstanceNo string, blockStorageSize string) (result server.BlockStorageVolumeSizeResponse, err error)
	Create(ctx context.Context, serverInstanceNo string, zoneCode string, blockStorageName string, blockStorageDiskDetailTypeCode server.BlockStorageDiskDetailTypeCode, blockStorageSnapshotInstanceNo string, blockStorageSize string, blockStorageDescription string) (result server.BlockStorageInstanceResponse, err error)
	Delete(ctx context.Context, blockStorageInstanceNoListN string) (result server.BlockStorageInstanceResponse, err error)
	Detach(ctx context.Context, blockStorageInstanceNoListN string) (result server.BlockStorageInstanceResponse, err error)
	GetDetail(ctx context.Context, blockStorageInstanceNo string) (result server.BlockStorageInstanceDetailResponse, err error)
	GetList(ctx context.Context, serverInstanceNo string, blockStorageTypeCodeListN server.BlockStorageTypeCode, blockStorageInstanceStatusCode server.BlockStorageInstanceStatusCode, pageNo string, pageSize string, blockStorageSize string, blockStorageInstanceNoListN string, blockStorageName string, serverName string, connectionInfo string, blockStorageDiskTypeCode server.BlockStorageDiskTypeCode, blockStorageDiskDetailTypeCode server.BlockStorageDiskDetailTypeCode, zoneCode string) (result server.BlockStorageInstanceListResponse, err error)
}

var _ BlockStorageClientAPI = (*server.BlockStorageClient)(nil)
