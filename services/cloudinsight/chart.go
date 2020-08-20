package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ChartClient is the cloud Insight Client
type ChartClient struct {
	BaseClient
}

// NewChartClient creates an instance of the ChartClient client.
func NewChartClient() ChartClient {
	return NewChartClientWithBaseURI(DefaultBaseURI)
}

// NewChartClientWithBaseURI creates an instance of the ChartClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewChartClientWithBaseURI(baseURI string) ChartClient {
	return ChartClient{NewWithBaseURI(baseURI)}
}

// Preview 간단하게 Metric 데이터와 함께 Preview Chart를 조회합니다.
// Parameters:
// parameters - metric information for the data to be retrieved
func (client ChartClient) Preview(ctx context.Context, parameters ChartDataWidgetRequest) (result WidgetMetricInfoResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ChartClient.Preview")
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
			Constraints: []validation.Constraint{{Target: "parameters.MetricsInfo", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.ChartClient", "Preview", err.Error())
	}

	req, err := client.PreviewPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ChartClient", "Preview", nil, "Failure preparing request")
		return
	}

	resp, err := client.PreviewSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.ChartClient", "Preview", resp, "Failure sending request")
		return
	}

	result, err = client.PreviewResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ChartClient", "Preview", resp, "Failure responding to request")
	}

	return
}

// PreviewPreparer prepares the Preview request.
func (client ChartClient) PreviewPreparer(ctx context.Context, parameters ChartDataWidgetRequest) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/chart/preview"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PreviewSender sends the Preview request. The method will close the
// http.Response Body if it receives an error.
func (client ChartClient) PreviewSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PreviewResponder handles the response to the Preview request. The method always
// closes the http.Response Body.
func (client ChartClient) PreviewResponder(resp *http.Response) (result WidgetMetricInfoResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}