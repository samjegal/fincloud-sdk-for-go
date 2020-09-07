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

// ACGClient is the server Client
type ACGClient struct {
	BaseClient
}

// NewACGClient creates an instance of the ACGClient client.
func NewACGClient() ACGClient {
	return NewACGClientWithBaseURI(DefaultBaseURI)
}

// NewACGClientWithBaseURI creates an instance of the ACGClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewACGClientWithBaseURI(baseURI string) ACGClient {
	return ACGClient{NewWithBaseURI(baseURI)}
}

// Create aCG를 생성
// Parameters:
// vpcNo - VPC 번호
// accessControlGroupName - ACG 이름
// accessControlGroupDescription - ACG 설명
func (client ACGClient) Create(ctx context.Context, vpcNo string, accessControlGroupName string, accessControlGroupDescription string) (result AccessControlGroupResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, vpcNo, accessControlGroupName, accessControlGroupDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ACGClient) CreatePreparer(ctx context.Context, vpcNo string, accessControlGroupName string, accessControlGroupDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(accessControlGroupName) > 0 {
		queryParameters["accessControlGroupName"] = autorest.Encode("query", accessControlGroupName)
	}
	if len(accessControlGroupDescription) > 0 {
		queryParameters["accessControlGroupDescription"] = autorest.Encode("query", accessControlGroupDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createAccessControlGroup")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ACGClient) CreateResponder(resp *http.Response) (result AccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete aCG를 삭제
// Parameters:
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
func (client ACGClient) Delete(ctx context.Context, vpcNo string, accessControlGroupNo string) (result AccessControlGroupResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vpcNo, accessControlGroupNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ACGClient) DeletePreparer(ctx context.Context, vpcNo string, accessControlGroupNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", "json"),
		"vpcNo":                autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteAccessControlGroup")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ACGClient) DeleteResponder(resp *http.Response) (result AccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail ACG 상세 정보를 조회
// Parameters:
// accessControlGroupNo - ACG 번호
func (client ACGClient) GetDetail(ctx context.Context, accessControlGroupNo string) (result AccessControlGroupDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, accessControlGroupNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client ACGClient) GetDetailPreparer(ctx context.Context, accessControlGroupNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getAccessControlGroupDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client ACGClient) GetDetailResponder(resp *http.Response) (result AccessControlGroupDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList ACG 리스트를 조회
// Parameters:
// accessControlGroupNoListN - ACG 번호 리스트
// accessControlGroupName - ACG 이름
// accessControlGroupStatusCode - ACG 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// vpcNo - VPC 번호
func (client ACGClient) GetList(ctx context.Context, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (result AccessControlGroupListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, accessControlGroupNoListN, accessControlGroupName, accessControlGroupStatusCode, pageNo, pageSize, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ACGClient) GetListPreparer(ctx context.Context, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(accessControlGroupNoListN) > 0 {
		queryParameters["accessControlGroupNoList.N"] = autorest.Encode("query", accessControlGroupNoListN)
	}
	if len(accessControlGroupName) > 0 {
		queryParameters["accessControlGroupName"] = autorest.Encode("query", accessControlGroupName)
	}
	if len(string(accessControlGroupStatusCode)) > 0 {
		queryParameters["accessControlGroupStatusCode"] = autorest.Encode("query", accessControlGroupStatusCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getAccessControlGroupList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ACGClient) GetListResponder(resp *http.Response) (result AccessControlGroupListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRuleList ACG Rule 리스트를 조회
// Parameters:
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleTypeCode - ACG Rule 유형 코드
func (client ACGClient) GetRuleList(ctx context.Context, accessControlGroupNo string, accessControlGroupRuleTypeCode AccessControlGroupRuleTypeCode) (result AccessControlGroupRuleListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGClient.GetRuleList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRuleListPreparer(ctx, accessControlGroupNo, accessControlGroupRuleTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetRuleList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRuleListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetRuleList", resp, "Failure sending request")
		return
	}

	result, err = client.GetRuleListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGClient", "GetRuleList", resp, "Failure responding to request")
	}

	return
}

// GetRuleListPreparer prepares the GetRuleList request.
func (client ACGClient) GetRuleListPreparer(ctx context.Context, accessControlGroupNo string, accessControlGroupRuleTypeCode AccessControlGroupRuleTypeCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(string(accessControlGroupRuleTypeCode)) > 0 {
		queryParameters["accessControlGroupRuleTypeCode"] = autorest.Encode("query", accessControlGroupRuleTypeCode)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/getAccessControlGroupRuleList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupRuleList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRuleListSender sends the GetRuleList request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetRuleListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetRuleListResponder handles the response to the GetRuleList request. The method always
// closes the http.Response Body.
func (client ACGClient) GetRuleListResponder(resp *http.Response) (result AccessControlGroupRuleListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
