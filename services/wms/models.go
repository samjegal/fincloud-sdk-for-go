package wms

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/wms"

// DetailParameter ...
type DetailParameter struct {
	// ScriptID - 스크립트 ID
	ScriptID *string `json:"scriptId,omitempty"`
	// Interval - 모니터링 주기
	Interval *int32 `json:"interval,omitempty"`
	// RequestTimeout - 요청 제한 시간 (millisecond)
	RequestTimeout *int32 `json:"requestTimeout,omitempty"`
	// RunTimeout - 실행 제한 시간 (millisecond)
	RunTimeout *int32 `json:"runTimeout,omitempty"`
	// ServiceType - 서비스 유형. Possible values include: 'PC', 'MOBILE'
	ServiceType ServiceType `json:"serviceType,omitempty"`
	// StatusType - 모니터링 상태. Possible values include: 'StatusTypeON', 'StatusTypeOFF'
	StatusType StatusType `json:"statusType,omitempty"`
	// MonitoringType - 모니터링 유형. Possible values include: 'MonitoringTypeURL', 'MonitoringTypeSCENARIO'
	MonitoringType MonitoringType `json:"monitoringType,omitempty"`
	// AlarmStatus - 알람 상태. Possible values include: 'ON', 'OFF'
	AlarmStatus AlarmStatus `json:"alarmStatus,omitempty"`
	// AlarmTrigger - 알람 발송 조건
	AlarmTrigger *int32 `json:"alarmTrigger,omitempty"`
	// AlarmInterval - 알람 발송 간격 (단위:시)
	AlarmInterval *int32 `json:"alarmInterval,omitempty"`
	// URL - URL
	URL *string `json:"url,omitempty"`
	// TotalCount - 모니터링 전체 횟수
	TotalCount *int32 `json:"totalCount,omitempty"`
	// SuccessCount - 모니터링 결과가 정상인 횟수
	SuccessCount *int32 `json:"successCount,omitempty"`
	// ErrorCount - 오류가 감지된 횟수
	ErrorCount *int32 `json:"errorCount,omitempty"`
	// AvgLoadTime - URL 평균 Load time
	AvgLoadTime *int32 `json:"avgLoadTime,omitempty"`
	// RecentMonitoringTime - 최근 모니터링 시간
	RecentMonitoringTime *float64 `json:"recentMonitoringTime,omitempty"`
	// CreateDate - 생성 시간
	CreateDate *float64 `json:"createDate,omitempty"`
	// Locations - 국가 정보
	Locations *[]LocationInfoParameter `json:"locations,omitempty"`
	// Filters - 필터 정보
	Filters *[]FilterParameter `json:"filters,omitempty"`
}

// FilterParameter ...
type FilterParameter struct {
	// Type - 등록된 필터 타입. Possible values include: 'FilterTypeURL', 'FilterTypeUrlprefix', 'FilterTypeJs', 'FilterTypeJsprefix'
	Type FilterType `json:"type,omitempty"`
	// Text - 등록된 필터 값
	Text *string `json:"text,omitempty"`
}

// ListParameter ...
type ListParameter struct {
	// ScriptID - 스크립트 ID
	ScriptID *string `json:"scriptId,omitempty"`
	// Interval - 모니터링 주기
	Interval *int32 `json:"interval,omitempty"`
	// RequestTimeout - 요청 제한 시간 (millisecond)
	RequestTimeout *int32 `json:"requestTimeout,omitempty"`
	// RunTimeout - 실행 제한 시간 (millisecond)
	RunTimeout *int32 `json:"runTimeout,omitempty"`
	// ServiceType - 서비스 유형. Possible values include: 'PC', 'MOBILE'
	ServiceType ServiceType `json:"serviceType,omitempty"`
	// StatusType - 모니터링 상태. Possible values include: 'StatusTypeON', 'StatusTypeOFF'
	StatusType StatusType `json:"statusType,omitempty"`
	// MonitoringType - 모니터링 유형. Possible values include: 'MonitoringTypeURL', 'MonitoringTypeSCENARIO'
	MonitoringType MonitoringType `json:"monitoringType,omitempty"`
	// AlarmStatus - 알람 상태. Possible values include: 'ON', 'OFF'
	AlarmStatus AlarmStatus `json:"alarmStatus,omitempty"`
	// AlarmTrigger - 알람 발송 조건
	AlarmTrigger *int32 `json:"alarmTrigger,omitempty"`
	// AlarmInterval - 알람 발송 간격 (단위:시)
	AlarmInterval *int32 `json:"alarmInterval,omitempty"`
	// URL - URL
	URL *string `json:"url,omitempty"`
	// RecentMonitoringTime - 최근 모니터링 시간
	RecentMonitoringTime *float64 `json:"recentMonitoringTime,omitempty"`
	// Locations - 국가 정보
	Locations *[]LocationInfoParameter `json:"locations,omitempty"`
}

// LocationInfoParameter ...
type LocationInfoParameter struct {
	// LocationCode - 국가명
	LocationCode *string `json:"locationCode,omitempty"`
	// LocationName - 국가 코드. Possible values include: 'KR', 'US', 'JP', 'HK', 'SG'
	LocationName LocationName `json:"locationName,omitempty"`
	// LocationAddTime - 추가된 시간
	LocationAddTime *float64 `json:"locationAddTime,omitempty"`
}

// ResultContentDataParameter ...
type ResultContentDataParameter struct {
	// StepNo - 스텝 번호
	StepNo *int32 `json:"stepNo,omitempty"`
	// StepResult - 스텝 결과. Possible values include: 'StepResultSUCCESS', 'StepResultFAIL'
	StepResult StepResult `json:"stepResult,omitempty"`
	// StepName - 스텝 이름
	StepName *string `json:"stepName,omitempty"`
	// StepEvent - 이벤트 유형. Possible values include: 'Click', 'Text', 'FindText', 'FindObject'
	StepEvent StepEvent `json:"stepEvent,omitempty"`
	// StepTarget - 이벤트 대상
	StepTarget *string `json:"stepTarget,omitempty"`
	// StepText - 입력 텍스트
	StepText *string   `json:"stepText,omitempty"`
	Logs     *[]string `json:"logs,omitempty"`
}

// ResultContentLogParameter ...
type ResultContentLogParameter struct {
	// ErrorType - 오류 유형. Possible values include: 'URL', 'JAVASCRIPT'
	ErrorType ErrorType `json:"errorType,omitempty"`
	// ErrorCode - 오류 코드
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// ErrorDesc - 오류 메시지
	ErrorDesc *string `json:"errorDesc,omitempty"`
	// ErrorURL - 오류 URL
	ErrorURL *string `json:"errorUrl,omitempty"`
	// HTTPStatusCode - HTTP 상태 코드
	HTTPStatusCode *string `json:"httpStatusCode,omitempty"`
	// HTTPStatusText - HTTP 상태 메세지
	HTTPStatusText *string `json:"httpStatusText,omitempty"`
}

// ResultContentlParameter ...
type ResultContentlParameter struct {
	// MonitoringTime - 모니터링 시작 시간
	MonitoringTime *float64 `json:"monitoringTime,omitempty"`
	// RunTime - 모니터링 구동 시간
	RunTime *int32 `json:"runTime,omitempty"`
	// LoadTime - URL Load time
	LoadTime *int32 `json:"loadTime,omitempty"`
	// RequestCount - 서비스 접속 후 다운로드한 컨텐츠 수
	RequestCount *int32 `json:"requestCount,omitempty"`
	// ErrorCount - 감지된 오류수
	ErrorCount *int32 `json:"errorCount,omitempty"`
	// Result - 모니터링 결과. Possible values include: 'SUCCESS', 'FAIL'
	Result ResultEnum `json:"result,omitempty"`
	// LocationName - 국가 코드. Possible values include: 'KR', 'US', 'JP', 'HK', 'SG'
	LocationName LocationName `json:"locationName,omitempty"`
	// LocationCode - 국가명
	LocationCode *string `json:"locationCode,omitempty"`
	// Logs - 모니터링 유형이 URL인 경우
	Logs *ResultContentLogParameter `json:"logs,omitempty"`
	// Data - 모니터링 유형이 SCENARIO인 경우
	Data *ResultContentDataParameter `json:"data,omitempty"`
}

// ReturnParameter ...
type ReturnParameter struct {
	autorest.Response `json:"-"`
	// Status - 오류 코드
	Status *int32 `json:"status,omitempty"`
	// Message - 오류 메시지
	Message *string `json:"message,omitempty"`
}

// SetObject ...
type SetObject struct {
	autorest.Response `json:"-"`
	Value             interface{} `json:"value,omitempty"`
}
