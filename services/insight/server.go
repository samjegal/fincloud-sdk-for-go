package insight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServerClient is the cloud Insight Client
type ServerClient struct {
	BaseClient
}

// NewServerClient creates an instance of the ServerClient client.
func NewServerClient() ServerClient {
	return NewServerClientWithBaseURI(DefaultBaseURI)
}

// NewServerClientWithBaseURI creates an instance of the ServerClient client.
func NewServerClientWithBaseURI(baseURI string) ServerClient {
	return ServerClient{NewWithBaseURI(baseURI)}
}

// CreatTop target metric (mem_usert/avg_cpu_user_rto/fs_usert)
// Parameters:
// query - target metric (mem_usert/avg_cpu_user_rto/fs_usert)
func (client ServerClient) CreatTop(ctx context.Context, query string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerClient.CreatTop")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatTopPreparer(ctx, query)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.ServerClient", "CreatTop", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreatTopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.ServerClient", "CreatTop", resp, "Failure sending request")
		return
	}

	result, err = client.CreatTopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.ServerClient", "CreatTop", resp, "Failure responding to request")
	}

	return
}

// CreatTopPreparer prepares the CreatTop request.
func (client ServerClient) CreatTopPreparer(ctx context.Context, query string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(query) > 0 {
		queryParameters["query"] = autorest.Encode("query", query)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/servers/top"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreatTopSender sends the CreatTop request. The method will close the
// http.Response Body if it receives an error.
func (client ServerClient) CreatTopSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreatTopResponder handles the response to the CreatTop request. The method always
// closes the http.Response Body.
func (client ServerClient) CreatTopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
