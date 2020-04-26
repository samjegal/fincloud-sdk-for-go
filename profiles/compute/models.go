// +build go1.9

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/samjegal/fincloud-sdk-for-go/tools/profileBuilder

package compute

import original "github.com/samjegal/fincloud-sdk-for-go/services/compute"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BlockStorageInstanceStatusCode = original.BlockStorageInstanceStatusCode

const (
	ATTAC BlockStorageInstanceStatusCode = original.ATTAC
)

type DiskType2DetailCode = original.DiskType2DetailCode

const (
	HDD DiskType2DetailCode = original.HDD
	SSD DiskType2DetailCode = original.SSD
)

type NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCode

const (
	NOTUSED NetworkInterfaceStatusCode = original.NOTUSED
	USED    NetworkInterfaceStatusCode = original.USED
)

type ProtocolTypeCode = original.ProtocolTypeCode

const (
	ICMP ProtocolTypeCode = original.ICMP
	TCP  ProtocolTypeCode = original.TCP
	UDP  ProtocolTypeCode = original.UDP
)

type SecurityRuleStatusCode = original.SecurityRuleStatusCode

const (
	INIT SecurityRuleStatusCode = original.INIT
	RUN  SecurityRuleStatusCode = original.RUN
	SET  SecurityRuleStatusCode = original.SET
)

type ServerInstanceStatusCode = original.ServerInstanceStatusCode

const (
	ServerInstanceStatusCodeCREAT ServerInstanceStatusCode = original.ServerInstanceStatusCodeCREAT
	ServerInstanceStatusCodeINIT  ServerInstanceStatusCode = original.ServerInstanceStatusCodeINIT
	ServerInstanceStatusCodeNSTOP ServerInstanceStatusCode = original.ServerInstanceStatusCodeNSTOP
	ServerInstanceStatusCodeRUN   ServerInstanceStatusCode = original.ServerInstanceStatusCodeRUN
)

type ServerInstanceStatusName = original.ServerInstanceStatusName

const (
	Booting      ServerInstanceStatusName = original.Booting
	Creating     ServerInstanceStatusName = original.Creating
	Init         ServerInstanceStatusName = original.Init
	Running      ServerInstanceStatusName = original.Running
	Settingup    ServerInstanceStatusName = original.Settingup
	Shuttingdown ServerInstanceStatusName = original.Shuttingdown
	Stopped      ServerInstanceStatusName = original.Stopped
	Terminating  ServerInstanceStatusName = original.Terminating
)

type ServerOperationCode = original.ServerOperationCode

const (
	NULL  ServerOperationCode = original.NULL
	RESTA ServerOperationCode = original.RESTA
	SHTDN ServerOperationCode = original.SHTDN
	START ServerOperationCode = original.START
	TERMT ServerOperationCode = original.TERMT
)

type ServerStatusCode = original.ServerStatusCode

const (
	ServerStatusCodeNOTUSED ServerStatusCode = original.ServerStatusCodeNOTUSED
	ServerStatusCodeSET     ServerStatusCode = original.ServerStatusCodeSET
	ServerStatusCodeUNSET   ServerStatusCode = original.ServerStatusCodeUNSET
)

type StatusCode = original.StatusCode

const (
	StatusCodeCREATING StatusCode = original.StatusCodeCREATING
	StatusCodeRUN      StatusCode = original.StatusCodeRUN
	StatusCodeTERMTING StatusCode = original.StatusCodeTERMTING
)

type StorageStatusCode = original.StorageStatusCode

const (
	StorageStatusCodeATTAC StorageStatusCode = original.StorageStatusCodeATTAC
	StorageStatusCodeCREAT StorageStatusCode = original.StorageStatusCodeCREAT
)

type StorageStatusName = original.StorageStatusName

const (
	Attached    StorageStatusName = original.Attached
	Attaching   StorageStatusName = original.Attaching
	Detaching   StorageStatusName = original.Detaching
	Initialized StorageStatusName = original.Initialized
)

type BaseClient = original.BaseClient
type BlockStorageClient = original.BlockStorageClient
type BlockStorageContentParameter = original.BlockStorageContentParameter
type ErrorMessageParameter = original.ErrorMessageParameter
type ErrorMessageProperties = original.ErrorMessageProperties
type InboundRuleClient = original.InboundRuleClient
type InitScriptClient = original.InitScriptClient
type InitScriptContentProperties = original.InitScriptContentProperties
type InitScriptDetailParameters = original.InitScriptDetailParameters
type InitScriptListParameter = original.InitScriptListParameter
type InitScriptNumberListParameter = original.InitScriptNumberListParameter
type InitScriptNumberParameter = original.InitScriptNumberParameter
type InitScriptParameter = original.InitScriptParameter
type LoginKeyClient = original.LoginKeyClient
type LoginKeyContentParameter = original.LoginKeyContentParameter
type LoginKeyContentProperties = original.LoginKeyContentProperties
type LoginKeyListParameter = original.LoginKeyListParameter
type NetworkInterfaceAttachableContentParameter = original.NetworkInterfaceAttachableContentParameter
type NetworkInterfaceAttachableListParameter = original.NetworkInterfaceAttachableListParameter
type NetworkInterfaceAttachableListProperties = original.NetworkInterfaceAttachableListProperties
type NetworkInterfaceClient = original.NetworkInterfaceClient
type NetworkInterfaceContentParameter = original.NetworkInterfaceContentParameter
type NetworkInterfaceContentProperties = original.NetworkInterfaceContentProperties
type NetworkInterfaceListParameter = original.NetworkInterfaceListParameter
type NetworkInterfaceParameter = original.NetworkInterfaceParameter
type NetworkInterfaceSearchFilterProperties = original.NetworkInterfaceSearchFilterProperties
type NetworkInterfaceSearchParameter = original.NetworkInterfaceSearchParameter
type NetworkInterfaceSecurityGroupsProperties = original.NetworkInterfaceSecurityGroupsProperties
type NetworkInterfaceSubnetParameter = original.NetworkInterfaceSubnetParameter
type NetworkInterfaceValidationParameter = original.NetworkInterfaceValidationParameter
type OutboundRuleClient = original.OutboundRuleClient
type PublicIPAddressClient = original.PublicIPAddressClient
type PublicIPAddressSearchFilterParameter = original.PublicIPAddressSearchFilterParameter
type PublicIPAddressSearchParameter = original.PublicIPAddressSearchParameter
type PublicIPAddressSearchProperties = original.PublicIPAddressSearchProperties
type PublicIPAddressServerInstanceParameter = original.PublicIPAddressServerInstanceParameter
type PublicIPAddressServerListParameter = original.PublicIPAddressServerListParameter
type PublicIPAddressServerListProperties = original.PublicIPAddressServerListProperties
type PublicIPAddressSummaryParameter = original.PublicIPAddressSummaryParameter
type PublicIPAddressSummaryProperties = original.PublicIPAddressSummaryProperties
type PublicIPSearchFilterProperties = original.PublicIPSearchFilterProperties
type ReadCloser = original.ReadCloser
type RootPasswordClient = original.RootPasswordClient
type RootPasswordContentParameter = original.RootPasswordContentParameter
type SecurityGroupClient = original.SecurityGroupClient
type SecurityGroupContentListParameter = original.SecurityGroupContentListParameter
type SecurityGroupContentNameParameter = original.SecurityGroupContentNameParameter
type SecurityGroupListContentParameter = original.SecurityGroupListContentParameter
type SecurityGroupListParameter = original.SecurityGroupListParameter
type SecurityGroupNameParameter = original.SecurityGroupNameParameter
type SecurityGroupNumberParameter = original.SecurityGroupNumberParameter
type SecurityGroupParameter = original.SecurityGroupParameter
type SecurityGroupRuleClient = original.SecurityGroupRuleClient
type SecurityGroupRuleContentParameter = original.SecurityGroupRuleContentParameter
type SecurityGroupRuleProperties = original.SecurityGroupRuleProperties
type SecurityGroupRulesProperties = original.SecurityGroupRulesProperties
type ServerAttachableNetworkInterfaceContentParameterProperties = original.ServerAttachableNetworkInterfaceContentParameterProperties
type ServerBlockStorageContentParameterProperties = original.ServerBlockStorageContentParameterProperties
type ServerClient = original.ServerClient
type ServerContentContractRestrictionParameter = original.ServerContentContractRestrictionParameter
type ServerContentContractRestrictionParameterProperties = original.ServerContentContractRestrictionParameterProperties
type ServerContentParameter = original.ServerContentParameter
type ServerContractRestrictionClient = original.ServerContractRestrictionClient
type ServerDetailParameter = original.ServerDetailParameter
type ServerImageClient = original.ServerImageClient
type ServerImageParameter = original.ServerImageParameter
type ServerInstanceListParameter = original.ServerInstanceListParameter
type ServerInstanceStatus = original.ServerInstanceStatus
type ServerListAdditionalParameterMapProperties = original.ServerListAdditionalParameterMapProperties
type ServerListNetworkInterfaceParameterProperties = original.ServerListNetworkInterfaceParameterProperties
type ServerListParameter = original.ServerListParameter
type ServerListParametersProperties = original.ServerListParametersProperties
type ServerNetworkInterfaceContentParameterProperties = original.ServerNetworkInterfaceContentParameterProperties
type ServerNetworkInterfaceProperties = original.ServerNetworkInterfaceProperties
type ServerNetworkInterfaceSecurityGroupProperties = original.ServerNetworkInterfaceSecurityGroupProperties
type ServerParameter = original.ServerParameter
type ServerRootPasswordContentParameterProperties = original.ServerRootPasswordContentParameterProperties
type ServerSearchParameter = original.ServerSearchParameter
type StorageAddableParameter = original.StorageAddableParameter
type StorageAddableProperties = original.StorageAddableProperties
type StorageClient = original.StorageClient
type StorageDetachAndDeleteParameter = original.StorageDetachAndDeleteParameter
type StorageListParameter = original.StorageListParameter
type StorageListProperties = original.StorageListProperties
type StorageParameter = original.StorageParameter

func New() BaseClient {
	return original.New()
}
func NewBlockStorageClient() BlockStorageClient {
	return original.NewBlockStorageClient()
}
func NewBlockStorageClientWithBaseURI(baseURI string) BlockStorageClient {
	return original.NewBlockStorageClientWithBaseURI(baseURI)
}
func NewInboundRuleClient() InboundRuleClient {
	return original.NewInboundRuleClient()
}
func NewInboundRuleClientWithBaseURI(baseURI string) InboundRuleClient {
	return original.NewInboundRuleClientWithBaseURI(baseURI)
}
func NewInitScriptClient() InitScriptClient {
	return original.NewInitScriptClient()
}
func NewInitScriptClientWithBaseURI(baseURI string) InitScriptClient {
	return original.NewInitScriptClientWithBaseURI(baseURI)
}
func NewLoginKeyClient() LoginKeyClient {
	return original.NewLoginKeyClient()
}
func NewLoginKeyClientWithBaseURI(baseURI string) LoginKeyClient {
	return original.NewLoginKeyClientWithBaseURI(baseURI)
}
func NewNetworkInterfaceClient() NetworkInterfaceClient {
	return original.NewNetworkInterfaceClient()
}
func NewNetworkInterfaceClientWithBaseURI(baseURI string) NetworkInterfaceClient {
	return original.NewNetworkInterfaceClientWithBaseURI(baseURI)
}
func NewOutboundRuleClient() OutboundRuleClient {
	return original.NewOutboundRuleClient()
}
func NewOutboundRuleClientWithBaseURI(baseURI string) OutboundRuleClient {
	return original.NewOutboundRuleClientWithBaseURI(baseURI)
}
func NewPublicIPAddressClient() PublicIPAddressClient {
	return original.NewPublicIPAddressClient()
}
func NewPublicIPAddressClientWithBaseURI(baseURI string) PublicIPAddressClient {
	return original.NewPublicIPAddressClientWithBaseURI(baseURI)
}
func NewRootPasswordClient() RootPasswordClient {
	return original.NewRootPasswordClient()
}
func NewRootPasswordClientWithBaseURI(baseURI string) RootPasswordClient {
	return original.NewRootPasswordClientWithBaseURI(baseURI)
}
func NewSecurityGroupClient() SecurityGroupClient {
	return original.NewSecurityGroupClient()
}
func NewSecurityGroupClientWithBaseURI(baseURI string) SecurityGroupClient {
	return original.NewSecurityGroupClientWithBaseURI(baseURI)
}
func NewSecurityGroupRuleClient() SecurityGroupRuleClient {
	return original.NewSecurityGroupRuleClient()
}
func NewSecurityGroupRuleClientWithBaseURI(baseURI string) SecurityGroupRuleClient {
	return original.NewSecurityGroupRuleClientWithBaseURI(baseURI)
}
func NewServerClient() ServerClient {
	return original.NewServerClient()
}
func NewServerClientWithBaseURI(baseURI string) ServerClient {
	return original.NewServerClientWithBaseURI(baseURI)
}
func NewServerContractRestrictionClient() ServerContractRestrictionClient {
	return original.NewServerContractRestrictionClient()
}
func NewServerContractRestrictionClientWithBaseURI(baseURI string) ServerContractRestrictionClient {
	return original.NewServerContractRestrictionClientWithBaseURI(baseURI)
}
func NewServerImageClient() ServerImageClient {
	return original.NewServerImageClient()
}
func NewServerImageClientWithBaseURI(baseURI string) ServerImageClient {
	return original.NewServerImageClientWithBaseURI(baseURI)
}
func NewStorageClient() StorageClient {
	return original.NewStorageClient()
}
func NewStorageClientWithBaseURI(baseURI string) StorageClient {
	return original.NewStorageClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleBlockStorageInstanceStatusCodeValues() []BlockStorageInstanceStatusCode {
	return original.PossibleBlockStorageInstanceStatusCodeValues()
}
func PossibleDiskType2DetailCodeValues() []DiskType2DetailCode {
	return original.PossibleDiskType2DetailCodeValues()
}
func PossibleNetworkInterfaceStatusCodeValues() []NetworkInterfaceStatusCode {
	return original.PossibleNetworkInterfaceStatusCodeValues()
}
func PossibleProtocolTypeCodeValues() []ProtocolTypeCode {
	return original.PossibleProtocolTypeCodeValues()
}
func PossibleSecurityRuleStatusCodeValues() []SecurityRuleStatusCode {
	return original.PossibleSecurityRuleStatusCodeValues()
}
func PossibleServerInstanceStatusCodeValues() []ServerInstanceStatusCode {
	return original.PossibleServerInstanceStatusCodeValues()
}
func PossibleServerInstanceStatusNameValues() []ServerInstanceStatusName {
	return original.PossibleServerInstanceStatusNameValues()
}
func PossibleServerInstanceStatusValues() []ServerInstanceStatus {
	return original.PossibleServerInstanceStatusValues()
}
func PossibleServerOperationCodeValues() []ServerOperationCode {
	return original.PossibleServerOperationCodeValues()
}
func PossibleServerStatusCodeValues() []ServerStatusCode {
	return original.PossibleServerStatusCodeValues()
}
func PossibleStatusCodeValues() []StatusCode {
	return original.PossibleStatusCodeValues()
}
func PossibleStorageStatusCodeValues() []StorageStatusCode {
	return original.PossibleStorageStatusCodeValues()
}
func PossibleStorageStatusNameValues() []StorageStatusName {
	return original.PossibleStorageStatusNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
