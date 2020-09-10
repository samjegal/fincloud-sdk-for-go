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

// ServerClient is the cloud Insight Client
type ServerClient struct {
	BaseClient
}

// NewServerClient creates an instance of the ServerClient client.
func NewServerClient() ServerClient {
	return NewServerClientWithBaseURI(DefaultBaseURI)
}

// NewServerClientWithBaseURI creates an instance of the ServerClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServerClientWithBaseURI(baseURI string) ServerClient {
	return ServerClient{NewWithBaseURI(baseURI)}
}

// GetTop 사용자의 Server 중 CPU, Memory, File system 별 사용량 top5에 해당하는 server를 조회합니다.
// Parameters:
// query - target metric (mem_usert/avg_cpu_user_rto/fs_usert)
func (client ServerClient) GetTop(ctx context.Context, query SeverTargetMetric) (result ListServerTopMetricParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerClient.GetTop")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTopPreparer(ctx, query)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ServerClient", "GetTop", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTopSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.ServerClient", "GetTop", resp, "Failure sending request")
		return
	}

	result, err = client.GetTopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ServerClient", "GetTop", resp, "Failure responding to request")
	}

	return
}

// GetTopPreparer prepares the GetTop request.
func (client ServerClient) GetTopPreparer(ctx context.Context, query SeverTargetMetric) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"query": autorest.Encode("query", query),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/servers/top")+"?"+common.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/servers/top"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTopSender sends the GetTop request. The method will close the
// http.Response Body if it receives an error.
func (client ServerClient) GetTopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTopResponder handles the response to the GetTop request. The method always
// closes the http.Response Body.
func (client ServerClient) GetTopResponder(resp *http.Response) (result ListServerTopMetricParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
