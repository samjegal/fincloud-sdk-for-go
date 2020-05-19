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

// DeviceClient is the SENS Client
type DeviceClient struct {
	BaseClient
}

// NewDeviceClient creates an instance of the DeviceClient client.
func NewDeviceClient() DeviceClient {
	return NewDeviceClientWithBaseURI(DefaultBaseURI)
}

// NewDeviceClientWithBaseURI creates an instance of the DeviceClient client.
func NewDeviceClientWithBaseURI(baseURI string) DeviceClient {
	return DeviceClient{NewWithBaseURI(baseURI)}
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
			Constraints: []validation.Constraint{{Target: "parameter.UserID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameter.UserID", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
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

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users", pathParameters),
		autorest.WithJSON(parameter))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DeviceClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
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

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users/{userId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DeviceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
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

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/users/{userId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeviceClient) GetResponder(resp *http.Response) (result PushUserResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
