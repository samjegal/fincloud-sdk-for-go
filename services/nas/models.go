package nas

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/nas"

// VolumeAccessControlAddResponse ...
type VolumeAccessControlAddResponse struct {
	autorest.Response                 `json:"-"`
	AddNasVolumeAccessControlResponse *VolumeAccessControlAddResponseAddNasVolumeAccessControlResponse `json:"addNasVolumeAccessControlResponse,omitempty"`
}

// VolumeAccessControlAddResponseAddNasVolumeAccessControlResponse ...
type VolumeAccessControlAddResponseAddNasVolumeAccessControlResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeAccessControlRemoveResponse ...
type VolumeAccessControlRemoveResponse struct {
	autorest.Response                    `json:"-"`
	RemoveNasVolumeAccessControlResponse *VolumeAccessControlRemoveResponseRemoveNasVolumeAccessControlResponse `json:"removeNasVolumeAccessControlResponse,omitempty"`
}

// VolumeAccessControlRemoveResponseRemoveNasVolumeAccessControlResponse ...
type VolumeAccessControlRemoveResponseRemoveNasVolumeAccessControlResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeAccessControlSetResponse ...
type VolumeAccessControlSetResponse struct {
	autorest.Response                 `json:"-"`
	SetNasVolumeAccessControlResponse *VolumeAccessControlSetResponseSetNasVolumeAccessControlResponse `json:"setNasVolumeAccessControlResponse,omitempty"`
}

// VolumeAccessControlSetResponseSetNasVolumeAccessControlResponse ...
type VolumeAccessControlSetResponseSetNasVolumeAccessControlResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeAllotmentProtocolType ...
type VolumeAllotmentProtocolType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// VolumeInstanceDetailResponse ...
type VolumeInstanceDetailResponse struct {
	autorest.Response                  `json:"-"`
	GetNasVolumeInstanceDetailResponse *VolumeInstanceDetailResponseGetNasVolumeInstanceDetailResponse `json:"getNasVolumeInstanceDetailResponse,omitempty"`
}

// VolumeInstanceDetailResponseGetNasVolumeInstanceDetailResponse ...
type VolumeInstanceDetailResponseGetNasVolumeInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeInstanceList ...
type VolumeInstanceList struct {
	NasVolumeInstanceNo              *string                      `json:"nasVolumeInstanceNo,omitempty"`
	NasVolumeInstanceStatus          *VolumeInstanceStatus        `json:"nasVolumeInstanceStatus,omitempty"`
	NasVolumeInstanceOperation       *VolumeInstanceOperation     `json:"nasVolumeInstanceOperation,omitempty"`
	NasVolumeInstanceStatusName      *string                      `json:"nasVolumeInstanceStatusName,omitempty"`
	CreateDate                       *string                      `json:"createDate,omitempty"`
	NasVolumeDescription             *string                      `json:"nasVolumeDescription,omitempty"`
	MountInformation                 *string                      `json:"mountInformation,omitempty"`
	VolumeAllotmentProtocolType      *VolumeAllotmentProtocolType `json:"volumeAllotmentProtocolType,omitempty"`
	VolumeName                       *string                      `json:"volumeName,omitempty"`
	VolumeTotalSize                  *float64                     `json:"volumeTotalSize,omitempty"`
	VolumeSize                       *float64                     `json:"volumeSize,omitempty"`
	VolumeUseSize                    *float64                     `json:"volumeUseSize,omitempty"`
	VolumeUseRatio                   *float64                     `json:"volumeUseRatio,omitempty"`
	SnapshotVolumeConfigurationRatio *float64                     `json:"snapshotVolumeConfigurationRatio,omitempty"`
	SnapshotVolumeSize               *float64                     `json:"snapshotVolumeSize,omitempty"`
	SnapshotVolumeUseSize            *float64                     `json:"snapshotVolumeUseSize,omitempty"`
	SnapshotVolumeUseRatio           *float64                     `json:"snapshotVolumeUseRatio,omitempty"`
	IsSnapshotConfiguration          *bool                        `json:"isSnapshotConfiguration,omitempty"`
	IsEventConfiguration             *bool                        `json:"isEventConfiguration,omitempty"`
	RegionCode                       *string                      `json:"regionCode,omitempty"`
	ZoneCode                         *string                      `json:"zoneCode,omitempty"`
	NasVolumeServerInstanceNoList    *VolumeServerInstanceNoList  `json:"nasVolumeServerInstanceNoList,omitempty"`
	IsEncryptedVolume                *bool                        `json:"isEncryptedVolume,omitempty"`
}

// VolumeInstanceListResponse ...
type VolumeInstanceListResponse struct {
	autorest.Response                `json:"-"`
	GetNasVolumeInstanceListResponse *VolumeInstanceListResponseGetNasVolumeInstanceListResponse `json:"getNasVolumeInstanceListResponse,omitempty"`
}

// VolumeInstanceListResponseGetNasVolumeInstanceListResponse ...
type VolumeInstanceListResponseGetNasVolumeInstanceListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeInstanceOperation ...
type VolumeInstanceOperation struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// VolumeInstancesCreateResponse ...
type VolumeInstancesCreateResponse struct {
	autorest.Response               `json:"-"`
	CreateNasVolumeInstanceResponse *VolumeInstancesCreateResponseCreateNasVolumeInstanceResponse `json:"createNasVolumeInstanceResponse,omitempty"`
}

// VolumeInstancesCreateResponseCreateNasVolumeInstanceResponse ...
type VolumeInstancesCreateResponseCreateNasVolumeInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeInstancesDeleteResponse ...
type VolumeInstancesDeleteResponse struct {
	autorest.Response                `json:"-"`
	DeleteNasVolumeInstancesResponse *VolumeInstancesDeleteResponseDeleteNasVolumeInstancesResponse `json:"deleteNasVolumeInstancesResponse,omitempty"`
}

// VolumeInstancesDeleteResponseDeleteNasVolumeInstancesResponse ...
type VolumeInstancesDeleteResponseDeleteNasVolumeInstancesResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}

// VolumeInstanceStatus ...
type VolumeInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// VolumeServerInstanceNoList ...
type VolumeServerInstanceNoList struct {
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
}

// VolumeSizeResponse ...
type VolumeSizeResponse struct {
	autorest.Response           `json:"-"`
	ChangeNasVolumeSizeResponse *VolumeSizeResponseChangeNasVolumeSizeResponse `json:"changeNasVolumeSizeResponse,omitempty"`
}

// VolumeSizeResponseChangeNasVolumeSizeResponse ...
type VolumeSizeResponseChangeNasVolumeSizeResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}
