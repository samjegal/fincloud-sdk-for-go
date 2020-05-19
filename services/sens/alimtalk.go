package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AlimtalkClient is the SENS Client
type AlimtalkClient struct {
	BaseClient
}

// NewAlimtalkClient creates an instance of the AlimtalkClient client.
func NewAlimtalkClient() AlimtalkClient {
	return NewAlimtalkClientWithBaseURI(DefaultBaseURI)
}

// NewAlimtalkClientWithBaseURI creates an instance of the AlimtalkClient client.
func NewAlimtalkClientWithBaseURI(baseURI string) AlimtalkClient {
	return AlimtalkClient{NewWithBaseURI(baseURI)}
}

// Create 알림톡 발송 요청
// Parameters:
// serviceID - serviceId
func (client AlimtalkClient) Create(ctx context.Context, serviceID string, request AlimTalkMessageRequestParameter) (result AlimTalkMessageResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlimtalkClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.TemplateCode", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "request.PlusFriendID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "request.Messages", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sens.AlimtalkClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, serviceID, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AlimtalkClient) CreatePreparer(ctx context.Context, serviceID string, request AlimTalkMessageRequestParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/alimtalk/v2/services/{serviceId}/messages", pathParameters),
		autorest.WithJSON(request))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AlimtalkClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AlimtalkClient) CreateResponder(resp *http.Response) (result AlimTalkMessageResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteReservation 예약 메시지 삭제
// Parameters:
// serviceID - serviceId
// reserveID - reserveId
func (client AlimtalkClient) DeleteReservation(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlimtalkClient.DeleteReservation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteReservationPreparer(ctx, serviceID, reserveID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteReservation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteReservationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteReservation", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteReservationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteReservation", resp, "Failure responding to request")
	}

	return
}

// DeleteReservationPreparer prepares the DeleteReservation request.
func (client AlimtalkClient) DeleteReservationPreparer(ctx context.Context, serviceID string, reserveID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reserveId": autorest.Encode("path", reserveID),
		"serviceId": serviceID,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/alimtalk/v2/services/{serviceId}/reservations/{reserveId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteReservationSender sends the DeleteReservation request. The method will close the
// http.Response Body if it receives an error.
func (client AlimtalkClient) DeleteReservationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteReservationResponder handles the response to the DeleteReservation request. The method always
// closes the http.Response Body.
func (client AlimtalkClient) DeleteReservationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteSchedule 스케쥴 메시지 삭제
// Parameters:
// serviceID - serviceId
// scheduleCode - scheduleCode
// messageID - messageId
func (client AlimtalkClient) DeleteSchedule(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlimtalkClient.DeleteSchedule")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteSchedulePreparer(ctx, serviceID, scheduleCode, messageID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteSchedule", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteScheduleSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteSchedule", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteScheduleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "DeleteSchedule", resp, "Failure responding to request")
	}

	return
}

// DeleteSchedulePreparer prepares the DeleteSchedule request.
func (client AlimtalkClient) DeleteSchedulePreparer(ctx context.Context, serviceID string, scheduleCode string, messageID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"messageId":    autorest.Encode("path", messageID),
		"scheduleCode": autorest.Encode("path", scheduleCode),
		"serviceId":    serviceID,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/alimtalk/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteScheduleSender sends the DeleteSchedule request. The method will close the
// http.Response Body if it receives an error.
func (client AlimtalkClient) DeleteScheduleSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteScheduleResponder handles the response to the DeleteSchedule request. The method always
// closes the http.Response Body.
func (client AlimtalkClient) DeleteScheduleResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetRequest 알림톡 발송 요청 조회
// Parameters:
// serviceID - serviceId
// requestID - 요청 아이디
func (client AlimtalkClient) GetRequest(ctx context.Context, serviceID string, requestID string) (result AlimTalkRequestResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlimtalkClient.GetRequest")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRequestPreparer(ctx, serviceID, requestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetRequest", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRequestSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetRequest", resp, "Failure sending request")
		return
	}

	result, err = client.GetRequestResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetRequest", resp, "Failure responding to request")
	}

	return
}

// GetRequestPreparer prepares the GetRequest request.
func (client AlimtalkClient) GetRequestPreparer(ctx context.Context, serviceID string, requestID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	queryParameters := map[string]interface{}{
		"requestId": autorest.Encode("query", requestID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/alimtalk/v2/services/{serviceId}/messages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRequestSender sends the GetRequest request. The method will close the
// http.Response Body if it receives an error.
func (client AlimtalkClient) GetRequestSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetRequestResponder handles the response to the GetRequest request. The method always
// closes the http.Response Body.
func (client AlimtalkClient) GetRequestResponder(resp *http.Response) (result AlimTalkRequestResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetResult 알림톡 발송 결과 조회
// Parameters:
// serviceID - serviceId
// messageID - messageId
func (client AlimtalkClient) GetResult(ctx context.Context, serviceID string, messageID string) (result AlimTalkResultResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlimtalkClient.GetResult")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetResultPreparer(ctx, serviceID, messageID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetResult", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetResult", resp, "Failure sending request")
		return
	}

	result, err = client.GetResultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.AlimtalkClient", "GetResult", resp, "Failure responding to request")
	}

	return
}

// GetResultPreparer prepares the GetResult request.
func (client AlimtalkClient) GetResultPreparer(ctx context.Context, serviceID string, messageID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"messageId": autorest.Encode("path", messageID),
		"serviceId": serviceID,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/alimtalk/v2/services/{serviceId}/messages/{messageId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResultSender sends the GetResult request. The method will close the
// http.Response Body if it receives an error.
func (client AlimtalkClient) GetResultSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResultResponder handles the response to the GetResult request. The method always
// closes the http.Response Body.
func (client AlimtalkClient) GetResultResponder(resp *http.Response) (result AlimTalkResultResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
