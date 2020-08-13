package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/server"

// AccessControlGroupDetailResponse ...
type AccessControlGroupDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupList - ACG 리스트
	AccessControlGroupList *[]AccessControlGroupList `json:"accessControlGroupList,omitempty"`
}

// AccessControlGroupInboundRuleResponse ...
type AccessControlGroupInboundRuleResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupRuleList - ACG 룰 리스트
	AccessControlGroupRuleList *[]AccessControlGroupRuleList `json:"accessControlGroupRuleList,omitempty"`
}

// AccessControlGroupList ...
type AccessControlGroupList struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// IsDefault - 기본 여부
	IsDefault *bool `json:"isDefault,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// AccessControlGroupStatus - ACG 상태
	AccessControlGroupStatus *AccessControlGroupStatus `json:"accessControlGroupStatus,omitempty"`
	// AccessControlGroupDescription - ACG 설명
	AccessControlGroupDescription *string `json:"accessControlGroupDescription,omitempty"`
}

// AccessControlGroupListResponse ...
type AccessControlGroupListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupList - ACG 리스트
	AccessControlGroupList *[]AccessControlGroupList `json:"accessControlGroupList,omitempty"`
}

// AccessControlGroupNoList ...
type AccessControlGroupNoList struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`
}

// AccessControlGroupOutboundRuleResponse ...
type AccessControlGroupOutboundRuleResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupRuleList - ACG 룰 리스트
	AccessControlGroupRuleList *[]AccessControlGroupRuleList `json:"accessControlGroupRuleList,omitempty"`
}

// AccessControlGroupResponse ...
type AccessControlGroupResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupList - ACG 리스트
	AccessControlGroupList *[]AccessControlGroupList `json:"accessControlGroupList,omitempty"`
}

// AccessControlGroupRuleList ...
type AccessControlGroupRuleList struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`
	// ProtocolType - 프로토콜 타입
	ProtocolType *ProtocolType `json:"protocolType,omitempty"`
	// IPBlock - 차단 여부
	IPBlock *string `json:"ipBlock,omitempty"`
	// AccessControlGroupSequence - ACG 시퀀스
	AccessControlGroupSequence *string `json:"accessControlGroupSequence,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// AccessControlGroupRuleType - ACG 룰 타입
	AccessControlGroupRuleType *AccessControlGroupRuleType `json:"accessControlGroupRuleType,omitempty"`
	// AccessControlGroupDescription - ACG 설명
	AccessControlGroupDescription *string `json:"accessControlGroupDescription,omitempty"`
}

// AccessControlGroupRuleListResponse ...
type AccessControlGroupRuleListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// AccessControlGroupRuleList - ACG 룰 리스트
	AccessControlGroupRuleList *[]AccessControlGroupRuleList `json:"accessControlGroupRuleList,omitempty"`
}

// AccessControlGroupRuleType ...
type AccessControlGroupRuleType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// AccessControlGroupStatus ...
type AccessControlGroupStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BaseBlockStorageDiskDetailType ...
type BaseBlockStorageDiskDetailType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BaseBlockStorageDiskType ...
type BaseBlockStorageDiskType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageDiskDetailType ...
type BlockStorageDiskDetailType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageDiskType ...
type BlockStorageDiskType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageInstanceDetailResponse ...
type BlockStorageInstanceDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// BlockStorageInstanceList - 블럭 스토리지 인스턴스 리스트
	BlockStorageInstanceList *[]BlockStorageInstanceList `json:"blockStorageInstanceList,omitempty"`
}

// BlockStorageInstanceList ...
type BlockStorageInstanceList struct {
	// BlockStorageInstanceNo - 블럭 스토리지 인스턴스 번호
	BlockStorageInstanceNo *string `json:"blockStorageInstanceNo,omitempty"`
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// BlockStorageName - 블럭 스토리지 이름
	BlockStorageName *string `json:"blockStorageName,omitempty"`
	// BlockStorageType - 블럭 스토리지 타입
	BlockStorageType *BlockStorageType `json:"blockStorageType,omitempty"`
	// BlockStorageSize - 블럭 스토리지 크기
	BlockStorageSize *float64 `json:"blockStorageSize,omitempty"`
	// DeviceName - 장치 이름
	DeviceName *string `json:"deviceName,omitempty"`
	// BlockStorageProductCode - 블럭 스토리지 제품 코드
	BlockStorageProductCode *string `json:"blockStorageProductCode,omitempty"`
	// BlockStorageInstanceStatus - 블럭 스토리지 인스턴스 상태
	BlockStorageInstanceStatus *BlockStorageInstanceStatus `json:"blockStorageInstanceStatus,omitempty"`
	// BlockStorageInstanceOperation - 블럭 스토리지 인스턴스 운영
	BlockStorageInstanceOperation *BlockStorageInstanceOperation `json:"blockStorageInstanceOperation,omitempty"`
	// BlockStorageInstanceStatusName - 블럭 스토리지 인스턴스 상태 이름
	BlockStorageInstanceStatusName *string `json:"blockStorageInstanceStatusName,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// BlockStorageDescription - 블럭 스토리지 설명
	BlockStorageDescription *string `json:"blockStorageDescription,omitempty"`
	// BlockStorageDiskType - 블럭 스토리지 디스크 타입
	BlockStorageDiskType *BlockStorageDiskType `json:"blockStorageDiskType,omitempty"`
	// BlockStorageDiskDetailType - 블럭 스토리지 디스크 상세 타입
	BlockStorageDiskDetailType *BlockStorageDiskDetailType `json:"blockStorageDiskDetailType,omitempty"`
	// MaxIopsThroughput - 최대 IOPS 성녕
	MaxIopsThroughput *float64 `json:"maxIopsThroughput,omitempty"`
	// IsEncryptedVolume - 볼륨 암호화 여부
	IsEncryptedVolume *bool `json:"isEncryptedVolume,omitempty"`
	// ZoneCode - ZONE 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
}

// BlockStorageInstanceListResponse ...
type BlockStorageInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// BlockStorageInstanceList - 블럭 스토리지 인스턴스 리스트
	BlockStorageInstanceList *[]BlockStorageInstanceList `json:"blockStorageInstanceList,omitempty"`
}

// BlockStorageInstanceOperation ...
type BlockStorageInstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageInstanceResponse ...
type BlockStorageInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// BlockStorageInstanceList - 블럭 스토리지 인스턴스 리스트
	BlockStorageInstanceList *[]BlockStorageInstanceList `json:"blockStorageInstanceList,omitempty"`
}

// BlockStorageInstanceStatus ...
type BlockStorageInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageSnapshotInstanceDetailResponse ...
type BlockStorageSnapshotInstanceDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - 블럭 스토리 스냅샷 인스턴스 리스트
	VpcList *[]BlockStorageSnapshotInstanceList `json:"vpcList,omitempty"`
}

// BlockStorageSnapshotInstanceList ...
type BlockStorageSnapshotInstanceList struct {
	// BlockStorageSnapshotInstanceNo - 블럭 스토리지 스냅샷 인스턴스 번호
	BlockStorageSnapshotInstanceNo *string `json:"blockStorageSnapshotInstanceNo,omitempty"`
	// BlockStorageSnapshotName - 블럭 스토리지 스냅샷 이름
	BlockStorageSnapshotName *string `json:"blockStorageSnapshotName,omitempty"`
	// BlockStorageSnapshotVolumeSize - 블럭 스토리지 스냅샷 볼륨 크기
	BlockStorageSnapshotVolumeSize *float64 `json:"blockStorageSnapshotVolumeSize,omitempty"`
	// OriginalBlockStorageInstanceNo - 원본 블럭 스토리지 인스턴스 번호
	OriginalBlockStorageInstanceNo *string `json:"originalBlockStorageInstanceNo,omitempty"`
	// BlockStorageSnapshotInstanceStatus - 블럭 스토리지 스냅샷 인스턴스 상태
	BlockStorageSnapshotInstanceStatus *BlockStorageSnapshotInstanceStatus `json:"blockStorageSnapshotInstanceStatus,omitempty"`
	// BlockStorageSnapshotInstanceOperation - 블럭 스토리지 스냅샷 인스턴스 운영
	BlockStorageSnapshotInstanceOperation *BlockStorageSnapshotInstanceOperation `json:"blockStorageSnapshotInstanceOperation,omitempty"`
	// BlockStorageSnapshotInstanceStatusName - 블럭 스토리지 스냅샷 인스턴스 상태 이름
	BlockStorageSnapshotInstanceStatusName *string `json:"blockStorageSnapshotInstanceStatusName,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// IsEncryptedOriginalBlockStorageVolume - 원본 블럭 스토리지 볼륨 암호화 여부
	IsEncryptedOriginalBlockStorageVolume *bool `json:"isEncryptedOriginalBlockStorageVolume,omitempty"`
	// BlockStorageSnapshotDescription - 블럭 스토리지 스냅샷 설명
	BlockStorageSnapshotDescription *string `json:"blockStorageSnapshotDescription,omitempty"`
}

// BlockStorageSnapshotInstanceListResponse ...
type BlockStorageSnapshotInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - 블럭 스토리 스냅샷 인스턴스 리스트
	VpcList *[]BlockStorageSnapshotInstanceList `json:"vpcList,omitempty"`
}

// BlockStorageSnapshotInstanceOperation ...
type BlockStorageSnapshotInstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageSnapshotInstanceResponse ...
type BlockStorageSnapshotInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - 블럭 스토리 스냅샷 인스턴스 리스트
	VpcList *[]BlockStorageSnapshotInstanceList `json:"vpcList,omitempty"`
}

// BlockStorageSnapshotInstanceStatus ...
type BlockStorageSnapshotInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageType ...
type BlockStorageType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// BlockStorageVolumeSizeResponse ...
type BlockStorageVolumeSizeResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// BlockStorageInstanceList - 블럭 스토리지 인스턴스 리스트
	BlockStorageInstanceList *[]BlockStorageInstanceList `json:"blockStorageInstanceList,omitempty"`
}

// CreateLoginKeyResponse ...
type CreateLoginKeyResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// KeyName - 키 이름
	KeyName *string `json:"keyName,omitempty"`
	// PrivateKey - 비밀키
	PrivateKey *string `json:"privateKey,omitempty"`
}

// DeleteLoginKeysResponse ...
type DeleteLoginKeysResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
}

// ImageProductListResponse ...
type ImageProductListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ProductList - 제품 리스트
	ProductList *[]ProductList `json:"productList,omitempty"`
}

// InfraResourceType ...
type InfraResourceType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// InitScriptDetailResponse ...
type InitScriptDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// InitScriptList - 초기화 스크립트 리스트
	InitScriptList *[]InitScriptList `json:"initScriptList,omitempty"`
}

// InitScriptList ...
type InitScriptList struct {
	// InitScriptNo - 초기화 스크립트 번호
	InitScriptNo *string `json:"initScriptNo,omitempty"`
	// InitScriptName - 초기화 스크립트 이름
	InitScriptName *string `json:"initScriptName,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// InitScriptDescription - 초기화 스크립트 설명
	InitScriptDescription *string `json:"initScriptDescription,omitempty"`
	// InitScriptContent - 초기화 스크립트 컨텐츠
	InitScriptContent *string `json:"initScriptContent,omitempty"`
	// OsType - 운영체제 타입
	OsType *OsType `json:"osType,omitempty"`
}

// InitScriptListResponse ...
type InitScriptListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// InitScriptList - 초기화 스크립트 리스트
	InitScriptList *[]InitScriptList `json:"initScriptList,omitempty"`
}

// InitScriptResponse ...
type InitScriptResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// InitScriptList - 초기화 스크립트 리스트
	InitScriptList *[]InitScriptList `json:"initScriptList,omitempty"`
}

// InstanceDetailResponse ...
type InstanceDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// InstanceList ...
type InstanceList struct {
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// ServerName - 서버 이름
	ServerName *string `json:"serverName,omitempty"`
	// ServerDescription - 서버 설명
	ServerDescription *string `json:"serverDescription,omitempty"`
	// CPUCount - CPU 코어 개수
	CPUCount *int32 `json:"cpuCount,omitempty"`
	// MemorySize - 메모리 크기
	MemorySize *float64 `json:"memorySize,omitempty"`
	// PlatformType - 플랫폼 타입
	PlatformType *PlatformType `json:"platformType,omitempty"`
	// LoginKeyName - 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// PublicIPInstanceNo - 공인 IP 인스턴스 번호
	PublicIPInstanceNo *string `json:"publicIpInstanceNo,omitempty"`
	// PublicIP - 공인 IP 주소
	PublicIP *string `json:"publicIp,omitempty"`
	// ServerInstanceStatus - 서버 인스턴스 상태
	ServerInstanceStatus *InstanceStatus `json:"serverInstanceStatus,omitempty"`
	// ServerInstanceOperation - 서버 인스턴스 운영
	ServerInstanceOperation *InstanceOperation `json:"serverInstanceOperation,omitempty"`
	// ServerInstanceStatusName - 서버 인스턴스 상태 이름
	ServerInstanceStatusName *string `json:"serverInstanceStatusName,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// Uptime - 운영 시간
	Uptime *string `json:"uptime,omitempty"`
	// ServerImageProductCode - 서버 이미지 제품 코드
	ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`
	// ServerProductCode - 서버 제품 코드
	ServerProductCode *string `json:"serverProductCode,omitempty"`
	// IsProtectServerTermination - 서버 반납 방지 여부
	IsProtectServerTermination *bool `json:"isProtectServerTermination,omitempty"`
	// ZoneCode - ZONE 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// NetworkInterfaceNoList - 네트워크 인터페이스 번호 리스트
	NetworkInterfaceNoList *NetworkInterfaceNoList `json:"networkInterfaceNoList,omitempty"`
	// InitScriptNo - 초기화 스크립트 번호
	InitScriptNo *string `json:"initScriptNo,omitempty"`
	// ServerInstanceType - 서버 인스턴스 타입
	ServerInstanceType *InstanceTypeType `json:"serverInstanceType,omitempty"`
	// BaseBlockStorageDiskType - 기본 블럭 스토리지 디스크 타입
	BaseBlockStorageDiskType *BaseBlockStorageDiskType `json:"baseBlockStorageDiskType,omitempty"`
	// BaseBlockStorageDiskDetailType - 기본 블럭 스토리지 디스크 상세 타입
	BaseBlockStorageDiskDetailType *BaseBlockStorageDiskDetailType `json:"baseBlockStorageDiskDetailType,omitempty"`
	// PlacementGroupNo - 물리 배치 그룹 번호
	PlacementGroupNo *string `json:"placementGroupNo,omitempty"`
	// PlacementGroupName - 물리 배치 그룹 이름
	PlacementGroupName *string `json:"placementGroupName,omitempty"`
}

// InstanceListResponse ...
type InstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// InstanceOperation ...
type InstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceSpecResponse ...
type InstanceSpecResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// InstancesResponse ...
type InstancesResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// InstanceStatus ...
type InstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceType ...
type InstanceType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// InstanceTypeType ...
type InstanceTypeType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// LoginKeyList ...
type LoginKeyList struct {
	// Fingerprint - 리턴 코드
	Fingerprint *string `json:"fingerprint,omitempty"`
	// KeyName - 리턴 코드
	KeyName *string `json:"keyName,omitempty"`
	// CreateDate - 리턴 코드
	CreateDate *string `json:"createDate,omitempty"`
}

// LoginKeyListResponse ...
type LoginKeyListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// LoginKeyList - 로그인 키 리스트
	LoginKeyList *[]LoginKeyList `json:"loginKeyList,omitempty"`
}

// LoginKeyResponse ...
type LoginKeyResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// LoginKeyList - 로그인 키 리스트
	LoginKeyList *[]LoginKeyList `json:"loginKeyList,omitempty"`
}

// MemberServerImageInstanceDetailResponse ...
type MemberServerImageInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// MemberServerImageInstanceList - 서버 이미지 인스턴스 리스트
	MemberServerImageInstanceList *[]MemberServerImageInstanceList `json:"memberServerImageInstanceList,omitempty"`
}

// MemberServerImageInstanceList ...
type MemberServerImageInstanceList struct {
	// MemberServerImageInstanceNo - 서버 이미지 인스턴스 번호
	MemberServerImageInstanceNo *string `json:"memberServerImageInstanceNo,omitempty"`
	// MemberServerImageName - 서버 이미지 이름
	MemberServerImageName *string `json:"memberServerImageName,omitempty"`
	// MemberServerImageDescription - 서버 이미지 설명
	MemberServerImageDescription *string `json:"memberServerImageDescription,omitempty"`
	// OriginalServerInstanceNo - 서버 이미지 번호
	OriginalServerInstanceNo *string `json:"originalServerInstanceNo,omitempty"`
	// OriginalServerImageProductCode - 서버 이미지 제품 코드
	OriginalServerImageProductCode *string `json:"originalServerImageProductCode,omitempty"`
	// MemberServerImageInstanceStatus - 서버 이미지 인스턴스 상태
	MemberServerImageInstanceStatus *MemberServerImageInstanceStatus `json:"memberServerImageInstanceStatus,omitempty"`
	// MemberServerImageInstanceOperation - 서버 이미지 인스턴스 운영
	MemberServerImageInstanceOperation *MemberServerImageInstanceOperation `json:"memberServerImageInstanceOperation,omitempty"`
	// MemberServerImageInstanceStatusName - 서버 이미지 인스턴스 상태 이름
	MemberServerImageInstanceStatusName *string `json:"memberServerImageInstanceStatusName,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
}

// MemberServerImageInstanceListResponse ...
type MemberServerImageInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// MemberServerImageInstanceList - 서버 이미지 인스턴스 리스트
	MemberServerImageInstanceList *[]MemberServerImageInstanceList `json:"memberServerImageInstanceList,omitempty"`
}

// MemberServerImageInstanceOperation ...
type MemberServerImageInstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// MemberServerImageInstanceResponse ...
type MemberServerImageInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// MemberServerImageInstanceList - 서버 이미지 인스턴스 리스트
	MemberServerImageInstanceList *[]MemberServerImageInstanceList `json:"memberServerImageInstanceList,omitempty"`
}

// MemberServerImageInstanceStatus ...
type MemberServerImageInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkInterfaceAccessControlGroupResponse ...
type NetworkInterfaceAccessControlGroupResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkInterfaceList - 네트워크 인터페이스 리스트
	NetworkInterfaceList *[]NetworkInterfaceList `json:"networkInterfaceList,omitempty"`
}

// NetworkInterfaceDetailResponse ...
type NetworkInterfaceDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkInterfaceList - 네트워크 인터페이스 리스트
	NetworkInterfaceList *[]NetworkInterfaceList `json:"networkInterfaceList,omitempty"`
}

// NetworkInterfaceList ...
type NetworkInterfaceList struct {
	// NetworkInterfaceNo - 네트워크 인터페이스 번호
	NetworkInterfaceNo *string `json:"networkInterfaceNo,omitempty"`
	// NetworkInterfaceName - 네트워크 인터페이스 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// DeleteOnTermination - 반납 후 삭제 여부
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty"`
	// IsDefault - 기본 여부
	IsDefault *bool `json:"isDefault,omitempty"`
	// DeviceName - 장치 이름
	DeviceName *string `json:"deviceName,omitempty"`
	// NetworkInterfaceStatus - 네트워크 인터페이스 상태
	NetworkInterfaceStatus *NetworkInterfaceStatus `json:"networkInterfaceStatus,omitempty"`
	// InstanceType - 인스턴스 타입
	InstanceType *InstanceType `json:"instanceType,omitempty"`
	// InstanceNo - 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// IP - IP 주소
	IP *string `json:"ip,omitempty"`
	// AccessControlGroupNoList - ACG 번호 리스트
	AccessControlGroupNoList *AccessControlGroupNoList `json:"accessControlGroupNoList,omitempty"`
}

// NetworkInterfaceListResponse ...
type NetworkInterfaceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkInterfaceList - 네트워크 인터페이스 리스트
	NetworkInterfaceList *[]NetworkInterfaceList `json:"networkInterfaceList,omitempty"`
}

// NetworkInterfaceNoList ...
type NetworkInterfaceNoList struct {
	// NetworkInterfaceNo - 네트워크 인터페이스 번호
	NetworkInterfaceNo *string `json:"networkInterfaceNo,omitempty"`
}

// NetworkInterfaceResponse ...
type NetworkInterfaceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkInterfaceList - 네트워크 인터페이스 리스트
	NetworkInterfaceList *[]NetworkInterfaceList `json:"networkInterfaceList,omitempty"`
}

// NetworkInterfaceStatus ...
type NetworkInterfaceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// OsType ...
type OsType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PlacementGroupDetailResponse ...
type PlacementGroupDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PlacementGroupList - 물리 배치 그룹 리스트
	PlacementGroupList *[]PlacementGroupList `json:"placementGroupList,omitempty"`
}

// PlacementGroupList ...
type PlacementGroupList struct {
	// PlacementGroupNo - 믈리 배치 그룹 번호
	PlacementGroupNo *string `json:"placementGroupNo,omitempty"`
	// PlacementGroupName - 물리 배치 이름
	PlacementGroupName *string `json:"placementGroupName,omitempty"`
	// PlacementGroupType - 물리 배치 그룹 타입
	PlacementGroupType *PlacementGroupType `json:"placementGroupType,omitempty"`
}

// PlacementGroupListResponse ...
type PlacementGroupListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PlacementGroupList - 물리 배치 그룹 리스트
	PlacementGroupList *[]PlacementGroupList `json:"placementGroupList,omitempty"`
}

// PlacementGroupResponse ...
type PlacementGroupResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PlacementGroupList - 물리 배치 그룹 리스트
	PlacementGroupList *[]PlacementGroupList `json:"placementGroupList,omitempty"`
}

// PlacementGroupServerInstanceResponse ...
type PlacementGroupServerInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// PlacementGroupType ...
type PlacementGroupType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PlatformType ...
type PlatformType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// ProductList ...
type ProductList struct {
	// ProductCode - 제품 코드
	ProductCode *string `json:"productCode,omitempty"`
	// ProductName - 제품 이름
	ProductName *string `json:"productName,omitempty"`
	// ProductType - 제품 타입
	ProductType *ProductType `json:"productType,omitempty"`
	// ProductDescription - 제품 설명
	ProductDescription *string `json:"productDescription,omitempty"`
	// InfraResourceType - 인프라 리소스 타입
	InfraResourceType *InfraResourceType `json:"infraResourceType,omitempty"`
	// CPUCount - CPU 코어 개수
	CPUCount *float64 `json:"cpuCount,omitempty"`
	// MemorySize - 메모리 크기
	MemorySize *float64 `json:"memorySize,omitempty"`
	// BaseBlockStorageSize - 기본 블럭 스토리지 크기
	BaseBlockStorageSize *float64 `json:"baseBlockStorageSize,omitempty"`
	// PlatformType - 플랫폼 타입
	PlatformType *PlatformType `json:"platformType,omitempty"`
	// OsInformation - OS 정보
	OsInformation *string `json:"osInformation,omitempty"`
	// DbKindCode - 데이터베이스 종류 코드
	DbKindCode *string `json:"dbKindCode,omitempty"`
	// AddBlockStorageSize - 추가 블럭 스토리지 크기
	AddBlockStorageSize *float64 `json:"addBlockStorageSize,omitempty"`
	// GenerationCode - 생성 코드
	GenerationCode *string `json:"generationCode,omitempty"`
}

// ProductListResponse ...
type ProductListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ProductList - 제품 리스트
	ProductList *[]ProductList `json:"productList,omitempty"`
}

// ProductType ...
type ProductType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// ProtocolType ...
type ProtocolType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PublicIPFromServerInstanceResponse ...
type PublicIPFromServerInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PublicIPInstanceList - 공인 IP 인스턴스 리스트
	PublicIPInstanceList *[]PublicIPInstanceList `json:"publicIpInstanceList,omitempty"`
}

// PublicIPInstanceDetailResponse ...
type PublicIPInstanceDetailResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PublicIPInstanceList - 공인 IP 인스턴스 리스트
	PublicIPInstanceList *[]PublicIPInstanceList `json:"publicIpInstanceList,omitempty"`
}

// PublicIPInstanceList ...
type PublicIPInstanceList struct {
	// PublicIPInstanceNo - 공인 IP 인스턴스 번호
	PublicIPInstanceNo *string `json:"publicIpInstanceNo,omitempty"`
	// PublicIP - 공인 IP
	PublicIP *string `json:"publicIp,omitempty"`
	// PublicIPDescription - 공인 IP 설명
	PublicIPDescription *string `json:"publicIpDescription,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// PublicIPInstanceStatusName - 공인 IP 인스턴스 상태 이름
	PublicIPInstanceStatusName *string `json:"publicIpInstanceStatusName,omitempty"`
	// PublicIPInstanceStatus - 공인 IP 인스턴스 상태
	PublicIPInstanceStatus *PublicIPInstanceStatus `json:"publicIpInstanceStatus,omitempty"`
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// ServerName - 서버 이름
	ServerName *string `json:"serverName,omitempty"`
	// PrivateIP - 사설 IP
	PrivateIP *string `json:"privateIp,omitempty"`
	// PublicIPInstanceOperation - 공인 IP 인스턴스 운영
	PublicIPInstanceOperation *PublicIPInstanceOperation `json:"publicIpInstanceOperation,omitempty"`
}

// PublicIPInstanceListResponse ...
type PublicIPInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PublicIPInstanceList - 공인 IP 인스턴스 리스트
	PublicIPInstanceList *[]PublicIPInstanceList `json:"publicIpInstanceList,omitempty"`
}

// PublicIPInstanceOperation ...
type PublicIPInstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PublicIPInstanceResponse ...
type PublicIPInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PublicIPInstanceList - 공인 IP 인스턴스 리스트
	PublicIPInstanceList *[]PublicIPInstanceList `json:"publicIpInstanceList,omitempty"`
}

// PublicIPInstanceStatus ...
type PublicIPInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PublicIPTargetServerInstanceListResponse ...
type PublicIPTargetServerInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ServerInstanceList - 서버 인스턴스 리스트
	ServerInstanceList *[]InstanceList `json:"serverInstanceList,omitempty"`
}

// PublicIPWithServerInstanceResponse ...
type PublicIPWithServerInstanceResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// PublicIPInstanceList - 공인 IP 인스턴스 리스트
	PublicIPInstanceList *[]PublicIPInstanceList `json:"publicIpInstanceList,omitempty"`
}

// RegionList ...
type RegionList struct {
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// RegionName - 리전 이름
	RegionName *string `json:"regionName,omitempty"`
}

// RegionListResponse ...
type RegionListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RegionList - 리전 리스트
	RegionList *[]RegionList `json:"regionList,omitempty"`
}

// RootPassword ...
type RootPassword struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RootPassword - 관리자 비밀번호
	RootPassword *string `json:"rootPassword,omitempty"`
}

// RootPasswordServerInstanceList ...
type RootPasswordServerInstanceList struct {
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// RootPassword - 관리자 비밀번호
	RootPassword *string `json:"rootPassword,omitempty"`
}

// RootPasswordServerInstanceListResponse ...
type RootPasswordServerInstanceListResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RootPasswordServerInstanceList - 관리자 비밀번호 리스트
	RootPasswordServerInstanceList *[]RootPasswordServerInstanceList `json:"rootPasswordServerInstanceList,omitempty"`
}

// ZoneList ...
type ZoneList struct {
	// ZoneName - ZONE 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// ZoneCode - ZONE 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
}

// ZoneListResponse ...
type ZoneListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// ZoneList - ZONE 리스트
	ZoneList *[]ZoneList `json:"zoneList,omitempty"`
}
