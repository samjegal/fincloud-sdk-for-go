package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// ImageClient is the server Client
type ImageClient struct {
	BaseClient
}

// NewImageClient creates an instance of the ImageClient client.
func NewImageClient() ImageClient {
	return NewImageClientWithBaseURI(DefaultBaseURI)
}

// NewImageClientWithBaseURI creates an instance of the ImageClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewImageClientWithBaseURI(baseURI string) ImageClient {
	return ImageClient{NewWithBaseURI(baseURI)}
}

// Create 회원 서버 이미지 인스턴스를 생성
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// memberServerImageName - 회원 서버 이미지 이름
// memberServerImageDescription - 회원 서버 이미지 설명
func (client ImageClient) Create(ctx context.Context, serverInstanceNo string, memberServerImageName string, memberServerImageDescription string) (result MemberServerImageInstanceCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImageClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, serverInstanceNo, memberServerImageName, memberServerImageDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ImageClient) CreatePreparer(ctx context.Context, serverInstanceNo string, memberServerImageName string, memberServerImageDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(memberServerImageName) > 0 {
		queryParameters["memberServerImageName"] = autorest.Encode("query", memberServerImageName)
	}
	if len(memberServerImageDescription) > 0 {
		queryParameters["memberServerImageDescription"] = autorest.Encode("query", memberServerImageDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createMemberServerImageInstance")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createMemberServerImageInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ImageClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ImageClient) CreateResponder(resp *http.Response) (result MemberServerImageInstanceCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 회원 서버 이미지 인스턴스를 삭제
// Parameters:
// memberServerImageInstanceNoListN - 회원 서버 이미지 인스턴스 번호 리스트
func (client ImageClient) Delete(ctx context.Context, memberServerImageInstanceNoListN string) (result MemberServerImageInstanceDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImageClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, memberServerImageInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ImageClient) DeletePreparer(ctx context.Context, memberServerImageInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(memberServerImageInstanceNoListN) > 0 {
		queryParameters["memberServerImageInstanceNoList.N"] = autorest.Encode("query", memberServerImageInstanceNoListN)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteMemberServerImageInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteMemberServerImageInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ImageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ImageClient) DeleteResponder(resp *http.Response) (result MemberServerImageInstanceDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 회원 서버 이미지 인스턴스 상세 정보를 조회
// Parameters:
// memberServerImageInstanceNo - 회원 서버 이미지 인스턴스 번호
func (client ImageClient) GetDetail(ctx context.Context, memberServerImageInstanceNo string) (result MemberServerImageInstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImageClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, memberServerImageInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client ImageClient) GetDetailPreparer(ctx context.Context, memberServerImageInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"memberServerImageInstanceNo": autorest.Encode("query", memberServerImageInstanceNo),
		"responseFormatType":          autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getMemberServerImageInstanceDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getMemberServerImageInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client ImageClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client ImageClient) GetDetailResponder(resp *http.Response) (result MemberServerImageInstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 회원 서버 이미지 인스턴스 리스트를 조회
// Parameters:
// memberServerImageName - 회원 서버 이미지 이름
// memberServerImageInstanceStatusCode - 회원 서버 이미지 인스턴스 상태 코드
// memberServerImageInstanceNoListN - 회원 서버 이미지 인스턴스 번호 리스트
// platformTypeCodeListN - 플랫폼 유형 코드 리스트
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
func (client ImageClient) GetList(ctx context.Context, memberServerImageName string, memberServerImageInstanceStatusCode MemberServerImageInstanceStatusCode, memberServerImageInstanceNoListN string, platformTypeCodeListN PlatformTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder) (result MemberServerImageInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImageClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, memberServerImageName, memberServerImageInstanceStatusCode, memberServerImageInstanceNoListN, platformTypeCodeListN, pageNo, pageSize, sortedBy, sortingOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ImageClient) GetListPreparer(ctx context.Context, memberServerImageName string, memberServerImageInstanceStatusCode MemberServerImageInstanceStatusCode, memberServerImageInstanceNoListN string, platformTypeCodeListN PlatformTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(memberServerImageName) > 0 {
		queryParameters["memberServerImageName"] = autorest.Encode("query", memberServerImageName)
	}
	if len(string(memberServerImageInstanceStatusCode)) > 0 {
		queryParameters["memberServerImageInstanceStatusCode"] = autorest.Encode("query", memberServerImageInstanceStatusCode)
	}
	if len(memberServerImageInstanceNoListN) > 0 {
		queryParameters["memberServerImageInstanceNoList.N"] = autorest.Encode("query", memberServerImageInstanceNoListN)
	}
	if len(string(platformTypeCodeListN)) > 0 {
		queryParameters["platformTypeCodeList.N"] = autorest.Encode("query", platformTypeCodeListN)
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
	} else {
		queryParameters["sortingOrder"] = autorest.Encode("query", "ASC")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getMemberServerImageInstanceList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getMemberServerImageInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ImageClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ImageClient) GetListResponder(resp *http.Response) (result MemberServerImageInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
