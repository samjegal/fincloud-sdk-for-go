package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NetworkInterfaceClient is the server Client
type NetworkInterfaceClient struct {
	BaseClient
}

// NewNetworkInterfaceClient creates an instance of the NetworkInterfaceClient client.
func NewNetworkInterfaceClient() NetworkInterfaceClient {
	return NewNetworkInterfaceClientWithBaseURI(DefaultBaseURI)
}

// NewNetworkInterfaceClientWithBaseURI creates an instance of the NetworkInterfaceClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewNetworkInterfaceClientWithBaseURI(baseURI string) NetworkInterfaceClient {
	return NetworkInterfaceClient{NewWithBaseURI(baseURI)}
}

// AddACG 네트워크 인터페이스의 ACG를 추가
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// networkInterfaceNo - 네트워크 인터페이스 번호
// accessControlGroupNoListN - ACG 번호 리스트
// regionCode - REGION 코드
func (client NetworkInterfaceClient) AddACG(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (result NetworkInterfaceAccessControlGroupResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.AddACG")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddACGPreparer(ctx, responseFormatType, networkInterfaceNo, accessControlGroupNoListN, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "AddACG", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddACGSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "AddACG", resp, "Failure sending request")
		return
	}

	result, err = client.AddACGResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "AddACG", resp, "Failure responding to request")
	}

	return
}

// AddACGPreparer prepares the AddACG request.
func (client NetworkInterfaceClient) AddACGPreparer(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNoList.N": autorest.Encode("query", accessControlGroupNoListN),
		"networkInterfaceNo":         autorest.Encode("query", networkInterfaceNo),
		"responseFormatType":         autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/addNetworkInterfaceAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddACGSender sends the AddACG request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) AddACGSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddACGResponder handles the response to the AddACG request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) AddACGResponder(resp *http.Response) (result NetworkInterfaceAccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Attach 네트워크 인터페이스를 할당
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// subnetNo - 서브넷 번호
// networkInterfaceNo - 네트워크 인터페이스 번호
// serverInstanceNo - 서버 인스턴스 번호
// regionCode - REGION 코드
func (client NetworkInterfaceClient) Attach(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (result NetworkInterfaceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Attach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachPreparer(ctx, responseFormatType, subnetNo, networkInterfaceNo, serverInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Attach", nil, "Failure preparing request")
		return
	}

	resp, err := client.AttachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Attach", resp, "Failure sending request")
		return
	}

	result, err = client.AttachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Attach", resp, "Failure responding to request")
	}

	return
}

// AttachPreparer prepares the Attach request.
func (client NetworkInterfaceClient) AttachPreparer(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkInterfaceNo": autorest.Encode("query", networkInterfaceNo),
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
		"subnetNo":           autorest.Encode("query", subnetNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/attachNetworkInterface"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachSender sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) AttachSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AttachResponder handles the response to the Attach request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) AttachResponder(resp *http.Response) (result NetworkInterfaceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 네트워크 인터페이스를 생성
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// vpcNo - VPC 번호
// subnetNo - 서브넷 번호
// networkInterfaceName - 네트워크 인터페이스 이름
// accessControlGroupNoListN - ACG 번호 리스트
// serverInstanceNo - 서버 인스턴스 번호
// IP - IP 주소
// networkInterfaceDescription - 네트워크 인터페이스 설명
func (client NetworkInterfaceClient) Create(ctx context.Context, responseFormatType string, regionCode string, vpcNo string, subnetNo string, networkInterfaceName string, accessControlGroupNoListN string, serverInstanceNo string, IP string, networkInterfaceDescription string) (result NetworkInterfaceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, responseFormatType, regionCode, vpcNo, subnetNo, networkInterfaceName, accessControlGroupNoListN, serverInstanceNo, IP, networkInterfaceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client NetworkInterfaceClient) CreatePreparer(ctx context.Context, responseFormatType string, regionCode string, vpcNo string, subnetNo string, networkInterfaceName string, accessControlGroupNoListN string, serverInstanceNo string, IP string, networkInterfaceDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}
	if len(subnetNo) > 0 {
		queryParameters["subnetNo"] = autorest.Encode("query", subnetNo)
	}
	if len(networkInterfaceName) > 0 {
		queryParameters["networkInterfaceName"] = autorest.Encode("query", networkInterfaceName)
	}
	if len(accessControlGroupNoListN) > 0 {
		queryParameters["accessControlGroupNoList.N"] = autorest.Encode("query", accessControlGroupNoListN)
	}
	if len(serverInstanceNo) > 0 {
		queryParameters["serverInstanceNo"] = autorest.Encode("query", serverInstanceNo)
	}
	if len(IP) > 0 {
		queryParameters["ip"] = autorest.Encode("query", IP)
	}
	if len(networkInterfaceDescription) > 0 {
		queryParameters["networkInterfaceDescription"] = autorest.Encode("query", networkInterfaceDescription)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/createNetworkInterface"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) CreateResponder(resp *http.Response) (result NetworkInterfaceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 네트워크 인터페이스 삭제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// networkInterfaceNo - 네트워크 인터페이스 번호
func (client NetworkInterfaceClient) Delete(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNo string) (result NetworkInterfaceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, responseFormatType, regionCode, networkInterfaceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NetworkInterfaceClient) DeletePreparer(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkInterfaceNo) > 0 {
		queryParameters["networkInterfaceNo"] = autorest.Encode("query", networkInterfaceNo)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/deleteNetworkInterface"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) DeleteResponder(resp *http.Response) (result NetworkInterfaceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Detach 네트워크 인터페이스를 할당 해제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// subnetNo - 서브넷 번호
// networkInterfaceNo - 네트워크 인터페이스 번호
// serverInstanceNo - 서버 인스턴스 번호
// regionCode - REGION 코드
func (client NetworkInterfaceClient) Detach(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (result NetworkInterfaceResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Detach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetachPreparer(ctx, responseFormatType, subnetNo, networkInterfaceNo, serverInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Detach", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Detach", resp, "Failure sending request")
		return
	}

	result, err = client.DetachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "Detach", resp, "Failure responding to request")
	}

	return
}

// DetachPreparer prepares the Detach request.
func (client NetworkInterfaceClient) DetachPreparer(ctx context.Context, responseFormatType string, subnetNo string, networkInterfaceNo string, serverInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkInterfaceNo": autorest.Encode("query", networkInterfaceNo),
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"serverInstanceNo":   autorest.Encode("query", serverInstanceNo),
		"subnetNo":           autorest.Encode("query", subnetNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/detachNetworkInterface"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetachSender sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) DetachSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetachResponder handles the response to the Detach request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) DetachResponder(resp *http.Response) (result NetworkInterfaceResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 네트워크 인터페이스 상세 정보를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// networkInterfaceNo - 네트워크 인터페이스 번호
// regionCode - REGION 코드
func (client NetworkInterfaceClient) GetDetail(ctx context.Context, responseFormatType string, networkInterfaceNo string, regionCode string) (result NetworkInterfaceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, responseFormatType, networkInterfaceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client NetworkInterfaceClient) GetDetailPreparer(ctx context.Context, responseFormatType string, networkInterfaceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkInterfaceNo": autorest.Encode("query", networkInterfaceNo),
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/getNetworkInterfaceDetail"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) GetDetailResponder(resp *http.Response) (result NetworkInterfaceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 네트워크 인터페이스 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// networkInterfaceNoListN - 네트워크 인터페이스 번호 리스트
// IP - IP 주소
// networkInterfaceName - 네트워크 인터페이스 이름
// serverInstanceName - 서버 인스턴스 이름
// subnetName - 서브넷 이름
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// networkInterfaceStatusCode - 네트워크 인터페이스 상태 코드
func (client NetworkInterfaceClient) GetList(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNoListN string, IP string, networkInterfaceName string, serverInstanceName string, subnetName string, pageNo string, pageSize string, networkInterfaceStatusCode NetworkInterfaceStatusCode) (result NetworkInterfaceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, regionCode, networkInterfaceNoListN, IP, networkInterfaceName, serverInstanceName, subnetName, pageNo, pageSize, networkInterfaceStatusCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client NetworkInterfaceClient) GetListPreparer(ctx context.Context, responseFormatType string, regionCode string, networkInterfaceNoListN string, IP string, networkInterfaceName string, serverInstanceName string, subnetName string, pageNo string, pageSize string, networkInterfaceStatusCode NetworkInterfaceStatusCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkInterfaceNoListN) > 0 {
		queryParameters["networkInterfaceNoList.N"] = autorest.Encode("query", networkInterfaceNoListN)
	}
	if len(IP) > 0 {
		queryParameters["ip"] = autorest.Encode("query", IP)
	}
	if len(networkInterfaceName) > 0 {
		queryParameters["networkInterfaceName"] = autorest.Encode("query", networkInterfaceName)
	}
	if len(serverInstanceName) > 0 {
		queryParameters["serverInstanceName"] = autorest.Encode("query", serverInstanceName)
	}
	if len(subnetName) > 0 {
		queryParameters["subnetName"] = autorest.Encode("query", subnetName)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(string(networkInterfaceStatusCode)) > 0 {
		queryParameters["networkInterfaceStatusCode"] = autorest.Encode("query", networkInterfaceStatusCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/getNetworkInterfaceList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) GetListResponder(resp *http.Response) (result NetworkInterfaceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveACG 네트워크 인터페이스의 ACG를 제거
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// networkInterfaceNo - 네트워크 인터페이스 번호
// accessControlGroupNoListN - ACG 번호 리스트
// regionCode - REGION 코드
func (client NetworkInterfaceClient) RemoveACG(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (result NetworkInterfaceAccessControlGroupResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.RemoveACG")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveACGPreparer(ctx, responseFormatType, networkInterfaceNo, accessControlGroupNoListN, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "RemoveACG", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveACGSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "RemoveACG", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveACGResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.NetworkInterfaceClient", "RemoveACG", resp, "Failure responding to request")
	}

	return
}

// RemoveACGPreparer prepares the RemoveACG request.
func (client NetworkInterfaceClient) RemoveACGPreparer(ctx context.Context, responseFormatType string, networkInterfaceNo string, accessControlGroupNoListN string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNoList.N": autorest.Encode("query", accessControlGroupNoListN),
		"networkInterfaceNo":         autorest.Encode("query", networkInterfaceNo),
		"responseFormatType":         autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/removeNetworkInterfaceAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveACGSender sends the RemoveACG request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) RemoveACGSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveACGResponder handles the response to the RemoveACG request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) RemoveACGResponder(resp *http.Response) (result NetworkInterfaceAccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
