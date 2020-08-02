package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RootPasswordClient is the server Client
type RootPasswordClient struct {
	BaseClient
}

// NewRootPasswordClient creates an instance of the RootPasswordClient client.
func NewRootPasswordClient() RootPasswordClient {
	return NewRootPasswordClientWithBaseURI(DefaultBaseURI)
}

// NewRootPasswordClientWithBaseURI creates an instance of the RootPasswordClient client.
func NewRootPasswordClientWithBaseURI(baseURI string) RootPasswordClient {
	return RootPasswordClient{NewWithBaseURI(baseURI)}
}

// Get 서버 인스턴스의 루트 패스워드를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// serverInstanceNo - 서버 인스턴스 번호
// regionCode - REGION 코드
// privateKey - 개인키
func (client RootPasswordClient) Get(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string, privateKey string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RootPasswordClient.Get")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, responseFormatType, serverInstanceNo, regionCode, privateKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = resp
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
func (client RootPasswordClient) GetPreparer(ctx context.Context, responseFormatType string, serverInstanceNo string, regionCode string, privateKey string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(privateKey) > 0 {
		queryParameters["privateKey"] = autorest.Encode("query", privateKey)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRootPassword"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RootPasswordClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RootPasswordClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetList 서버 인스턴스 리스트의 루트 패스워드를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// rootPasswordServerInstanceListNserverInstanceNo - 서버 인스턴스 번호
// regionCode - REGION 코드
// rootPasswordServerInstanceListNprivateKey - 개인키
func (client RootPasswordClient) GetList(ctx context.Context, responseFormatType string, rootPasswordServerInstanceListNserverInstanceNo string, regionCode string, rootPasswordServerInstanceListNprivateKey string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RootPasswordClient.GetList")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, rootPasswordServerInstanceListNserverInstanceNo, regionCode, rootPasswordServerInstanceListNprivateKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.RootPasswordClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = resp
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
func (client RootPasswordClient) GetListPreparer(ctx context.Context, responseFormatType string, rootPasswordServerInstanceListNserverInstanceNo string, regionCode string, rootPasswordServerInstanceListNprivateKey string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"rootPasswordServerInstanceList.N.serverInstanceNo": autorest.Encode("query", rootPasswordServerInstanceListNserverInstanceNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(rootPasswordServerInstanceListNprivateKey) > 0 {
		queryParameters["rootPasswordServerInstanceList.N.privateKey"] = autorest.Encode("query", rootPasswordServerInstanceListNprivateKey)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRootPasswordServerInstanceList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RootPasswordClient) GetListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RootPasswordClient) GetListResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
