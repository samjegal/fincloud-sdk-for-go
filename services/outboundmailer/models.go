package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/outboundmailer"

// AddressBookDeleteAddressRequest ...
type AddressBookDeleteAddressRequest struct {
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`
}

// AddressBookDeleteRelationRequest ...
type AddressBookDeleteRelationRequest struct {
	// GroupName - 수신자 그룹명
	GroupName      *string   `json:"groupName,omitempty"`
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`
}

// AddressBookGenerateRequest ...
type AddressBookGenerateRequest struct {
	// Groups - 수신자 이메일 주소를 포함한 수신자 그룹목록
	Groups *[]AddressBookGenerateRequestGroupsItem `json:"groups,omitempty"`
}

// AddressBookGenerateRequestGroupsItem ...
type AddressBookGenerateRequestGroupsItem struct {
	// GroupName - 수신자 그룹명 (값이 전달되지 않으면 이메일 주소만 등록하는 것으로 판단)
	GroupName      *string   `json:"groupName,omitempty"`
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`
}

// AddressBookInitResponse ...
type AddressBookInitResponse struct {
	autorest.Response `json:"-"`
	// DeletedAddressCount - 총 삭제 이메일 주소 개수
	DeletedAddressCount *int32 `json:"deletedAddressCount,omitempty"`
}

// AddressBookResponse ...
type AddressBookResponse struct {
	autorest.Response `json:"-"`
	// TotalAddressCount - 총 수신자 이메일 주소 수
	TotalAddressCount *int32 `json:"totalAddressCount,omitempty"`
	// Groups - 수신자 그룹 목록
	Groups *[]AddressBookResponseGroupsItem `json:"groups,omitempty"`
}

// AddressBookResponseGroupsItem ...
type AddressBookResponseGroupsItem struct {
	// Sid - 수신자 그룹 ID
	Sid *int32 `json:"sid,omitempty"`
	// GroupName - 수신자 그룹명
	GroupName *string `json:"groupName,omitempty"`
	// AddressCount - 수신자 그룹에 포함된 이메일 주소 개수
	AddressCount *int32 `json:"addressCount,omitempty"`
}

// EmailListResponse ...
type EmailListResponse struct {
	// TotalElements - 총 레코드 개수
	TotalElements *int32 `json:"totalElements,omitempty"`
	// TotalPages - 총 페이지 개수
	TotalPages *int32 `json:"totalPages,omitempty"`
	// NumberOfElements - 현재 페이지의 레코드 개수
	NumberOfElements *int32 `json:"numberOfElements,omitempty"`
	// First - 첫번째 페이지 여부
	First *bool `json:"first,omitempty"`
	// Last - 마지막 페이지 여부
	Last *bool `json:"last,omitempty"`
	// Number - 현재 페이지 index (0부터 시작)
	Number *int32 `json:"number,omitempty"`
	// Size - 한페이지에 당 레코드 개수
	Size    *int32                          `json:"size,omitempty"`
	Sort    *[]EmailListResponseSortItem    `json:"sort,omitempty"`
	Content *[]EmailListResponseContentItem `json:"content,omitempty"`
}

// EmailListResponseContentItem ...
type EmailListResponseContentItem struct {
	// RequestID - Email 발송 요청 ID (각 요청을 구분하는 ID, 한번에 여러건에 메일 발송을 요청할 경우 requestId가 여러개의 mailId를 포함할 수 있다.
	RequestID *string `json:"requestId,omitempty"`
	// RequestDate - 요청일시
	RequestDate *NesDateTime `json:"requestDate,omitempty"`
	// MailID - Email ID (각 메일 한 건을 구분하는 ID)
	MailID *string `json:"mailId,omitempty"`
	// TemplateSid - 템플릿 ID
	TemplateSid *int32 `json:"templateSid,omitempty"`
	// TemplateName - 템플릿 제목
	TemplateName *string `json:"templateName,omitempty"`
	// Title - Mail 제목
	Title *string `json:"title,omitempty"`
	// EmailStatus - Email 발송상태
	EmailStatus *EmailListResponseContentItemEmailStatus `json:"emailStatus,omitempty"`
	// SenderAddress - 발송자 Email 주소
	SenderAddress *string `json:"senderAddress,omitempty"`
	// SenderName - 발송자 이름
	SenderName *string `json:"senderName,omitempty"`
	// SendDate - 발송완료 일시
	SendDate *NesDateTime `json:"sendDate,omitempty"`
	// ReservationStatus - 예약발송 여부
	ReservationStatus *EmailListResponseContentItemReservationStatus `json:"reservationStatus,omitempty"`
	// Advertising - 광고메일 여부
	Advertising *bool `json:"advertising,omitempty"`
	// RepresentativeRecipient - 수신자 목록 축약표현 (참조/숨은참조 포함하여 총 10명에게 발송된 메일의 경우 xxx@domain.com(9) 의 형태로 표시됨)
	RepresentativeRecipient *string `json:"representativeRecipient,omitempty"`
}

// EmailListResponseContentItemEmailStatus email 발송상태
type EmailListResponseContentItemEmailStatus struct {
	// Label - 상태 명
	Label *string `json:"label,omitempty"`
	// Code - 상태 코드 (P: 발송준비중, R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
	Code *string `json:"code,omitempty"`
}

// EmailListResponseContentItemReservationStatus 예약발송 여부
type EmailListResponseContentItemReservationStatus struct {
	// Label - 예약 발송 여부 (예약발송 , 즉시발송)
	Label *string `json:"label,omitempty"`
	// Code - 상태 코드 (Y: 예약발송, N: 즉시발송)
	Code *string `json:"code,omitempty"`
}

// EmailListResponseSortItem ...
type EmailListResponseSortItem struct {
	// Direction - 정렬 방향 (ASC|DESC)
	Direction *string `json:"direction,omitempty"`
	// Property - 정렬 기준 필드명
	Property *string `json:"property,omitempty"`
	// IgnoreCase - 대소문자 구분하여 정렬할지 여부
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// NullHandling - null 처리 방식 (NATIVE: data 처리로직에 맡김 , NULLS_FIRST : null 값이 앞으로, NULLS_LAST: null 값이 뒤로)
	NullHandling *string `json:"nullHandling,omitempty"`
	// Ascending - 정렬 방향 Ascending(ASC) 인지 여부
	Ascending *bool `json:"ascending,omitempty"`
	// Descending - 정렬 방향 Descending(DESC) 인지 여부
	Descending *bool `json:"descending,omitempty"`
}

// EmailResponse ...
type EmailResponse struct {
	// RequestID - Email 발송 요청 ID (각 요청을 구분하는 ID, 한번에 여러건에 메일 발송을 요청할 경우 requestId가 여러개의 mailId를 포함할 수 있다.
	RequestID *string `json:"requestId,omitempty"`
	// RequesterIP - Email 발송 요청자 IP
	RequesterIP *string `json:"requesterIp,omitempty"`
	// RequestDate - 요청일시
	RequestDate *NesDateTime `json:"requestDate,omitempty"`
	// MailID - Email ID (각 메일 한 건을 구분하는 ID)
	MailID *string `json:"mailId,omitempty"`
	// Title - 메일 제목
	Title *string `json:"title,omitempty"`
	// Body - 메일 내용
	Body *string `json:"body,omitempty"`
	// TemplateSid - 템플릿 ID
	TemplateSid *int32 `json:"templateSid,omitempty"`
	// TemplateName - 템플릿 제목
	TemplateName *string `json:"templateName,omitempty"`
	// EmailStatus - Email 발송상태
	EmailStatus *EmailResponseEmailStatus `json:"emailStatus,omitempty"`
	// SenderAddress - 발송자 Email 주소
	SenderAddress *string `json:"senderAddress,omitempty"`
	// SenderName - 발송자 이름
	SenderName *string `json:"senderName,omitempty"`
	// SendDate - 발송 일시
	SendDate *NesDateTime `json:"sendDate,omitempty"`
	// ReservationDate - 예약 발송 일시
	ReservationDate *NesDateTime `json:"reservationDate,omitempty"`
	// Advertising - 광고메일 여부
	Advertising *bool `json:"advertising,omitempty"`
	// ReferencesHeader - References 헤더 (다음의 형태가 되어야 함 <unique_id@domain.com>), 네이버 메일에서는 References 헤더에 따라 메일을 모아 볼 수 있음. 특정 메일을 모아서 보기 위해서는 Unique 한 값이 입력 되어야만 함. 값이 중복되는 경우 같은 메일 쓰레드로 판단하여 메일을 묶어서 노출 됨.하나의 값만 입력해도 가능(References 헤더의 최상단 값만으로만 판단).
	ReferencesHeader *string `json:"referencesHeader,omitempty"`
	// AttachFiles -  첨부파일
	AttachFiles *[]EmailResponseAttachFilesItem `json:"attachFiles,omitempty"`
	// Recipients -  수신자 목록
	Recipients *[]EmailResponseRecipientsItem `json:"recipients,omitempty"`
}

// EmailResponseAttachFilesItem ...
type EmailResponseAttachFilesItem struct {
	// FileID - File ID
	FileID *string `json:"fileId,omitempty"`
	// FileName - File name
	FileName *string `json:"fileName,omitempty"`
	// FileSize - File 크기
	FileSize *int32 `json:"fileSize,omitempty"`
}

// EmailResponseEmailStatus email 발송상태
type EmailResponseEmailStatus struct {
	// Label - 상태 명
	Label *string `json:"label,omitempty"`
	// Code - 상태 코드 (R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
	Code *string `json:"code,omitempty"`
}

// EmailResponseRecipientsItem ...
type EmailResponseRecipientsItem struct {
	// Address - 수신자 Email 주소
	Address *string `json:"address,omitempty"`
	// Name - 수신자 명
	Name *string `json:"name,omitempty"`
	// Type - 수신자 유형 (R: 수신자, C: 참조인, B: 숨은참조)
	Type *EmailResponseRecipientsItemType `json:"type,omitempty"`
	// Received - 수신 여부
	Received *bool `json:"received,omitempty"`
	// ReceivedDate - 수신 일시
	ReceivedDate *NesDateTime `json:"receivedDate,omitempty"`
	// Status - Email 발송상태
	Status *EmailResponseRecipientsItemStatus `json:"status,omitempty"`
	// RetryCount - 재발송 시도 횟수
	RetryCount *int32 `json:"retryCount,omitempty"`
	// SendResultMessage - 발송 결과 상세 메시지
	SendResultMessage *string `json:"sendResultMessage,omitempty"`
}

// EmailResponseRecipientsItemStatus email 발송상태
type EmailResponseRecipientsItemStatus struct {
	// Label - 상태 명
	Label *string `json:"label,omitempty"`
	// Code - 상태 코드 (R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
	Code *string `json:"code,omitempty"`
}

// EmailResponseRecipientsItemType 수신자 유형 (R: 수신자, C: 참조인, B: 숨은참조)
type EmailResponseRecipientsItemType struct {
	// Label - 수신자 유형명 (수신자, 참조인, 숨은참조)
	Label *string `json:"label,omitempty"`
	// Code - 수신자 유형 코드 (R, C, B)
	Code *string `json:"code,omitempty"`
}

// EmailSendListResponse ...
type EmailSendListResponse struct {
	// TotalElements - 총 레코드 개수
	TotalElements *int32 `json:"totalElements,omitempty"`
	// TotalPages - 총 페이지 개수
	TotalPages *int32 `json:"totalPages,omitempty"`
	// NumberOfElements - 현재 페이지의 레코드 개수
	NumberOfElements *int32 `json:"numberOfElements,omitempty"`
	// First - 첫번째 페이지 여부
	First *bool `json:"first,omitempty"`
	// Last - 마지막 페이지 여부
	Last *bool `json:"last,omitempty"`
	// Number - 현재 페이지 index (0부터 시작)
	Number *int32 `json:"number,omitempty"`
	// Size - 한페이지에 당 레코드 개수
	Size    *int32                              `json:"size,omitempty"`
	Sort    *[]EmailSendListResponseSortItem    `json:"sort,omitempty"`
	Content *[]EmailSendListResponseContentItem `json:"content,omitempty"`
}

// EmailSendListResponseContentItem ...
type EmailSendListResponseContentItem struct {
	// RequestID - Email 발송 요청 ID (각 요청을 구분하는 ID, 한번에 여러건에 메일 발송을 요청할 경우 requestId가 여러개의 mailId를 포함할 수 있다.
	RequestID *string `json:"requestId,omitempty"`
	// RequestDate - 요청일시
	RequestDate *NesDateTime `json:"requestDate,omitempty"`
	// TemplateSid - 템플릿 ID
	TemplateSid *int32 `json:"templateSid,omitempty"`
	// TemplateName - 템플릿 제목
	TemplateName *string                                      `json:"templateName,omitempty"`
	EmailStatus  *EmailSendListResponseContentItemEmailStatus `json:"emailStatus,omitempty"`
	// SenderAddress - 발송자 Email 주소
	SenderAddress *string `json:"senderAddress,omitempty"`
	// SenderName - 발송자 이름
	SenderName *string `json:"senderName,omitempty"`
	// DispatchType - Email 발송요청 도구 (CONSOLE : NCP Console 화면을 통한 요청, API : API 호출을 통한 요청)
	DispatchType *string `json:"dispatchType,omitempty"`
	// SendDate - 발송완료 일시
	SendDate *NesDateTime `json:"sendDate,omitempty"`
	// ReservationDate - 예약 발송 일시
	ReservationDate *NesDateTime `json:"reservationDate,omitempty"`
	// RequestCount - 발송 메일 건수 (개별 발송일 경우 recipientCount 와 동일)
	RequestCount *int32 `json:"requestCount,omitempty"`
	// RecipientCount - 발송 수신자/참조/숨은참조자 수 총합(개별 발송일 경우 requestCount 와 동일하지만 일반발송의 경우 requestCount 보다 같거나 크다)
	RecipientCount *int32 `json:"recipientCount,omitempty"`
}

// EmailSendListResponseContentItemEmailStatus ...
type EmailSendListResponseContentItemEmailStatus struct {
	// Label - 상태 명
	Label *string `json:"label,omitempty"`
	// Code - 상태 코드 (P: 발송준비중, R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
	Code *string `json:"code,omitempty"`
}

// EmailSendListResponseSortItem ...
type EmailSendListResponseSortItem struct {
	// Direction - 정렬 방향 (ASC|DESC)
	Direction *string `json:"direction,omitempty"`
	// Property - 정렬 기준 필드명
	Property *string `json:"property,omitempty"`
	// IgnoreCase - 대소문자 구분하여 정렬할지 여부
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// NullHandling - null 처리 방식 (NATIVE: data 처리로직에 맡김 , NULLS_FIRST : null 값이 앞으로, NULLS_LAST: null 값이 뒤로)
	NullHandling *string `json:"nullHandling,omitempty"`
	// Ascending - 정렬 방향 Ascending(ASC) 인지 여부
	Ascending *bool `json:"ascending,omitempty"`
	// Descending - 정렬 방향 Descending(DESC) 인지 여부
	Descending *bool `json:"descending,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError ...
type ErrorResponseError struct {
	// Code - 에러 코드, ( 77001:METHOD_NOT_ALLOWED, 77002:UNSUPPORTED_MEDIA_TYPE, 77101:로그인 정보 오류, 77102:BAD_REQUEST, 77103:리소스가 존재하지 않습니다., 77201:권한없음, 77202:Email 상품 사용신청 하지 않음, 77301:기본 프로젝트가 존재하지 않음, 77302:외부 시스템 API 연동 오류, 77303:그외 INTERNAL_SERVER_ERROR )
	Code *string `json:"code,omitempty"`
	// ErrorCode - API Gateway Error Code
	ErrorCode *string `json:"errorCode,omitempty"`
	// Message - 상세 에러 메세지
	Message *string `json:"message,omitempty"`
}

// FileUploadResponse ...
type FileUploadResponse struct {
	autorest.Response `json:"-"`
	// TempRequestID - 임시 요청 ID
	TempRequestID *string `json:"tempRequestId,omitempty"`
	// Files -  File 목록
	Files *[]FileUploadResponseFilesItem `json:"files,omitempty"`
}

// FileUploadResponseFilesItem ...
type FileUploadResponseFilesItem struct {
	// FileID - File ID
	FileID *string `json:"fileId,omitempty"`
	// FileName - File name
	FileName *string `json:"fileName,omitempty"`
	// FileSize - File 크기
	FileSize *int64 `json:"fileSize,omitempty"`
}

// MailRequestParameter ...
type MailRequestParameter struct {
	// SenderAddress - 발송자 Email 주소
	SenderAddress *string `json:"senderAddress,omitempty"`
	// SenderName - 발송자 이름
	SenderName *string `json:"senderName,omitempty"`
	// TemplateSid - 템플릿 ID
	TemplateSid *int32 `json:"templateSid,omitempty"`
	// Title - 메일 제목
	Title *string `json:"title,omitempty"`
	// Body - 메일 본문
	Body *string `json:"body,omitempty"`
	// Individual - 개인발송여부
	Individual *bool `json:"individual,omitempty"`
	// ConfirmAndSend - 확인 후 발송 여부
	ConfirmAndSend *bool `json:"confirmAndSend,omitempty"`
	// Advertising - 광고메일여부
	Advertising *bool `json:"advertising,omitempty"`
	// Parameters - 치환 파라미터
	Parameters interface{} `json:"parameters,omitempty"`
	// ReferencesHeader - 특정 메일을 모아서 보기 위해 네이버 메일에서 지원하는 기능
	ReferencesHeader *string `json:"referencesHeader,omitempty"`
	// ReservationUtc - 예약 발송 일시
	ReservationUtc *float64 `json:"reservationUtc,omitempty"`
	// ReservationDateTime - 예약 발송 일시, reservationUtc 값이 우선
	ReservationDateTime *string `json:"reservationDateTime,omitempty"`
	// AttachFileIds - 첨부파일 ID 목록
	AttachFileIds *[]string `json:"attachFileIds,omitempty"`
	// Recipients - 수신자 목록
	Recipients *[]RecipientForRequest `json:"recipients,omitempty"`
	// RecipientGroupFilter - 수신자 그룹 조합 발송 조건
	RecipientGroupFilter interface{} `json:"recipientGroupFilter,omitempty"`
	// UseBasicUnsubscribeMsg - 광고 메일일 경우 기본 수신 거부 문구 사용 여부
	UseBasicUnsubscribeMsg *bool `json:"useBasicUnsubscribeMsg,omitempty"`
	// UnsubscribeMessage - 사용자 정의 수신 거부 문구
	UnsubscribeMessage *string `json:"unsubscribeMessage,omitempty"`
}

// MailResponseParameter ...
type MailResponseParameter struct {
	autorest.Response `json:"-"`
	// RequestID - 메일 발송 요청 ID
	RequestID *string `json:"requestId,omitempty"`
	// Count - 메일 요청 건수
	Count *int32 `json:"count,omitempty"`
}

// NesDateTime ...
type NesDateTime struct {
	// Utc - 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수
	Utc *int64 `json:"utc,omitempty"`
	// FormattedDate - 'yyyy-MM-dd' UTC+09:00 형태의 날짜
	FormattedDate *string `json:"formattedDate,omitempty"`
	// FormattedDateTime - 'yyyy-MM-dd HH:mm:ss SSS' UTC+09:00 형태의 시간
	FormattedDateTime *string `json:"formattedDateTime,omitempty"`
}

// RecipientForRequest 수신자 요청 정보
type RecipientForRequest struct {
	// Address - 수신자 Email 주소
	Address *string `json:"address,omitempty"`
	// Name - 수신자 이름
	Name *string `json:"name,omitempty"`
	// Type - 수신자 유형 (R: 수신자, C: 참조인, B: 숨은참조)
	Type *string `json:"type,omitempty"`
	// Parameters - 치환 파라미터 (수신자별로 적용), '치환 ID' 를 key로, '치환 ID에 맵핑되는 값' 을 value로 가지는 Map 형태의 Object
	Parameters *RecipientForRequestParameters `json:"parameters,omitempty"`
}

// RecipientForRequestParameters 치환 파라미터 (수신자별로 적용), '치환 ID' 를 key로, '치환 ID에 맵핑되는 값' 을 value로 가지는 Map 형태의
// Object
type RecipientForRequestParameters struct {
	// CustomerName - 고객 이름
	CustomerName *string `json:"customer_name,omitempty"`
	// BEFOREGRADE - 등급 (BRONZE, SILVER, GOLD)
	BEFOREGRADE *string `json:"BEFORE_GRADE,omitempty"`
	// AFTERGRADE - 등급 (BRONZE, SILVER, GOLD)
	AFTERGRADE *string `json:"AFTER_GRADE,omitempty"`
}

// RecipientGroupFilter 수신자 그룹 조합 필터
type RecipientGroupFilter struct {
	// AndFilter - AND 조합 여부 (true : AND 조합, false : OR 조합)
	AndFilter *bool `json:"andFilter,omitempty"`
	// Groups - 수신자 그룹명 목록
	Groups *[]string `json:"groups,omitempty"`
}

// SetObject ...
type SetObject struct {
	autorest.Response `json:"-"`
	Value             interface{} `json:"value,omitempty"`
}

// UnsubscribeListResponse ...
type UnsubscribeListResponse struct {
	// TotalElements - 총 레코드 개수
	TotalElements *int32 `json:"totalElements,omitempty"`
	// TotalPages - 총 페이지 개수
	TotalPages *int32 `json:"totalPages,omitempty"`
	// NumberOfElements - 현재 페이지의 레코드 개수
	NumberOfElements *int32 `json:"numberOfElements,omitempty"`
	// First - 첫번째 페이지 여부
	First *bool `json:"first,omitempty"`
	// Last - 마지막 페이지 여부
	Last *bool `json:"last,omitempty"`
	// Number - 현재 페이지 index (0부터 시작)
	Number *int32 `json:"number,omitempty"`
	// Size - 한페이지에 당 레코드 개수
	Size    *int32                                `json:"size,omitempty"`
	Sort    *[]UnsubscribeListResponseSortItem    `json:"sort,omitempty"`
	Content *[]UnsubscribeListResponseContentItem `json:"content,omitempty"`
}

// UnsubscribeListResponseContentItem ...
type UnsubscribeListResponseContentItem struct {
	// Address - Email 주소
	Address *string `json:"address,omitempty"`
	// BlockDate - 수신거부 등록일
	BlockDate *UnsubscribeListResponseContentItemBlockDate `json:"blockDate,omitempty"`
	// IsRegByManager - 관리자 등록 여부 (true: 관리자가 등록, false: 메일 수신자가 링크를 통해 수신거부 등록)
	IsRegByManager *bool `json:"isRegByManager,omitempty"`
}

// UnsubscribeListResponseContentItemBlockDate 수신거부 등록일
type UnsubscribeListResponseContentItemBlockDate struct {
	// Utc - 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수
	Utc *int64 `json:"utc,omitempty"`
	// FormattedDate - 'yyyy-MM-dd' UTC+09:00 형태의 날짜
	FormattedDate *string `json:"formattedDate,omitempty"`
	// FormattedDateTime - 'yyyy-MM-dd HH:mm:ss SSS' UTC+09:00 형태의 시간
	FormattedDateTime *string `json:"formattedDateTime,omitempty"`
}

// UnsubscribeListResponseSortItem ...
type UnsubscribeListResponseSortItem struct {
	// Direction - 정렬 방향 (ASC|DESC)
	Direction *string `json:"direction,omitempty"`
	// Property - 정렬 기준 필드명
	Property *string `json:"property,omitempty"`
	// IgnoreCase - 대소문자 구분하여 정렬할지 여부
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// NullHandling - null 처리 방식 (NATIVE: data 처리로직에 맡김 , NULLS_FIRST : null 값이 앞으로, NULLS_LAST: null 값이 뒤로)
	NullHandling *string `json:"nullHandling,omitempty"`
	// Ascending - 정렬 방향 Ascending(ASC) 인지 여부
	Ascending *bool `json:"ascending,omitempty"`
	// Descending - 정렬 방향 Descending(DESC) 인지 여부
	Descending *bool `json:"descending,omitempty"`
}
