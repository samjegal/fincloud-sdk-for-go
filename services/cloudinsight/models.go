package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/cloudinsight"

// AsgPolicyParameter auto Scaling Group Policy
type AsgPolicyParameter struct {
	// AutoScalingGroupNo - Auto Scaling Group id
	AutoScalingGroupNo *int64 `json:"autoScalingGroupNo,omitempty"`
	// AutoScalingPolicyNo - Auto Scaling Policy id
	AutoScalingPolicyNo *int64 `json:"autoScalingPolicyNo,omitempty"`
	// PolicyName - Auto Scaling Policy name
	PolicyName *string `json:"policyName,omitempty"`
}

// ChartDataWidgetRequest metric 정보로 Preview Chart 조회
type ChartDataWidgetRequest struct {
	// MetricsInfo - 데이터 조회를 위한 Metric 정보
	MetricsInfo *[]WidgetMetricInfoParameter `json:"metricsInfo,omitempty"`
	// PeriodEnd - 조회 종료 시간, default: 현재시간
	PeriodEnd *int64 `json:"periodEnd,omitempty"`
	// PeriodStart - 조회 시작 시간, default: 현재시간으로부터 1시간 이전
	PeriodStart *int64 `json:"periodStart,omitempty"`
	// Type - 위젯 타입. Possible values include: 'Line', 'Area', 'Index', 'Text', 'Table'
	Type ChartDataWidgetType `json:"type,omitempty"`
}

// CollectorRequest ...
type CollectorRequest struct {
	// CwKey - Schema 생성시 발급받은 키
	CwKey *string `json:"cw_key,omitempty"`
	// Data - 검색할 subnet 페이지 번호
	Data interface{} `json:"data,omitempty"`
}

// CollectorResponse ...
type CollectorResponse struct {
	autorest.Response `json:"-"`
	// Result - 결과 정보
	Result *string `json:"result,omitempty"`
	// DataID - 데이터 아이디
	DataID *float64 `json:"dataId,omitempty"`
	// Message - 결과 메시지
	Message *string `json:"message,omitempty"`
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
	Port *string `json:"port,omitempty"`
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
	AutoScalingGroupNo *string `json:"autoScalingGroupNo,omitempty"`
}

// EventRuleResponse event rule search response, used by NAP to check whether the given Notification GroupNo
// has been used by rules
type EventRuleResponse struct {
	autorest.Response `json:"-"`
	GrpNoHasNotiRule  *bool `json:"grpNoHasNotiRule,omitempty"`
	// RuleDtos - Events rule list
	RuleDtos *[]RuleGroup `json:"ruleDtos,omitempty"`
}

// EventSearchDimensionParameter ...
type EventSearchDimensionParameter struct {
	// InstanceNo - 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// Type - 인스턴스 타입
	Type *string `json:"type,omitempty"`
}

// EventSearchRequest event 검색 시 요청
type EventSearchRequest struct {
	// RuleID - 조회하려는 Event Rule id
	RuleID *float64 `json:"ruleId,omitempty"`
	// EventID - 조회하려는 EVent id
	EventID *string `json:"eventId,omitempty"`
	// Query - 이벤트 쿼리
	Query *string `json:"query,omitempty"`
	// StartTime - 조회 시작 시간
	StartTime *float64 `json:"startTime,omitempty"`
	// EndTime - 조회 종료 시간
	EndTime *float64 `json:"endTime,omitempty"`
	// PageSize - 페이지 크기
	PageSize *int32 `json:"pageSize,omitempty"`
	// PageNum - 페이지 번호
	PageNum               *int32 `json:"pageNum,omitempty"`
	OnlyFetchUnCloseEvent *bool  `json:"onlyFetchUnCloseEvent,omitempty"`
}

// EventSearchResultParameter ...
type EventSearchResultParameter struct {
	// EventID - 이벤트 아이디
	EventID *string `json:"eventId,omitempty"`
	// DetectValue - 검출값
	DetectValue *int32 `json:"detectValue,omitempty"`
	// ProdKey - 상품 키
	ProdKey *string `json:"prodKey,omitempty"`
	// Criteria - 영역
	Criteria *float64 `json:"criteria,omitempty"`
	// Cycles - 사이클
	Cycles *int32 `json:"cycles,omitempty"`
	// Operator - 운영
	Operator *string `json:"operator,omitempty"`
	// NotificationGroups - 알람 그룹
	NotificationGroups *string `json:"notificationGroups,omitempty"`
	// EventLevel - 이벤트 레벨
	EventLevel *string `json:"eventLevel,omitempty"`
	// Metric - 메트릭 종류
	Metric *string `json:"metric,omitempty"`
	// EventCause - 이벤트 원인
	EventCause *string `json:"eventCause,omitempty"`
	// RuleName - 룰 이름
	RuleName *string `json:"ruleName,omitempty"`
	// StartTime - 시작 시간
	StartTime *float64 `json:"startTime,omitempty"`
	// Interval - 인터벌 타임
	Interval *string `json:"interval,omitempty"`
	// Describe - 설명
	Describe *string `json:"describe,omitempty"`
	// EndTime - 종료 시간
	EndTime *int32 `json:"endTime,omitempty"`
	// Calc - 계산값
	Calc *string `json:"calc,omitempty"`
	// RuleID - 룰 아이디
	RuleID     *float64                       `json:"ruleId,omitempty"`
	Dimensions *EventSearchDimensionParameter `json:"dimensions,omitempty"`
}

// FilePluginDetailResponse ...
type FilePluginDetailResponse struct {
	autorest.Response `json:"-"`
	// ConfigList - 파일 리스트
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// FilePluginParameter ...
type FilePluginParameter struct {
	// ConfigList - 파일 리스트
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// FilePluginRequest ...
type FilePluginRequest struct {
	// ConfigList - 파일 리스트
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// Int64 ...
type Int64 struct {
	autorest.Response `json:"-"`
	Value             *int64 `json:"value,omitempty"`
}

// ListEventSearchResultParameter ...
type ListEventSearchResultParameter struct {
	autorest.Response `json:"-"`
	Value             *[]EventSearchResultParameter `json:"value,omitempty"`
}

// ListFilePluginParameter ...
type ListFilePluginParameter struct {
	autorest.Response `json:"-"`
	Value             *[]FilePluginParameter `json:"value,omitempty"`
}

// ListListFloat64 ...
type ListListFloat64 struct {
	autorest.Response `json:"-"`
	Value             *[][]float64 `json:"value,omitempty"`
}

// ListMultipleDataParameter ...
type ListMultipleDataParameter struct {
	autorest.Response `json:"-"`
	Value             *[]MultipleDataParameter `json:"value,omitempty"`
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

// ListSchemaExtendedStatusParameter ...
type ListSchemaExtendedStatusParameter struct {
	autorest.Response `json:"-"`
	Value             *[]SchemaExtendedStatusParameter `json:"value,omitempty"`
}

// ListServerTopMetricParameter ...
type ListServerTopMetricParameter struct {
	autorest.Response `json:"-"`
	Value             *[]ServerTopMetricParameter `json:"value,omitempty"`
}

// ListString ...
type ListString struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// MetricInfoParameter 데이터 조회 요청
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

// MonitorGroupItemParameter 감시 대상
type MonitorGroupItemParameter struct {
	// Nrn - NRN이란 네이버 클라우드 플랫폼의 모든 리소스를 일정한 규칙으로 표현하기 위한 정보이며 상품별로 구성요소가 다를 수 있습니다.
	Nrn *string `json:"nrn,omitempty"`
	// ResourceID - 감시 대상 Id
	ResourceID *string `json:"resourceId,omitempty"`
}

// MonitorGroupRequest 감시 대상 그룹 생성/수정 시 필요
type MonitorGroupRequest struct {
	// GroupDesc - 감시 대상 그룹 설명
	GroupDesc *string `json:"groupDesc,omitempty"`
	// GroupName - 감시 대상 그룹 이름
	GroupName *string `json:"groupName,omitempty"`
	// ID - 감시 대상 그룹 Id
	ID *string `json:"id,omitempty"`
	// MonitorGroupItemList - 감시 대상 지정
	MonitorGroupItemList *[]MonitorGroupItemParameter `json:"monitorGroupItemList,omitempty"`
	// ProdKey - 상품의 cw_key
	ProdKey *string `json:"prodKey,omitempty"`
	// IsTemporaryGroup - 감시 대상 그룹 생성 여부, false일 경우 감시 대상 그룹 생성 없이 Event Rule을 생성
	IsTemporaryGroup *bool `json:"isTemporaryGroup,omitempty"`
}

// MultipleDataParameter ...
type MultipleDataParameter struct {
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

// PortPluginDetailResponse ...
type PortPluginDetailResponse struct {
	autorest.Response `json:"-"`
	// ConfigList - 포트 번호 리스트
	ConfigList *[]int32 `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// PortPluginParameter ...
type PortPluginParameter struct {
	// ConfigList - 포트 번호 리스트
	ConfigList *[]int32 `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// PortPluginRequest ...
type PortPluginRequest struct {
	// ConfigList - 포트 번호 리스트
	ConfigList *[]int32 `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// ProcessPluginDetailResponse ...
type ProcessPluginDetailResponse struct {
	autorest.Response `json:"-"`
	// ConfigList - 프로세스 이름
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// ProcessPluginParameter ...
type ProcessPluginParameter struct {
	// ConfigList - 프로세스 이름
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// ProcessPluginRequest ...
type ProcessPluginRequest struct {
	// ConfigList - 프로세스 이름
	ConfigList *[]string `json:"configList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
}

// QueryMultipleRequest cloud Insight에서 수집한 여러개의 time-series 데이터를 쿼리합니다.
type QueryMultipleRequest struct {
	// MetricInfoList - 요청할 Metric 정보
	MetricInfoList *[]MetricInfoParameter `json:"metricInfoList,omitempty"`
	// TimeStart - 최초 조회 시간
	TimeStart *float64 `json:"timeStart,omitempty"`
	// TimeEnd - 마지막 조회 시간
	TimeEnd *float64 `json:"timeEnd,omitempty"`
}

// QueryRequest ...
type QueryRequest struct {
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

// RecipientNotificationParameter recipient Notification
type RecipientNotificationParameter struct {
	// GroupName - Recipient Notification의 이름
	GroupName *string `json:"groupName,omitempty"`
	// GroupNum - Recipient Notification id
	GroupNum *int32 `json:"groupNum,omitempty"`
	// NotifyTypes - Notify Type
	NotifyTypes *[]RecipientNotificationNotifyType `json:"notifyTypes,omitempty"`
}

// RuleGroup ...
type RuleGroup struct {
	ActionSubAccountID *int32                          `json:"actionSubAccountId,omitempty"`
	ActionUserType     *string                         `json:"actionUserType,omitempty"`
	AsgPolicyMap       map[string][]AsgPolicyParameter `json:"asgPolicyMap"`
	CfTriggers         *[]string                       `json:"cfTriggers,omitempty"`
	CreateTime         *int64                          `json:"createTime,omitempty"`
	DomainCode         *string                         `json:"domainCode,omitempty"`
	EventTime          *int64                          `json:"eventTime,omitempty"`
	GroupDesc          *string                         `json:"groupDesc,omitempty"`
	GroupName          *string                         `json:"groupName,omitempty"`
	ID                 *int64                          `json:"id,omitempty"`
	MbrNo              *string                         `json:"mbrNo,omitempty"`
	ProdKey            *string                         `json:"prodKey,omitempty"`
	// ProductDataStatus - Possible values include: 'CREATED', 'UPDATED', 'DELETED'
	ProductDataStatus      RuleGroupProductDataStatus        `json:"productDataStatus,omitempty"`
	RecipientNotifications *[]RecipientNotificationParameter `json:"recipientNotifications,omitempty"`
	RegionCode             *string                           `json:"regionCode,omitempty"`
	RelatedNrn             *[]string                         `json:"relatedNrn,omitempty"`
	// RuleVersion - Possible values include: 'V1', 'V2'
	RuleVersion RuleGroupVersion `json:"ruleVersion,omitempty"`
	SourceIP    *string          `json:"sourceIp,omitempty"`
	SourceType  *string          `json:"sourceType,omitempty"`
	// Status - Possible values include: 'OK', 'VIOLATED', 'INSUFFICIENT'
	Status           RuleGroupStatus             `json:"status,omitempty"`
	SuspendRuleItems *[]SuspendRuleItemParameter `json:"suspendRuleItems,omitempty"`
	UpdateTime       *int64                      `json:"updateTime,omitempty"`
}

// MarshalJSON is the custom marshaler for RuleGroup.
func (rg RuleGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if rg.ActionSubAccountID != nil {
		objectMap["actionSubAccountId"] = rg.ActionSubAccountID
	}
	if rg.ActionUserType != nil {
		objectMap["actionUserType"] = rg.ActionUserType
	}
	if rg.AsgPolicyMap != nil {
		objectMap["asgPolicyMap"] = rg.AsgPolicyMap
	}
	if rg.CfTriggers != nil {
		objectMap["cfTriggers"] = rg.CfTriggers
	}
	if rg.CreateTime != nil {
		objectMap["createTime"] = rg.CreateTime
	}
	if rg.DomainCode != nil {
		objectMap["domainCode"] = rg.DomainCode
	}
	if rg.EventTime != nil {
		objectMap["eventTime"] = rg.EventTime
	}
	if rg.GroupDesc != nil {
		objectMap["groupDesc"] = rg.GroupDesc
	}
	if rg.GroupName != nil {
		objectMap["groupName"] = rg.GroupName
	}
	if rg.ID != nil {
		objectMap["id"] = rg.ID
	}
	if rg.MbrNo != nil {
		objectMap["mbrNo"] = rg.MbrNo
	}
	if rg.ProdKey != nil {
		objectMap["prodKey"] = rg.ProdKey
	}
	if rg.ProductDataStatus != "" {
		objectMap["productDataStatus"] = rg.ProductDataStatus
	}
	if rg.RecipientNotifications != nil {
		objectMap["recipientNotifications"] = rg.RecipientNotifications
	}
	if rg.RegionCode != nil {
		objectMap["regionCode"] = rg.RegionCode
	}
	if rg.RelatedNrn != nil {
		objectMap["relatedNrn"] = rg.RelatedNrn
	}
	if rg.RuleVersion != "" {
		objectMap["ruleVersion"] = rg.RuleVersion
	}
	if rg.SourceIP != nil {
		objectMap["sourceIp"] = rg.SourceIP
	}
	if rg.SourceType != nil {
		objectMap["sourceType"] = rg.SourceType
	}
	if rg.Status != "" {
		objectMap["status"] = rg.Status
	}
	if rg.SuspendRuleItems != nil {
		objectMap["suspendRuleItems"] = rg.SuspendRuleItems
	}
	if rg.UpdateTime != nil {
		objectMap["updateTime"] = rg.UpdateTime
	}
	return json.Marshal(objectMap)
}

// RuleGroupRequest event Rule 생성/수정 시 필요
type RuleGroupRequest struct {
	// AsgPolicyList - only
	AsgPolicyList *[]AsgPolicyParameter `json:"asgPolicyList,omitempty"`
	// CfTriggers - 추후 제공될 예정입니다. 현재는 미사용 상태
	CfTriggers *[]string `json:"cfTriggers,omitempty"`
	// GroupDesc - Event Rule 설명
	GroupDesc *string `json:"groupDesc,omitempty"`
	// GroupName - Event Rule 이름
	GroupName *string `json:"groupName,omitempty"`
	// ID - Event Rule id
	ID *string `json:"id,omitempty"`
	// MetricsGroupKey - 감시 항목 그룹 id, 여러개 입력 가능
	MetricsGroupKey *[]string `json:"metricsGroupKey,omitempty"`
	// MonitorGroupKey - 감시 대상 그룹 id, 여러개 입력 가능
	MonitorGroupKey *[]string `json:"monitorGroupKey,omitempty"`
	// ProdKey - 상품의 cw_key
	ProdKey *string `json:"prodKey,omitempty"`
	// RecipientNotifications - 통보 대상 그룹, 여러개 입력 가능
	RecipientNotifications *[]RecipientNotificationParameter `json:"recipientNotifications,omitempty"`
	// SuspendRuleItems - Event Rule 중 비활성화 시킬 목록 설정
	SuspendRuleItems *[]SuspendRuleItemParameter `json:"suspendRuleItems,omitempty"`
}

// ScehmaUpdateResponse ...
type ScehmaUpdateResponse struct {
	autorest.Response `json:"-"`
	// Msg - 응답 메시지
	Msg *string `json:"msg,omitempty"`
}

// SchemaExtendedDisableResponse ...
type SchemaExtendedDisableResponse struct {
	autorest.Response `json:"-"`
	// Status - 응답 상태
	Status *string `json:"status,omitempty"`
}

// SchemaExtendedEnableResponse ...
type SchemaExtendedEnableResponse struct {
	autorest.Response `json:"-"`
	// Status - 응답 상태
	Status *string `json:"status,omitempty"`
}

// SchemaExtendedStatusParameter ...
type SchemaExtendedStatusParameter struct {
	// Enabled - Extended 설정 여부
	Enabled *string `json:"enabled,omitempty"`
	// InstanceID - 인스턴스 아이디
	InstanceID *string `json:"instanceId,omitempty"`
}

// SchemaFieldParameter 사용자 정의 스키마 필드 파라미터
type SchemaFieldParameter struct {
	// Aggregations - Aggregations to be performed on this metric, default to ALL aggregations if null. The values will be ignored for dimensions
	Aggregations map[string][]SchemaFieldaAggregation `json:"aggregations"`
	// Counter - Describe whether the field is of COUNTER type, default to false if missing
	Counter *bool `json:"counter,omitempty"`
	// DataType - Data type of the field, only STRING, LONG and FLOAT are available. Possible values include: 'STRING', 'INTEGER', 'FLOAT', 'LONG', 'TIMESTAMP', 'ARRAY'
	DataType SchemaFieldDataType `json:"dataType,omitempty"`
	// DefaultMetric - Describe whether the metric is default ,only use if metric = true, to show in dashboard summary api, default to false if missing
	DefaultMetric *bool `json:"defaultMetric,omitempty"`
	// Desc - Description of this field
	Desc *string `json:"desc,omitempty"`
	// Dimension - Describe whether the field is part of dimensions, default to false if missing
	Dimension *bool `json:"dimension,omitempty"`
	// DimensionType - Dimension type of this field. Possible values include: 'DEFAULT', 'CUSTOM'
	DimensionType SchemaFieldDimensionType `json:"dimensionType,omitempty"`
	IDDimension   *bool                    `json:"idDimension,omitempty"`
	// Metric - Describe whether the field is part of metrics, default to false if missing
	Metric *bool `json:"metric,omitempty"`
	// MetricType - Metric type of this field. Possible values include: 'SchemaFieldMetricTypeBASIC', 'SchemaFieldMetricTypeEXTENDED', 'SchemaFieldMetricTypeCUSTOM'
	MetricType SchemaFieldMetricType `json:"metricType,omitempty"`
	// Name - Name of the field
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for SchemaFieldParameter.
func (sfp SchemaFieldParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sfp.Aggregations != nil {
		objectMap["aggregations"] = sfp.Aggregations
	}
	if sfp.Counter != nil {
		objectMap["counter"] = sfp.Counter
	}
	if sfp.DataType != "" {
		objectMap["dataType"] = sfp.DataType
	}
	if sfp.DefaultMetric != nil {
		objectMap["defaultMetric"] = sfp.DefaultMetric
	}
	if sfp.Desc != nil {
		objectMap["desc"] = sfp.Desc
	}
	if sfp.Dimension != nil {
		objectMap["dimension"] = sfp.Dimension
	}
	if sfp.DimensionType != "" {
		objectMap["dimensionType"] = sfp.DimensionType
	}
	if sfp.IDDimension != nil {
		objectMap["idDimension"] = sfp.IDDimension
	}
	if sfp.Metric != nil {
		objectMap["metric"] = sfp.Metric
	}
	if sfp.MetricType != "" {
		objectMap["metricType"] = sfp.MetricType
	}
	if sfp.Name != nil {
		objectMap["name"] = sfp.Name
	}
	return json.Marshal(objectMap)
}

// SchemaFieldRequestParameter 사용자 정의 스키마 필드 요청 파라미터
type SchemaFieldRequestParameter struct {
	// Aggregations - Aggregations to be performed on this metric, default to ALL aggregations if null. The values will be ignored for dimensions
	Aggregations map[string][]SchemaFieldaAggregation `json:"aggregations"`
	// Counter - Describe whether the field is of COUNTER type, default to false if missing
	Counter *bool `json:"counter,omitempty"`
	// DataType - Data type of the field, only STRING, LONG and FLOAT are available. Possible values include: 'STRING', 'INTEGER', 'FLOAT', 'LONG', 'TIMESTAMP', 'ARRAY'
	DataType SchemaFieldDataType `json:"dataType,omitempty"`
	// DefaultMetric - Describe whether the metric is default ,only use if metric = true, to show in dashboard summary api, default to false if missing
	DefaultMetric *bool `json:"defaultMetric,omitempty"`
	// Desc - Description of this field
	Desc *string `json:"desc,omitempty"`
	// Dimension - Describe whether the field is part of dimensions, default to false if missing
	Dimension   *bool `json:"dimension,omitempty"`
	IDDimension *bool `json:"idDimension,omitempty"`
	// Metric - Describe whether the field is part of metrics, default to false if missing
	Metric *bool `json:"metric,omitempty"`
	// MetricType - Metric type of this field, default to BASIC if not set. Possible values include: 'SchemaFieldMetricTypeBASIC', 'SchemaFieldMetricTypeEXTENDED', 'SchemaFieldMetricTypeCUSTOM'
	MetricType SchemaFieldMetricType `json:"metricType,omitempty"`
	// Name - Name of the field
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for SchemaFieldRequestParameter.
func (sfrp SchemaFieldRequestParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sfrp.Aggregations != nil {
		objectMap["aggregations"] = sfrp.Aggregations
	}
	if sfrp.Counter != nil {
		objectMap["counter"] = sfrp.Counter
	}
	if sfrp.DataType != "" {
		objectMap["dataType"] = sfrp.DataType
	}
	if sfrp.DefaultMetric != nil {
		objectMap["defaultMetric"] = sfrp.DefaultMetric
	}
	if sfrp.Desc != nil {
		objectMap["desc"] = sfrp.Desc
	}
	if sfrp.Dimension != nil {
		objectMap["dimension"] = sfrp.Dimension
	}
	if sfrp.IDDimension != nil {
		objectMap["idDimension"] = sfrp.IDDimension
	}
	if sfrp.Metric != nil {
		objectMap["metric"] = sfrp.Metric
	}
	if sfrp.MetricType != "" {
		objectMap["metricType"] = sfrp.MetricType
	}
	if sfrp.Name != nil {
		objectMap["name"] = sfrp.Name
	}
	return json.Marshal(objectMap)
}

// SchemaRegisterResponse ...
type SchemaRegisterResponse struct {
	autorest.Response `json:"-"`
	// CwKey - Generated key for this product, not required when registering product
	CwKey *string `json:"cw_key,omitempty"`
}

// SchemaRequest ...
type SchemaRequest struct {
	// CwKey - Generated key for this product, not required when registering product
	CwKey *string `json:"cw_key,omitempty"`
	// Secret - Secret of the product, for response display ONLY, not required in request
	Secret *string `json:"secret,omitempty"`
	// ProdName - Metric Product 이름
	ProdName *string `json:"prodName,omitempty"`
	// Fields - Metric 필드 정의
	Fields *[]SchemaFieldRequestParameter `json:"fields,omitempty"`
}

// SchemaResponse 사용자 정의 스키마
type SchemaResponse struct {
	autorest.Response `json:"-"`
	// ProdName - Name of the product. Must NOT be null or empty
	ProdName *string `json:"prodName,omitempty"`
	// CwKey - Generated key for this product, not required when registering product
	CwKey *string `json:"cw_key,omitempty"`
	// Fields - Field list of the schema. Must NOT be null or empty
	Fields *[]SchemaFieldParameter `json:"fields,omitempty"`
	// Secret - Secret of the product, for response display ONLY, not required in request
	Secret *string `json:"secret,omitempty"`
}

// SearchEventCountConsoleRequest event rule information
type SearchEventCountConsoleRequest struct {
	// StartTime - 조회 시작 시간
	StartTime *int64 `json:"startTime,omitempty"`
	// EndTime - 조회 종료 시간
	EndTime *int64 `json:"endTime,omitempty"`
}

// SearchEventCountConsoleResponse console useage, Event search response
type SearchEventCountConsoleResponse struct {
	autorest.Response `json:"-"`
	// ClosedRecords - Page size of the request
	ClosedRecords *int32 `json:"closedRecords,omitempty"`
	// OpenRecords - Page size of the request
	OpenRecords *int32 `json:"openRecords,omitempty"`
	// TotalRecords - OpenEvent Count
	TotalRecords *int32 `json:"totalRecords,omitempty"`
}

// ServerTopMetricParameter ...
type ServerTopMetricParameter struct {
	// HostName - 호스트 이름
	HostName *string `json:"hostName,omitempty"`
	// InstanceNo - 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// AvgCPUUsedRto - CPU 평균 사용량
	AvgCPUUsedRto *string `json:"avg_cpu_used_rto,omitempty"`
	// AvgFsUsert - 파일 시스템 사용량
	AvgFsUsert *string `json:"avg_fs_usert,omitempty"`
	// MemUsert - 메모리 사용량
	MemUsert *string `json:"mem_usert,omitempty"`
}

// String ...
type String struct {
	autorest.Response `json:"-"`
	Value             *string `json:"value,omitempty"`
}

// SuspendRuleItemParameter ...
type SuspendRuleItemParameter struct {
	// MetricGroupItemID - 감시 항목 id
	MetricGroupItemID *string `json:"metricGroupItemId,omitempty"`
	// ResourceID - 감시 대상 id
	ResourceID *string `json:"resourceId,omitempty"`
}

// WidgetMetricInfoParameter widget 내 Metric 정보
type WidgetMetricInfoParameter struct {
	// Color - Widget에서 보여지는 Chart의 color
	Color *string `json:"color,omitempty"`
	// Data - 조회된 결과
	Data *[]interface{} `json:"data,omitempty"`
	// Dimensions - 조회하려는 데이터의 Dimension 이름
	Dimensions map[string]*string `json:"dimensions"`
	// DisplayName - Widget에서 보여지는 Metric의 이름
	DisplayName *string `json:"displayName,omitempty"`
	ID          *string `json:"id,omitempty"`
	// Metric - 조회하려는 Metric 이름
	Metric  *string                                `json:"metric,omitempty"`
	Options map[string][]ChartDataWidgetInfoOption `json:"options"`
	// Period - 조회하려는 집계 주기. Possible values include: 'Min1', 'Min5', 'Min30', 'Hour2', 'Day1'
	Period ChartDataWidgetInfoPeriod `json:"period,omitempty"`
	// ProdKey - 상품의 cw_key
	ProdKey *string `json:"prodKey,omitempty"`
	// ProdName - 상품의 이름
	ProdName *string `json:"prodName,omitempty"`
	// Statistic - 조회하려는 집계 함수. Possible values include: 'ChartDataWidgetInfoStatisticCOUNT', 'ChartDataWidgetInfoStatisticSUM', 'ChartDataWidgetInfoStatisticMAX', 'ChartDataWidgetInfoStatisticMIN', 'ChartDataWidgetInfoStatisticAVG', 'ChartDataWidgetInfoStatisticCOUNTER'
	Statistic ChartDataWidgetInfoStatistic `json:"statistic,omitempty"`
}

// MarshalJSON is the custom marshaler for WidgetMetricInfoParameter.
func (wmip WidgetMetricInfoParameter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if wmip.Color != nil {
		objectMap["color"] = wmip.Color
	}
	if wmip.Data != nil {
		objectMap["data"] = wmip.Data
	}
	if wmip.Dimensions != nil {
		objectMap["dimensions"] = wmip.Dimensions
	}
	if wmip.DisplayName != nil {
		objectMap["displayName"] = wmip.DisplayName
	}
	if wmip.ID != nil {
		objectMap["id"] = wmip.ID
	}
	if wmip.Metric != nil {
		objectMap["metric"] = wmip.Metric
	}
	if wmip.Options != nil {
		objectMap["options"] = wmip.Options
	}
	if wmip.Period != "" {
		objectMap["period"] = wmip.Period
	}
	if wmip.ProdKey != nil {
		objectMap["prodKey"] = wmip.ProdKey
	}
	if wmip.ProdName != nil {
		objectMap["prodName"] = wmip.ProdName
	}
	if wmip.Statistic != "" {
		objectMap["statistic"] = wmip.Statistic
	}
	return json.Marshal(objectMap)
}

// WidgetMetricInfoResponse ...
type WidgetMetricInfoResponse struct {
	autorest.Response `json:"-"`
	// Color - Widget에서 보여지는 Chart의 color
	Color *string `json:"color,omitempty"`
	// Data - 조회된 결과
	Data *[]interface{} `json:"data,omitempty"`
	// Dimensions - 조회하려는 데이터의 Dimension 이름
	Dimensions map[string]*string `json:"dimensions"`
	// DisplayName - Widget에서 보여지는 Metric의 이름
	DisplayName *string `json:"displayName,omitempty"`
	// Metric - 조회하려는 Metric 이름
	Metric *string `json:"metric,omitempty"`
	// Period - 조회하려는 집계 주기. Possible values include: 'WidgetMetricInfoPeriodMin1', 'WidgetMetricInfoPeriodMin5', 'WidgetMetricInfoPeriodMin30', 'WidgetMetricInfoPeriodHour2', 'WidgetMetricInfoPeriodDay1'
	Period WidgetMetricInfoPeriod `json:"period,omitempty"`
	// ProdKey - 상품의 cw_key
	ProdKey *string `json:"prodKey,omitempty"`
	// ProdName - 상품의 이름
	ProdName *string `json:"prodName,omitempty"`
	// QueryAggregation - 메트릭 쿼리 집계. Possible values include: 'WidgetMetricInfoQueryAggregationAVG', 'WidgetMetricInfoQueryAggregationCOUNT', 'WidgetMetricInfoQueryAggregationFIRST', 'WidgetMetricInfoQueryAggregationLAST', 'WidgetMetricInfoQueryAggregationMAX', 'WidgetMetricInfoQueryAggregationMIMMAX', 'WidgetMetricInfoQueryAggregationMIMMIN', 'WidgetMetricInfoQueryAggregationMIN', 'WidgetMetricInfoQueryAggregationMULT', 'WidgetMetricInfoQueryAggregationNONE', 'WidgetMetricInfoQueryAggregationSUM', 'WidgetMetricInfoQueryAggregationZIMSUM'
	QueryAggregation WidgetMetricInfoQueryAggregation `json:"queryAggregation,omitempty"`
	// Statistic - 메트릭 통계. Possible values include: 'WidgetMetricInfoStatisticCOUNT', 'WidgetMetricInfoStatisticSUM', 'WidgetMetricInfoStatisticMAX', 'WidgetMetricInfoStatisticMIN', 'WidgetMetricInfoStatisticAVG', 'WidgetMetricInfoStatisticCOUNTER'
	Statistic WidgetMetricInfoStatistic `json:"statistic,omitempty"`
}

// MarshalJSON is the custom marshaler for WidgetMetricInfoResponse.
func (wmir WidgetMetricInfoResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if wmir.Color != nil {
		objectMap["color"] = wmir.Color
	}
	if wmir.Data != nil {
		objectMap["data"] = wmir.Data
	}
	if wmir.Dimensions != nil {
		objectMap["dimensions"] = wmir.Dimensions
	}
	if wmir.DisplayName != nil {
		objectMap["displayName"] = wmir.DisplayName
	}
	if wmir.Metric != nil {
		objectMap["metric"] = wmir.Metric
	}
	if wmir.Period != "" {
		objectMap["period"] = wmir.Period
	}
	if wmir.ProdKey != nil {
		objectMap["prodKey"] = wmir.ProdKey
	}
	if wmir.ProdName != nil {
		objectMap["prodName"] = wmir.ProdName
	}
	if wmir.QueryAggregation != "" {
		objectMap["queryAggregation"] = wmir.QueryAggregation
	}
	if wmir.Statistic != "" {
		objectMap["statistic"] = wmir.Statistic
	}
	return json.Marshal(objectMap)
}
