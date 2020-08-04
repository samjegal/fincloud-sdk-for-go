package serverapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/server"
)

// ACGClientAPI contains the set of methods on the ACGClient type.
type ACGClientAPI interface {
	Create(ctx context.Context, responseFormatType string, vpcNo string, regionCode string, accessControlGroupName string, accessControlGroupDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode server.AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (result autorest.Response, err error)
	GetRuleList(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string, accessControlGroupRuleTypeCode server.AccessControlGroupRuleTypeCode) (result autorest.Response, err error)
}

var _ ACGClientAPI = (*server.ACGClient)(nil)

// ACGInboundClientAPI contains the set of methods on the ACGInboundClient type.
type ACGInboundClientAPI interface {
	AddRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result autorest.Response, err error)
	RemoveRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result autorest.Response, err error)
}

var _ ACGInboundClientAPI = (*server.ACGInboundClient)(nil)

// ACGOutboundClientAPI contains the set of methods on the ACGOutboundClient type.
type ACGOutboundClientAPI interface {
	AddRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result autorest.Response, err error)
	RemoveRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode server.ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result autorest.Response, err error)
}

var _ ACGOutboundClientAPI = (*server.ACGOutboundClient)(nil)

// RegionClientAPI contains the set of methods on the RegionClient type.
type RegionClientAPI interface {
	GetList(ctx context.Context, responseFormatType string) (result autorest.Response, err error)
}

var _ RegionClientAPI = (*server.RegionClient)(nil)

// ZoneClientAPI contains the set of methods on the ZoneClient type.
type ZoneClientAPI interface {
	GetList(ctx context.Context, responseFormatType string, regionCode string) (result autorest.Response, err error)
}

var _ ZoneClientAPI = (*server.ZoneClient)(nil)

// ImageProductClientAPI contains the set of methods on the ImageProductClient type.
type ImageProductClientAPI interface {
	GetList(ctx context.Context, responseFormatType string, regionCode string, blockStorageSize string, exclusionProductCode string, productCode string, platformTypeCodeListN server.PlatformTypeCode) (result autorest.Response, err error)
}

var _ ImageProductClientAPI = (*server.ImageProductClient)(nil)

// ProductClientAPI contains the set of methods on the ProductClient type.
type ProductClientAPI interface {
	GetList(ctx context.Context, responseFormatType string, serverImageProductCode string, regionCode string, zoneCode string, exclusionProductCode string, productCode string, generationCode string) (result autorest.Response, err error)
}

var _ ProductClientAPI = (*server.ProductClient)(nil)

// InitScriptClientAPI contains the set of methods on the InitScriptClient type.
type InitScriptClientAPI interface {
	Create(ctx context.Context, responseFormatType string, initScriptContent string, regionCode string, osTypeCode server.OsTypeCode, initScriptName string, initScriptDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, regionCode string, initScriptNoListN string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, regionCode string, initScriptNo string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, osTypeCode server.OsTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder server.SortingOrder, initScriptName string, initScriptNoListN string) (result autorest.Response, err error)
}

var _ InitScriptClientAPI = (*server.InitScriptClient)(nil)

// LoginKeyClientAPI contains the set of methods on the LoginKeyClient type.
type LoginKeyClientAPI interface {
	Create(ctx context.Context, responseFormatType string, keyName string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, keyNameListN string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, pageNo string, pageSize string) (result autorest.Response, err error)
	Import(ctx context.Context, responseFormatType string, publicKey string, keyName string) (result autorest.Response, err error)
}

var _ LoginKeyClientAPI = (*server.LoginKeyClient)(nil)

// NetworkInterfaceClientAPI contains the set of methods on the NetworkInterfaceClient type.
type NetworkInterfaceClientAPI interface {
	AddACG(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (result autorest.Response, err error)
	Attach(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, regionCode string, vpcNo string, subnetNo string, networkInterfaceName string, accessControlGroupNoListN string, serverInstanceNo string, IP string, networkInterfaceDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNo string) (result autorest.Response, err error)
	Detach(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, networkInterfaceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNoListN string, IP string, networkInterfaceName string, serverInstanceName string, subnetName string, pageNo string, pageSize string, networkInterfaceStatusCode server.NetworkInterfaceStatusCode) (result autorest.Response, err error)
	RemoveACG(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (result autorest.Response, err error)
}

var _ NetworkInterfaceClientAPI = (*server.NetworkInterfaceClient)(nil)

// PlacementGroupClientAPI contains the set of methods on the PlacementGroupClient type.
type PlacementGroupClientAPI interface {
	Add(ctx context.Context, responseFormatType string, placementGroupNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, regionCode string, placementGroupName string, placementGroupTypeCode string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, placementGroupNo string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, regionCode string, placementGroupNo string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, placementGroupName string, placementGroupNoListN string) (result autorest.Response, err error)
	Remove(ctx context.Context, responseFormatType string, placementGroupNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
}

var _ PlacementGroupClientAPI = (*server.PlacementGroupClient)(nil)

// PublicIPClientAPI contains the set of methods on the PublicIPClient type.
type PublicIPClientAPI interface {
	Associate(ctx context.Context, responseFormatType string, publicIPInstanceNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, regionCode string, serverInstanceNo string, publicIPDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, publicIPInstanceNo string, regionCode string) (result autorest.Response, err error)
	Disassociate(ctx context.Context, responseFormatType string, regionCode string, publicIPInstanceNo string, serverInstanceNo string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, publicIPInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, publicIPInstanceNoListN string, publicIP string, privateIP string, isAssociated *bool) (result autorest.Response, err error)
	GetTargetList(ctx context.Context, responseFormatType string, publicIPInstanceNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
}

var _ PublicIPClientAPI = (*server.PublicIPClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	ChangeSpec(ctx context.Context, responseFormatType string, serverInstanceNo string, serverProductCode string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, vpcNo string, subnetNo string, networkInterfaceListNnetworkInterfaceOrder string, networkInterfaceListNaccessControlGroupNoListN string, regionCode string, memberServerImageInstanceNo string, serverImageProductCode string, serverProductCode string, isEncryptedBaseBlockStorageVolume *bool, feeSystemTypeCode server.FeeSystemTypeCode, serverCreateCount string, serverCreateStartNo string, serverName string, networkInterfaceListNnetworkInterfaceNo string, networkInterfaceListNsubnetNo string, networkInterfaceListNip string, placementGroupNo string, isProtectServerTermination *bool, serverDescription string, initScriptNo string, loginKeyName string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetInstanceList(ctx context.Context, responseFormatType string, regionCode string, serverInstanceNoListN string, vpcNo string, pageNo string, serverInstanceStatusCode string, baseBlockStorageDiskTypeCode string, baseBlockStorageDiskDetailTypeCode string, serverName string, IP string, sortedBy string, sortingOrder string, placementGroupNoListN string) (result server.InstanceListResponse, err error)
	Reboot(ctx context.Context, responseFormatType string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	Start(ctx context.Context, responseFormatType string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	Stop(ctx context.Context, responseFormatType string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	Terminate(ctx context.Context, responseFormatType string, serverInstanceNoListN string, regionCode string) (result autorest.Response, err error)
}

var _ ClientAPI = (*server.Client)(nil)

// RootPasswordClientAPI contains the set of methods on the RootPasswordClient type.
type RootPasswordClientAPI interface {
	Get(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string, privateKey string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, rootPasswordServerInstanceListNserverInstanceNo string, regionCode string, rootPasswordServerInstanceListNprivateKey string) (result autorest.Response, err error)
}

var _ RootPasswordClientAPI = (*server.RootPasswordClient)(nil)

// ImageClientAPI contains the set of methods on the ImageClient type.
type ImageClientAPI interface {
	Create(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string, memberServerImageName string, memberServerImageDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, regionCode string, memberServerImageInstanceNoListN string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, memberServerImageInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, memberServerImageName string, memberServerImageInstanceStatusCode server.MemberServerImageInstanceStatusCode, memberServerImageInstanceNoListN string, platformTypeCodeListN server.PlatformTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder server.SortingOrder) (result autorest.Response, err error)
}

var _ ImageClientAPI = (*server.ImageClient)(nil)

// SnapshotClientAPI contains the set of methods on the SnapshotClient type.
type SnapshotClientAPI interface {
	Create(ctx context.Context, responseFormatType string, originalBlockStorageInstanceNo string, regionCode string, blockStorageSnapshotName string, blockStorageSnapshotDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, blockStorageSnapshotInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, blockStorageSnapshotInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, blockStorageSnapshotInstanceNoListN string, originalBlockStorageInstanceNoListN string, blockStorageSnapshotInstanceStatusCode server.BlockStorageSnapshotInstanceStatusCode, pageNo string, pageSize string, blockStorageSnapshotVolumeSize string, isEncryptedOriginalBlockStorageVolume *bool, blockStorageSnapshotName string, sortedBy string, sortingOrder server.SortingOrder) (result autorest.Response, err error)
}

var _ SnapshotClientAPI = (*server.SnapshotClient)(nil)

// BlockStorageClientAPI contains the set of methods on the BlockStorageClient type.
type BlockStorageClientAPI interface {
	Attach(ctx context.Context, responseFormatType string, blockStorageInstanceNo string, serverInstanceNo string, regionCode string) (result autorest.Response, err error)
	ChangeVolumeSize(ctx context.Context, responseFormatType string, blockStorageInstanceNo string, blockStorageSize string, regionCode string) (result autorest.Response, err error)
	Create(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string, zoneCode string, blockStorageName string, blockStorageDiskDetailTypeCode server.BlockStorageDiskDetailTypeCode, blockStorageSnapshotInstanceNo string, blockStorageSize string, blockStorageDescription string) (result autorest.Response, err error)
	Delete(ctx context.Context, responseFormatType string, blockStorageInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	Detach(ctx context.Context, responseFormatType string, blockStorageInstanceNoListN string, regionCode string) (result autorest.Response, err error)
	GetDetail(ctx context.Context, responseFormatType string, blockStorageInstanceNo string, regionCode string) (result autorest.Response, err error)
	GetList(ctx context.Context, responseFormatType string, regionCode string, serverInstanceNo string, blockStorageTypeCodeListN server.BlockStorageTypeCode, blockStorageInstanceStatusCode server.BlockStorageInstanceStatusCode, pageNo string, pageSize string, blockStorageSize string, blockStorageInstanceNoListN string, blockStorageName string, serverName string, connectionInfo string, blockStorageDiskTypeCode server.BlockStorageDiskTypeCode, blockStorageDiskDetailTypeCode server.BlockStorageDiskDetailTypeCode, zoneCode string) (result autorest.Response, err error)
}

var _ BlockStorageClientAPI = (*server.BlockStorageClient)(nil)
