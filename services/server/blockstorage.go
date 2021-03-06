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

// BlockStorageClient is the server Client
type BlockStorageClient struct {
	BaseClient
}

// NewBlockStorageClient creates an instance of the BlockStorageClient client.
func NewBlockStorageClient() BlockStorageClient {
	return NewBlockStorageClientWithBaseURI(DefaultBaseURI)
}

func NewBlockStorageClientWithKey(accessKey string, secretKey string) BlockStorageClient {
	return NewBlockStorageClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewBlockStorageClientWithBaseURI creates an instance of the BlockStorageClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBlockStorageClientWithBaseURI(baseURI string) BlockStorageClient {
	return BlockStorageClient{NewWithBaseURI(baseURI)}
}

func NewBlockStorageClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) BlockStorageClient {
	return BlockStorageClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Attach 블록스토리지 인스턴스를 할당
// Parameters:
// blockStorageInstanceNo - 블록스토리지 인스턴스 번호
// serverInstanceNo - 서버 인스턴스 번호
func (client BlockStorageClient) Attach(ctx context.Context, blockStorageInstanceNo string, serverInstanceNo string) (result BlockStorageInstanceAttachResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.Attach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachPreparer(ctx, blockStorageInstanceNo, serverInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Attach", nil, "Failure preparing request")
		return
	}

	resp, err := client.AttachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Attach", resp, "Failure sending request")
		return
	}

	result, err = client.AttachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Attach", resp, "Failure responding to request")
	}

	return
}

// AttachPreparer prepares the Attach request.
func (client BlockStorageClient) AttachPreparer(ctx context.Context, blockStorageInstanceNo string, serverInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNo": autorest.Encode("query", blockStorageInstanceNo),
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNo":       autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/attachBlockStorageInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/attachBlockStorageInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachSender sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) AttachSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AttachResponder handles the response to the Attach request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) AttachResponder(resp *http.Response) (result BlockStorageInstanceAttachResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ChangeVolumeSize 블록스토리지 볼륨 사이즈를 변경
// Parameters:
// blockStorageInstanceNo - 블록스토리지 인스턴스 번호
// blockStorageSize - 블록스토리지 사이즈
func (client BlockStorageClient) ChangeVolumeSize(ctx context.Context, blockStorageInstanceNo string, blockStorageSize string) (result BlockStorageVolumeSizeResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.ChangeVolumeSize")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ChangeVolumeSizePreparer(ctx, blockStorageInstanceNo, blockStorageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "ChangeVolumeSize", nil, "Failure preparing request")
		return
	}

	resp, err := client.ChangeVolumeSizeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "ChangeVolumeSize", resp, "Failure sending request")
		return
	}

	result, err = client.ChangeVolumeSizeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "ChangeVolumeSize", resp, "Failure responding to request")
	}

	return
}

// ChangeVolumeSizePreparer prepares the ChangeVolumeSize request.
func (client BlockStorageClient) ChangeVolumeSizePreparer(ctx context.Context, blockStorageInstanceNo string, blockStorageSize string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNo": autorest.Encode("query", blockStorageInstanceNo),
		"blockStorageSize":       autorest.Encode("query", blockStorageSize),
		"responseFormatType":     autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/changeBlockStorageVolumeSize")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/changeBlockStorageVolumeSize"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ChangeVolumeSizeSender sends the ChangeVolumeSize request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) ChangeVolumeSizeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ChangeVolumeSizeResponder handles the response to the ChangeVolumeSize request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) ChangeVolumeSizeResponder(resp *http.Response) (result BlockStorageVolumeSizeResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 블록스토리지 인스턴스를 생성
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// zoneCode - ZONE 코드
// blockStorageName - 블록스토리지 이름
// blockStorageDiskDetailTypeCode - 블록스토리지 디스크 상세 유형 코드
// blockStorageSnapshotInstanceNo - 블록스토리지 스냅샷 인스턴스 번호
// blockStorageSize - 블록스토리지 사이즈
// blockStorageDescription - 블록스토리지 설명
func (client BlockStorageClient) Create(ctx context.Context, serverInstanceNo string, zoneCode string, blockStorageName string, blockStorageDiskDetailTypeCode BlockStorageDiskDetailTypeCode, blockStorageSnapshotInstanceNo string, blockStorageSize string, blockStorageDescription string) (result BlockStorageInstanceCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, serverInstanceNo, zoneCode, blockStorageName, blockStorageDiskDetailTypeCode, blockStorageSnapshotInstanceNo, blockStorageSize, blockStorageDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client BlockStorageClient) CreatePreparer(ctx context.Context, serverInstanceNo string, zoneCode string, blockStorageName string, blockStorageDiskDetailTypeCode BlockStorageDiskDetailTypeCode, blockStorageSnapshotInstanceNo string, blockStorageSize string, blockStorageDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}
	if len(blockStorageName) > 0 {
		queryParameters["blockStorageName"] = autorest.Encode("query", blockStorageName)
	}
	if len(string(blockStorageDiskDetailTypeCode)) > 0 {
		queryParameters["blockStorageDiskDetailTypeCode"] = autorest.Encode("query", blockStorageDiskDetailTypeCode)
	}
	if len(blockStorageSnapshotInstanceNo) > 0 {
		queryParameters["blockStorageSnapshotInstanceNo"] = autorest.Encode("query", blockStorageSnapshotInstanceNo)
	}
	if len(blockStorageSize) > 0 {
		queryParameters["blockStorageSize"] = autorest.Encode("query", blockStorageSize)
	}
	if len(blockStorageDescription) > 0 {
		queryParameters["blockStorageDescription"] = autorest.Encode("query", blockStorageDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/createBlockStorageInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createBlockStorageInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) CreateResponder(resp *http.Response) (result BlockStorageInstanceCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 블록스토리지 인스턴스를 삭제
// Parameters:
// blockStorageInstanceNoListN - 블록스토리지 인스턴스 번호 리스트
func (client BlockStorageClient) Delete(ctx context.Context, blockStorageInstanceNoListN string) (result BlockStorageInstanceDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, blockStorageInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BlockStorageClient) DeletePreparer(ctx context.Context, blockStorageInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNoList.N": autorest.Encode("query", blockStorageInstanceNoListN),
		"responseFormatType":           autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/deleteBlockStorageInstances")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteBlockStorageInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) DeleteResponder(resp *http.Response) (result BlockStorageInstanceDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Detach 블록스토리지 인스턴스를 할당 해제
// Parameters:
// blockStorageInstanceNoListN - 블록스토리지 인스턴스 번호 리스트
func (client BlockStorageClient) Detach(ctx context.Context, blockStorageInstanceNoListN string) (result BlockStorageInstanceDetachResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.Detach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetachPreparer(ctx, blockStorageInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Detach", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Detach", resp, "Failure sending request")
		return
	}

	result, err = client.DetachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "Detach", resp, "Failure responding to request")
	}

	return
}

// DetachPreparer prepares the Detach request.
func (client BlockStorageClient) DetachPreparer(ctx context.Context, blockStorageInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNoList.N": autorest.Encode("query", blockStorageInstanceNoListN),
		"responseFormatType":           autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/detachBlockStorageInstances")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/detachBlockStorageInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetachSender sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) DetachSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetachResponder handles the response to the Detach request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) DetachResponder(resp *http.Response) (result BlockStorageInstanceDetachResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 블록스토리지 인스턴스 상세 정보를 조회
// Parameters:
// blockStorageInstanceNo - 블록스토리지 인스턴스 번호
func (client BlockStorageClient) GetDetail(ctx context.Context, blockStorageInstanceNo string) (result BlockStorageInstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, blockStorageInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client BlockStorageClient) GetDetailPreparer(ctx context.Context, blockStorageInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNo": autorest.Encode("query", blockStorageInstanceNo),
		"responseFormatType":     autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getBlockStorageInstanceDetail")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getBlockStorageInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) GetDetailResponder(resp *http.Response) (result BlockStorageInstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 블록스토리지 인스턴스 리스트를 조회
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// blockStorageTypeCodeListN - 블록스토리지 유형 코드 리스트
// blockStorageInstanceStatusCode - 블록스토리지 인스턴스 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// blockStorageSize - 블록스토리지 사이즈
// blockStorageInstanceNoListN - 블록스토리지 인스턴스 번호 리스트
// blockStorageName - 블록스토리지 이름
// serverName - 서버 이름
// connectionInfo - 연결 정보
// blockStorageDiskTypeCode - 블록스토리지 디스크 유형 코드
// blockStorageDiskDetailTypeCode - 블록스토리지 디스크 상세 유형 코드
// zoneCode - ZONE 코드
func (client BlockStorageClient) GetList(ctx context.Context, serverInstanceNo string, blockStorageTypeCodeListN BlockStorageTypeCode, blockStorageInstanceStatusCode BlockStorageInstanceStatusCode, pageNo string, pageSize string, blockStorageSize string, blockStorageInstanceNoListN string, blockStorageName string, serverName string, connectionInfo string, blockStorageDiskTypeCode BlockStorageDiskTypeCode, blockStorageDiskDetailTypeCode BlockStorageDiskDetailTypeCode, zoneCode string) (result BlockStorageInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, serverInstanceNo, blockStorageTypeCodeListN, blockStorageInstanceStatusCode, pageNo, pageSize, blockStorageSize, blockStorageInstanceNoListN, blockStorageName, serverName, connectionInfo, blockStorageDiskTypeCode, blockStorageDiskDetailTypeCode, zoneCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.BlockStorageClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client BlockStorageClient) GetListPreparer(ctx context.Context, serverInstanceNo string, blockStorageTypeCodeListN BlockStorageTypeCode, blockStorageInstanceStatusCode BlockStorageInstanceStatusCode, pageNo string, pageSize string, blockStorageSize string, blockStorageInstanceNoListN string, blockStorageName string, serverName string, connectionInfo string, blockStorageDiskTypeCode BlockStorageDiskTypeCode, blockStorageDiskDetailTypeCode BlockStorageDiskDetailTypeCode, zoneCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(serverInstanceNo) > 0 {
		queryParameters["serverInstanceNo"] = autorest.Encode("query", serverInstanceNo)
	}
	if len(string(blockStorageTypeCodeListN)) > 0 {
		queryParameters["blockStorageTypeCodeList.N"] = autorest.Encode("query", blockStorageTypeCodeListN)
	}
	if len(string(blockStorageInstanceStatusCode)) > 0 {
		queryParameters["blockStorageInstanceStatusCode"] = autorest.Encode("query", blockStorageInstanceStatusCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(blockStorageSize) > 0 {
		queryParameters["blockStorageSize"] = autorest.Encode("query", blockStorageSize)
	}
	if len(blockStorageInstanceNoListN) > 0 {
		queryParameters["blockStorageInstanceNoList.N"] = autorest.Encode("query", blockStorageInstanceNoListN)
	}
	if len(blockStorageName) > 0 {
		queryParameters["blockStorageName"] = autorest.Encode("query", blockStorageName)
	}
	if len(serverName) > 0 {
		queryParameters["serverName"] = autorest.Encode("query", serverName)
	}
	if len(connectionInfo) > 0 {
		queryParameters["connectionInfo"] = autorest.Encode("query", connectionInfo)
	}
	if len(string(blockStorageDiskTypeCode)) > 0 {
		queryParameters["blockStorageDiskTypeCode"] = autorest.Encode("query", blockStorageDiskTypeCode)
	}
	if len(string(blockStorageDiskDetailTypeCode)) > 0 {
		queryParameters["blockStorageDiskDetailTypeCode"] = autorest.Encode("query", blockStorageDiskDetailTypeCode)
	}
	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getBlockStorageInstanceList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getBlockStorageInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) GetListResponder(resp *http.Response) (result BlockStorageInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
