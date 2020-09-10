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

// CollectorClient is the cloud Insight Client
type CollectorClient struct {
	BaseClient
}

// NewCollectorClient creates an instance of the CollectorClient client.
func NewCollectorClient() CollectorClient {
	return NewCollectorClientWithBaseURI(DefaultBaseURI)
}

// NewCollectorClientWithBaseURI creates an instance of the CollectorClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCollectorClientWithBaseURI(baseURI string) CollectorClient {
	return CollectorClient{NewWithBaseURI(baseURI)}
}

// SendMethod JSON 데이터를 Cloud Insight Collector로 보냅니다.
// Parameters:
// parameters - cloud Insight Custom 메트릭 데이터
func (client CollectorClient) SendMethod(ctx context.Context, parameters CollectorRequest) (result CollectorResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CollectorClient.SendMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SendMethodPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.CollectorClient", "SendMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.SendMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.CollectorClient", "SendMethod", resp, "Failure sending request")
		return
	}

	result, err = client.SendMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.CollectorClient", "SendMethod", resp, "Failure responding to request")
	}

	return
}

// SendMethodPreparer prepares the SendMethod request.
func (client CollectorClient) SendMethodPreparer(ctx context.Context, parameters CollectorRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_collector/real/"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_collector/real/"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SendMethodSender sends the SendMethod request. The method will close the
// http.Response Body if it receives an error.
func (client CollectorClient) SendMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SendMethodResponder handles the response to the SendMethod request. The method always
// closes the http.Response Body.
func (client CollectorClient) SendMethodResponder(resp *http.Response) (result CollectorResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
