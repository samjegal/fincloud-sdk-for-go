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

// MessageClient is the SENS Client
type MessageClient struct {
	BaseClient
}

// NewMessageClient creates an instance of the MessageClient client.
func NewMessageClient() MessageClient {
	return NewMessageClientWithBaseURI(DefaultBaseURI)
}

// NewMessageClientWithBaseURI creates an instance of the MessageClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMessageClientWithBaseURI(baseURI string) MessageClient {
	return MessageClient{NewWithBaseURI(baseURI)}
}

// Delete 예약 메시지 삭제
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// reserveID - 예약 발송 요청 조회시 반환되는 메시지 식별자
func (client MessageClient) Delete(ctx context.Context, serviceID string, reserveID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MessageClient.Delete")
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
		err = autorest.NewErrorWithError(err, "sens.MessageClient", "Delete", nil, "Failure preparing request")
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
	pathParameters := map[string]interface{}{
		"reserveId": autorest.Encode("path", reserveID),
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/reservations/{reserveId}", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/reservations/{reserveId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MessageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MessageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
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
		ctx = tracing.StartSpan(ctx, fqdn+"/MessageClient.GetResult")
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
		err = autorest.NewErrorWithError(err, "sens.MessageClient", "GetResult", nil, "Failure preparing request")
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
	pathParameters := map[string]interface{}{
		"requestId": autorest.Encode("path", requestID),
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/messages/{requestId}", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/messages/{requestId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResultSender sends the GetResult request. The method will close the
// http.Response Body if it receives an error.
func (client MessageClient) GetResultSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResultResponder handles the response to the GetResult request. The method always
// closes the http.Response Body.
func (client MessageClient) GetResultResponder(resp *http.Response) (result PushMessageResultResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SendMethod 메시지 발송
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// parameter - 생성할 채널
func (client MessageClient) SendMethod(ctx context.Context, serviceID string, parameter PushMessageRequestParameter) (result PushMessageResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MessageClient.SendMethod")
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
			Constraints: []validation.Constraint{{Target: "parameter.Message", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameter.Message.Apns", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameter.Message.Apns.I18n", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameter.Message.Apns.I18n.Default", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "parameter.Message.Gcm", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameter.Message.Gcm.I18n", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameter.Message.Gcm.I18n.Default", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameter.Message.Default", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameter.Message.Default.I18n", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameter.Message.Default.I18n.Default", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
				}},
				{Target: "parameter.Target", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameter.Target.Type", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("sens.MessageClient", "SendMethod", err.Error())
	}

	req, err := client.SendMethodPreparer(ctx, serviceID, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.MessageClient", "SendMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.MessageClient", "SendMethod", resp, "Failure sending request")
		return
	}

	result, err = client.SendMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.MessageClient", "SendMethod", resp, "Failure responding to request")
	}

	return
}

// SendMethodPreparer prepares the SendMethod request.
func (client MessageClient) SendMethodPreparer(ctx context.Context, serviceID string, parameter PushMessageRequestParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/messages", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/messages", pathParameters),
		autorest.WithJSON(parameter),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendMethodSender sends the SendMethod request. The method will close the
// http.Response Body if it receives an error.
func (client MessageClient) SendMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SendMethodResponder handles the response to the SendMethod request. The method always
// closes the http.Response Body.
func (client MessageClient) SendMethodResponder(resp *http.Response) (result PushMessageResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
