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

// DataClient is the cloud Insight Client
type DataClient struct {
	BaseClient
}

// NewDataClient creates an instance of the DataClient client.
func NewDataClient() DataClient {
	return NewDataClientWithBaseURI(DefaultBaseURI)
}

// NewDataClientWithBaseURI creates an instance of the DataClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataClientWithBaseURI(baseURI string) DataClient {
	return DataClient{NewWithBaseURI(baseURI)}
}

// Query cloud Insight에서 수집한 time-series 데이터를 쿼리합니다.
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client DataClient) Query(ctx context.Context, parameters QueryRequest) (result ListListFloat64, err error) {
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
func (client DataClient) QueryPreparer(ctx context.Context, parameters QueryRequest) (*http.Request, error) {
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
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryResponder handles the response to the Query request. The method always
// closes the http.Response Body.
func (client DataClient) QueryResponder(resp *http.Response) (result ListListFloat64, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// QueryMultiple cloud Insight에서 수집한 여러개의 time-series 데이터를 쿼리합니다.
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client DataClient) QueryMultiple(ctx context.Context, parameters QueryMultipleRequest) (result ListMultipleDataParameter, err error) {
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
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.MetricInfoList", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimeStart", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.TimeEnd", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.DataClient", "QueryMultiple", err.Error())
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
func (client DataClient) QueryMultiplePreparer(ctx context.Context, parameters QueryMultipleRequest) (*http.Request, error) {
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
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryMultipleResponder handles the response to the QueryMultiple request. The method always
// closes the http.Response Body.
func (client DataClient) QueryMultipleResponder(resp *http.Response) (result ListMultipleDataParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
