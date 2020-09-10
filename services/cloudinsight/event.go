package cloudinsight

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

// EventClient is the cloud Insight Client
type EventClient struct {
	BaseClient
}

// NewEventClient creates an instance of the EventClient client.
func NewEventClient() EventClient {
	return NewEventClientWithBaseURI(DefaultBaseURI)
}

// NewEventClientWithBaseURI creates an instance of the EventClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEventClientWithBaseURI(baseURI string) EventClient {
	return EventClient{NewWithBaseURI(baseURI)}
}

// SearchByID event id와 Rule id로 Event를 상세 조회합니다.
// Parameters:
// parameters - event search criteria
func (client EventClient) SearchByID(ctx context.Context, parameters EventSearchRequest) (result ListEventSearchResultParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventClient.SearchByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RuleID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.EventID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.StartTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.EndTime", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.EventClient", "SearchByID", err.Error())
	}

	req, err := client.SearchByIDPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchByID", resp, "Failure sending request")
		return
	}

	result, err = client.SearchByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchByID", resp, "Failure responding to request")
	}

	return
}

// SearchByIDPreparer prepares the SearchByID request.
func (client EventClient) SearchByIDPreparer(ctx context.Context, parameters EventSearchRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw/api/event/searchById"), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw/api/event/searchById"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchByIDSender sends the SearchByID request. The method will close the
// http.Response Body if it receives an error.
func (client EventClient) SearchByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SearchByIDResponder handles the response to the SearchByID request. The method always
// closes the http.Response Body.
func (client EventClient) SearchByIDResponder(resp *http.Response) (result ListEventSearchResultParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SearchEventCount search event count console
func (client EventClient) SearchEventCount(ctx context.Context, parameters SearchEventCountConsoleRequest) (result SearchEventCountConsoleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventClient.SearchEventCount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.StartTime", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.EndTime", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.EventClient", "SearchEventCount", err.Error())
	}

	req, err := client.SearchEventCountPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchEventCount", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchEventCountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchEventCount", resp, "Failure sending request")
		return
	}

	result, err = client.SearchEventCountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.EventClient", "SearchEventCount", resp, "Failure responding to request")
	}

	return
}

// SearchEventCountPreparer prepares the SearchEventCount request.
func (client EventClient) SearchEventCountPreparer(ctx context.Context, parameters SearchEventCountConsoleRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/event/searchEventCountConsole"), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/event/searchEventCountConsole"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchEventCountSender sends the SearchEventCount request. The method will close the
// http.Response Body if it receives an error.
func (client EventClient) SearchEventCountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SearchEventCountResponder handles the response to the SearchEventCount request. The method always
// closes the http.Response Body.
func (client EventClient) SearchEventCountResponder(resp *http.Response) (result SearchEventCountConsoleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
