package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/sens"

// AlimTalkMessageRequestParameter ...
type AlimTalkMessageRequestParameter struct {
	TemplateCode *string                                        `json:"templateCode,omitempty"`
	PlusFriendID *string                                        `json:"plusFriendId,omitempty"`
	Messages     *[]AlimTalkMessageRequestParameterMessagesItem `json:"messages,omitempty"`
	// ScheduleCode - 등록하려는 스케줄 코드
	ScheduleCode *string `json:"scheduleCode,omitempty"`
	// ReserveTime - 예약시간 (yyyy-MM-dd HH:mm)
	ReserveTime *string `json:"reserveTime,omitempty"`
	// ReserveTimeZone - 예약시간 타임존 (Area/Name. IANA time zone database) *default: Asia/Seoul
	ReserveTimeZone *string `json:"reserveTimeZone,omitempty"`
}

// AlimTalkMessageRequestParameterMessagesItem ...
type AlimTalkMessageRequestParameterMessagesItem struct {
	CountryCode *string                                                       `json:"countryCode,omitempty"`
	To          *string                                                       `json:"to,omitempty"`
	Content     *string                                                       `json:"content,omitempty"`
	Buttons     *[]AlimTalkMessageRequestParameterMessagesPropertiesItemsItem `json:"buttons,omitempty"`
}

// AlimTalkMessageRequestParameterMessagesPropertiesItemsItem ...
type AlimTalkMessageRequestParameterMessagesPropertiesItemsItem struct {
	Type          *string `json:"type,omitempty"`
	Name          *string `json:"name,omitempty"`
	LinkMobile    *string `json:"linkMobile,omitempty"`
	LinkPc        *string `json:"linkPc,omitempty"`
	SchemeIos     *string `json:"schemeIos,omitempty"`
	SchemeAndroid *string `json:"schemeAndroid,omitempty"`
}

// AlimTalkMessageResponseParameter ...
type AlimTalkMessageResponseParameter struct {
	autorest.Response `json:"-"`
	RequestID         *uuid.UUID                                      `json:"requestId,omitempty"`
	RequestTime       *string                                         `json:"requestTime,omitempty"`
	Messages          *[]AlimTalkMessageResponseParameterMessagesItem `json:"messages,omitempty"`
}

// AlimTalkMessageResponseParameterMessagesItem ...
type AlimTalkMessageResponseParameterMessagesItem struct {
	MessageID         *uuid.UUID `json:"messageId,omitempty"`
	RequestStatusCode *string    `json:"requestStatusCode,omitempty"`
	RequestStatusName *string    `json:"requestStatusName,omitempty"`
	RequestStatusDesc *string    `json:"requestStatusDesc,omitempty"`
}

// AlimTalkRequestResponseParameter ...
type AlimTalkRequestResponseParameter struct {
	autorest.Response `json:"-"`
	RequestID         *uuid.UUID                                      `json:"requestId,omitempty"`
	StatusCode        *string                                         `json:"statusCode,omitempty"`
	StatusName        *string                                         `json:"statusName,omitempty"`
	Messages          *[]AlimTalkRequestResponseParameterMessagesItem `json:"messages,omitempty"`
}

// AlimTalkRequestResponseParameterMessagesItem ...
type AlimTalkRequestResponseParameterMessagesItem struct {
	RequestTime       *string    `json:"requestTime,omitempty"`
	MessageID         *uuid.UUID `json:"messageId,omitempty"`
	CountryCode       *uuid.UUID `json:"countryCode,omitempty"`
	To                *uuid.UUID `json:"to,omitempty"`
	PlusFriendID      *uuid.UUID `json:"plusFriendId,omitempty"`
	TemplateCode      *uuid.UUID `json:"templateCode,omitempty"`
	RequestStatusCode *string    `json:"requestStatusCode,omitempty"`
	RequestStatusName *string    `json:"requestStatusName,omitempty"`
	RequestStatusDesc *string    `json:"requestStatusDesc,omitempty"`
	MessageStatusCode *string    `json:"messageStatusCode,omitempty"`
	MessageStatusName *string    `json:"messageStatusName,omitempty"`
	MessageStatusDesc *string    `json:"messageStatusDesc,omitempty"`
}

// AlimTalkResultResponseParameter ...
type AlimTalkResultResponseParameter struct {
	autorest.Response `json:"-"`
	MessageID         *uuid.UUID `json:"messageId,omitempty"`
	RequestID         *uuid.UUID `json:"requestId,omitempty"`
	RequestTime       *string    `json:"requestTime,omitempty"`
	CompleteTime      *string    `json:"completeTime,omitempty"`
	PlusFriendID      *string    `json:"plusFriendId,omitempty"`
	TemplateCode      *string    `json:"templateCode,omitempty"`
	CountryCode       *string    `json:"countryCode,omitempty"`
	To                *string    `json:"to,omitempty"`
	Content           *string    `json:"content,omitempty"`
	RequestStatusCode *string    `json:"requestStatusCode,omitempty"`
	RequestStatusName *string    `json:"requestStatusName,omitempty"`
	RequestStatusDesc *string    `json:"requestStatusDesc,omitempty"`
	MessageStatusCode *string    `json:"messageStatusCode,omitempty"`
	MessageStatusName *string    `json:"messageStatusName,omitempty"`
	MessageStatusDesc *string    `json:"messageStatusDesc,omitempty"`
}

// PushChannelRequestParameter ...
type PushChannelRequestParameter struct {
	ChannelName *string `json:"channelName,omitempty"`
	ChannelDesc *string `json:"channelDesc,omitempty"`
}

// PushMessageRequestParameter ...
type PushMessageRequestParameter struct {
	Message *PushMessageRequestParameterMessage `json:"message,omitempty"`
	// MessageType - (NOTIF|AD)
	MessageType *string                            `json:"messageType,omitempty"`
	Target      *PushMessageRequestParameterTarget `json:"target,omitempty"`
	// ScheduleCode - 등록하려는 스케줄 코드
	ScheduleCode *string `json:"scheduleCode,omitempty"`
	// ReserveTime - 예약시간 (yyyy-MM-dd HH:mm)
	ReserveTime *string `json:"reserveTime,omitempty"`
	// ReserveTimeZone - 예약시간 타임존 (Area/Name. IANA time zone database) *default: Asia/Seoul
	ReserveTimeZone *string `json:"reserveTimeZone,omitempty"`
}

// PushMessageRequestParameterMessage ...
type PushMessageRequestParameterMessage struct {
	Apns    *PushMessageRequestParameterMessageApns    `json:"apns,omitempty"`
	Gcm     *PushMessageRequestParameterMessageGcm     `json:"gcm,omitempty"`
	Default *PushMessageRequestParameterMessageDefault `json:"default,omitempty"`
}

// PushMessageRequestParameterMessageApns ...
type PushMessageRequestParameterMessageApns struct {
	Content *string                                     `json:"content,omitempty"`
	Custom  interface{}                                 `json:"custom,omitempty"`
	Option  interface{}                                 `json:"option,omitempty"`
	I18n    *PushMessageRequestParameterMessageApnsI18n `json:"i18n,omitempty"`
}

// PushMessageRequestParameterMessageApnsI18n ...
type PushMessageRequestParameterMessageApnsI18n struct {
	Default *PushMessageRequestParameterMessageApnsI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageRequestParameterMessageApnsI18nlanguage `json:"[language],omitempty"`
}

// PushMessageRequestParameterMessageApnsI18nDefault ...
type PushMessageRequestParameterMessageApnsI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterMessageApnsI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageRequestParameterMessageApnsI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterMessageDefault ...
type PushMessageRequestParameterMessageDefault struct {
	Content *string                                        `json:"content,omitempty"`
	Custom  interface{}                                    `json:"custom,omitempty"`
	Option  interface{}                                    `json:"option,omitempty"`
	I18n    *PushMessageRequestParameterMessageDefaultI18n `json:"i18n,omitempty"`
}

// PushMessageRequestParameterMessageDefaultI18n ...
type PushMessageRequestParameterMessageDefaultI18n struct {
	Default *PushMessageRequestParameterMessageDefaultI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageRequestParameterMessageDefaultI18nlanguage `json:"[language],omitempty"`
}

// PushMessageRequestParameterMessageDefaultI18nDefault ...
type PushMessageRequestParameterMessageDefaultI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterMessageDefaultI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageRequestParameterMessageDefaultI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterMessageGcm ...
type PushMessageRequestParameterMessageGcm struct {
	Content *string                                    `json:"content,omitempty"`
	Custom  interface{}                                `json:"custom,omitempty"`
	Option  interface{}                                `json:"option,omitempty"`
	I18n    *PushMessageRequestParameterMessageGcmI18n `json:"i18n,omitempty"`
}

// PushMessageRequestParameterMessageGcmI18n ...
type PushMessageRequestParameterMessageGcmI18n struct {
	Default *PushMessageRequestParameterMessageGcmI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageRequestParameterMessageGcmI18nlanguage `json:"[language],omitempty"`
}

// PushMessageRequestParameterMessageGcmI18nDefault ...
type PushMessageRequestParameterMessageGcmI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterMessageGcmI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageRequestParameterMessageGcmI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageRequestParameterTarget ...
type PushMessageRequestParameterTarget struct {
	// DeviceType - (APNS|GCM)
	DeviceType *string   `json:"deviceType,omitempty"`
	To         *[]string `json:"to,omitempty"`
	// Type - (all|channel|user)
	Type *string `json:"type,omitempty"`
	// Country - 국가코드 (ISO 3166-1 alpha-2, 대문자)
	Country *[]string `json:"country,omitempty"`
}

// PushMessageResponseParameter ...
type PushMessageResponseParameter struct {
	autorest.Response `json:"-"`
	RequestID         *uuid.UUID `json:"requestId,omitempty"`
	RequestTime       *date.Time `json:"requestTime,omitempty"`
	// StatusCode - 202 - 요청성공<br>400 - bad request<br>401 - unauthorized<br>404 - not found<br>500 - internal server error
	StatusCode *string `json:"statusCode,omitempty"`
	// StatusName - (success|fail)
	StatusName    *string `json:"statusName,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
	ErrorMessage  *string `json:"errorMessage,omitempty"`
}

// PushMessageResultResponseParameter ...
type PushMessageResultResponseParameter struct {
	autorest.Response `json:"-"`
	RequestID         *uuid.UUID `json:"requestId,omitempty"`
	RequestTime       *date.Time `json:"requestTime,omitempty"`
	// StatusCode - 202 - 요청성공<br>400 - bad request<br>401 - unauthorized<br>404 - not found<br>500 - internal server error
	StatusCode *string `json:"statusCode,omitempty"`
	// StatusName - (success|fail)
	StatusName    *string `json:"statusName,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
	// MessageStatusCode - 200 - 발송성공<br>400 - bad request<br>404 - no targets<br>500 - internal server error
	MessageStatusCode *string `json:"messageStatusCode,omitempty"`
	// MessageStatusName - (success|processing|fail)
	MessageStatusName    *string                                    `json:"messageStatusName,omitempty"`
	MessageStatusMessage *string                                    `json:"messageStatusMessage,omitempty"`
	ErrorMessage         *string                                    `json:"errorMessage,omitempty"`
	CompleteTime         *date.Time                                 `json:"completeTime,omitempty"`
	TargetCount          *int32                                     `json:"targetCount,omitempty"`
	SentCount            *int32                                     `json:"sentCount,omitempty"`
	Message              *PushMessageResultResponseParameterMessage `json:"message,omitempty"`
	// MessageType - (NOTIF|AD)
	MessageType *string                                   `json:"messageType,omitempty"`
	Target      *PushMessageResultResponseParameterTarget `json:"target,omitempty"`
	// ScheduleCode - 스케줄 코드
	ScheduleCode *string `json:"scheduleCode,omitempty"`
}

// PushMessageResultResponseParameterMessage ...
type PushMessageResultResponseParameterMessage struct {
	Apns    *PushMessageResultResponseParameterMessageApns    `json:"apns,omitempty"`
	Gcm     *PushMessageResultResponseParameterMessageGcm     `json:"gcm,omitempty"`
	Default *PushMessageResultResponseParameterMessageDefault `json:"default,omitempty"`
}

// PushMessageResultResponseParameterMessageApns ...
type PushMessageResultResponseParameterMessageApns struct {
	Content *string                                            `json:"content,omitempty"`
	Custom  interface{}                                        `json:"custom,omitempty"`
	Option  interface{}                                        `json:"option,omitempty"`
	I18n    *PushMessageResultResponseParameterMessageApnsI18n `json:"i18n,omitempty"`
}

// PushMessageResultResponseParameterMessageApnsI18n ...
type PushMessageResultResponseParameterMessageApnsI18n struct {
	Default *PushMessageResultResponseParameterMessageApnsI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageResultResponseParameterMessageApnsI18nlanguage `json:"[language],omitempty"`
}

// PushMessageResultResponseParameterMessageApnsI18nDefault ...
type PushMessageResultResponseParameterMessageApnsI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterMessageApnsI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageResultResponseParameterMessageApnsI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterMessageDefault ...
type PushMessageResultResponseParameterMessageDefault struct {
	Content *string                                               `json:"content,omitempty"`
	Custom  interface{}                                           `json:"custom,omitempty"`
	Option  interface{}                                           `json:"option,omitempty"`
	I18n    *PushMessageResultResponseParameterMessageDefaultI18n `json:"i18n,omitempty"`
}

// PushMessageResultResponseParameterMessageDefaultI18n ...
type PushMessageResultResponseParameterMessageDefaultI18n struct {
	Default *PushMessageResultResponseParameterMessageDefaultI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageResultResponseParameterMessageDefaultI18nlanguage `json:"[language],omitempty"`
}

// PushMessageResultResponseParameterMessageDefaultI18nDefault ...
type PushMessageResultResponseParameterMessageDefaultI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterMessageDefaultI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageResultResponseParameterMessageDefaultI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterMessageGcm ...
type PushMessageResultResponseParameterMessageGcm struct {
	Content *string                                           `json:"content,omitempty"`
	Custom  interface{}                                       `json:"custom,omitempty"`
	Option  interface{}                                       `json:"option,omitempty"`
	I18n    *PushMessageResultResponseParameterMessageGcmI18n `json:"i18n,omitempty"`
}

// PushMessageResultResponseParameterMessageGcmI18n ...
type PushMessageResultResponseParameterMessageGcmI18n struct {
	Default *PushMessageResultResponseParameterMessageGcmI18nDefault `json:"default,omitempty"`
	// language - [language]: ISO 639-1, 소문자
	language *PushMessageResultResponseParameterMessageGcmI18nlanguage `json:"[language],omitempty"`
}

// PushMessageResultResponseParameterMessageGcmI18nDefault ...
type PushMessageResultResponseParameterMessageGcmI18nDefault struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterMessageGcmI18nlanguage [language]: ISO 639-1, 소문자
type PushMessageResultResponseParameterMessageGcmI18nlanguage struct {
	Content *string `json:"content,omitempty"`
}

// PushMessageResultResponseParameterTarget ...
type PushMessageResultResponseParameterTarget struct {
	// DeviceType - (APNS|GCM)
	DeviceType *string   `json:"deviceType,omitempty"`
	To         *[]string `json:"to,omitempty"`
	// Type - (all|channel|user)
	Type *string `json:"type,omitempty"`
	// Country - 국가코드 (ISO 3166-1 alpha-2, 대문자)
	Country *[]string `json:"country,omitempty"`
	// Timezone - Number of seconds away from UTC
	Timezone *[]int32 `json:"timezone,omitempty"`
}

// PushScheduleFetchAllParameter ...
type PushScheduleFetchAllParameter struct {
	autorest.Response `json:"-"`
	Content           *[]PushScheduleMessageFetchParameter `json:"content,omitempty"`
	TotalElements     *int32                               `json:"totalElements,omitempty"`
	TotalPages        *int32                               `json:"totalPages,omitempty"`
	First             *bool                                `json:"first,omitempty"`
	Last              *bool                                `json:"last,omitempty"`
	NumberOfElements  *int32                               `json:"numberOfElements,omitempty"`
	Size              *int32                               `json:"size,omitempty"`
	Number            *int32                               `json:"number,omitempty"`
}

// PushScheduleFetchParameter ...
type PushScheduleFetchParameter struct {
	autorest.Response `json:"-"`
	// ScheduleID - 스케줄 아이디
	ScheduleID *string `json:"scheduleId,omitempty"`
	// ScheduleCode - 스케줄 코드 (알파벳/숫자/-/_)
	ScheduleCode *string `json:"scheduleCode,omitempty"`
	// ScheduleDesc - 스케줄 설명
	ScheduleDesc *string `json:"scheduleDesc,omitempty"`
	// ScheduleTimeZone - 스케줄 타임존. (Area/Name. IANA time zone database) *default: Asia/Seoul, 현지시간기준: LOCALTIME
	ScheduleTimeZone *string `json:"scheduleTimeZone,omitempty"`
	// ScheduleTime - 스케줄 시/분 (HH:mm:ss) *초단위는 무시
	ScheduleTime *string `json:"scheduleTime,omitempty"`
	// StartDate - 시작일 (yyyy-MM-dd) *default: 오늘
	StartDate *string `json:"startDate,omitempty"`
	// EndDate - 종료일 (yyyy-MM-dd) *오늘부터 최대6개월
	EndDate *string `json:"endDate,omitempty"`
	// DayOfWeeks - ALL(매일), MON, TUE, WED, ..., SUN
	DayOfWeeks *[]string `json:"dayOfWeeks,omitempty"`
	// ScheduleStatus - 스케줄 상태. (ACTIVE, INACTIVE)
	ScheduleStatus *string `json:"scheduleStatus,omitempty"`
}

// PushScheduleMessageFetchParameter ...
type PushScheduleMessageFetchParameter struct {
	// MessageID - 메시지 아이디
	MessageID *string `json:"messageId,omitempty"`
	// ScheduleID - 스케줄 아이디
	ScheduleID *string `json:"scheduleId,omitempty"`
	// ServiceID - 서비스 아이디
	ServiceID *string `json:"serviceId,omitempty"`
	// Message - 메시지 요청 JSON string
	Message *string `json:"message,omitempty"`
}

// PushScheduleRegisterParameter ...
type PushScheduleRegisterParameter struct {
	// ScheduleCode - 스케줄 코드. (알파벳/숫자/-/_)
	ScheduleCode *string `json:"scheduleCode,omitempty"`
	// ScheduleDesc - 스케줄 설명
	ScheduleDesc *string `json:"scheduleDesc,omitempty"`
	// ScheduleTimeZone - 스케줄 타임존. (Area/Name. IANA time zone database) *default: Asia/Seoul, 현지시간기준: LOCALTIME
	ScheduleTimeZone *string `json:"scheduleTimeZone,omitempty"`
	// ScheduleTime - 스케줄 시/분 (HH:mm:ss) *초단위는 무시
	ScheduleTime *string `json:"scheduleTime,omitempty"`
	// StartDate - 시작일 (yyyy-MM-dd) *default: 오늘
	StartDate *string `json:"startDate,omitempty"`
	// EndDate - 종료일 (yyyy-MM-dd) *오늘부터 최대6개월
	EndDate *string `json:"endDate,omitempty"`
	// DayOfWeeks - ALL(매일), MON, TUE, WED, ..., SUN
	DayOfWeeks *[]string `json:"dayOfWeeks,omitempty"`
}

// PushUserRequestParameter ...
type PushUserRequestParameter struct {
	UserID *string `json:"userId,omitempty"`
	// DeviceType - (APNS|GCM)
	DeviceType              *string `json:"deviceType,omitempty"`
	DeviceToken             *string `json:"deviceToken,omitempty"`
	ChannelName             *string `json:"channelName,omitempty"`
	IsNotificationAgreement *bool   `json:"isNotificationAgreement,omitempty"`
	IsAdAgreement           *bool   `json:"isAdAgreement,omitempty"`
	IsNightAdAgreement      *bool   `json:"isNightAdAgreement,omitempty"`
	// Country - 국가코드 (ISO 3166-1 alpha-2, 대문자) *default: KR
	Country *string `json:"country,omitempty"`
	// Language - 언어 (ISO 639-1, 소문자) *default: ko
	Language *string `json:"language,omitempty"`
	// Timezone - Number of seconds away from UTC. Example: -28800 *default: 32400
	Timezone *int32 `json:"timezone,omitempty"`
}

// PushUserResponseParameter ...
type PushUserResponseParameter struct {
	autorest.Response         `json:"-"`
	UserID                    *string    `json:"userId,omitempty"`
	ChannelName               *string    `json:"channelName,omitempty"`
	IsNotificationAgreement   *bool      `json:"isNotificationAgreement,omitempty"`
	IsAdAgreement             *bool      `json:"isAdAgreement,omitempty"`
	NotificationAgreementTime *date.Time `json:"notificationAgreementTime,omitempty"`
	AdAgreementTime           *date.Time `json:"adAgreementTime,omitempty"`
	NightAdAgreementTime      *date.Time `json:"nightAdAgreementTime,omitempty"`
	// Country - 국가코드 (ISO 3166-1 alpha-2, 대문자)
	Country *string `json:"country,omitempty"`
	// Language - 언어 (ISO 639-1, 소문자)
	Language *string `json:"language,omitempty"`
	// Timezone - Number of seconds away from UTC
	Timezone *int32                                  `json:"timezone,omitempty"`
	Devices  *[]PushUserResponseParameterDevicesItem `json:"devices,omitempty"`
}

// PushUserResponseParameterDevicesItem ...
type PushUserResponseParameterDevicesItem struct {
	// DeviceType - (APNS|GCM)
	DeviceType  *string `json:"deviceType,omitempty"`
	DeviceToken *string `json:"deviceToken,omitempty"`
}

// SMSMessage ...
type SMSMessage struct {
	// MessageID - 메시지 아이디
	MessageID *string `json:"messageId,omitempty"`
	// From - 발신번호
	From *string `json:"from,omitempty"`
	// To - 수신번호
	To *string `json:"to,omitempty"`
	// ContentType - 메시지 컨텐츠 구분(COMM: 일반미시지, AD: 광고메시지)
	ContentType *string `json:"contentType,omitempty"`
	// CountryCode - 국가번호
	CountryCode *string `json:"countryCode,omitempty"`
	// RequestID - 요청 아이디
	RequestID *string `json:"requestId,omitempty"`
	// StatusCode - 단말 수신 상태 결과 코드
	StatusCode *string `json:"statusCode,omitempty"`
	// StatusName - 단말 수신 상태 결과명
	StatusName *string `json:"statusName,omitempty"`
	// StatusMessage - 단말 수신 상태 결과 메시지
	StatusMessage *string `json:"statusMessage,omitempty"`
	// Status - 발송 처리 상태(READY: 대기, PROCESSING: 처리중, COMPLETED: 처리완료)
	Status *string `json:"status,omitempty"`
	// CompleteTime - 발송 완료 시간
	CompleteTime *string `json:"completeTime,omitempty"`
	// Content - 메시지 컨텐츠 내용
	Content *string `json:"content,omitempty"`
	// TelcoCode - 통신사코드
	TelcoCode *string `json:"telcoCode,omitempty"`
}

// SMSMessageParameter ...
type SMSMessageParameter struct {
	autorest.Response `json:"-"`
	// RequestID - 요청 아이디
	RequestID *string `json:"requestId,omitempty"`
	// RequestTime - 요청 시간
	RequestTime *string `json:"requestTime,omitempty"`
	// StatusCode - 발송 요청 상태 코드
	StatusCode *string `json:"statusCode,omitempty"`
	// StatusName - 발송 요청 상태명
	StatusName *string `json:"statusName,omitempty"`
	// StatusMessage - 발송 요청 상태 메시지
	StatusMessage *string       `json:"statusMessage,omitempty"`
	Messages      *[]SMSMessage `json:"messages,omitempty"`
}

// SMSMessageRequestParameter ...
type SMSMessageRequestParameter struct {
	// Type - 메시지 구분. (SMS, LMS, MMS)
	Type *string `json:"type,omitempty"`
	// CountryCode - 국가번호
	CountryCode *string `json:"countryCode,omitempty"`
	// From - 사전 등록된 발신번호
	From *string `json:"from,omitempty"`
	// Subject - 기본 메시지 제목. (LMS, MMS에서만 사용가능)
	Subject *string `json:"subject,omitempty"`
	// ContentType - 메시지 컨텐츠 구분. (COMM: 일반미시지, AD: 광고메시지)
	ContentType *string `json:"contentType,omitempty"`
	// Content - 기본 메시지 컨텐츠 내용
	Content *string `json:"content,omitempty"`
	// ReserveTime - 예약 발송 시간
	ReserveTime *string `json:"reserveTime,omitempty"`
	// ReserveTimeZone - 예약 시간 시간대 (default: Asia/Seoul)
	ReserveTimeZone *string `json:"reserveTimeZone,omitempty"`
	// ScheduleCode - 발송할 스케줄 코드
	ScheduleCode *string                                   `json:"scheduleCode,omitempty"`
	Messages     *[]SMSMessageRequestParameterMessagesItem `json:"messages,omitempty"`
	Files        *[]SMSMessageRequestParameterFilesItem    `json:"files,omitempty"`
}

// SMSMessageRequestParameterFilesItem ...
type SMSMessageRequestParameterFilesItem struct {
	// Name - 파일 이름(jpg, jpeg 확장자를 가진 파일 이름)
	Name *string `json:"name,omitempty"`
	// Body - 파일을 Base64로 인코딩한 값
	Body *string `json:"body,omitempty"`
}

// SMSMessageRequestParameterMessagesItem ...
type SMSMessageRequestParameterMessagesItem struct {
	// Subject - 메시지 제목(정의하지 않으면 기본 메시지 제목을 사용)
	Subject *string `json:"subject,omitempty"`
	// Content - 메시지 컨텐츠 내용(정의하지 않으면 기본 메시지 컨텐츠 내용을 사용)
	Content *string `json:"content,omitempty"`
	// To - 수신번호
	To *string `json:"to,omitempty"`
}
