package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LoginKeyClient is the server Client
type LoginKeyClient struct {
	BaseClient
}

// NewLoginKeyClient creates an instance of the LoginKeyClient client.
func NewLoginKeyClient() LoginKeyClient {
	return NewLoginKeyClientWithBaseURI(DefaultBaseURI)
}

// NewLoginKeyClientWithBaseURI creates an instance of the LoginKeyClient client.
func NewLoginKeyClientWithBaseURI(baseURI string) LoginKeyClient {
	return LoginKeyClient{NewWithBaseURI(baseURI)}
}

// Create 로그인 키를 생성
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// keyName - 키 이름
func (client LoginKeyClient) Create(ctx context.Context, responseFormatType string, keyName string) (result CreateLoginKeyResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoginKeyClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, responseFormatType, keyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client LoginKeyClient) CreatePreparer(ctx context.Context, responseFormatType string, keyName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(keyName) > 0 {
		queryParameters["keyName"] = autorest.Encode("query", keyName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createLoginKey"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client LoginKeyClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client LoginKeyClient) CreateResponder(resp *http.Response) (result CreateLoginKeyResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 로그인 키를 삭제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// keyNameListN - 키 이름 리스트
func (client LoginKeyClient) Delete(ctx context.Context, responseFormatType string, keyNameListN string) (result DeleteLoginKeysResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoginKeyClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, responseFormatType, keyNameListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LoginKeyClient) DeletePreparer(ctx context.Context, responseFormatType string, keyNameListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"keyNameList.N":      autorest.Encode("query", keyNameListN),
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteLoginKeys"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LoginKeyClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LoginKeyClient) DeleteResponder(resp *http.Response) (result DeleteLoginKeysResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 로그인 키 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
func (client LoginKeyClient) GetList(ctx context.Context, responseFormatType string, pageNo string, pageSize string) (result LoginKeyListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoginKeyClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, pageNo, pageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client LoginKeyClient) GetListPreparer(ctx context.Context, responseFormatType string, pageNo string, pageSize string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getLoginKeyList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client LoginKeyClient) GetListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client LoginKeyClient) GetListResponder(resp *http.Response) (result LoginKeyListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Import 사용자가 생성한 로그인 키를 Import
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// publicKey - 공개키
// keyName - 키 이름
func (client LoginKeyClient) Import(ctx context.Context, responseFormatType string, publicKey string, keyName string) (result LoginKeyResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoginKeyClient.Import")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ImportPreparer(ctx, responseFormatType, publicKey, keyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Import", nil, "Failure preparing request")
		return
	}

	resp, err := client.ImportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Import", resp, "Failure sending request")
		return
	}

	result, err = client.ImportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.LoginKeyClient", "Import", resp, "Failure responding to request")
	}

	return
}

// ImportPreparer prepares the Import request.
func (client LoginKeyClient) ImportPreparer(ctx context.Context, responseFormatType string, publicKey string, keyName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"publicKey":          autorest.Encode("query", publicKey),
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(keyName) > 0 {
		queryParameters["keyName"] = autorest.Encode("query", keyName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/importLoginKey"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImportSender sends the Import request. The method will close the
// http.Response Body if it receives an error.
func (client LoginKeyClient) ImportSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ImportResponder handles the response to the Import request. The method always
// closes the http.Response Body.
func (client LoginKeyClient) ImportResponder(resp *http.Response) (result LoginKeyResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
