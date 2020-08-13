package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InitScriptClient is the server Client
type InitScriptClient struct {
	BaseClient
}

// NewInitScriptClient creates an instance of the InitScriptClient client.
func NewInitScriptClient() InitScriptClient {
	return NewInitScriptClientWithBaseURI(DefaultBaseURI)
}

// NewInitScriptClientWithBaseURI creates an instance of the InitScriptClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInitScriptClientWithBaseURI(baseURI string) InitScriptClient {
	return InitScriptClient{NewWithBaseURI(baseURI)}
}

// Create 초기화 스크립트를 생성
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// initScriptContent - 초기화 스크립트 내용
// regionCode - REGION 코드
// osTypeCode - OS 유형 코드
// initScriptName - 초기화 스크립트 이름
// initScriptDescription - 초기화 스크립트 설명
func (client InitScriptClient) Create(ctx context.Context, responseFormatType string, initScriptContent string, regionCode string, osTypeCode OsTypeCode, initScriptName string, initScriptDescription string) (result InitScriptResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InitScriptClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, responseFormatType, initScriptContent, regionCode, osTypeCode, initScriptName, initScriptDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client InitScriptClient) CreatePreparer(ctx context.Context, responseFormatType string, initScriptContent string, regionCode string, osTypeCode OsTypeCode, initScriptName string, initScriptDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"initScriptContent":  autorest.Encode("query", initScriptContent),
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(string(osTypeCode)) > 0 {
		queryParameters["osTypeCode"] = autorest.Encode("query", osTypeCode)
	}
	if len(initScriptName) > 0 {
		queryParameters["initScriptName"] = autorest.Encode("query", initScriptName)
	}
	if len(initScriptDescription) > 0 {
		queryParameters["initScriptDescription"] = autorest.Encode("query", initScriptDescription)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/createInitScript"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client InitScriptClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client InitScriptClient) CreateResponder(resp *http.Response) (result InitScriptResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 초기화 스크립트를 삭제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// initScriptNoListN - 초기화 스크립트 번호 리스트
func (client InitScriptClient) Delete(ctx context.Context, responseFormatType string, regionCode string, initScriptNoListN string) (result InitScriptResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InitScriptClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, responseFormatType, regionCode, initScriptNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InitScriptClient) DeletePreparer(ctx context.Context, responseFormatType string, regionCode string, initScriptNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(initScriptNoListN) > 0 {
		queryParameters["initScriptNoList.N"] = autorest.Encode("query", initScriptNoListN)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/deleteInitScripts"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InitScriptClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InitScriptClient) DeleteResponder(resp *http.Response) (result InitScriptResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 초기화 스크립트 상세 정보를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// initScriptNo - 초기화 스크립트 번호
func (client InitScriptClient) GetDetail(ctx context.Context, responseFormatType string, regionCode string, initScriptNo string) (result InitScriptDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InitScriptClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, responseFormatType, regionCode, initScriptNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client InitScriptClient) GetDetailPreparer(ctx context.Context, responseFormatType string, regionCode string, initScriptNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(initScriptNo) > 0 {
		queryParameters["initScriptNo"] = autorest.Encode("query", initScriptNo)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/getInitScriptDetail"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client InitScriptClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client InitScriptClient) GetDetailResponder(resp *http.Response) (result InitScriptDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 초기화 스크립트 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// osTypeCode - OS 유형 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
// initScriptName - 초기화 스크립트 이름
// initScriptNoListN - 초기화 스크립트 번호 리스트
func (client InitScriptClient) GetList(ctx context.Context, responseFormatType string, regionCode string, osTypeCode OsTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder, initScriptName string, initScriptNoListN string) (result InitScriptListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InitScriptClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, regionCode, osTypeCode, pageNo, pageSize, sortedBy, sortingOrder, initScriptName, initScriptNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.InitScriptClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client InitScriptClient) GetListPreparer(ctx context.Context, responseFormatType string, regionCode string, osTypeCode OsTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder, initScriptName string, initScriptNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(string(osTypeCode)) > 0 {
		queryParameters["osTypeCode"] = autorest.Encode("query", osTypeCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(sortedBy) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(string(sortingOrder)) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	}
	if len(initScriptName) > 0 {
		queryParameters["initScriptName"] = autorest.Encode("query", initScriptName)
	}
	if len(initScriptNoListN) > 0 {
		queryParameters["initScriptNoList.N"] = autorest.Encode("query", initScriptNoListN)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/getInitScriptList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client InitScriptClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client InitScriptClient) GetListResponder(resp *http.Response) (result InitScriptListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
