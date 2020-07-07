package geolocation

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/geolocation"

// Parameter ...
type Parameter struct {
	// Country - 국가코드 EU(유럽), AP(아시아, 호주), A1(Anonymous Proxy), A2(Satellite provider), O1(Other Country)를 포함
	Country *string `json:"country,omitempty"`
	// Code - 한국의 경우에 행정자치부에서 정하는 행정구역 코드
	Code *string `json:"code,omitempty"`
	// R1 - 도, 광역시, 주
	R1 *string `json:"r1,omitempty"`
	// R2 - 시, 군, 구
	R2 *string `json:"r2,omitempty"`
	// R3 - 동, 면, 읍
	R3 *string `json:"r3,omitempty"`
	// Lat - 위도
	Lat *float64 `json:"lat,omitempty"`
	// Long - 경도
	Long *float64 `json:"long,omitempty"`
	// Net - 통신사 이름
	Net *string `json:"net,omitempty"`
}

// QuotaParameter ...
type QuotaParameter struct {
	autorest.Response `json:"-"`
	// Resource - 사용량 리소스 정보
	Resource interface{} `json:"resource,omitempty"`
}

// QuotaResourceParameter ...
type QuotaResourceParameter struct {
	// Type - 제한 또는 무제한 사용 상태 반환 (Unlimited 또는 Limited)
	Type *string `json:"type,omitempty"`
	// Quota - 설정된 Quota 값 (무제한 사용일 경우, -1로 응답)
	Quota *int32 `json:"quota,omitempty"`
}

// ResultParameter ...
type ResultParameter struct {
	autorest.Response `json:"-"`
	// ReturnCode - 리턴 코드
	ReturnCode *int32 `json:"returnCode,omitempty"`
	// RequestID - GeoLocation 요청 UUID
	RequestID *string `json:"requestId,omitempty"`
	// GeoLocation - GeoLocation 정보
	GeoLocation interface{} `json:"geoLocation,omitempty"`
}
