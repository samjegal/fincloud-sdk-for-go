package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ACGClient is the server Client
type ACGClient struct {
	BaseClient
}

// NewACGClient creates an instance of the ACGClient client.
func NewACGClient() ACGClient {
	return NewACGClientWithBaseURI(DefaultBaseURI)
}

// NewACGClientWithBaseURI creates an instance of the ACGClient client.
func NewACGClientWithBaseURI(baseURI string) ACGClient {
	return ACGClient{NewWithBaseURI(baseURI)}
}

// Create aCG를 생성
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcNo - VPC 번호
// regionCode - REGION 코드
// accessControlGroupName - ACG 이름
// accessControlGroupDescription - ACG 설명
func (client ACGClient) Create(ctx context.Context, responseFormatType string, vpcNo string, regionCode string, accessControlGroupName string, accessControlGroupDescription string) (result AccessControlGroupResponse, err error) {
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
	req, err := client.CreatePreparer(ctx, responseFormatType, vpcNo, regionCode, accessControlGroupName, accessControlGroupDescription)
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
func (client ACGClient) CreatePreparer(ctx context.Context, responseFormatType string, vpcNo string, regionCode string, accessControlGroupName string, accessControlGroupDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(accessControlGroupName) > 0 {
		queryParameters["accessControlGroupName"] = autorest.Encode("query", accessControlGroupName)
	}
	if len(accessControlGroupDescription) > 0 {
		queryParameters["accessControlGroupDescription"] = autorest.Encode("query", accessControlGroupDescription)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ACGClient) CreateResponder(resp *http.Response) (result AccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete aCG를 삭제
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// regionCode - REGION 코드
func (client ACGClient) Delete(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, regionCode string) (result AccessControlGroupResponse, err error) {
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
	req, err := client.DeletePreparer(ctx, responseFormatType, vpcNo, accessControlGroupNo, regionCode)
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
func (client ACGClient) DeletePreparer(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", responseFormatType),
		"vpcNo":                autorest.Encode("query", vpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteAccessControlGroup"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ACGClient) DeleteResponder(resp *http.Response) (result AccessControlGroupResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail ACG 상세 정보를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// accessControlGroupNo - ACG 번호
// regionCode - REGION 코드
func (client ACGClient) GetDetail(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string) (result AccessControlGroupDetailResponse, err error) {
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
	req, err := client.GetDetailPreparer(ctx, responseFormatType, accessControlGroupNo, regionCode)
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
func (client ACGClient) GetDetailPreparer(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupDetail"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client ACGClient) GetDetailResponder(resp *http.Response) (result AccessControlGroupDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList ACG 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// accessControlGroupNoListN - ACG 번호 리스트
// accessControlGroupName - ACG 이름
// accessControlGroupStatusCode - ACG 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// vpcNo - VPC 번호
func (client ACGClient) GetList(ctx context.Context, responseFormatType string, regionCode string, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (result AccessControlGroupListResponse, err error) {
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
	req, err := client.GetListPreparer(ctx, responseFormatType, regionCode, accessControlGroupNoListN, accessControlGroupName, accessControlGroupStatusCode, pageNo, pageSize, vpcNo)
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
func (client ACGClient) GetListPreparer(ctx context.Context, responseFormatType string, regionCode string, accessControlGroupNoListN string, accessControlGroupName string, accessControlGroupStatusCode AccessControlGroupStatusCode, pageNo string, pageSize string, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
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

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ACGClient) GetListResponder(resp *http.Response) (result AccessControlGroupListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRuleList ACG Rule 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// accessControlGroupNo - ACG 번호
// regionCode - REGION 코드
// accessControlGroupRuleTypeCode - ACG Rule 유형 코드
func (client ACGClient) GetRuleList(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string, accessControlGroupRuleTypeCode AccessControlGroupRuleTypeCode) (result AccessControlGroupRuleListResponse, err error) {
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
	req, err := client.GetRuleListPreparer(ctx, responseFormatType, accessControlGroupNo, regionCode, accessControlGroupRuleTypeCode)
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
func (client ACGClient) GetRuleListPreparer(ctx context.Context, responseFormatType string, accessControlGroupNo string, regionCode string, accessControlGroupRuleTypeCode AccessControlGroupRuleTypeCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo": autorest.Encode("query", accessControlGroupNo),
		"responseFormatType":   autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(string(accessControlGroupRuleTypeCode)) > 0 {
		queryParameters["accessControlGroupRuleTypeCode"] = autorest.Encode("query", accessControlGroupRuleTypeCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getAccessControlGroupRuleList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRuleListSender sends the GetRuleList request. The method will close the
// http.Response Body if it receives an error.
func (client ACGClient) GetRuleListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetRuleListResponder handles the response to the GetRuleList request. The method always
// closes the http.Response Body.
func (client ACGClient) GetRuleListResponder(resp *http.Response) (result AccessControlGroupRuleListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
