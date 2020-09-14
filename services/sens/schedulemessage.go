package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/fincloud-sdk-for-go/common"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// ScheduleMessageClient is the SENS Client
type ScheduleMessageClient struct {
	BaseClient
}

// NewScheduleMessageClient creates an instance of the ScheduleMessageClient client.
func NewScheduleMessageClient() ScheduleMessageClient {
	return NewScheduleMessageClientWithBaseURI(DefaultBaseURI)
}

func NewScheduleMessageClientWithKey(accessKey string, secretKey string) ScheduleMessageClient {
	return NewScheduleMessageClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewScheduleMessageClientWithBaseURI creates an instance of the ScheduleMessageClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewScheduleMessageClientWithBaseURI(baseURI string) ScheduleMessageClient {
	return ScheduleMessageClient{NewWithBaseURI(baseURI)}
}

func NewScheduleMessageClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) ScheduleMessageClient {
	return ScheduleMessageClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Delete 스케쥴 메시지 삭제
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
// messageID - 스케줄 발송 요청 조회시 반환되는 메시지 식별자
func (client ScheduleMessageClient) Delete(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleMessageClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, serviceID, scheduleCode, messageID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ScheduleMessageClient) DeletePreparer(ctx context.Context, serviceID string, scheduleCode string, messageID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"messageId":    autorest.Encode("path", messageID),
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleMessageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScheduleMessageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get 스케쥴 메시지 단건 조회
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
// messageID - 스케줄 발송 요청 조회시 반환되는 메시지 식별자
func (client ScheduleMessageClient) Get(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result PushScheduleFetchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleMessageClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceID, scheduleCode, messageID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScheduleMessageClient) GetPreparer(ctx context.Context, serviceID string, scheduleCode string, messageID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"messageId":    autorest.Encode("path", messageID),
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleMessageClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScheduleMessageClient) GetResponder(resp *http.Response) (result PushScheduleFetchParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List 스케쥴 메시지 조회
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
// page - 페이지 번호 (0부터 시작)
// size - 페이지 사이즈 (기본 20)
func (client ScheduleMessageClient) List(ctx context.Context, serviceID string, scheduleCode string, page *int32, size *int32) (result PushScheduleFetchAllParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleMessageClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, serviceID, scheduleCode, page, size)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleMessageClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScheduleMessageClient) ListPreparer(ctx context.Context, serviceID string, scheduleCode string, page *int32, size *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	queryParameters := map[string]interface{}{}

	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages", pathParameters)+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}/messages", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleMessageClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScheduleMessageClient) ListResponder(resp *http.Response) (result PushScheduleFetchAllParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
