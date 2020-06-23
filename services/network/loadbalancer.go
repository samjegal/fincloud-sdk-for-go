package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LoadBalancerClient is the network Client
type LoadBalancerClient struct {
	BaseClient
}

// NewLoadBalancerClient creates an instance of the LoadBalancerClient client.
func NewLoadBalancerClient() LoadBalancerClient {
	return NewLoadBalancerClientWithBaseURI(DefaultBaseURI)
}

// NewLoadBalancerClientWithBaseURI creates an instance of the LoadBalancerClient client.
func NewLoadBalancerClientWithBaseURI(baseURI string) LoadBalancerClient {
	return LoadBalancerClient{NewWithBaseURI(baseURI)}
}

// Search 로드밸런서 검색
// Parameters:
// parameters - 로드밸런서 검색 데이터
func (client LoadBalancerClient) Search(ctx context.Context, parameters LoadBalancerSearchParameter) (result LoadBalancerSearchListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerClient.Search")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SearchPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "Search", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "Search", resp, "Failure sending request")
		return
	}

	result, err = client.SearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "Search", resp, "Failure responding to request")
	}

	return
}

// SearchPreparer prepares the Search request.
func (client LoadBalancerClient) SearchPreparer(ctx context.Context, parameters LoadBalancerSearchParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/instances/search"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchSender sends the Search request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerClient) SearchSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// SearchResponder handles the response to the Search request. The method always
// closes the http.Response Body.
func (client LoadBalancerClient) SearchResponder(resp *http.Response) (result LoadBalancerSearchListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
