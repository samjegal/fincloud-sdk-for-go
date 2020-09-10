package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/fincloud-sdk-for-go/common"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
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

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/getMetricGroupItemsId")+"?"+common.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/getMetricGroupItemsId"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
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

// SearchList 감시대상 그룹에서 조회가능한 항목(metric) 리스트를 조회합니다.
// Parameters:
// parameters - search metric list
func (client MetricClient) SearchList(ctx context.Context, parameters MetricListRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricClient.SearchList")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SearchListPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "SearchList", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchListSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "SearchList", resp, "Failure sending request")
		return
	}

	result, err = client.SearchListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MetricClient", "SearchList", resp, "Failure responding to request")
	}

	return
}

// SearchListPreparer prepares the SearchList request.
func (client MetricClient) SearchListPreparer(ctx context.Context, parameters MetricListRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/metric/search"), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/metric/search"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchListSender sends the SearchList request. The method will close the
// http.Response Body if it receives an error.
func (client MetricClient) SearchListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SearchListResponder handles the response to the SearchList request. The method always
// closes the http.Response Body.
func (client MetricClient) SearchListResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
