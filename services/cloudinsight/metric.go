package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MetricClient is the cloud Insight Client
type MetricClient struct {
	BaseClient
}

// NewMetricClient creates an instance of the MetricClient client.
func NewMetricClient() MetricClient {
	return NewMetricClientWithBaseURI(DefaultBaseURI)
}

// NewMetricClientWithBaseURI creates an instance of the MetricClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMetricClientWithBaseURI(baseURI string) MetricClient {
	return MetricClient{NewWithBaseURI(baseURI)}
}

// GetGroupItemsID 감시 항목 id를 할당받습니다.
// Parameters:
// count - metricGroupItem id 개수
func (client MetricClient) GetGroupItemsID(ctx context.Context, count int32) (result ListString, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricClient.GetGroupItemsID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetGroupItemsIDPreparer(ctx, count)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "GetGroupItemsID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGroupItemsIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "GetGroupItemsID", resp, "Failure sending request")
		return
	}

	result, err = client.GetGroupItemsIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "GetGroupItemsID", resp, "Failure responding to request")
	}

	return
}

// GetGroupItemsIDPreparer prepares the GetGroupItemsID request.
func (client MetricClient) GetGroupItemsIDPreparer(ctx context.Context, count int32) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"count": autorest.Encode("query", count),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/getMetricGroupItemsId"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetGroupItemsIDSender sends the GetGroupItemsID request. The method will close the
// http.Response Body if it receives an error.
func (client MetricClient) GetGroupItemsIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetGroupItemsIDResponder handles the response to the GetGroupItemsID request. The method always
// closes the http.Response Body.
func (client MetricClient) GetGroupItemsIDResponder(resp *http.Response) (result ListString, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
