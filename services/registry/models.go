package registry

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/registry"

// RepositoryStatus enumerates the values for repository status.
type RepositoryStatus string

const (
	// Disconnected ...
	Disconnected RepositoryStatus = "disconnected"
	// Running ...
	Running RepositoryStatus = "running"
)

// PossibleRepositoryStatusValues returns an array of possible values for the RepositoryStatus const type.
func PossibleRepositoryStatusValues() []RepositoryStatus {
	return []RepositoryStatus{Disconnected, Running}
}

// ImageTagResultImagesItem ...
type ImageTagResultImagesItem struct {
	// Size - 이미지 사이즈
	Size *float64 `json:"size,omitempty"`
	// Architecture - CPU 아키텍쳐
	Architecture *string `json:"architecture,omitempty"`
	// Variant - 사용하지 않음
	Variant *string `json:"variant,omitempty"`
	// Features - 사용하지 않음
	Features *string `json:"features,omitempty"`
	// Os - 운영체제
	Os *string `json:"os,omitempty"`
	// OsVersion - 사용하지 않음
	OsVersion *string `json:"os_version,omitempty"`
	// OsFeatures - 사용하지 않음
	OsFeatures *string `json:"os_features,omitempty"`
	// Created - 이미지 최초 생성일
	Created *float64 `json:"created,omitempty"`
}

// ImageTagResultItem ...
type ImageTagResultItem struct {
	// Name - 이미지 태그 이름
	Name *string `json:"name,omitempty"`
	// FullSize - 이미지 태그 사이즈
	FullSize *float64                    `json:"full_size,omitempty"`
	Images   *[]ImageTagResultImagesItem `json:"images,omitempty"`
	// ID - 이미지 태그의 아이디
	ID *float64 `json:"id,omitempty"`
	// Repository - 레지스트리 아이디
	Repository *float64 `json:"repository,omitempty"`
	// Creator - 해당 태그의 이미지를 생성한 사람의 id
	Creator *string `json:"creator,omitempty"`
	// LastUpdater - 해당 태그의 이미지를 업데이트 한 사람 id
	LastUpdater *string `json:"last_updater,omitempty"`
	// LastUpdated - 마지막 업데이트 일자
	LastUpdated *float64 `json:"last_updated,omitempty"`
	// ImageID - 해당 태그의 아이디
	ImageID *float64 `json:"image_id,omitempty"`
	// V2 - 이미지에 대한 v2 스펙 적용 유무
	V2 *bool `json:"v2,omitempty"`
}

// MessageResponse ...
type MessageResponse struct {
	autorest.Response `json:"-"`
	// ReturnCode - 반환 코드
	ReturnCode *float64 `json:"returnCode,omitempty"`
	// ReturnMessage - 반환 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
}

// Permissions ...
type Permissions struct {
	Read  *bool `json:"read,omitempty"`
	Write *bool `json:"write,omitempty"`
	Admin *bool `json:"admin,omitempty"`
}

// RepositoryGetListResponse ...
type RepositoryGetListResponse struct {
	Count    *float64                `json:"count,omitempty"`
	Next     *float64                `json:"next,omitempty"`
	Previous *float64                `json:"previous,omitempty"`
	Results  *[]RepositoryResultItem `json:"results,omitempty"`
}

// RepositoryImageDetailResponse ...
type RepositoryImageDetailResponse struct {
	// User - 사용하지 않음
	User *string `json:"user,omitempty"`
	// Name - 이미지 이름
	Name *string `json:"name,omitempty"`
	// Namespace - 레지스트리 이름
	Namespace *string `json:"namespace,omitempty"`
	// RepositoryType - 사용하지 않음
	RepositoryType *string `json:"repository_type,omitempty"`
	// Status - 사용하지 않음
	Status *float64 `json:"status,omitempty"`
	// Description - 이미지에 대한short description
	Description *string `json:"description,omitempty"`
	// IsPrivate - 사용하지 않음
	IsPrivate *bool `json:"is_private,omitempty"`
	// IsAutomated - 사용하지 않음
	IsAutomated *bool `json:"is_automated,omitempty"`
	// CanEdit - 사용하지 않음
	CanEdit *bool `json:"can_edit,omitempty"`
	// StarCount - 사용하지 않음
	StarCount *float64 `json:"star_count,omitempty"`
	// PullCount - 이미지에 대한 총 Pull count 값
	PullCount *float64 `json:"pull_count,omitempty"`
	// LastUpdated - 이미지의 최근 변경일
	LastUpdated *float64 `json:"last_updated,omitempty"`
	// HasStarred - 사용하지 않음
	HasStarred *bool `json:"has_starred,omitempty"`
	// FullDescription - 이미지에 대한 상세 description
	FullDescription *string `json:"full_description,omitempty"`
	// Affiliation - 사용하지 않음
	Affiliation *string `json:"affiliation,omitempty"`
	// Permissions - 사용하지 않음
	Permissions *Permissions `json:"permissions,omitempty"`
}

// RepositoryImageItem ...
type RepositoryImageItem struct {
	// User - 사용하지 않음
	User *string `json:"user,omitempty"`
	// Name - 이미지 이름
	Name *string `json:"name,omitempty"`
	// Namespace - 레지스트리 이름
	Namespace *string `json:"namespace,omitempty"`
	// RepositoryType - 사용하지 않음
	RepositoryType *string `json:"repository_type,omitempty"`
	// Status - 사용하지 않음
	Status *float64 `json:"status,omitempty"`
	// Description - 이미지에 대한 short description
	Description *string `json:"description,omitempty"`
	// IsPrivate - 사용하지 않음
	IsPrivate *bool `json:"is_private,omitempty"`
	// IsAutomated - 사용하지 않음
	IsAutomated *bool `json:"is_automated,omitempty"`
	// CanEdit - 사용하지 않음
	CanEdit *bool `json:"can_edit,omitempty"`
	// StarCount - 사용하지 않음
	StarCount *float64 `json:"star_count,omitempty"`
	// PullCount - 이미지에 대한 총 Pull count 값
	PullCount *float64 `json:"pull_count,omitempty"`
	// LastUpdated - 이미지의 최종 변경일
	LastUpdated *float64 `json:"last_updated,omitempty"`
}

// RepositoryImageListResponse ...
type RepositoryImageListResponse struct {
	Count    *float64               `json:"count,omitempty"`
	Next     *float64               `json:"next,omitempty"`
	Previous *float64               `json:"previous,omitempty"`
	Results  *[]RepositoryImageItem `json:"results,omitempty"`
}

// RepositoryImageTagListResponse ...
type RepositoryImageTagListResponse struct {
	Count    *float64              `json:"count,omitempty"`
	Next     *float64              `json:"next,omitempty"`
	Previous *float64              `json:"previous,omitempty"`
	Results  *[]ImageTagResultItem `json:"results,omitempty"`
}

// RepositoryImageTagReferenceResponse ...
type RepositoryImageTagReferenceResponse struct {
	// Name - 이미지 태그 이름
	Name *string `json:"name,omitempty"`
	// FullSize - 이미지 태그 사이즈
	FullSize *float64                    `json:"full_size,omitempty"`
	Images   *[]ImageTagResultImagesItem `json:"images,omitempty"`
	// ID - 이미지 태그의 아이디
	ID *float64 `json:"id,omitempty"`
	// Repository - 레지스트리 아이디
	Repository *float64 `json:"repository,omitempty"`
	// Creator - 해당 태그의 이미지를 생성한 사람의 id
	Creator *string `json:"creator,omitempty"`
	// LastUpdater - 해당 태그의 이미지를 업데이트 한 사람 id
	LastUpdater *string `json:"last_updater,omitempty"`
	// LastUpdated - 마지막 업데이트 일자
	LastUpdated *float64 `json:"last_updated,omitempty"`
	// ImageID - 해당 태그의 아이디
	ImageID *float64 `json:"image_id,omitempty"`
	// V2 - 이미지에 대한 v2 스펙 적용 유무
	V2 *bool `json:"v2,omitempty"`
}

// RepositoryRequest ...
type RepositoryRequest struct {
	// Bucket - 생성될 레지스트리와 연동될 Object Storage의 Bucket 이름
	Bucket *string `json:"bucket,omitempty"`
}

// RepositoryResultItem ...
type RepositoryResultItem struct {
	// Name - Registry 이름
	Name *string `json:"name,omitempty"`
	// Bucket - 이미지가 저장될 Object Storage의 Bucket 이름
	Bucket *string `json:"bucket,omitempty"`
	// EndPoint - Registry에 접근 하기위해 docker client가 사용할 엔드포인트
	EndPoint *string `json:"end_point,omitempty"`
	// Usage - Registry가 이용한 Object Storage의 크기
	Usage *string `json:"usage,omitempty"`
	// Created - Registry가 생성된 일시의 타임스탬프
	Created *float64 `json:"created,omitempty"`
	// Status - 현재 Registry의 상태. Possible values include: 'Running', 'Disconnected'
	Status RepositoryStatus `json:"status,omitempty"`
}

// SetObject ...
type SetObject struct {
	autorest.Response `json:"-"`
	Value             interface{} `json:"value,omitempty"`
}

// UpdateRepositoryImageRequest ...
type UpdateRepositoryImageRequest struct {
	// Description - 이미지에 대한 짧은 설명 작성
	Description *string `json:"description,omitempty"`
	// FullDescription - 이미지에 대한 상세 설명 작성
	FullDescription *string `json:"full_description,omitempty"`
}
