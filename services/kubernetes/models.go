package kubernetes

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/kubernetes"

// AutoscaleOptionParameter ...
type AutoscaleOptionParameter struct {
	// Enabled - 오토스케일 가능여부
	Enabled *bool `json:"enabled,omitempty"`
	// Max - 오토스케일 가능 최대 노드 수
	Max *float64 `json:"max,omitempty"`
	// Min - 오토스케일 가능 최소 노드 수
	Min *float64 `json:"min,omitempty"`
}

// AutoscalerUpdateParameter ...
type AutoscalerUpdateParameter struct {
	// Enabled - 오토스케일 가능여부
	Enabled *bool `json:"enabled,omitempty"`
	// Max - 오토스케일 가능 최대 노드 수
	Max *float64 `json:"max,omitempty"`
	// Min - 오토스케일 가능 최소 노드 수
	Min *float64 `json:"min,omitempty"`
}

// ClusterParamter ...
type ClusterParamter struct {
	// UUID - 클러스터 UUID
	UUID *string `json:"uuid,omitempty"`
	// AcgName - 클러스터 ACG 이름
	AcgName *string `json:"acgName,omitempty"`
	// Name - 클러스터 이름
	Name *string `json:"name,omitempty"`
	// Capacity - 클러스터 용량
	Capacity *string `json:"capacity,omitempty"`
	// ClusterType - 클러스터 타입
	ClusterType *string `json:"clusterType,omitempty"`
	// NodeCount - 등록된 노드 총 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// NodeMaxCount - 사용할 수 있는 노드의 최대 개수
	NodeMaxCount *float64 `json:"nodeMaxCount,omitempty"`
	// CPUCount - cpu 개수
	CPUCount *float64 `json:"cpuCount,omitempty"`
	// MemorySize - 메모리 용량
	MemorySize *float64 `json:"memorySize,omitempty"`
	// CreatedAt - 생성 일자
	CreatedAt *string `json:"createdAt,omitempty"`
	// Endpoint - Control Plane API 주소
	Endpoint *string `json:"endpoint,omitempty"`
	// K8sVersion - 쿠버네티스 버전
	K8sVersion *string `json:"k8sVersion,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// Status - 클러스터의 상태
	Status *string `json:"status,omitempty"`
	// SubnetLbName - 로드밸런서 전용 서브넷 이름
	SubnetLbName *string `json:"subnetLbName,omitempty"`
	// SubnetLbNo - 로드밸런서 전용 서브넷 번호
	SubnetLbNo *float64 `json:"subnetLbNo,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// SubnetNoList - 서브넷 번호 목록
	SubnetNoList *[]string `json:"subnetNoList,omitempty"`
	// UpdatedAt - 최근 업데이트 일자
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// VpcName - VCP 이름
	VpcName *string `json:"vpcName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *float64 `json:"vpcNo,omitempty"`
	// ZoneCode - 금융존 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// LoginKeyName - 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// NodePool - 노드풀
	NodePool *[]NodePoolResponse `json:"nodePool,omitempty"`
}

// ClusterRequest ...
type ClusterRequest struct {
	// Name - 클러스터 이름
	Name *string `json:"name,omitempty"`
	// ClusterType - 클러스터 타입
	ClusterType *string `json:"clusterType,omitempty"`
	// LoginKeyName - 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// K8sVersion - 쿠버네티스 버전
	K8sVersion *string `json:"k8sVersion,omitempty"`
	// RegionCode - region의 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *float64 `json:"vpcNo,omitempty"`
	// SubnetNoList - 서브넷 번호 목록
	SubnetNoList *[]float64 `json:"subnetNoList,omitempty"`
	// SubnetLbNo - 로드밸런서 전용 서브넷 번호
	SubnetLbNo *float64 `json:"subnetLbNo,omitempty"`
	// DefaultNodePool - 기본 노드풀
	DefaultNodePool *DefaultNodePoolParameter `json:"defaultNodePool,omitempty"`
	// NodePool - 추가 노드풀
	NodePool *[]NodePoolParameter `json:"nodePool,omitempty"`
}

// ClusterResponse ...
type ClusterResponse struct {
	autorest.Response `json:"-"`
	// Cluster - 클러스터
	Cluster *ClusterParamter `json:"cluster,omitempty"`
}

// ClustersListResponse ...
type ClustersListResponse struct {
	autorest.Response `json:"-"`
	// Clusters - 클러스터 목록
	Clusters *[]ClusterParamter `json:"clusters,omitempty"`
}

// ConfigResponse ...
type ConfigResponse struct {
	autorest.Response `json:"-"`
	// Kubeconfig - 쿠버네티스 설정 정보
	Kubeconfig *string `json:"kubeconfig,omitempty"`
}

// DefaultNodePoolParameter ...
type DefaultNodePoolParameter struct {
	// NodeCount - 노드 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// ProductCode - 상품 코드
	ProductCode *string `json:"productCode,omitempty"`
}

// NodePoolParameter ...
type NodePoolParameter struct {
	// Name - 노드풀 이름
	Name *string `json:"name,omitempty"`
	// NodeCount - 노드 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// ProductCode - 상품 코드
	ProductCode *string `json:"productCode,omitempty"`
}

// NodePoolRequest ...
type NodePoolRequest struct {
	// Name - 노드풀 이름
	Name *string `json:"name,omitempty"`
	// NodeCount - 등록 될 노드 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// ProductCode - 상품 코드
	ProductCode *string `json:"productCode,omitempty"`
}

// NodePoolResponse ...
type NodePoolResponse struct {
	autorest.Response `json:"-"`
	// InstanceNo - 인스턴스 번호
	InstanceNo *float64 `json:"instanceNo,omitempty"`
	// IsDefault - 기본 POOL 여부
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name - 노드풀 이름
	Name *string `json:"name,omitempty"`
	// NodeCount - 노드 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// ProductCode - 상품 코드
	ProductCode *string `json:"productCode,omitempty"`
	// Status - 노드풀 상태
	Status *string `json:"status,omitempty"`
	// Autoscale - 자동 확장설정
	Autoscale *AutoscaleOptionParameter `json:"autoscale,omitempty"`
}

// NodePoolUpdateRequest ...
type NodePoolUpdateRequest struct {
	// NodeCount - 노드 개수
	NodeCount *float64 `json:"nodeCount,omitempty"`
	// Autoscale - 오토스케일
	Autoscale *AutoscalerUpdateParameter `json:"autoscale,omitempty"`
}

// WorkerNodeParameter ...
type WorkerNodeParameter struct {
	// ID - 워커노드 ID
	ID *string `json:"id,omitempty"`
	// Name - 워커노드 이름
	Name *string `json:"name,omitempty"`
	// ServerName - 워커노드 서버이름
	ServerName *string `json:"serverName,omitempty"`
	// ServerSpec - 워커노드 서버 스펙
	ServerSpec *string `json:"serverSpec,omitempty"`
	// PrivateIP - 비공인 IP
	PrivateIP *string `json:"privateIp,omitempty"`
	// PublicIP - 공인 IP
	PublicIP *string `json:"publicIp,omitempty"`
	// ReturnProtectionYn - 반납 보호 설정
	ReturnProtectionYn *string `json:"returnProtectionYn,omitempty"`
	// Status - 워커노드 현재 상태
	Status *string `json:"status,omitempty"`
	// StatusCode - 상태 코드
	StatusCode *string `json:"statusCode,omitempty"`
	// StatusIcon - 상태 아이콘
	StatusIcon *string `json:"statusIcon,omitempty"`
	// StatusColor - 상태 색깔
	StatusColor *string `json:"statusColor,omitempty"`
	// StatusName - 상태 이름
	StatusName *string `json:"statusName,omitempty"`
	// ServerImageName - 서버 이미지 이름
	ServerImageName *string `json:"serverImageName,omitempty"`
	// CPUCount - CPU 개수
	CPUCount *float64 `json:"cpuCount,omitempty"`
	// MemorySize - 총 메모리 용량
	MemorySize *float64 `json:"memorySize,omitempty"`
	// SpecCode - 스펙 코드
	SpecCode *string `json:"specCode,omitempty"`
	// LoginKeyName - 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// K8sStatus - 쿠버네티스 상태
	K8sStatus *string `json:"k8sStatus,omitempty"`
	// DockerVersion - 도커 버전
	DockerVersion *string `json:"dockerVersion,omitempty"`
	// KernelVersion - 커널 버전
	KernelVersion *string `json:"kernelVersion,omitempty"`
	// NodePoolName - 노드풀 이름
	NodePoolName *string `json:"nodePoolName,omitempty"`
}

// WorkerNodeResponse ...
type WorkerNodeResponse struct {
	autorest.Response `json:"-"`
	// Nodes - 워커노드 목록
	Nodes *[]WorkerNodeParameter `json:"nodes,omitempty"`
}
