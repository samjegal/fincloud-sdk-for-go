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

// Client is the server Client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// ChangeSpec 서버 인스턴스 스펙을 변경
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
// serverProductCode - 서버 상품 코드
func (client Client) ChangeSpec(ctx context.Context, serverInstanceNo string, serverProductCode string) (result InstanceSpecResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ChangeSpec")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ChangeSpecPreparer(ctx, serverInstanceNo, serverProductCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "ChangeSpec", nil, "Failure preparing request")
		return
	}

	resp, err := client.ChangeSpecSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "ChangeSpec", resp, "Failure sending request")
		return
	}

	result, err = client.ChangeSpecResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "ChangeSpec", resp, "Failure responding to request")
	}

	return
}

// ChangeSpecPreparer prepares the ChangeSpec request.
func (client Client) ChangeSpecPreparer(ctx context.Context, serverInstanceNo string, serverProductCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
		"serverProductCode":  autorest.Encode("query", serverProductCode),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/changeServerInstanceSpec")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/changeServerInstanceSpec"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ChangeSpecSender sends the ChangeSpec request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ChangeSpecSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ChangeSpecResponder handles the response to the ChangeSpec request. The method always
// closes the http.Response Body.
func (client Client) ChangeSpecResponder(resp *http.Response) (result InstanceSpecResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 서버 인스턴스를 생성
// Parameters:
// vpcNo - VPC 번호
// subnetNo - 서브넷 번호
// networkInterfaceListNnetworkInterfaceOrder - 네트워크 인터페이스 순서
// networkInterfaceListNaccessControlGroupNoListN - ACG 번호 리스트
// memberServerImageInstanceNo - 회원 서버 이미지 인스턴스 번호
// serverImageProductCode - 서버 이미지 상품 코드
// serverProductCode - 서버 상품 코드
// isEncryptedBaseBlockStorageVolume - 기본 블록스토리지 볼륨 암호화 여부
// feeSystemTypeCode - 요금제 유형 코드
// serverCreateCount - 서버 생성 개수
// serverCreateStartNo - 서버 생성 시작 번호
// serverName - 서버 이름
// networkInterfaceListNnetworkInterfaceNo - 네트워크 인터페이스 번호
// networkInterfaceListNsubnetNo - 서브넷 번호
// networkInterfaceListNip - IP 주소
// placementGroupNo - 물리 배치 그룹 번호
// isProtectServerTermination - 반납 보호 여부
// serverDescription - 서버 설명
// initScriptNo - 초기화 스크립트 번호
// loginKeyName - 로그인 키 이름
func (client Client) Create(ctx context.Context, vpcNo string, subnetNo string, networkInterfaceListNnetworkInterfaceOrder string, networkInterfaceListNaccessControlGroupNoListN string, memberServerImageInstanceNo string, serverImageProductCode string, serverProductCode string, isEncryptedBaseBlockStorageVolume *bool, feeSystemTypeCode FeeSystemTypeCode, serverCreateCount string, serverCreateStartNo string, serverName string, networkInterfaceListNnetworkInterfaceNo string, networkInterfaceListNsubnetNo string, networkInterfaceListNip string, placementGroupNo string, isProtectServerTermination *bool, serverDescription string, initScriptNo string, loginKeyName string) (result InstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, vpcNo, subnetNo, networkInterfaceListNnetworkInterfaceOrder, networkInterfaceListNaccessControlGroupNoListN, memberServerImageInstanceNo, serverImageProductCode, serverProductCode, isEncryptedBaseBlockStorageVolume, feeSystemTypeCode, serverCreateCount, serverCreateStartNo, serverName, networkInterfaceListNnetworkInterfaceNo, networkInterfaceListNsubnetNo, networkInterfaceListNip, placementGroupNo, isProtectServerTermination, serverDescription, initScriptNo, loginKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, vpcNo string, subnetNo string, networkInterfaceListNnetworkInterfaceOrder string, networkInterfaceListNaccessControlGroupNoListN string, memberServerImageInstanceNo string, serverImageProductCode string, serverProductCode string, isEncryptedBaseBlockStorageVolume *bool, feeSystemTypeCode FeeSystemTypeCode, serverCreateCount string, serverCreateStartNo string, serverName string, networkInterfaceListNnetworkInterfaceNo string, networkInterfaceListNsubnetNo string, networkInterfaceListNip string, placementGroupNo string, isProtectServerTermination *bool, serverDescription string, initScriptNo string, loginKeyName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkInterfaceList.N.accessControlGroupNoList.N": autorest.Encode("query", networkInterfaceListNaccessControlGroupNoListN),
		"networkInterfaceList.N.networkInterfaceOrder":      autorest.Encode("query", networkInterfaceListNnetworkInterfaceOrder),
		"responseFormatType":                                autorest.Encode("query", "json"),
		"subnetNo":                                          autorest.Encode("query", subnetNo),
		"vpcNo":                                             autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(memberServerImageInstanceNo) > 0 {
		queryParameters["memberServerImageInstanceNo"] = autorest.Encode("query", memberServerImageInstanceNo)
	}
	if len(serverImageProductCode) > 0 {
		queryParameters["serverImageProductCode"] = autorest.Encode("query", serverImageProductCode)
	}
	if len(serverProductCode) > 0 {
		queryParameters["serverProductCode"] = autorest.Encode("query", serverProductCode)
	}
	if isEncryptedBaseBlockStorageVolume != nil {
		queryParameters["isEncryptedBaseBlockStorageVolume"] = autorest.Encode("query", *isEncryptedBaseBlockStorageVolume)
	}
	if len(string(feeSystemTypeCode)) > 0 {
		queryParameters["feeSystemTypeCode"] = autorest.Encode("query", feeSystemTypeCode)
	}
	if len(serverCreateCount) > 0 {
		queryParameters["serverCreateCount"] = autorest.Encode("query", serverCreateCount)
	}
	if len(serverCreateStartNo) > 0 {
		queryParameters["serverCreateStartNo"] = autorest.Encode("query", serverCreateStartNo)
	}
	if len(serverName) > 0 {
		queryParameters["serverName"] = autorest.Encode("query", serverName)
	}
	if len(networkInterfaceListNnetworkInterfaceNo) > 0 {
		queryParameters["networkInterfaceList.N.networkInterfaceNo"] = autorest.Encode("query", networkInterfaceListNnetworkInterfaceNo)
	}
	if len(networkInterfaceListNsubnetNo) > 0 {
		queryParameters["networkInterfaceList.N.subnetNo"] = autorest.Encode("query", networkInterfaceListNsubnetNo)
	}
	if len(networkInterfaceListNip) > 0 {
		queryParameters["networkInterfaceList.N.ip"] = autorest.Encode("query", networkInterfaceListNip)
	}
	if len(placementGroupNo) > 0 {
		queryParameters["placementGroupNo"] = autorest.Encode("query", placementGroupNo)
	}
	if isProtectServerTermination != nil {
		queryParameters["isProtectServerTermination"] = autorest.Encode("query", *isProtectServerTermination)
	}
	if len(serverDescription) > 0 {
		queryParameters["serverDescription"] = autorest.Encode("query", serverDescription)
	}
	if len(initScriptNo) > 0 {
		queryParameters["initScriptNo"] = autorest.Encode("query", initScriptNo)
	}
	if len(loginKeyName) > 0 {
		queryParameters["loginKeyName"] = autorest.Encode("query", loginKeyName)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createServerInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createServerInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result InstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 서버 인스턴스 상세 정보를 조회
// Parameters:
// serverInstanceNo - 서버 인스턴스 번호
func (client Client) GetDetail(ctx context.Context, serverInstanceNo string) (result InstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, serverInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client Client) GetDetailPreparer(ctx context.Context, serverInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getServerInstanceDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getServerInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client Client) GetDetailResponder(resp *http.Response) (result InstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInstanceList 서버 인스턴스 리스트를 조회
// Parameters:
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
// vpcNo - VPC 번호
// pageNo - 페이지 번호
// serverInstanceStatusCode - 서버 인스턴스 상태 코드
// baseBlockStorageDiskTypeCode - 기본 블록스토리지 디스크 유형 코드
// baseBlockStorageDiskDetailTypeCode - 기본 블록스토리지 디스크 상세 유형 코드
// serverName - 서버 이름
// IP - IP 주소
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
// placementGroupNoListN - 물리 배치 그룹 번호 리스트
func (client Client) GetInstanceList(ctx context.Context, serverInstanceNoListN string, vpcNo string, pageNo string, serverInstanceStatusCode string, baseBlockStorageDiskTypeCode string, baseBlockStorageDiskDetailTypeCode string, serverName string, IP string, sortedBy string, sortingOrder string, placementGroupNoListN string) (result InstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetInstanceList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetInstanceListPreparer(ctx, serverInstanceNoListN, vpcNo, pageNo, serverInstanceStatusCode, baseBlockStorageDiskTypeCode, baseBlockStorageDiskDetailTypeCode, serverName, IP, sortedBy, sortingOrder, placementGroupNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "GetInstanceList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInstanceListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "GetInstanceList", resp, "Failure sending request")
		return
	}

	result, err = client.GetInstanceListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "GetInstanceList", resp, "Failure responding to request")
	}

	return
}

// GetInstanceListPreparer prepares the GetInstanceList request.
func (client Client) GetInstanceListPreparer(ctx context.Context, serverInstanceNoListN string, vpcNo string, pageNo string, serverInstanceStatusCode string, baseBlockStorageDiskTypeCode string, baseBlockStorageDiskDetailTypeCode string, serverName string, IP string, sortedBy string, sortingOrder string, placementGroupNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}

	queryParameters["responseFormatType"] = autorest.Encode("query", "json")
	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(serverInstanceNoListN) > 0 {
		queryParameters["serverInstanceNoList.N"] = autorest.Encode("query", serverInstanceNoListN)
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(serverInstanceStatusCode) > 0 {
		queryParameters["serverInstanceStatusCode"] = autorest.Encode("query", serverInstanceStatusCode)
	}
	if len(baseBlockStorageDiskTypeCode) > 0 {
		queryParameters["baseBlockStorageDiskTypeCode"] = autorest.Encode("query", baseBlockStorageDiskTypeCode)
	}
	if len(baseBlockStorageDiskDetailTypeCode) > 0 {
		queryParameters["baseBlockStorageDiskDetailTypeCode"] = autorest.Encode("query", baseBlockStorageDiskDetailTypeCode)
	}
	if len(serverName) > 0 {
		queryParameters["serverName"] = autorest.Encode("query", serverName)
	}
	if len(IP) > 0 {
		queryParameters["ip"] = autorest.Encode("query", IP)
	}
	if len(sortedBy) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(sortingOrder) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	} else {
		queryParameters["sortingOrder"] = autorest.Encode("query", "ASC")
	}
	if len(placementGroupNoListN) > 0 {
		queryParameters["placementGroupNoList.N"] = autorest.Encode("query", placementGroupNoListN)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getServerInstanceList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getServerInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInstanceListSender sends the GetInstanceList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetInstanceListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInstanceListResponder handles the response to the GetInstanceList request. The method always
// closes the http.Response Body.
func (client Client) GetInstanceListResponder(resp *http.Response) (result InstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reboot 서버 인스턴스를 재시작
// Parameters:
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) Reboot(ctx context.Context, serverInstanceNoListN string) (result InstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Reboot")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RebootPreparer(ctx, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Reboot", nil, "Failure preparing request")
		return
	}

	resp, err := client.RebootSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "Reboot", resp, "Failure sending request")
		return
	}

	result, err = client.RebootResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Reboot", resp, "Failure responding to request")
	}

	return
}

// RebootPreparer prepares the Reboot request.
func (client Client) RebootPreparer(ctx context.Context, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/rebootServerInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/rebootServerInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RebootSender sends the Reboot request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RebootSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RebootResponder handles the response to the Reboot request. The method always
// closes the http.Response Body.
func (client Client) RebootResponder(resp *http.Response) (result InstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start 서버 인스턴스를 시작
// Parameters:
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) Start(ctx context.Context, serverInstanceNoListN string) (result InstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Start")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client Client) StartPreparer(ctx context.Context, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/startServerInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/startServerInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client Client) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client Client) StartResponder(resp *http.Response) (result InstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop 서버 인스턴스를 종료
// Parameters:
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) Stop(ctx context.Context, serverInstanceNoListN string) (result InstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Stop")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client Client) StopPreparer(ctx context.Context, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/stopServerInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/stopServerInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client Client) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client Client) StopResponder(resp *http.Response) (result InstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Terminate 서버 인스턴스를 반납
// Parameters:
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) Terminate(ctx context.Context, serverInstanceNoListN string) (result InstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Terminate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TerminatePreparer(ctx, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Terminate", nil, "Failure preparing request")
		return
	}

	resp, err := client.TerminateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.Client", "Terminate", resp, "Failure sending request")
		return
	}

	result, err = client.TerminateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.Client", "Terminate", resp, "Failure responding to request")
	}

	return
}

// TerminatePreparer prepares the Terminate request.
func (client Client) TerminatePreparer(ctx context.Context, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/terminateServerInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/terminateServerInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TerminateSender sends the Terminate request. The method will close the
// http.Response Body if it receives an error.
func (client Client) TerminateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// TerminateResponder handles the response to the Terminate request. The method always
// closes the http.Response Body.
func (client Client) TerminateResponder(resp *http.Response) (result InstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
