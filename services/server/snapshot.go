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

// SnapshotClient is the server Client
type SnapshotClient struct {
	BaseClient
}

// NewSnapshotClient creates an instance of the SnapshotClient client.
func NewSnapshotClient() SnapshotClient {
	return NewSnapshotClientWithBaseURI(DefaultBaseURI)
}

func NewSnapshotClientWithKey(accessKey string, secretKey string) SnapshotClient {
	return NewSnapshotClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewSnapshotClientWithBaseURI creates an instance of the SnapshotClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSnapshotClientWithBaseURI(baseURI string) SnapshotClient {
	return SnapshotClient{NewWithBaseURI(baseURI)}
}

func NewSnapshotClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) SnapshotClient {
	return SnapshotClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Create 블록스토리지 스냅샷 인스턴스를 생성
// Parameters:
// originalBlockStorageInstanceNo - 원본 블록스토리지 인스턴스 번호
// blockStorageSnapshotName - 블록스토리지 스냅샷 이름
// blockStorageSnapshotDescription - 블록스토리지 스냅샷 설명
func (client SnapshotClient) Create(ctx context.Context, originalBlockStorageInstanceNo string, blockStorageSnapshotName string, blockStorageSnapshotDescription string) (result BlockStorageSnapshotInstanceCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, originalBlockStorageInstanceNo, blockStorageSnapshotName, blockStorageSnapshotDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SnapshotClient) CreatePreparer(ctx context.Context, originalBlockStorageInstanceNo string, blockStorageSnapshotName string, blockStorageSnapshotDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"originalBlockStorageInstanceNo": autorest.Encode("query", originalBlockStorageInstanceNo),
		"responseFormatType":             autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(blockStorageSnapshotName) > 0 {
		queryParameters["blockStorageSnapshotName"] = autorest.Encode("query", blockStorageSnapshotName)
	}
	if len(blockStorageSnapshotDescription) > 0 {
		queryParameters["blockStorageSnapshotDescription"] = autorest.Encode("query", blockStorageSnapshotDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/createBlockStorageSnapshotInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createBlockStorageSnapshotInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SnapshotClient) CreateResponder(resp *http.Response) (result BlockStorageSnapshotInstanceCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 블록스토리지 스냅샷 인스턴스 번호 리스트
// Parameters:
// blockStorageSnapshotInstanceNoListN - 블록스토리지 스냅샷 인스턴스 번호 리스트
func (client SnapshotClient) Delete(ctx context.Context, blockStorageSnapshotInstanceNoListN string) (result BlockStorageSnapshotInstanceDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, blockStorageSnapshotInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SnapshotClient) DeletePreparer(ctx context.Context, blockStorageSnapshotInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageSnapshotInstanceNoList.N": autorest.Encode("query", blockStorageSnapshotInstanceNoListN),
		"responseFormatType":                   autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/deleteBlockStorageSnapshotInstances")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteBlockStorageSnapshotInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SnapshotClient) DeleteResponder(resp *http.Response) (result BlockStorageSnapshotInstanceDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 블록스토리지 스냅샷 인스턴스 상세 정보를 조회
// Parameters:
// blockStorageSnapshotInstanceNo - 블록스토리지 스냅샷 인스턴스 번호
func (client SnapshotClient) GetDetail(ctx context.Context, blockStorageSnapshotInstanceNo string) (result BlockStorageSnapshotInstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, blockStorageSnapshotInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client SnapshotClient) GetDetailPreparer(ctx context.Context, blockStorageSnapshotInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageSnapshotInstanceNo": autorest.Encode("query", blockStorageSnapshotInstanceNo),
		"responseFormatType":             autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getBlockStorageSnapshotInstanceDetail")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getBlockStorageSnapshotInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client SnapshotClient) GetDetailResponder(resp *http.Response) (result BlockStorageSnapshotInstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 블록스토리지 스냅샷 인스턴스 리스트를 조회
// Parameters:
// blockStorageSnapshotInstanceNoListN - 블록스토리지 스냅샷 인스턴스 번호 리스트
// originalBlockStorageInstanceNoListN - 원본 블록스토리지 인스턴스 번호 리스트
// blockStorageSnapshotInstanceStatusCode - 블록스토리지 스냅샷 인스턴스 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// blockStorageSnapshotVolumeSize - 블록스토리지 스냅샷 블록 사이즈
// isEncryptedOriginalBlockStorageVolume - 원본 블록스토리지 볼륨 암호화 여부
// blockStorageSnapshotName - 블록스토리지 스냅샷 이름
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
func (client SnapshotClient) GetList(ctx context.Context, blockStorageSnapshotInstanceNoListN string, originalBlockStorageInstanceNoListN string, blockStorageSnapshotInstanceStatusCode BlockStorageSnapshotInstanceStatusCode, pageNo string, pageSize string, blockStorageSnapshotVolumeSize string, isEncryptedOriginalBlockStorageVolume *bool, blockStorageSnapshotName string, sortedBy string, sortingOrder SortingOrder) (result BlockStorageSnapshotInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SnapshotClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, blockStorageSnapshotInstanceNoListN, originalBlockStorageInstanceNoListN, blockStorageSnapshotInstanceStatusCode, pageNo, pageSize, blockStorageSnapshotVolumeSize, isEncryptedOriginalBlockStorageVolume, blockStorageSnapshotName, sortedBy, sortingOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.SnapshotClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client SnapshotClient) GetListPreparer(ctx context.Context, blockStorageSnapshotInstanceNoListN string, originalBlockStorageInstanceNoListN string, blockStorageSnapshotInstanceStatusCode BlockStorageSnapshotInstanceStatusCode, pageNo string, pageSize string, blockStorageSnapshotVolumeSize string, isEncryptedOriginalBlockStorageVolume *bool, blockStorageSnapshotName string, sortedBy string, sortingOrder SortingOrder) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(blockStorageSnapshotInstanceNoListN) > 0 {
		queryParameters["blockStorageSnapshotInstanceNoList.N"] = autorest.Encode("query", blockStorageSnapshotInstanceNoListN)
	}
	if len(originalBlockStorageInstanceNoListN) > 0 {
		queryParameters["originalBlockStorageInstanceNoList.N"] = autorest.Encode("query", originalBlockStorageInstanceNoListN)
	}
	if len(string(blockStorageSnapshotInstanceStatusCode)) > 0 {
		queryParameters["blockStorageSnapshotInstanceStatusCode"] = autorest.Encode("query", blockStorageSnapshotInstanceStatusCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(blockStorageSnapshotVolumeSize) > 0 {
		queryParameters["blockStorageSnapshotVolumeSize"] = autorest.Encode("query", blockStorageSnapshotVolumeSize)
	}
	if isEncryptedOriginalBlockStorageVolume != nil {
		queryParameters["isEncryptedOriginalBlockStorageVolume"] = autorest.Encode("query", *isEncryptedOriginalBlockStorageVolume)
	}
	if len(blockStorageSnapshotName) > 0 {
		queryParameters["blockStorageSnapshotName"] = autorest.Encode("query", blockStorageSnapshotName)
	}
	if len(sortedBy) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(string(sortingOrder)) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getBlockStorageSnapshotInstanceList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getBlockStorageSnapshotInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client SnapshotClient) GetListResponder(resp *http.Response) (result BlockStorageSnapshotInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
