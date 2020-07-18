package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/cloudinsight"

// AggregationType enumerates the values for aggregation type.
type AggregationType string

const (
	// AVG ...
	AVG AggregationType = "AVG"
	// COUNT ...
	COUNT AggregationType = "COUNT"
	// MAX ...
	MAX AggregationType = "MAX"
	// MIN ...
	MIN AggregationType = "MIN"
	// SUM ...
	SUM AggregationType = "SUM"
)

// PossibleAggregationTypeValues returns an array of possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{AVG, COUNT, MAX, MIN, SUM}
}

// DimensionType enumerates the values for dimension type.
type DimensionType string

const (
	// Agent ...
	Agent DimensionType = "agent"
	// CPU ...
	CPU DimensionType = "cpu"
	// Disk ...
	Disk DimensionType = "disk"
	// Fs ...
	Fs DimensionType = "fs"
	// Memory ...
	Memory DimensionType = "memory"
	// Ntwk ...
	Ntwk DimensionType = "ntwk"
	// PluginFile ...
	PluginFile DimensionType = "plugin_file"
	// PluginPort ...
	PluginPort DimensionType = "plugin_port"
	// PluginProcess ...
	PluginProcess DimensionType = "plugin_process"
	// Process ...
	Process DimensionType = "process"
	// Svr ...
	Svr DimensionType = "svr"
)

// PossibleDimensionTypeValues returns an array of possible values for the DimensionType const type.
func PossibleDimensionTypeValues() []DimensionType {
	return []DimensionType{Agent, CPU, Disk, Fs, Memory, Ntwk, PluginFile, PluginPort, PluginProcess, Process, Svr}
}

// QueryAggregationType enumerates the values for query aggregation type.
type QueryAggregationType string

const (
	// QueryAggregationTypeAVG ...
	QueryAggregationTypeAVG QueryAggregationType = "AVG"
	// QueryAggregationTypeCOUNT ...
	QueryAggregationTypeCOUNT QueryAggregationType = "COUNT"
	// QueryAggregationTypeFIRST ...
	QueryAggregationTypeFIRST QueryAggregationType = "FIRST"
	// QueryAggregationTypeLAST ...
	QueryAggregationTypeLAST QueryAggregationType = "LAST"
	// QueryAggregationTypeMAX ...
	QueryAggregationTypeMAX QueryAggregationType = "MAX"
	// QueryAggregationTypeMIN ...
	QueryAggregationTypeMIN QueryAggregationType = "MIN"
	// QueryAggregationTypeMULT ...
	QueryAggregationTypeMULT QueryAggregationType = "MULT"
	// QueryAggregationTypeNONE ...
	QueryAggregationTypeNONE QueryAggregationType = "NONE"
	// QueryAggregationTypeSUM ...
	QueryAggregationTypeSUM QueryAggregationType = "SUM"
)

// PossibleQueryAggregationTypeValues returns an array of possible values for the QueryAggregationType const type.
func PossibleQueryAggregationTypeValues() []QueryAggregationType {
	return []QueryAggregationType{QueryAggregationTypeAVG, QueryAggregationTypeCOUNT, QueryAggregationTypeFIRST, QueryAggregationTypeLAST, QueryAggregationTypeMAX, QueryAggregationTypeMIN, QueryAggregationTypeMULT, QueryAggregationTypeNONE, QueryAggregationTypeSUM}
}

// QueryIntervalTime enumerates the values for query interval time.
type QueryIntervalTime string

const (
	// DAY1 ...
	DAY1 QueryIntervalTime = "DAY1"
	// HOUR2 ...
	HOUR2 QueryIntervalTime = "HOUR2"
	// MIN1 ...
	MIN1 QueryIntervalTime = "MIN1"
	// MIN30 ...
	MIN30 QueryIntervalTime = "MIN30"
	// MIN5 ...
	MIN5 QueryIntervalTime = "MIN5"
)

// PossibleQueryIntervalTimeValues returns an array of possible values for the QueryIntervalTime const type.
func PossibleQueryIntervalTimeValues() []QueryIntervalTime {
	return []QueryIntervalTime{DAY1, HOUR2, MIN1, MIN30, MIN5}
}

// SchemaDataType enumerates the values for schema data type.
type SchemaDataType string

const (
	// FLOAT ...
	FLOAT SchemaDataType = "FLOAT"
	// INTEGER ...
	INTEGER SchemaDataType = "INTEGER"
	// LONG ...
	LONG SchemaDataType = "LONG"
	// STRING ...
	STRING SchemaDataType = "STRING"
)

// PossibleSchemaDataTypeValues returns an array of possible values for the SchemaDataType const type.
func PossibleSchemaDataTypeValues() []SchemaDataType {
	return []SchemaDataType{FLOAT, INTEGER, LONG, STRING}
}

// CollectorParameter ...
type CollectorParameter struct {
	// Data - 검색할 subnet 페이지 번호
	Data interface{} `json:"data,omitempty"`
	// CwKey - Schema 생성시 발급받은 키
	CwKey *string `json:"cw_key,omitempty"`
}

// DataInfoParameter ...
type DataInfoParameter struct {
	// Aggregation - 일괄 처리. Possible values include: 'MIN', 'MAX', 'SUM', 'COUNT', 'AVG'
	Aggregation AggregationType `json:"aggregation,omitempty"`
	// Dimensions - Query Dimension 데이터
	Dimensions *DimensionResultParameter `json:"dimensions,omitempty"`
	Dps        *[][]float64              `json:"dps,omitempty"`
	// Interval - 조회 시간 간격. Possible values include: 'MIN1', 'MIN5', 'MIN30', 'HOUR2', 'DAY1'
	Interval QueryIntervalTime `json:"interval,omitempty"`
	// Metric - Metric 이름
	Metric *string `json:"metric,omitempty"`
	// ProductName - Product 이름
	ProductName *string `json:"productName,omitempty"`
}

// DimensionParameter ...
type DimensionParameter struct {
	// Type - Dimension 타입. Possible values include: 'CPU', 'Disk', 'Fs', 'Memory', 'Ntwk', 'Process', 'Svr', 'PluginProcess', 'PluginFile', 'PluginPort', 'Agent'
	Type DimensionType `json:"type,omitempty"`
	// InstanceNo - 서버 인스턴스
	InstanceNo *string `json:"instanceNo,omitempty"`
	// CPUIdx - CPU 인덱스
	CPUIdx *int32 `json:"cpu_idx,omitempty"`
	// DiskIdx - 디스크 인덱스
	DiskIdx *string `json:"disk_idx,omitempty"`
	// MntNm - 파일시스템 마운트 정보
	MntNm *string `json:"mnt_nm,omitempty"`
	// NicDesc - 네트워크 인터페이스 정보
	NicDesc *string `json:"nic_desc,omitempty"`
	// ProcName - 프로세스 이름
	ProcName *string `json:"proc_name,omitempty"`
	// Path - 프로세스 경로
	Path *string `json:"path,omitempty"`
	// Port - 포트 번호
	Port *int32 `json:"port,omitempty"`
	// LoadBalancerAddress - 로드밸런서 주소
	LoadBalancerAddress *string `json:"loadBalancerAddress,omitempty"`
	// LoadBalancerPort - 로드밸런서 포트
	LoadBalancerPort *string `json:"loadBalancerPort,omitempty"`
	// InstanceID - 인스턴스 번호
	InstanceID *string `json:"instanceId,omitempty"`
	// LoadBalancerName - 로드밸런서 이름
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`
	// NetworkType - 네트워크 타입
	NetworkType *string `json:"networkType,omitempty"`
	// LayerType - 로드밸런서 레이어 타입
	LayerType *string `json:"layerType,omitempty"`
	// AutoScalingGroupNo - 오토 스케일링 그룹 번호
	AutoScalingGroupNo *int32 `json:"autoScalingGroupNo,omitempty"`
}

// DimensionResultParameter ...
type DimensionResultParameter struct {
	// Type - Dimension 타입. Possible values include: 'CPU', 'Disk', 'Fs', 'Memory', 'Ntwk', 'Process', 'Svr', 'PluginProcess', 'PluginFile', 'PluginPort', 'Agent'
	Type DimensionType `json:"type,omitempty"`
	// InstanceNo - 서버 인스턴스
	InstanceNo *string `json:"instanceNo,omitempty"`
	// CPUIdx - CPU 인덱스
	CPUIdx *string `json:"cpu_idx,omitempty"`
	// DiskIdx - 디스크 인덱스
	DiskIdx *string `json:"disk_idx,omitempty"`
	// MntNm - 파일시스템 마운트 정보
	MntNm *string `json:"mnt_nm,omitempty"`
	// NicDesc - 네트워크 인터페이스 정보
	NicDesc *string `json:"nic_desc,omitempty"`
	// ProcName - 프로세스 이름
	ProcName *string `json:"proc_name,omitempty"`
	// Path - 프로세스 경로
	Path *string `json:"path,omitempty"`
	// Port - 포트 번호
	Port *int32 `json:"port,omitempty"`
	// LoadBalancerAddress - 로드밸런서 주소
	LoadBalancerAddress *string `json:"loadBalancerAddress,omitempty"`
	// LoadBalancerPort - 로드밸런서 포트
	LoadBalancerPort *string `json:"loadBalancerPort,omitempty"`
	// InstanceID - 인스턴스 번호
	InstanceID *string `json:"instanceId,omitempty"`
	// LoadBalancerName - 로드밸런서 이름
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`
	// NetworkType - 네트워크 타입
	NetworkType *string `json:"networkType,omitempty"`
	// LayerType - 로드밸런서 레이어 타입
	LayerType *string `json:"layerType,omitempty"`
	// AutoScalingGroupNo - 오토 스케일링 그룹 번호
	AutoScalingGroupNo *int32 `json:"autoScalingGroupNo,omitempty"`
}

// ListDataInfoParameter ...
type ListDataInfoParameter struct {
	autorest.Response `json:"-"`
	Value             *[]DataInfoParameter `json:"value,omitempty"`
}

// ListListFloat64 ...
type ListListFloat64 struct {
	autorest.Response `json:"-"`
	Value             *[][]float64 `json:"value,omitempty"`
}

// ListPortPluginParameter ...
type ListPortPluginParameter struct {
	autorest.Response `json:"-"`
	Value             *[]PortPluginParameter `json:"value,omitempty"`
}

// ListProcessPluginParameter ...
type ListProcessPluginParameter struct {
	autorest.Response `json:"-"`
	Value             *[]ProcessPluginParameter `json:"value,omitempty"`
}

// MetricInfoParameter ...
type MetricInfoParameter struct {
	// Prodkey - Schema 생성시 발급받은 키
	Prodkey *string `json:"prodkey,omitempty"`
	// ProductName - Product 이름
	ProductName *string `json:"productName,omitempty"`
	// Metric - Metric 이름
	Metric *string `json:"metric,omitempty"`
	// Interval - 조회 시간 간격. Possible values include: 'MIN1', 'MIN5', 'MIN30', 'HOUR2', 'DAY1'
	Interval QueryIntervalTime `json:"interval,omitempty"`
	// Aggregation - 일괄 처리. Possible values include: 'MIN', 'MAX', 'SUM', 'COUNT', 'AVG'
	Aggregation AggregationType `json:"aggregation,omitempty"`
	// QueryAggregation - Query 일괄 처리. Possible values include: 'QueryAggregationTypeAVG', 'QueryAggregationTypeCOUNT', 'QueryAggregationTypeMIN', 'QueryAggregationTypeMAX', 'QueryAggregationTypeNONE', 'QueryAggregationTypeSUM', 'QueryAggregationTypeFIRST', 'QueryAggregationTypeLAST', 'QueryAggregationTypeMULT'
	QueryAggregation QueryAggregationType `json:"queryAggregation,omitempty"`
	// Dimensions - Query Dimension 데이터
	Dimensions *DimensionParameter `json:"dimensions,omitempty"`
}

// PortPluginParameter ...
type PortPluginParameter struct {
	autorest.Response `json:"-"`
	// ConfigList - 포트 번호 리스트
	ConfigList *[]int32 `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// ProcessPluginParameter ...
type ProcessPluginParameter struct {
	autorest.Response `json:"-"`
	// ConfigList - 프로세스 이름
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// QueryMultipleParameter ...
type QueryMultipleParameter struct {
	// MetricInfoList - 요청할 Metric 정보
	MetricInfoList *[]MetricInfoParameter `json:"metricInfoList,omitempty"`
	// TimeStart - 최초 조회 시간
	TimeStart *float64 `json:"timeStart,omitempty"`
	// TimeEnd - 마지막 조회 시간
	TimeEnd *float64 `json:"timeEnd,omitempty"`
}

// QueryParameter ...
type QueryParameter struct {
	// TimeEnd - 마지막 조회 시간
	TimeEnd *float64 `json:"timeEnd,omitempty"`
	// TimeStart - 최초 조회 시간
	TimeStart *float64 `json:"timeStart,omitempty"`
	// Metric - Metric 이름
	Metric *string `json:"metric,omitempty"`
	// Interval - 조회 시간 간격. Possible values include: 'MIN1', 'MIN5', 'MIN30', 'HOUR2', 'DAY1'
	Interval QueryIntervalTime `json:"interval,omitempty"`
	// Aggregation - 일괄 처리. Possible values include: 'MIN', 'MAX', 'SUM', 'COUNT', 'AVG'
	Aggregation AggregationType `json:"aggregation,omitempty"`
	// QueryAggregation - Query 일괄 처리. Possible values include: 'QueryAggregationTypeAVG', 'QueryAggregationTypeCOUNT', 'QueryAggregationTypeMIN', 'QueryAggregationTypeMAX', 'QueryAggregationTypeNONE', 'QueryAggregationTypeSUM', 'QueryAggregationTypeFIRST', 'QueryAggregationTypeLAST', 'QueryAggregationTypeMULT'
	QueryAggregation QueryAggregationType `json:"queryAggregation,omitempty"`
	// CwKey - Schema 생성시 발급받은 키
	CwKey *string `json:"cw_key,omitempty"`
	// ProductName - Product 이름
	ProductName *string `json:"productName,omitempty"`
	// Dimensions - Query Dimension 데이터
	Dimensions *DimensionParameter `json:"dimensions,omitempty"`
}

// SchemaFieldsParameter ...
type SchemaFieldsParameter struct {
	// Metric - Metric 유무
	Metric *bool `json:"metric,omitempty"`
	// DataType - Metric 데이터 타입. Possible values include: 'STRING', 'INTEGER', 'LONG', 'FLOAT'
	DataType SchemaDataType `json:"dataType,omitempty"`
	// Name - Metric 이름
	Name *string `json:"name,omitempty"`
	// DefaultMetric - Metric 기본 유무
	DefaultMetric *bool `json:"defaultMetric,omitempty"`
	// Dimension - Dimension 설정
	Dimension *bool `json:"dimension,omitempty"`
	// Desc - Dimension 설명
	Desc *string `json:"desc,omitempty"`
}

// SchemaParameter ...
type SchemaParameter struct {
	// ProdName - Metric Product 이름
	ProdName *string `json:"prodName,omitempty"`
	// Fields - Metric 필드 정의
	Fields *[]SchemaFieldsParameter `json:"fields,omitempty"`
}
