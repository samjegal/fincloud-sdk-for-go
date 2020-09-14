package kubernetes

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

// WorkerNodeClient is the kubernetes Client
type WorkerNodeClient struct {
	BaseClient
}

// NewWorkerNodeClient creates an instance of the WorkerNodeClient client.
func NewWorkerNodeClient() WorkerNodeClient {
	return NewWorkerNodeClientWithBaseURI(DefaultBaseURI)
}

func NewWorkerNodeClientWithKey(accessKey string, secretKey string) WorkerNodeClient {
	return NewWorkerNodeClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewWorkerNodeClientWithBaseURI creates an instance of the WorkerNodeClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkerNodeClientWithBaseURI(baseURI string) WorkerNodeClient {
	return WorkerNodeClient{NewWithBaseURI(baseURI)}
}

func NewWorkerNodeClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) WorkerNodeClient {
	return WorkerNodeClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Get 워커노드 조회
// Parameters:
// UUID - 클러스터 UUID
func (client WorkerNodeClient) Get(ctx context.Context, UUID string) (result WorkerNodeResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkerNodeClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, UUID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.WorkerNodeClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetes.WorkerNodeClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.WorkerNodeClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkerNodeClient) GetPreparer(ctx context.Context, UUID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"uuid": autorest.Encode("path", UUID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/clusters/{uuid}/nodes", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/clusters/{uuid}/nodes", pathParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkerNodeClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkerNodeClient) GetResponder(resp *http.Response) (result WorkerNodeResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
