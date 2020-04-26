package insight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CollectorClient is the cloud Insight Client
type CollectorClient struct {
	BaseClient
}

// NewCollectorClient creates an instance of the CollectorClient client.
func NewCollectorClient() CollectorClient {
	return NewCollectorClientWithBaseURI(DefaultBaseURI)
}

// NewCollectorClientWithBaseURI creates an instance of the CollectorClient client.
func NewCollectorClientWithBaseURI(baseURI string) CollectorClient {
	return CollectorClient{NewWithBaseURI(baseURI)}
}

// Push collector API
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client CollectorClient) Push(ctx context.Context, parameters CloudInsightCollectorParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CollectorClient.Push")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PushPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.CollectorClient", "Push", nil, "Failure preparing request")
		return
	}

	resp, err := client.PushSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.CollectorClient", "Push", resp, "Failure sending request")
		return
	}

	result, err = client.PushResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.CollectorClient", "Push", resp, "Failure responding to request")
	}

	return
}

// PushPreparer prepares the Push request.
func (client CollectorClient) PushPreparer(ctx context.Context, parameters CloudInsightCollectorParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_collector/real"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PushSender sends the Push request. The method will close the
// http.Response Body if it receives an error.
func (client CollectorClient) PushSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PushResponder handles the response to the Push request. The method always
// closes the http.Response Body.
func (client CollectorClient) PushResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
