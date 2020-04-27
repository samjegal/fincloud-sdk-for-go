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

// MessageClient is the SENS Client
type MessageClient struct {
    BaseClient
}
// NewMessageClient creates an instance of the MessageClient client.
func NewMessageClient() MessageClient {
    return NewMessageClientWithBaseURI(DefaultBaseURI, )
}

// NewMessageClientWithBaseURI creates an instance of the MessageClient client.
    func NewMessageClientWithBaseURI(baseURI string, ) MessageClient {
        return MessageClient{ NewWithBaseURI(baseURI, )}
    }

// Delete 예약 메시지 삭제
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // reserveID - 예약 발송 요청 조회시 반환되는 메시지 식별자
func (client MessageClient) Delete(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MessageClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, serviceID, reserveID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sens.MessageClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client MessageClient) DeletePreparer(ctx context.Context, serviceID string, reserveID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "reserveId": autorest.Encode("path",reserveID),
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/push/v2/services/{serviceId}/reservations/{reserveId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client MessageClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MessageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// GetResult 메시지 결과 조회
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // requestID - 메시지 발송시 반환되는 요청 식별자
func (client MessageClient) GetResult(ctx context.Context, serviceID string, requestID string) (result PushMessageResultResponseParameter, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MessageClient.GetResult")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetResultPreparer(ctx, serviceID, requestID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sens.MessageClient", "GetResult", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetResultSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "GetResult", resp, "Failure sending request")
            return
            }

            result, err = client.GetResultResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "GetResult", resp, "Failure responding to request")
            }

    return
    }

    // GetResultPreparer prepares the GetResult request.
    func (client MessageClient) GetResultPreparer(ctx context.Context, serviceID string, requestID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "requestId": autorest.Encode("path",requestID),
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/push/v2/services/{serviceId}/messages/{requestId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetResultSender sends the GetResult request. The method will close the
    // http.Response Body if it receives an error.
    func (client MessageClient) GetResultSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResultResponder handles the response to the GetResult request. The method always
// closes the http.Response Body.
func (client MessageClient) GetResultResponder(resp *http.Response) (result PushMessageResultResponseParameter, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Send 메시지 발송
    // Parameters:
        // serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
        // parameter - 생성할 채널
func (client MessageClient) Send(ctx context.Context, serviceID string, parameter PushMessageRequestParameter) (result PushMessageResponseParameter, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/MessageClient.Send")
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
             Constraints: []validation.Constraint{	{Target: "parameter.Message", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Apns", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Apns.I18n", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Apns.I18n.Default", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }},
            	{Target: "parameter.Message.Gcm", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Gcm.I18n", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Gcm.I18n.Default", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }},
            	{Target: "parameter.Message.Default", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Default.I18n", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameter.Message.Default.I18n.Default", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }},
            }},
            	{Target: "parameter.Target", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameter.Target.Type", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("sens.MessageClient", "Send", err.Error())
            }

                req, err := client.SendPreparer(ctx, serviceID, parameter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sens.MessageClient", "Send", nil , "Failure preparing request")
    return
    }

            resp, err := client.SendSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "Send", resp, "Failure sending request")
            return
            }

            result, err = client.SendResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sens.MessageClient", "Send", resp, "Failure responding to request")
            }

    return
    }

    // SendPreparer prepares the Send request.
    func (client MessageClient) SendPreparer(ctx context.Context, serviceID string, parameter PushMessageRequestParameter) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "serviceId": serviceID,
            }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=UTF-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/push/v2/services/{serviceId}/messages",pathParameters),
    autorest.WithJSON(parameter))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // SendSender sends the Send request. The method will close the
    // http.Response Body if it receives an error.
    func (client MessageClient) SendSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// SendResponder handles the response to the Send request. The method always
// closes the http.Response Body.
func (client MessageClient) SendResponder(resp *http.Response) (result PushMessageResponseParameter, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusBadRequest,http.StatusUnauthorized,http.StatusForbidden,http.StatusNotFound,http.StatusInternalServerError),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

