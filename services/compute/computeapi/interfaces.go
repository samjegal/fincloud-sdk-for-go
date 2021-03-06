package computeapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/compute"
	"io"
)

// SecurityGroupClientAPI contains the set of methods on the SecurityGroupClient type.
type SecurityGroupClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters compute.SecurityGroupNameParameter) (result compute.SecurityGroupListParameter, err error)
	Delete(ctx context.Context, parameters compute.SecurityGroupNumberParameter) (result autorest.Response, err error)
	List(ctx context.Context) (result compute.SecurityGroupListParameter, err error)
	ListByPrefixName(ctx context.Context, parameters compute.SecurityGroupNameParameter) (result compute.SecurityGroupContentListParameter, err error)
}

var _ SecurityGroupClientAPI = (*compute.SecurityGroupClient)(nil)

// InboundRuleClientAPI contains the set of methods on the InboundRuleClient type.
type InboundRuleClientAPI interface {
	Get(ctx context.Context, networkAcgNo string) (result compute.SecurityGroupRuleContentParameter, err error)
}

var _ InboundRuleClientAPI = (*compute.InboundRuleClient)(nil)

// OutboundRuleClientAPI contains the set of methods on the OutboundRuleClient type.
type OutboundRuleClientAPI interface {
	Get(ctx context.Context, networkAcgNo string) (result compute.SecurityGroupRuleContentParameter, err error)
}

var _ OutboundRuleClientAPI = (*compute.OutboundRuleClient)(nil)

// SecurityGroupRuleClientAPI contains the set of methods on the SecurityGroupRuleClient type.
type SecurityGroupRuleClientAPI interface {
	CreateOrUpdate(ctx context.Context, networkAcgNo string, parameters compute.SecurityGroupRulesProperties) (result autorest.Response, err error)
}

var _ SecurityGroupRuleClientAPI = (*compute.SecurityGroupRuleClient)(nil)

// ServerClientAPI contains the set of methods on the ServerClient type.
type ServerClientAPI interface {
	Create(ctx context.Context, parameter compute.ServerParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, parameter compute.ServerInstanceListParameter) (result autorest.Response, err error)
	Get(ctx context.Context, instanceNo string) (result compute.ServerDetailParameter, err error)
	IsReturn(ctx context.Context, instanceNoList int32) (result compute.ServerContentParameter, err error)
	List(ctx context.Context, parameter compute.ServerSearchParameter) (result compute.ServerListParameter, err error)
	Shutdown(ctx context.Context, parameter compute.ServerInstanceListParameter) (result autorest.Response, err error)
	Start(ctx context.Context, parameter compute.ServerInstanceListParameter) (result autorest.Response, err error)
}

var _ ServerClientAPI = (*compute.ServerClient)(nil)

// ServerContractRestrictionClientAPI contains the set of methods on the ServerContractRestrictionClient type.
type ServerContractRestrictionClientAPI interface {
	Get(ctx context.Context, productType2Code string, singleProductContractTypeCode string) (result compute.ServerContentContractRestrictionParameter, err error)
}

var _ ServerContractRestrictionClientAPI = (*compute.ServerContractRestrictionClient)(nil)

// ServerImageClientAPI contains the set of methods on the ServerImageClient type.
type ServerImageClientAPI interface {
	Create(ctx context.Context, instanceNo string, infraResourceTypeCode string, instanceName string, blockStorageUsageTypeCode string, parameter compute.ServerImageParameter) (result autorest.Response, err error)
}

var _ ServerImageClientAPI = (*compute.ServerImageClient)(nil)

// BlockStorageClientAPI contains the set of methods on the BlockStorageClient type.
type BlockStorageClientAPI interface {
	Get(ctx context.Context, instanceNo string) (result compute.BlockStorageContentParameter, err error)
}

var _ BlockStorageClientAPI = (*compute.BlockStorageClient)(nil)

// NetworkInterfaceClientAPI contains the set of methods on the NetworkInterfaceClient type.
type NetworkInterfaceClientAPI interface {
	Attach(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter compute.NetworkInterfaceSubnetParameter) (result compute.NetworkInterfaceAttachableContentParameter, err error)
	AttachableList(ctx context.Context, subnetNo int32) (result compute.NetworkInterfaceAttachableListParameter, err error)
	Create(ctx context.Context, parameters compute.NetworkInterfaceParameter) (result compute.ErrorMessageParameter, err error)
	Delete(ctx context.Context, networkInterfaceNo string) (result autorest.Response, err error)
	Detach(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter compute.NetworkInterfaceSubnetParameter) (result compute.NetworkInterfaceContentParameter, err error)
	List(ctx context.Context, parameters compute.NetworkInterfaceSearchParameter) (result compute.NetworkInterfaceListParameter, err error)
	ValidNetworkInterfaceIP(ctx context.Context, overlayIP string, subnetNo string) (result compute.NetworkInterfaceValidationParameter, err error)
}

var _ NetworkInterfaceClientAPI = (*compute.NetworkInterfaceClient)(nil)

// RootPasswordClientAPI contains the set of methods on the RootPasswordClient type.
type RootPasswordClientAPI interface {
	Get(ctx context.Context, instanceNo string, privateKeyFile io.ReadCloser) (result compute.RootPasswordContentParameter, err error)
}

var _ RootPasswordClientAPI = (*compute.RootPasswordClient)(nil)

// StorageClientAPI contains the set of methods on the StorageClient type.
type StorageClientAPI interface {
	Addable(ctx context.Context, diskType2DetailCode string, zoneNo string, blockStorageInstanceNo string) (result compute.StorageAddableParameter, err error)
	Create(ctx context.Context, parameters compute.StorageParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, parameters compute.StorageDetachAndDeleteParameter) (result autorest.Response, err error)
	Detach(ctx context.Context, parameters compute.StorageDetachAndDeleteParameter) (result autorest.Response, err error)
	List(ctx context.Context, pageNo string, pageSizeNo string, diskType2Code string, instanceStatusCode string) (result compute.StorageListParameter, err error)
}

var _ StorageClientAPI = (*compute.StorageClient)(nil)

// PublicIPAddressClientAPI contains the set of methods on the PublicIPAddressClient type.
type PublicIPAddressClientAPI interface {
	Assign(ctx context.Context, instanceNo string, parameters compute.PublicIPAddressServerInstanceParameter) (result compute.PublicIPAddressSearchParameter, err error)
	Create(ctx context.Context, parameters compute.PublicIPAddressServerInstanceParameter) (result compute.PublicIPAddressSearchParameter, err error)
	Delete(ctx context.Context, instanceNo string, parameters compute.PublicIPAddressServerInstanceParameter) (result compute.PublicIPAddressSearchParameter, err error)
	List(ctx context.Context, parameters compute.PublicIPAddressSearchFilterParameter) (result compute.PublicIPAddressSearchParameter, err error)
	Remove(ctx context.Context, instanceNo string, parameters compute.PublicIPAddressServerInstanceParameter) (result compute.PublicIPAddressSearchParameter, err error)
	ServerList(ctx context.Context) (result compute.PublicIPAddressServerListParameter, err error)
	Summary(ctx context.Context) (result compute.PublicIPAddressSummaryParameter, err error)
}

var _ PublicIPAddressClientAPI = (*compute.PublicIPAddressClient)(nil)

// LoginKeyClientAPI contains the set of methods on the LoginKeyClient type.
type LoginKeyClientAPI interface {
	Create(ctx context.Context, keyName string) (result compute.ReadCloser, err error)
	Delete(ctx context.Context, parameters compute.LoginKeyListParameter) (result autorest.Response, err error)
	List(ctx context.Context, infraResourceTypeCode string, instanceName string, blockStorageUsageTypeCode string) (result compute.LoginKeyContentParameter, err error)
}

var _ LoginKeyClientAPI = (*compute.LoginKeyClient)(nil)

// InitScriptClientAPI contains the set of methods on the InitScriptClient type.
type InitScriptClientAPI interface {
	Create(ctx context.Context, parameters compute.InitScriptParameter) (result autorest.Response, err error)
	Delete(ctx context.Context, parameters compute.InitScriptNumberListParameter) (result autorest.Response, err error)
	Get(ctx context.Context, initScriptNo string) (result compute.InitScriptDetailParameters, err error)
	List(ctx context.Context, pageNo string, pageSizeNo string) (result compute.InitScriptListParameter, err error)
	Update(ctx context.Context, initScriptNo string, parameters compute.InitScriptNumberParameter) (result autorest.Response, err error)
}

var _ InitScriptClientAPI = (*compute.InitScriptClient)(nil)
