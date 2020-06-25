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

// CheckName 로드밸런서 이름 적합성 검사
// Parameters:
// loadBalancerName - VPC 번호
func (client LoadBalancerClient) CheckName(ctx context.Context, loadBalancerName string) (result LoadBalancerCheckNameParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerClient.CheckName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNamePreparer(ctx, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "CheckName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "CheckName", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "CheckName", resp, "Failure responding to request")
	}

	return
}

// CheckNamePreparer prepares the CheckName request.
func (client LoadBalancerClient) CheckNamePreparer(ctx context.Context, loadBalancerName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"loadBalancerName": autorest.Encode("query", loadBalancerName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/check-load-balancer-name"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameSender sends the CheckName request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerClient) CheckNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckNameResponder handles the response to the CheckName request. The method always
// closes the http.Response Body.
func (client LoadBalancerClient) CheckNameResponder(resp *http.Response) (result LoadBalancerCheckNameParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
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

// ServerInstance 로드밸런서 적용 서버 정보
// Parameters:
// vpcNo - VPC 번호
// layerTypeCode - VPC 번호
func (client LoadBalancerClient) ServerInstance(ctx context.Context, vpcNo string, layerTypeCode string) (result LoadBalancerServerInstanceListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerClient.ServerInstance")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ServerInstancePreparer(ctx, vpcNo, layerTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ServerInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ServerInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ServerInstance", resp, "Failure sending request")
		return
	}

	result, err = client.ServerInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ServerInstance", resp, "Failure responding to request")
	}

	return
}

// ServerInstancePreparer prepares the ServerInstance request.
func (client LoadBalancerClient) ServerInstancePreparer(ctx context.Context, vpcNo string, layerTypeCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"layerTypeCode": autorest.Encode("query", layerTypeCode),
		"vpcNo":         autorest.Encode("query", vpcNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/server-instances"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ServerInstanceSender sends the ServerInstance request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerClient) ServerInstanceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ServerInstanceResponder handles the response to the ServerInstance request. The method always
// closes the http.Response Body.
func (client LoadBalancerClient) ServerInstanceResponder(resp *http.Response) (result LoadBalancerServerInstanceListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ZoneSubnet 로드밸런서 금융존 서브넷
// Parameters:
// vpcNo - VPC 번호
func (client LoadBalancerClient) ZoneSubnet(ctx context.Context, vpcNo string) (result LoadBalancerZoneSubnetParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerClient.ZoneSubnet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ZoneSubnetPreparer(ctx, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ZoneSubnet", nil, "Failure preparing request")
		return
	}

	resp, err := client.ZoneSubnetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ZoneSubnet", resp, "Failure sending request")
		return
	}

	result, err = client.ZoneSubnetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerClient", "ZoneSubnet", resp, "Failure responding to request")
	}

	return
}

// ZoneSubnetPreparer prepares the ZoneSubnet request.
func (client LoadBalancerClient) ZoneSubnetPreparer(ctx context.Context, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"vpcNo": autorest.Encode("query", vpcNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/zone-subnets"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ZoneSubnetSender sends the ZoneSubnet request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerClient) ZoneSubnetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ZoneSubnetResponder handles the response to the ZoneSubnet request. The method always
// closes the http.Response Body.
func (client LoadBalancerClient) ZoneSubnetResponder(resp *http.Response) (result LoadBalancerZoneSubnetParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
