package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ACGOutboundClient is the server Client
type ACGOutboundClient struct {
	BaseClient
}

// NewACGOutboundClient creates an instance of the ACGOutboundClient client.
func NewACGOutboundClient() ACGOutboundClient {
	return NewACGOutboundClientWithBaseURI(DefaultBaseURI)
}

// NewACGOutboundClientWithBaseURI creates an instance of the ACGOutboundClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewACGOutboundClientWithBaseURI(baseURI string) ACGOutboundClient {
	return ACGOutboundClient{NewWithBaseURI(baseURI)}
}

// AddRule ACG Outbound Rule을 추가
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// regionCode - REGION 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
// accessControlGroupRuleListNaccessControlGroupRuleDescription - ACG Rule 설명
func (client ACGOutboundClient) AddRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result AccessControlGroupOutboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGOutboundClient.AddRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddRulePreparer(ctx, responseFormatType, vpcNo, accessControlGroupNo, accessControlGroupRuleListNprotocolTypeCode, regionCode, accessControlGroupRuleListNipBlock, accessControlGroupRuleListNaccessControlGroupSequence, accessControlGroupRuleListNportRange, accessControlGroupRuleListNaccessControlGroupRuleDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "AddRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "AddRule", resp, "Failure sending request")
		return
	}

	result, err = client.AddRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "AddRule", resp, "Failure responding to request")
	}

	return
}

// AddRulePreparer prepares the AddRule request.
func (client ACGOutboundClient) AddRulePreparer(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo":                          autorest.Encode("query", accessControlGroupNo),
		"accessControlGroupRuleList.N.protocolTypeCode": autorest.Encode("query", accessControlGroupRuleListNprotocolTypeCode),
		"responseFormatType":                            autorest.Encode("query", responseFormatType),
		"vpcNo":                                         autorest.Encode("query", vpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(accessControlGroupRuleListNipBlock) > 0 {
		queryParameters["accessControlGroupRuleList.N.ipBlock"] = autorest.Encode("query", accessControlGroupRuleListNipBlock)
	}
	if len(accessControlGroupRuleListNaccessControlGroupSequence) > 0 {
		queryParameters["accessControlGroupRuleList.N.accessControlGroupSequence"] = autorest.Encode("query", accessControlGroupRuleListNaccessControlGroupSequence)
	}
	if len(accessControlGroupRuleListNportRange) > 0 {
		queryParameters["accessControlGroupRuleList.N.portRange"] = autorest.Encode("query", accessControlGroupRuleListNportRange)
	}
	if len(accessControlGroupRuleListNaccessControlGroupRuleDescription) > 0 {
		queryParameters["accessControlGroupRuleList.N.accessControlGroupRuleDescription"] = autorest.Encode("query", accessControlGroupRuleListNaccessControlGroupRuleDescription)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/addAccessControlGroupOutboundRule"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddRuleSender sends the AddRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGOutboundClient) AddRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddRuleResponder handles the response to the AddRule request. The method always
// closes the http.Response Body.
func (client ACGOutboundClient) AddRuleResponder(resp *http.Response) (result AccessControlGroupOutboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveRule ACG Outbound Rule을 제거
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// regionCode - REGION 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
func (client ACGOutboundClient) RemoveRule(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result AccessControlGroupOutboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGOutboundClient.RemoveRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveRulePreparer(ctx, responseFormatType, vpcNo, accessControlGroupNo, accessControlGroupRuleListNprotocolTypeCode, regionCode, accessControlGroupRuleListNipBlock, accessControlGroupRuleListNaccessControlGroupSequence, accessControlGroupRuleListNportRange)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "RemoveRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "RemoveRule", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGOutboundClient", "RemoveRule", resp, "Failure responding to request")
	}

	return
}

// RemoveRulePreparer prepares the RemoveRule request.
func (client ACGOutboundClient) RemoveRulePreparer(ctx context.Context, responseFormatType string, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, regionCode string, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo":                          autorest.Encode("query", accessControlGroupNo),
		"accessControlGroupRuleList.N.protocolTypeCode": autorest.Encode("query", accessControlGroupRuleListNprotocolTypeCode),
		"responseFormatType":                            autorest.Encode("query", responseFormatType),
		"vpcNo":                                         autorest.Encode("query", vpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(accessControlGroupRuleListNipBlock) > 0 {
		queryParameters["accessControlGroupRuleList.N.ipBlock"] = autorest.Encode("query", accessControlGroupRuleListNipBlock)
	}
	if len(accessControlGroupRuleListNaccessControlGroupSequence) > 0 {
		queryParameters["accessControlGroupRuleList.N.accessControlGroupSequence"] = autorest.Encode("query", accessControlGroupRuleListNaccessControlGroupSequence)
	}
	if len(accessControlGroupRuleListNportRange) > 0 {
		queryParameters["accessControlGroupRuleList.N.portRange"] = autorest.Encode("query", accessControlGroupRuleListNportRange)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/removeAccessControlGroupOutboundRule"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveRuleSender sends the RemoveRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGOutboundClient) RemoveRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveRuleResponder handles the response to the RemoveRule request. The method always
// closes the http.Response Body.
func (client ACGOutboundClient) RemoveRuleResponder(resp *http.Response) (result AccessControlGroupOutboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
