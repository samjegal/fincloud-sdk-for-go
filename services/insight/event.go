package insight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EventClient is the cloud Insight Client
type EventClient struct {
	BaseClient
}

// NewEventClient creates an instance of the EventClient client.
func NewEventClient() EventClient {
	return NewEventClientWithBaseURI(DefaultBaseURI)
}

// NewEventClientWithBaseURI creates an instance of the EventClient client.
func NewEventClientWithBaseURI(baseURI string) EventClient {
	return EventClient{NewWithBaseURI(baseURI)}
}

// SearchEventCount search event count console
func (client EventClient) SearchEventCount(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventClient.SearchEventCount")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SearchEventCountPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.EventClient", "SearchEventCount", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchEventCountSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.EventClient", "SearchEventCount", resp, "Failure sending request")
		return
	}

	result, err = client.SearchEventCountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.EventClient", "SearchEventCount", resp, "Failure responding to request")
	}

	return
}

// SearchEventCountPreparer prepares the SearchEventCount request.
func (client EventClient) SearchEventCountPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/event/searchEventCountConsole"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchEventCountSender sends the SearchEventCount request. The method will close the
// http.Response Body if it receives an error.
func (client EventClient) SearchEventCountSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// SearchEventCountResponder handles the response to the SearchEventCount request. The method always
// closes the http.Response Body.
func (client EventClient) SearchEventCountResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
