package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataClient is the cloud Insight Client
type DataClient struct {
	BaseClient
}

// NewDataClient creates an instance of the DataClient client.
func NewDataClient() DataClient {
	return NewDataClientWithBaseURI(DefaultBaseURI)
}

// NewDataClientWithBaseURI creates an instance of the DataClient client.
func NewDataClientWithBaseURI(baseURI string) DataClient {
	return DataClient{NewWithBaseURI(baseURI)}
}

// Preview metric information for the data to be retrieved
func (client DataClient) Preview(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.Preview")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PreviewPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Preview", nil, "Failure preparing request")
		return
	}

	resp, err := client.PreviewSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Preview", resp, "Failure sending request")
		return
	}

	result, err = client.PreviewResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Preview", resp, "Failure responding to request")
	}

	return
}

// PreviewPreparer prepares the Preview request.
func (client DataClient) PreviewPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/chart/preview"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PreviewSender sends the Preview request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) PreviewSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PreviewResponder handles the response to the Preview request. The method always
// closes the http.Response Body.
func (client DataClient) PreviewResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Query get widget data preview for dashboard widget
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client DataClient) Query(ctx context.Context, parameters QueryParameter) (result ListListFloat64, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.Query")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Query", nil, "Failure preparing request")
		return
	}

	resp, err := client.QuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Query", resp, "Failure sending request")
		return
	}

	result, err = client.QueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "Query", resp, "Failure responding to request")
	}

	return
}

// QueryPreparer prepares the Query request.
func (client DataClient) QueryPreparer(ctx context.Context, parameters QueryParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/query"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QuerySender sends the Query request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) QuerySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryResponder handles the response to the Query request. The method always
// closes the http.Response Body.
func (client DataClient) QueryResponder(resp *http.Response) (result ListListFloat64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// QueryMultiple query multiple metric data for a specific product with specified criteria
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client DataClient) QueryMultiple(ctx context.Context, parameters QueryMultipleParameter) (result ListDataInfoParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.QueryMultiple")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryMultiplePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "QueryMultiple", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryMultipleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "QueryMultiple", resp, "Failure sending request")
		return
	}

	result, err = client.QueryMultipleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.DataClient", "QueryMultiple", resp, "Failure responding to request")
	}

	return
}

// QueryMultiplePreparer prepares the QueryMultiple request.
func (client DataClient) QueryMultiplePreparer(ctx context.Context, parameters QueryMultipleParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/query/multiple"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryMultipleSender sends the QueryMultiple request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) QueryMultipleSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryMultipleResponder handles the response to the QueryMultiple request. The method always
// closes the http.Response Body.
func (client DataClient) QueryMultipleResponder(resp *http.Response) (result ListDataInfoParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
