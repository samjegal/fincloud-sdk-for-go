package server

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

// RootPasswordClient is the server Client
type RootPasswordClient struct {
	BaseClient
}

// NewRootPasswordClient creates an instance of the RootPasswordClient client.
func NewRootPasswordClient() RootPasswordClient {
	return NewRootPasswordClientWithBaseURI(DefaultBaseURI)
}

func NewRootPasswordClientWithKey(accessKey string, secretKey string) RootPasswordClient {
	return NewRootPasswordClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewRootPasswordClientWithBaseURI creates an instance of the RootPasswordClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRootPasswordClientWithBaseURI(baseURI string) RootPasswordClient {
	return RootPasswordClient{NewWithBaseURI(baseURI)}
}

func NewRootPasswordClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) RootPasswordClient {
	return RootPasswordClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Get 서버 인스턴스의 루트 패스워드를 조회
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// privateKey - 개인키
func (client RootPasswordClient) Get(ctx context.Context, serverInstanceNo string, privateKey string) (result RootPasswordResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RootPasswordClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, serverInstanceNo, privateKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RootPasswordClient) GetPreparer(ctx context.Context, serverInstanceNo string, privateKey string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(privateKey) > 0 {
		queryParameters["privateKey"] = autorest.Encode("query", privateKey)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/getRootPassword")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRootPassword"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RootPasswordClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RootPasswordClient) GetResponder(resp *http.Response) (result RootPasswordResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 서버 인스턴스 리스트의 루트 패스워드를 조회
// Parameters:
// rootPasswordServerInstanceListNserverInstanceNo - 서버 인스턴스 번호
// rootPasswordServerInstanceListNprivateKey - 개인키
func (client RootPasswordClient) GetList(ctx context.Context, rootPasswordServerInstanceListNserverInstanceNo string, rootPasswordServerInstanceListNprivateKey string) (result RootPasswordServerInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RootPasswordClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, rootPasswordServerInstanceListNserverInstanceNo, rootPasswordServerInstanceListNprivateKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client RootPasswordClient) GetListPreparer(ctx context.Context, rootPasswordServerInstanceListNserverInstanceNo string, rootPasswordServerInstanceListNprivateKey string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"rootPasswordServerInstanceList.N.serverInstanceNo": autorest.Encode("query", rootPasswordServerInstanceListNserverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(rootPasswordServerInstanceListNprivateKey) > 0 {
		queryParameters["rootPasswordServerInstanceList.N.privateKey"] = autorest.Encode("query", rootPasswordServerInstanceListNprivateKey)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/getRootPasswordServerInstanceList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRootPasswordServerInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RootPasswordClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RootPasswordClient) GetListResponder(resp *http.Response) (result RootPasswordServerInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
