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

// PublicIPClient is the server Client
type PublicIPClient struct {
	BaseClient
}

// NewPublicIPClient creates an instance of the PublicIPClient client.
func NewPublicIPClient() PublicIPClient {
	return NewPublicIPClientWithBaseURI(DefaultBaseURI)
}

// NewPublicIPClientWithBaseURI creates an instance of the PublicIPClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPublicIPClientWithBaseURI(baseURI string) PublicIPClient {
	return PublicIPClient{NewWithBaseURI(baseURI)}
}

// Associate 공인 IP를 서버 인스턴스에 할당
// Parameters:
// publicIPInstanceNo - 공인 IP 인스턴스 번호
// serverInstanceNo - 서버 인스턴스 번호
func (client PublicIPClient) Associate(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result PublicIPWithServerInstanceAssociateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.Associate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AssociatePreparer(ctx, publicIPInstanceNo, serverInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Associate", nil, "Failure preparing request")
		return
	}

	resp, err := client.AssociateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Associate", resp, "Failure sending request")
		return
	}

	result, err = client.AssociateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Associate", resp, "Failure responding to request")
	}

	return
}

// AssociatePreparer prepares the Associate request.
func (client PublicIPClient) AssociatePreparer(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"publicIpInstanceNo": autorest.Encode("query", publicIPInstanceNo),
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/associatePublicIpWithServerInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/associatePublicIpWithServerInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AssociateSender sends the Associate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) AssociateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AssociateResponder handles the response to the Associate request. The method always
// closes the http.Response Body.
func (client PublicIPClient) AssociateResponder(resp *http.Response) (result PublicIPWithServerInstanceAssociateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 공인 IP 인스턴스를 생성
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// publicIPDescription - 공인 IP 설명
func (client PublicIPClient) Create(ctx context.Context, serverInstanceNo string, publicIPDescription string) (result PublicIPInstanceCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, serverInstanceNo, publicIPDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PublicIPClient) CreatePreparer(ctx context.Context, serverInstanceNo string, publicIPDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(serverInstanceNo) > 0 {
		queryParameters["serverInstanceNo"] = autorest.Encode("query", serverInstanceNo)
	}
	if len(publicIPDescription) > 0 {
		queryParameters["publicIpDescription"] = autorest.Encode("query", publicIPDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/createPublicIpInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createPublicIpInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PublicIPClient) CreateResponder(resp *http.Response) (result PublicIPInstanceCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 공인 IP 인스턴스를 삭제
// Parameters:
// publicIPInstanceNo - 공인 IP 인스턴스 번호
func (client PublicIPClient) Delete(ctx context.Context, publicIPInstanceNo string) (result PublicIPInstanceDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, publicIPInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublicIPClient) DeletePreparer(ctx context.Context, publicIPInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"publicIpInstanceNo": autorest.Encode("query", publicIPInstanceNo),
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/deletePublicIpInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deletePublicIpInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublicIPClient) DeleteResponder(resp *http.Response) (result PublicIPInstanceDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Disassociate 공인 IP를 서버 인스턴스에서 할당 해제
// Parameters:
// publicIPInstanceNo - 공인 IP 인스턴스 번호
// serverInstanceNo - 서버 인스턴스 번호
func (client PublicIPClient) Disassociate(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result PublicIPFromServerInstanceDisassociateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.Disassociate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisassociatePreparer(ctx, publicIPInstanceNo, serverInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Disassociate", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisassociateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Disassociate", resp, "Failure sending request")
		return
	}

	result, err = client.DisassociateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "Disassociate", resp, "Failure responding to request")
	}

	return
}

// DisassociatePreparer prepares the Disassociate request.
func (client PublicIPClient) DisassociatePreparer(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(publicIPInstanceNo) > 0 {
		queryParameters["publicIpInstanceNo"] = autorest.Encode("query", publicIPInstanceNo)
	}
	if len(serverInstanceNo) > 0 {
		queryParameters["serverInstanceNo"] = autorest.Encode("query", serverInstanceNo)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/disassociatePublicIpFromServerInstance")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/disassociatePublicIpFromServerInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisassociateSender sends the Disassociate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) DisassociateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DisassociateResponder handles the response to the Disassociate request. The method always
// closes the http.Response Body.
func (client PublicIPClient) DisassociateResponder(resp *http.Response) (result PublicIPFromServerInstanceDisassociateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 공인 IP 인스턴스 상세 정보를 조회
// Parameters:
// publicIPInstanceNo - 공인 IP 인스턴스 번호
func (client PublicIPClient) GetDetail(ctx context.Context, publicIPInstanceNo string) (result PublicIPInstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, publicIPInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client PublicIPClient) GetDetailPreparer(ctx context.Context, publicIPInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"publicIpInstanceNo": autorest.Encode("query", publicIPInstanceNo),
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getPublicIpInstanceDetail")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getPublicIpInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client PublicIPClient) GetDetailResponder(resp *http.Response) (result PublicIPInstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 공인 IP 인스턴스 리스트를 조회
// Parameters:
// publicIPInstanceNoListN - 공인 IP 인스턴스 번호 리스트
// publicIP - 공인 IP 주소
// privateIP - 비공인 IP 주소
// isAssociated - 할당 여부
func (client PublicIPClient) GetList(ctx context.Context, publicIPInstanceNoListN string, publicIP string, privateIP string, isAssociated *bool) (result PublicIPInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, publicIPInstanceNoListN, publicIP, privateIP, isAssociated)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client PublicIPClient) GetListPreparer(ctx context.Context, publicIPInstanceNoListN string, publicIP string, privateIP string, isAssociated *bool) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(publicIPInstanceNoListN) > 0 {
		queryParameters["publicIpInstanceNoList.N"] = autorest.Encode("query", publicIPInstanceNoListN)
	}
	if len(publicIP) > 0 {
		queryParameters["publicIp"] = autorest.Encode("query", publicIP)
	}
	if len(privateIP) > 0 {
		queryParameters["privateIp"] = autorest.Encode("query", privateIP)
	}
	if isAssociated != nil {
		queryParameters["isAssociated"] = autorest.Encode("query", *isAssociated)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getPublicIpInstanceList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getPublicIpInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client PublicIPClient) GetListResponder(resp *http.Response) (result PublicIPInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTargetList 공인 IP 할당 가능한 서버 인스턴스 리스트를 조회
// Parameters:
// publicIPInstanceNo - 공인 IP 인스턴스 번호
// serverInstanceNo - 서버 인스턴스 번호
func (client PublicIPClient) GetTargetList(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (result PublicIPTargetServerInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPClient.GetTargetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTargetListPreparer(ctx, publicIPInstanceNo, serverInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetTargetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTargetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetTargetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetTargetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.PublicIPClient", "GetTargetList", resp, "Failure responding to request")
	}

	return
}

// GetTargetListPreparer prepares the GetTargetList request.
func (client PublicIPClient) GetTargetListPreparer(ctx context.Context, publicIPInstanceNo string, serverInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"publicIpInstanceNo": autorest.Encode("query", publicIPInstanceNo),
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/getPublicIpTargetServerInstanceList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getPublicIpTargetServerInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTargetListSender sends the GetTargetList request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPClient) GetTargetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTargetListResponder handles the response to the GetTargetList request. The method always
// closes the http.Response Body.
func (client PublicIPClient) GetTargetListResponder(resp *http.Response) (result PublicIPTargetServerInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
