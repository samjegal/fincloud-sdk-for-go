package geolocation

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// QuotaClient is the geoLocation Client
type QuotaClient struct {
	BaseClient
}

// NewQuotaClient creates an instance of the QuotaClient client.
func NewQuotaClient() QuotaClient {
	return NewQuotaClientWithBaseURI(DefaultBaseURI)
}

// NewQuotaClientWithBaseURI creates an instance of the QuotaClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQuotaClientWithBaseURI(baseURI string) QuotaClient {
	return QuotaClient{NewWithBaseURI(baseURI)}
}

// Get geoLocation 월간 최대 호출 가능량 조회
func (client QuotaClient) Get(ctx context.Context) (result QuotaParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client QuotaClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/geolocation/v2/quota"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QuotaClient) GetResponder(resp *http.Response) (result QuotaParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Set geoLocation 월간 최대 사용량(Quota) 수정
// Parameters:
// quota - unlimit 또는 제한량(ex:100)
func (client QuotaClient) Set(ctx context.Context, quota string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.Set")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SetPreparer(ctx, quota)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Set", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Set", resp, "Failure sending request")
		return
	}

	result, err = client.SetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.QuotaClient", "Set", resp, "Failure responding to request")
	}

	return
}

// SetPreparer prepares the Set request.
func (client QuotaClient) SetPreparer(ctx context.Context, quota string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"quota": autorest.Encode("query", quota),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/geolocation/v2/quota"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetSender sends the Set request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) SetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SetResponder handles the response to the Set request. The method always
// closes the http.Response Body.
func (client QuotaClient) SetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
