package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MailClient is the cloud Outbound Mailer Client
type MailClient struct {
	BaseClient
}

// NewMailClient creates an instance of the MailClient client.
func NewMailClient() MailClient {
	return NewMailClientWithBaseURI(DefaultBaseURI)
}

// NewMailClientWithBaseURI creates an instance of the MailClient client.
func NewMailClientWithBaseURI(baseURI string) MailClient {
	return MailClient{NewWithBaseURI(baseURI)}
}

// Create email 발송을 요청
// Parameters:
// parameter - 메일 발송 요청
func (client MailClient) Create(ctx context.Context, parameter MailRequestParameter) (result MailResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MailClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameter,
			Constraints: []validation.Constraint{{Target: "parameter.SenderAddress", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.Title", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.Body", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.Recipients", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("outboundmailer.MailClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MailClient) CreatePreparer(ctx context.Context, parameter MailRequestParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/mails"),
		autorest.WithJSON(parameter))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MailClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client MailClient) CreateResponder(resp *http.Response) (result MailResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get email 발송요청 상세 조회 발송 요청한 목록을 조회
// Parameters:
// mailID - 메일 Id
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client MailClient) Get(ctx context.Context, mailID string, xNCPLANG string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MailClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, mailID, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MailClient) GetPreparer(ctx context.Context, mailID string, xNCPLANG string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"mailId": autorest.Encode("path", mailID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/mails/{mailId}", pathParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MailClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MailClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetListByRequestID 발송 요청한 특정 요청ID에 의해 생성된 메일 목록을 조회
// Parameters:
// requestID - requestId
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
// mailID - email ID (각 메일 한 건을 구분하는 ID)
// page - 결과를 받고 싶은 페이지 index (0..N) (default:0)
// recipientAddress - 수신자 Email 주소
// sendStatus - email 발송 상태 (R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
// size - 한페이지에 당 레코드 개수 (default:10)
// sortParameter - 정렬기준 필드 (형식 : property(,asc|desc)) : 기본정렬 방향은 오름차순(asc)이며, 복수의 필드를 정렬기준으로 사용할 수 있다. 정렬가능한
// 필드는 다음과 같다. (id : mailId, createUtc : 생성일시, statusCode : 발송상태)
// title - 메일 제목
func (client MailClient) GetListByRequestID(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MailClient.GetListByRequestID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListByRequestIDPreparer(ctx, requestID, xNCPLANG, mailID, page, recipientAddress, sendStatus, size, sortParameter, title)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetListByRequestID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListByRequestIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetListByRequestID", resp, "Failure sending request")
		return
	}

	result, err = client.GetListByRequestIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetListByRequestID", resp, "Failure responding to request")
	}

	return
}

// GetListByRequestIDPreparer prepares the GetListByRequestID request.
func (client MailClient) GetListByRequestIDPreparer(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"requestId": autorest.Encode("path", requestID),
	}

	queryParameters := map[string]interface{}{}
	if len(mailID) > 0 {
		queryParameters["mailId"] = autorest.Encode("query", mailID)
	}
	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if len(recipientAddress) > 0 {
		queryParameters["recipientAddress"] = autorest.Encode("query", recipientAddress)
	}
	if sendStatus != nil && len(sendStatus) > 0 {
		queryParameters["sendStatus"] = sendStatus
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}
	if sortParameter != nil && len(sortParameter) > 0 {
		queryParameters["sort"] = sortParameter
	}
	if len(title) > 0 {
		queryParameters["title"] = autorest.Encode("query", title)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/mails/requests/{requestId}/mails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListByRequestIDSender sends the GetListByRequestID request. The method will close the
// http.Response Body if it receives an error.
func (client MailClient) GetListByRequestIDSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListByRequestIDResponder handles the response to the GetListByRequestID request. The method always
// closes the http.Response Body.
func (client MailClient) GetListByRequestIDResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRequestList 발송 요청한 목록을 조회
// Parameters:
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
// dispatchType - email 발송요청 도구 (CONSOLE : NCP Console 화면을 통한 요청, API : API 호출을 통한 요청)
// endDateTime - 다음과 같은 형태의 요청 종료 일시 ('yyyy-MM-dd HH:mm' UTC+09:00), endUtc 값이 우선한다.
// endUtc - 요청 종료 일시 ( 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수 ),  endDateTime 값보다 이 값이 우선
// 적용된다.
// mailID - email ID (각 메일 한 건을 구분하는 ID)
// page - 결과를 받고 싶은 페이지 index (0..N) (default:0)
// recipientAddress - 수신자 Email 주소
// senderAddress - 발송자 Email 주소
// sendStatus - email 발송 상태 (P: 발송준비중, R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
// size - 한페이지에 당 레코드 개수 (default:10)
// sortParameter - 정렬기준 필드 (형식 : property(,asc|desc)) : 기본정렬 방향은 오름차순(asc)이며, 복수의 필드를 정렬기준으로 사용할 수 있다. 정렬가능한
// 필드는 다음과 같다. (createUtc : 생성일시, recipientCount : 수신자 수, reservationUtc : 예약일시, sendUtc : 발송완료일시, statusCode :
// 발송상태)
// startDateTime - 다음과 같은 형태의 요청 시작 일시 ('yyyy-MM-dd HH:mm' UTC+09:00), startUtc 값이 우선한다.
// startUtc - 요청 시작 일시 ( 1970년 1월 1일 00:00:00 협정 세계시(UTC) 부터의 경과 시간을  1/1000초로 환산한 정수 ),  startDateTime 값보다 이
// 값이 우선 적용된다.
// templateSid - 템플릿 ID
// title - 메일 제목
func (client MailClient) GetRequestList(ctx context.Context, xNCPLANG string, dispatchType string, endDateTime string, endUtc *int64, mailID string, page *int32, recipientAddress string, senderAddress string, sendStatus []string, size *int32, sortParameter string, startDateTime string, startUtc *int64, templateSid *int32, title string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MailClient.GetRequestList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRequestListPreparer(ctx, xNCPLANG, dispatchType, endDateTime, endUtc, mailID, page, recipientAddress, senderAddress, sendStatus, size, sortParameter, startDateTime, startUtc, templateSid, title)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetRequestList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRequestListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetRequestList", resp, "Failure sending request")
		return
	}

	result, err = client.GetRequestListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetRequestList", resp, "Failure responding to request")
	}

	return
}

// GetRequestListPreparer prepares the GetRequestList request.
func (client MailClient) GetRequestListPreparer(ctx context.Context, xNCPLANG string, dispatchType string, endDateTime string, endUtc *int64, mailID string, page *int32, recipientAddress string, senderAddress string, sendStatus []string, size *int32, sortParameter string, startDateTime string, startUtc *int64, templateSid *int32, title string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(dispatchType) > 0 {
		queryParameters["dispatchType"] = autorest.Encode("query", dispatchType)
	}
	if len(endDateTime) > 0 {
		queryParameters["endDateTime"] = autorest.Encode("query", endDateTime)
	}
	if endUtc != nil {
		queryParameters["endUtc"] = autorest.Encode("query", *endUtc)
	}
	if len(mailID) > 0 {
		queryParameters["mailId"] = autorest.Encode("query", mailID)
	}
	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if len(recipientAddress) > 0 {
		queryParameters["recipientAddress"] = autorest.Encode("query", recipientAddress)
	}
	if len(senderAddress) > 0 {
		queryParameters["senderAddress"] = autorest.Encode("query", senderAddress)
	}
	if sendStatus != nil && len(sendStatus) > 0 {
		queryParameters["sendStatus"] = sendStatus
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
	if templateSid != nil {
		queryParameters["templateSid"] = autorest.Encode("query", *templateSid)
	}
	if len(title) > 0 {
		queryParameters["title"] = autorest.Encode("query", title)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/mails/requests"),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRequestListSender sends the GetRequestList request. The method will close the
// http.Response Body if it receives an error.
func (client MailClient) GetRequestListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetRequestListResponder handles the response to the GetRequestList request. The method always
// closes the http.Response Body.
func (client MailClient) GetRequestListResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatusByRequestID 발송 요청한 requestId 를 전달하여 해당 요청으로 발송중인 메일 현황과 상태별 개수를 조회
// Parameters:
// requestID - requestId
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
// mailID - email ID (각 메일 한 건을 구분하는 ID)
// page - 결과를 받고 싶은 페이지 index (0..N) (default:0)
// recipientAddress - 수신자 Email 주소
// sendStatus - email 발송 상태 (R: 발송준비, I: 발송중, S: 발송성공, F: 발송실패, U: 수신거부, C:발송취소, PF: 일부실패)
// size - 한페이지에 당 레코드 개수 (default:10)
// sortParameter - 정렬기준 필드 (형식 : property(,asc|desc)) : 기본정렬 방향은 오름차순(asc)이며, 복수의 필드를 정렬기준으로 사용할 수 있다. 정렬가능한
// 필드는 다음과 같다. (id : mailId, createUtc : 생성일시, statusCode : 발송상태)
// title - 메일 제목
func (client MailClient) GetStatusByRequestID(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MailClient.GetStatusByRequestID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatusByRequestIDPreparer(ctx, requestID, xNCPLANG, mailID, page, recipientAddress, sendStatus, size, sortParameter, title)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetStatusByRequestID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatusByRequestIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetStatusByRequestID", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatusByRequestIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.MailClient", "GetStatusByRequestID", resp, "Failure responding to request")
	}

	return
}

// GetStatusByRequestIDPreparer prepares the GetStatusByRequestID request.
func (client MailClient) GetStatusByRequestIDPreparer(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"requestId": autorest.Encode("path", requestID),
	}

	queryParameters := map[string]interface{}{}
	if len(mailID) > 0 {
		queryParameters["mailId"] = autorest.Encode("query", mailID)
	}
	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if len(recipientAddress) > 0 {
		queryParameters["recipientAddress"] = autorest.Encode("query", recipientAddress)
	}
	if sendStatus != nil && len(sendStatus) > 0 {
		queryParameters["sendStatus"] = sendStatus
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}
	if sortParameter != nil && len(sortParameter) > 0 {
		queryParameters["sort"] = sortParameter
	}
	if len(title) > 0 {
		queryParameters["title"] = autorest.Encode("query", title)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/mails/requests/{requestId}/status", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatusByRequestIDSender sends the GetStatusByRequestID request. The method will close the
// http.Response Body if it receives an error.
func (client MailClient) GetStatusByRequestIDSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetStatusByRequestIDResponder handles the response to the GetStatusByRequestID request. The method always
// closes the http.Response Body.
func (client MailClient) GetStatusByRequestIDResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
