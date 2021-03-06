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

// DeviceClient is the SENS Client
type DeviceClient struct {
	BaseClient
}

// NewDeviceClient creates an instance of the DeviceClient client.
func NewDeviceClient() DeviceClient {
	return NewDeviceClientWithBaseURI(DefaultBaseURI)
}

func NewDeviceClientWithKey(accessKey string, secretKey string) DeviceClient {
	return NewDeviceClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewDeviceClientWithBaseURI creates an instance of the DeviceClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDeviceClientWithBaseURI(baseURI string) DeviceClient {
	return DeviceClient{NewWithBaseURI(baseURI)}
}

func NewDeviceClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) DeviceClient {
	return DeviceClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Create 디바이스 등록
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
func (client DeviceClient) Create(ctx context.Context, serviceID string, parameter PushUserRequestParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceClient.Create")
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
			Constraints: []validation.Constraint{{Target: "parameter.UserID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.DeviceType", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.DeviceToken", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.IsNotificationAgreement", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.IsAdAgreement", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameter.IsNightAdAgreement", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sens.DeviceClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, serviceID, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DeviceClient) CreatePreparer(ctx context.Context, serviceID string, parameter PushUserRequestParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/users", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users", pathParameters),
		autorest.WithJSON(parameter),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DeviceClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete 디바이스 삭제
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// userID - 사용자를 식별하는 아이디. 최대 128자
func (client DeviceClient) Delete(ctx context.Context, serviceID string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, serviceID, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DeviceClient) DeletePreparer(ctx context.Context, serviceID string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
		"userId":    autorest.Encode("path", userID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/users/{userId}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users/{userId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeviceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get 디바이스 조회
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// userID - 사용자를 식별하는 아이디. 최대 128자
func (client DeviceClient) Get(ctx context.Context, serviceID string, userID string) (result PushUserResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serviceID, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.DeviceClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeviceClient) GetPreparer(ctx context.Context, serviceID string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
		"userId":    autorest.Encode("path", userID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/users/{userId}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users/{userId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeviceClient) GetResponder(resp *http.Response) (result PushUserResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
