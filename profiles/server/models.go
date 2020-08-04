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

package server

import original "github.com/samjegal/fincloud-sdk-for-go/services/server"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessControlGroupRuleTypeCode = original.AccessControlGroupRuleTypeCode

const (
	INBND AccessControlGroupRuleTypeCode = original.INBND
	OTBND AccessControlGroupRuleTypeCode = original.OTBND
)

type AccessControlGroupStatusCode = original.AccessControlGroupStatusCode

const (
	INIT     AccessControlGroupStatusCode = original.INIT
	RUN      AccessControlGroupStatusCode = original.RUN
	SET      AccessControlGroupStatusCode = original.SET
	TERMTING AccessControlGroupStatusCode = original.TERMTING
)

type BlockStorageDiskDetailTypeCode = original.BlockStorageDiskDetailTypeCode

const (
	HDD BlockStorageDiskDetailTypeCode = original.HDD
	SDD BlockStorageDiskDetailTypeCode = original.SDD
)

type BlockStorageDiskTypeCode = original.BlockStorageDiskTypeCode

const (
	NET BlockStorageDiskTypeCode = original.NET
)

type BlockStorageInstanceStatusCode = original.BlockStorageInstanceStatusCode

const (
	BlockStorageInstanceStatusCodeATTAC  BlockStorageInstanceStatusCode = original.BlockStorageInstanceStatusCodeATTAC
	BlockStorageInstanceStatusCodeCREATE BlockStorageInstanceStatusCode = original.BlockStorageInstanceStatusCodeCREATE
	BlockStorageInstanceStatusCodeINIT   BlockStorageInstanceStatusCode = original.BlockStorageInstanceStatusCodeINIT
)

type BlockStorageSnapshotInstanceStatusCode = original.BlockStorageSnapshotInstanceStatusCode

const (
	BlockStorageSnapshotInstanceStatusCodeATTAC BlockStorageSnapshotInstanceStatusCode = original.BlockStorageSnapshotInstanceStatusCodeATTAC
	BlockStorageSnapshotInstanceStatusCodeCREAT BlockStorageSnapshotInstanceStatusCode = original.BlockStorageSnapshotInstanceStatusCodeCREAT
	BlockStorageSnapshotInstanceStatusCodeINIT  BlockStorageSnapshotInstanceStatusCode = original.BlockStorageSnapshotInstanceStatusCodeINIT
)

type BlockStorageTypeCode = original.BlockStorageTypeCode

const (
	BASIC BlockStorageTypeCode = original.BASIC
	NSI   BlockStorageTypeCode = original.NSI
	SNAP  BlockStorageTypeCode = original.SNAP
	SVRBS BlockStorageTypeCode = original.SVRBS
)

type FeeSystemTypeCode = original.FeeSystemTypeCode

const (
	FXSUM FeeSystemTypeCode = original.FXSUM
	MTRAT FeeSystemTypeCode = original.MTRAT
)

type MemberServerImageInstanceStatusCode = original.MemberServerImageInstanceStatusCode

const (
	MemberServerImageInstanceStatusCodeCREATE MemberServerImageInstanceStatusCode = original.MemberServerImageInstanceStatusCodeCREATE
	MemberServerImageInstanceStatusCodeINIT   MemberServerImageInstanceStatusCode = original.MemberServerImageInstanceStatusCodeINIT
)

type NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCode

const (
	NetworkInterfaceStatusCodeNOTUSED NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCodeNOTUSED
	NetworkInterfaceStatusCodeSET     NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCodeSET
	NetworkInterfaceStatusCodeUNSET   NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCodeUNSET
	NetworkInterfaceStatusCodeUSED    NetworkInterfaceStatusCode = original.NetworkInterfaceStatusCodeUSED
)

type OsTypeCode = original.OsTypeCode

const (
	LNX OsTypeCode = original.LNX
	WND OsTypeCode = original.WND
)

type PlatformTypeCode = original.PlatformTypeCode

const (
	LNX32 PlatformTypeCode = original.LNX32
	LNX64 PlatformTypeCode = original.LNX64
	UBD64 PlatformTypeCode = original.UBD64
	UBS64 PlatformTypeCode = original.UBS64
	WND32 PlatformTypeCode = original.WND32
	WND64 PlatformTypeCode = original.WND64
)

type ProtocolTypeCode = original.ProtocolTypeCode

const (
	ICMP ProtocolTypeCode = original.ICMP
	TCP  ProtocolTypeCode = original.TCP
	UDP  ProtocolTypeCode = original.UDP
)

type SortingOrder = original.SortingOrder

const (
	ASC  SortingOrder = original.ASC
	DESC SortingOrder = original.DESC
)

type ACGClient = original.ACGClient
type ACGInboundClient = original.ACGInboundClient
type ACGOutboundClient = original.ACGOutboundClient
type BaseBlockStorageDiskDetailType = original.BaseBlockStorageDiskDetailType
type BaseBlockStorageDiskType = original.BaseBlockStorageDiskType
type BaseClient = original.BaseClient
type BlockStorageClient = original.BlockStorageClient
type Client = original.Client
type ImageClient = original.ImageClient
type ImageProductClient = original.ImageProductClient
type InitScriptClient = original.InitScriptClient
type InstanceList = original.InstanceList
type InstanceListResponse = original.InstanceListResponse
type InstanceOperation = original.InstanceOperation
type InstanceStatus = original.InstanceStatus
type InstanceType = original.InstanceType
type LoginKeyClient = original.LoginKeyClient
type NetworkInterfaceClient = original.NetworkInterfaceClient
type NetworkInterfaceNoList = original.NetworkInterfaceNoList
type PlacementGroupClient = original.PlacementGroupClient
type PlatformType = original.PlatformType
type ProductClient = original.ProductClient
type PublicIPClient = original.PublicIPClient
type RegionClient = original.RegionClient
type RootPasswordClient = original.RootPasswordClient
type SnapshotClient = original.SnapshotClient
type ZoneClient = original.ZoneClient

func New() BaseClient {
	return original.New()
}
func NewACGClient() ACGClient {
	return original.NewACGClient()
}
func NewACGClientWithBaseURI(baseURI string) ACGClient {
	return original.NewACGClientWithBaseURI(baseURI)
}
func NewACGInboundClient() ACGInboundClient {
	return original.NewACGInboundClient()
}
func NewACGInboundClientWithBaseURI(baseURI string) ACGInboundClient {
	return original.NewACGInboundClientWithBaseURI(baseURI)
}
func NewACGOutboundClient() ACGOutboundClient {
	return original.NewACGOutboundClient()
}
func NewACGOutboundClientWithBaseURI(baseURI string) ACGOutboundClient {
	return original.NewACGOutboundClientWithBaseURI(baseURI)
}
func NewBlockStorageClient() BlockStorageClient {
	return original.NewBlockStorageClient()
}
func NewBlockStorageClientWithBaseURI(baseURI string) BlockStorageClient {
	return original.NewBlockStorageClientWithBaseURI(baseURI)
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewImageClient() ImageClient {
	return original.NewImageClient()
}
func NewImageClientWithBaseURI(baseURI string) ImageClient {
	return original.NewImageClientWithBaseURI(baseURI)
}
func NewImageProductClient() ImageProductClient {
	return original.NewImageProductClient()
}
func NewImageProductClientWithBaseURI(baseURI string) ImageProductClient {
	return original.NewImageProductClientWithBaseURI(baseURI)
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
func NewPlacementGroupClient() PlacementGroupClient {
	return original.NewPlacementGroupClient()
}
func NewPlacementGroupClientWithBaseURI(baseURI string) PlacementGroupClient {
	return original.NewPlacementGroupClientWithBaseURI(baseURI)
}
func NewProductClient() ProductClient {
	return original.NewProductClient()
}
func NewProductClientWithBaseURI(baseURI string) ProductClient {
	return original.NewProductClientWithBaseURI(baseURI)
}
func NewPublicIPClient() PublicIPClient {
	return original.NewPublicIPClient()
}
func NewPublicIPClientWithBaseURI(baseURI string) PublicIPClient {
	return original.NewPublicIPClientWithBaseURI(baseURI)
}
func NewRegionClient() RegionClient {
	return original.NewRegionClient()
}
func NewRegionClientWithBaseURI(baseURI string) RegionClient {
	return original.NewRegionClientWithBaseURI(baseURI)
}
func NewRootPasswordClient() RootPasswordClient {
	return original.NewRootPasswordClient()
}
func NewRootPasswordClientWithBaseURI(baseURI string) RootPasswordClient {
	return original.NewRootPasswordClientWithBaseURI(baseURI)
}
func NewSnapshotClient() SnapshotClient {
	return original.NewSnapshotClient()
}
func NewSnapshotClientWithBaseURI(baseURI string) SnapshotClient {
	return original.NewSnapshotClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func NewZoneClient() ZoneClient {
	return original.NewZoneClient()
}
func NewZoneClientWithBaseURI(baseURI string) ZoneClient {
	return original.NewZoneClientWithBaseURI(baseURI)
}
func PossibleAccessControlGroupRuleTypeCodeValues() []AccessControlGroupRuleTypeCode {
	return original.PossibleAccessControlGroupRuleTypeCodeValues()
}
func PossibleAccessControlGroupStatusCodeValues() []AccessControlGroupStatusCode {
	return original.PossibleAccessControlGroupStatusCodeValues()
}
func PossibleBlockStorageDiskDetailTypeCodeValues() []BlockStorageDiskDetailTypeCode {
	return original.PossibleBlockStorageDiskDetailTypeCodeValues()
}
func PossibleBlockStorageDiskTypeCodeValues() []BlockStorageDiskTypeCode {
	return original.PossibleBlockStorageDiskTypeCodeValues()
}
func PossibleBlockStorageInstanceStatusCodeValues() []BlockStorageInstanceStatusCode {
	return original.PossibleBlockStorageInstanceStatusCodeValues()
}
func PossibleBlockStorageSnapshotInstanceStatusCodeValues() []BlockStorageSnapshotInstanceStatusCode {
	return original.PossibleBlockStorageSnapshotInstanceStatusCodeValues()
}
func PossibleBlockStorageTypeCodeValues() []BlockStorageTypeCode {
	return original.PossibleBlockStorageTypeCodeValues()
}
func PossibleFeeSystemTypeCodeValues() []FeeSystemTypeCode {
	return original.PossibleFeeSystemTypeCodeValues()
}
func PossibleMemberServerImageInstanceStatusCodeValues() []MemberServerImageInstanceStatusCode {
	return original.PossibleMemberServerImageInstanceStatusCodeValues()
}
func PossibleNetworkInterfaceStatusCodeValues() []NetworkInterfaceStatusCode {
	return original.PossibleNetworkInterfaceStatusCodeValues()
}
func PossibleOsTypeCodeValues() []OsTypeCode {
	return original.PossibleOsTypeCodeValues()
}
func PossiblePlatformTypeCodeValues() []PlatformTypeCode {
	return original.PossiblePlatformTypeCodeValues()
}
func PossibleProtocolTypeCodeValues() []ProtocolTypeCode {
	return original.PossibleProtocolTypeCodeValues()
}
func PossibleSortingOrderValues() []SortingOrder {
	return original.PossibleSortingOrderValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
