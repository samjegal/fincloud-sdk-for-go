package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UnsubscriberClient is the cloud Outbound Mailer Client
type UnsubscriberClient struct {
	BaseClient
}

// NewUnsubscriberClient creates an instance of the UnsubscriberClient client.
func NewUnsubscriberClient() UnsubscriberClient {
	return NewUnsubscriberClientWithBaseURI(DefaultBaseURI)
}

// NewUnsubscriberClientWithBaseURI creates an instance of the UnsubscriberClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUnsubscriberClientWithBaseURI(baseURI string) UnsubscriberClient {
	return UnsubscriberClient{NewWithBaseURI(baseURI)}
}

// Get 수신거부 이메일 주소 목록을 조회
// Parameters:
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
// emailAddress - email 주소
// endDateTime - 수신거부 등록일 검색 기준 종료 일시, 문자열 포멧('yyyy-MM-dd HH:mm' UTC+09:00), endUtc 값이 우선한다.
// endUtc - 수신거부 등록일 검색 기준 종료 일시, UTC ( 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수 ),
// endDateTime 값보다 이 값이 우선 적용된다.
// isRegByManager - 관리자 등록 여부
// page - 결과를 받고 싶은 페이지 index (0..N) (default:0)
// size - 한페이지에 당 레코드 개수 (default:10)
// sortParameter - 정렬기준 필드 (형식 : property(,asc|desc)) : 기본정렬 방향은 오름차순(asc)이며, 복수의 필드를 정렬기준으로 사용할 수 있다. 정렬가능한
// 필드는 다음과 같다. (createUtc : 등록일시, adminRegYn : 관리자 등록 여부 )
// startDateTime - 수신거부 등록일 검색 기준 시작 일시, 문자열 포멧('yyyy-MM-dd HH:mm' UTC+09:00), startUtc 값이 우선한다.
// startUtc - 수신거부 등록일 검색 기준 시작 일시, UTC( 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수 ),
// startDateTime 값보다 이 값이 우선 적용된다.
func (client UnsubscriberClient) Get(ctx context.Context, xNCPLANG string, emailAddress string, endDateTime string, endUtc *int64, isRegByManager *bool, page *int32, size *int32, sortParameter string, startDateTime string, startUtc *int64) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UnsubscriberClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, xNCPLANG, emailAddress, endDateTime, endUtc, isRegByManager, page, size, sortParameter, startDateTime, startUtc)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.UnsubscriberClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.UnsubscriberClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.UnsubscriberClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UnsubscriberClient) GetPreparer(ctx context.Context, xNCPLANG string, emailAddress string, endDateTime string, endUtc *int64, isRegByManager *bool, page *int32, size *int32, sortParameter string, startDateTime string, startUtc *int64) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(emailAddress) > 0 {
		queryParameters["emailAddress"] = autorest.Encode("query", emailAddress)
	}
	if len(endDateTime) > 0 {
		queryParameters["endDateTime"] = autorest.Encode("query", endDateTime)
	}
	if endUtc != nil {
		queryParameters["endUtc"] = autorest.Encode("query", *endUtc)
	}
	if isRegByManager != nil {
		queryParameters["isRegByManager"] = autorest.Encode("query", *isRegByManager)
	}
	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}
	if len(sortParameter) > 0 {
		queryParameters["sort"] = autorest.Encode("query", sortParameter)
	}
	if len(startDateTime) > 0 {
		queryParameters["startDateTime"] = autorest.Encode("query", startDateTime)
	}
	if startUtc != nil {
		queryParameters["startUtc"] = autorest.Encode("query", *startUtc)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/unsubscribers"),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UnsubscriberClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UnsubscriberClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
