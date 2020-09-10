package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/fincloud-sdk-for-go/common"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// ScheduleClient is the SENS Client
type ScheduleClient struct {
	BaseClient
}

// NewScheduleClient creates an instance of the ScheduleClient client.
func NewScheduleClient() ScheduleClient {
	return NewScheduleClientWithBaseURI(DefaultBaseURI)
}

// NewScheduleClientWithBaseURI creates an instance of the ScheduleClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewScheduleClientWithBaseURI(baseURI string) ScheduleClient {
	return ScheduleClient{NewWithBaseURI(baseURI)}
}

// Create 스케줄 등록
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
func (client ScheduleClient) Create(ctx context.Context, serviceID string, parameter PushScheduleRegisterParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameter,
			Constraints: []validation.Constraint{{Target: "parameter.ScheduleCode", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.ScheduleTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.EndDate", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.DayOfWeeks", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sens.ScheduleClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, serviceID, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ScheduleClient) CreatePreparer(ctx context.Context, serviceID string, parameter PushScheduleRegisterParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules", pathParameters),
		autorest.WithJSON(parameter),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ScheduleClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete 스케쥴 삭제
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
func (client ScheduleClient) Delete(ctx context.Context, serviceID string, scheduleCode string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, serviceID, scheduleCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ScheduleClient) DeletePreparer(ctx context.Context, serviceID string, scheduleCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScheduleClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get 스케쥴 단건 조회
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
func (client ScheduleClient) Get(ctx context.Context, serviceID string, scheduleCode string) (result PushScheduleFetchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceID, scheduleCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ScheduleClient) GetPreparer(ctx context.Context, serviceID string, scheduleCode string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScheduleClient) GetResponder(resp *http.Response) (result PushScheduleFetchParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List 스케줄 목록 조회
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// page - 페이지 (0부터 시작)
// scheduleCode - 등록하려는 스케줄 코드
// size - 페이지 크기 (기본 20)
// sortParameter - 정렬 항목 (asc|desc)
func (client ScheduleClient) List(ctx context.Context, serviceID string, page *int32, scheduleCode string, size *int32, sortParameter string) (result PushScheduleFetchAllParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, serviceID, page, scheduleCode, size, sortParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ScheduleClient) ListPreparer(ctx context.Context, serviceID string, page *int32, scheduleCode string, size *int32, sortParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	queryParameters := map[string]interface{}{}

	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if len(scheduleCode) > 0 {
		queryParameters["scheduleCode"] = autorest.Encode("query", scheduleCode)
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}
	if len(sortParameter) > 0 {
		queryParameters["sort"] = autorest.Encode("query", sortParameter)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules", pathParameters)+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ScheduleClient) ListResponder(resp *http.Response) (result PushScheduleFetchAllParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put 스케쥴 수정
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// scheduleCode - 등록하려는 스케줄 코드
func (client ScheduleClient) Put(ctx context.Context, serviceID string, scheduleCode string, parameter PushScheduleRegisterParameter) (result PushScheduleFetchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ScheduleClient.Put")
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
			Constraints: []validation.Constraint{{Target: "parameter.ScheduleCode", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.ScheduleTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.EndDate", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.DayOfWeeks", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sens.ScheduleClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, serviceID, scheduleCode, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ScheduleClient", "Put", resp, "Failure responding to request")
	}

	return
}

// PutPreparer prepares the Put request.
func (client ScheduleClient) PutPreparer(ctx context.Context, serviceID string, scheduleCode string, parameter PushScheduleRegisterParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("PUT", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=UTF-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/schedules/{scheduleCode}", pathParameters),
		autorest.WithJSON(parameter),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client ScheduleClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client ScheduleClient) PutResponder(resp *http.Response) (result PushScheduleFetchParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
