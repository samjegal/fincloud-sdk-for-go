package nas

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/nas"

// EncryptedVolume enumerates the values for encrypted volume.
type EncryptedVolume string

const (
	// False ...
	False EncryptedVolume = "false"
	// True ...
	True EncryptedVolume = "true"
)

// PossibleEncryptedVolumeValues returns an array of possible values for the EncryptedVolume const type.
func PossibleEncryptedVolumeValues() []EncryptedVolume {
	return []EncryptedVolume{False, True}
}

// EventConfiguration enumerates the values for event configuration.
type EventConfiguration string

const (
	// EventConfigurationFalse ...
	EventConfigurationFalse EventConfiguration = "false"
	// EventConfigurationTrue ...
	EventConfigurationTrue EventConfiguration = "true"
)

// PossibleEventConfigurationValues returns an array of possible values for the EventConfiguration const type.
func PossibleEventConfigurationValues() []EventConfiguration {
	return []EventConfiguration{EventConfigurationFalse, EventConfigurationTrue}
}

// SnapshotConfiguration enumerates the values for snapshot configuration.
type SnapshotConfiguration string

const (
	// SnapshotConfigurationFalse ...
	SnapshotConfigurationFalse SnapshotConfiguration = "false"
	// SnapshotConfigurationTrue ...
	SnapshotConfigurationTrue SnapshotConfiguration = "true"
)

// PossibleSnapshotConfigurationValues returns an array of possible values for the SnapshotConfiguration const type.
func PossibleSnapshotConfigurationValues() []SnapshotConfiguration {
	return []SnapshotConfiguration{SnapshotConfigurationFalse, SnapshotConfigurationTrue}
}

// SortedBy enumerates the values for sorted by.
type SortedBy string

const (
	// NasVolumeInstanceNo ...
	NasVolumeInstanceNo SortedBy = "nasVolumeInstanceNo"
	// VolumeName ...
	VolumeName SortedBy = "volumeName"
	// VolumeTotalSize ...
	VolumeTotalSize SortedBy = "volumeTotalSize"
)

// PossibleSortedByValues returns an array of possible values for the SortedBy const type.
func PossibleSortedByValues() []SortedBy {
	return []SortedBy{NasVolumeInstanceNo, VolumeName, VolumeTotalSize}
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

// VolumeAllotmentProtocolTypeCode enumerates the values for volume allotment protocol type code.
type VolumeAllotmentProtocolTypeCode string

const (
	// CIFS ...
	CIFS VolumeAllotmentProtocolTypeCode = "CIFS"
	// NFS ...
	NFS VolumeAllotmentProtocolTypeCode = "NFS"
)

// PossibleVolumeAllotmentProtocolTypeCodeValues returns an array of possible values for the VolumeAllotmentProtocolTypeCode const type.
func PossibleVolumeAllotmentProtocolTypeCodeValues() []VolumeAllotmentProtocolTypeCode {
	return []VolumeAllotmentProtocolTypeCode{CIFS, NFS}
}

// VolumeAccessControlResponse ...
type VolumeAccessControlResponse struct {
	autorest.Response `json:"-"`
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
	autorest.Response `json:"-"`
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
	autorest.Response `json:"-"`
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

// VolumeInstancesResponse ...
type VolumeInstancesResponse struct {
	autorest.Response `json:"-"`
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
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows             *int32                `json:"totalRows,omitempty"`
	NasVolumeInstanceList *[]VolumeInstanceList `json:"nasVolumeInstanceList,omitempty"`
}
