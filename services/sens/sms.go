package sens

// FINCLOUD_APACHE_NO_VERSION

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// SMSClient is the SENS Client
type SMSClient struct {
    BaseClient
}
// NewSMSClient creates an instance of the SMSClient client.
func NewSMSClient() SMSClient {
    return NewSMSClientWithBaseURI(DefaultBaseURI, )
}

// NewSMSClientWithBaseURI creates an instance of the SMSClient client.
    func NewSMSClientWithBaseURI(baseURI string, ) SMSClient {
        return SMSClient{ NewWithBaseURI(baseURI, )}
    }

// DeleteReservation 예약 메시지를 취소
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // reserveID - 예약 발송 요청 조회시 반환되는 메시지 식별자
func (client SMSClient) DeleteReservation(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SMSClient.DeleteReservation")
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
    err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteReservation", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteReservationSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteReservation", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteReservationResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteReservation", resp, "Failure responding to request")
            }

    return
    }

    // DeleteReservationPreparer prepares the DeleteReservation request.
    func (client SMSClient) DeleteReservationPreparer(ctx context.Context, serviceID string, reserveID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "reserveId": autorest.Encode("path",reserveID),
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/sms/v2/services/{serviceId}/reservations/{reserveId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteReservationSender sends the DeleteReservation request. The method will close the
    // http.Response Body if it receives an error.
    func (client SMSClient) DeleteReservationSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteReservationResponder handles the response to the DeleteReservation request. The method always
// closes the http.Response Body.
func (client SMSClient) DeleteReservationResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// DeleteSchedule 스케쥴 메시지 삭제
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // scheduleCode - 스케줄 등록시 사용한 코드
        // messageID - 스케줄 발송 요청 조회시 반환되는 메시지 식별자
func (client SMSClient) DeleteSchedule(ctx context.Context, serviceID string, scheduleCode string, messageID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SMSClient.DeleteSchedule")
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
    err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteSchedule", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteScheduleSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteSchedule", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteScheduleResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "DeleteSchedule", resp, "Failure responding to request")
            }

    return
    }

    // DeleteSchedulePreparer prepares the DeleteSchedule request.
    func (client SMSClient) DeleteSchedulePreparer(ctx context.Context, serviceID string, scheduleCode string, messageID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "messageId": autorest.Encode("path",messageID),
            "scheduleCode": autorest.Encode("path",scheduleCode),
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/sms/v2/services/{serviceId}/schedules/{scheduleCode}/messages/{messageId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteScheduleSender sends the DeleteSchedule request. The method will close the
    // http.Response Body if it receives an error.
    func (client SMSClient) DeleteScheduleSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteScheduleResponder handles the response to the DeleteSchedule request. The method always
// closes the http.Response Body.
func (client SMSClient) DeleteScheduleResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// GetRequest 메시지 발송 요청 조회
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // requestID - 발송 요청 아이디
func (client SMSClient) GetRequest(ctx context.Context, serviceID string, requestID string) (result SMSMessageParameter, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SMSClient.GetRequest")
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
    err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetRequest", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetRequestSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetRequest", resp, "Failure sending request")
            return
            }

            result, err = client.GetRequestResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetRequest", resp, "Failure responding to request")
            }

    return
    }

    // GetRequestPreparer prepares the GetRequest request.
    func (client SMSClient) GetRequestPreparer(ctx context.Context, serviceID string, requestID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "serviceId": serviceID,
            }

                    queryParameters := map[string]interface{} {
        "requestId": autorest.Encode("query",requestID),
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/sms/v2/services/{serviceId}/messages",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetRequestSender sends the GetRequest request. The method will close the
    // http.Response Body if it receives an error.
    func (client SMSClient) GetRequestSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetRequestResponder handles the response to the GetRequest request. The method always
// closes the http.Response Body.
func (client SMSClient) GetRequestResponder(resp *http.Response) (result SMSMessageParameter, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetResult 메시지 발송 결과 조회
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // messageID - 메시지 발송 요청 조회시 반환되는 메시지 식별자
func (client SMSClient) GetResult(ctx context.Context, serviceID string, messageID string) (result SMSMessageParameter, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SMSClient.GetResult")
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
    err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetResult", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetResultSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetResult", resp, "Failure sending request")
            return
            }

            result, err = client.GetResultResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "GetResult", resp, "Failure responding to request")
            }

    return
    }

    // GetResultPreparer prepares the GetResult request.
    func (client SMSClient) GetResultPreparer(ctx context.Context, serviceID string, messageID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "serviceId": serviceID,
            }

                    queryParameters := map[string]interface{} {
        "messageId": autorest.Encode("query",messageID),
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/sms/v2/services/{serviceId}/messages/{messageId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetResultSender sends the GetResult request. The method will close the
    // http.Response Body if it receives an error.
    func (client SMSClient) GetResultSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResultResponder handles the response to the GetResult request. The method always
// closes the http.Response Body.
func (client SMSClient) GetResultResponder(resp *http.Response) (result SMSMessageParameter, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Request SMS 메시지 발송 요청
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // parameter - SMS 메시지 요청 정보
func (client SMSClient) Request(ctx context.Context, serviceID string, parameter SMSMessageRequestParameter) (result SMSMessageParameter, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SMSClient.Request")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: parameter,
             Constraints: []validation.Constraint{	{Target: "parameter.Type", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameter.From", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameter.Content", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameter.ReserveTime", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.ReserveTime", Name: validation.Pattern, Rule: `^(19|20)\d{2}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[0-1]) (0[0-9]|1[0-9]|2[0-3]):([0-5][0-9])$`, Chain: nil },
            }},
            	{Target: "parameter.Messages", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("sens.SMSClient", "Request", err.Error())
            }

                req, err := client.RequestPreparer(ctx, serviceID, parameter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sens.SMSClient", "Request", nil , "Failure preparing request")
    return
    }

            resp, err := client.RequestSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "Request", resp, "Failure sending request")
            return
            }

            result, err = client.RequestResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.SMSClient", "Request", resp, "Failure responding to request")
            }

    return
    }

    // RequestPreparer prepares the Request request.
    func (client SMSClient) RequestPreparer(ctx context.Context, serviceID string, parameter SMSMessageRequestParameter) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/sms/v2/services/{serviceId}/messages",pathParameters),
    autorest.WithJSON(parameter))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RequestSender sends the Request request. The method will close the
    // http.Response Body if it receives an error.
    func (client SMSClient) RequestSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// RequestResponder handles the response to the Request request. The method always
// closes the http.Response Body.
func (client SMSClient) RequestResponder(resp *http.Response) (result SMSMessageParameter, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusTooManyRequests,http.StatusInternalServerError),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

