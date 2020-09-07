package sens

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// ChannelClient is the SENS Client
type ChannelClient struct {
	BaseClient
}

// NewChannelClient creates an instance of the ChannelClient client.
func NewChannelClient() ChannelClient {
	return NewChannelClientWithBaseURI(DefaultBaseURI)
}

// NewChannelClientWithBaseURI creates an instance of the ChannelClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewChannelClientWithBaseURI(baseURI string) ChannelClient {
	return ChannelClient{NewWithBaseURI(baseURI)}
}

// AddUser 채널에 사용자 추가
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// channelName - 사용자를 추가할 채널명
// userID - 디바이스 토큰 등록 시 바인딩한 사용자 아이디
func (client ChannelClient) AddUser(ctx context.Context, serviceID string, channelName string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChannelClient.AddUser")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddUserPreparer(ctx, serviceID, channelName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "AddUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddUserSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "AddUser", resp, "Failure sending request")
		return
	}

	result, err = client.AddUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "AddUser", resp, "Failure responding to request")
	}

	return
}

// AddUserPreparer prepares the AddUser request.
func (client ChannelClient) AddUserPreparer(ctx context.Context, serviceID string, channelName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName": autorest.Encode("path", channelName),
		"serviceId":   serviceID,
		"userId":      autorest.Encode("path", userID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/channels/{channelName}/users/{userId}", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/channels/{channelName}/users/{userId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddUserSender sends the AddUser request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelClient) AddUserSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddUserResponder handles the response to the AddUser request. The method always
// closes the http.Response Body.
func (client ChannelClient) AddUserResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create 채널 생성
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// channel - 생성할 채널명
func (client ChannelClient) Create(ctx context.Context, serviceID string, channel PushChannelRequestParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChannelClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: channel,
			Constraints: []validation.Constraint{{Target: "channel.ChannelName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sens.ChannelClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, serviceID, channel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ChannelClient) CreatePreparer(ctx context.Context, serviceID string, channel PushChannelRequestParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceId": serviceID,
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/channels", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/channels", pathParameters),
		autorest.WithJSON(channel),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ChannelClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteUser 채널에서 사용자 삭제
// Parameters:
// serviceID - 프로젝트 등록 시 발급받은 서비스 아이디
// channelName - 사용자를 추가할 채널명
// userID - 디바이스 토큰 등록 시 바인딩한 사용자 아이디
func (client ChannelClient) DeleteUser(ctx context.Context, serviceID string, channelName string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChannelClient.DeleteUser")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteUserPreparer(ctx, serviceID, channelName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "DeleteUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteUserSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "DeleteUser", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sens.ChannelClient", "DeleteUser", resp, "Failure responding to request")
	}

	return
}

// DeleteUserPreparer prepares the DeleteUser request.
func (client ChannelClient) DeleteUserPreparer(ctx context.Context, serviceID string, channelName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName": autorest.Encode("path", channelName),
		"serviceId":   serviceID,
		"userId":      autorest.Encode("path", userID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", autorest.GetPathParameters(DefaultBaseURI, "/push/v2/services/{serviceId}/channels/{channelName}/users/{userId}", pathParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/push/v2/services/{serviceId}/channels/{channelName}/users/{userId}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteUserSender sends the DeleteUser request. The method will close the
// http.Response Body if it receives an error.
func (client ChannelClient) DeleteUserSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteUserResponder handles the response to the DeleteUser request. The method always
// closes the http.Response Body.
func (client ChannelClient) DeleteUserResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
