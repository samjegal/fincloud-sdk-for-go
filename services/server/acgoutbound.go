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
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
// accessControlGroupRuleListNaccessControlGroupRuleDescription - ACG Rule 설명
func (client ACGOutboundClient) AddRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result AccessControlGroupOutboundRuleAddResponse, err error) {
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
	req, err := client.AddRulePreparer(ctx, vpcNo, accessControlGroupNo, accessControlGroupRuleListNprotocolTypeCode, accessControlGroupRuleListNipBlock, accessControlGroupRuleListNaccessControlGroupSequence, accessControlGroupRuleListNportRange, accessControlGroupRuleListNaccessControlGroupRuleDescription)
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
func (client ACGOutboundClient) AddRulePreparer(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo":                          autorest.Encode("query", accessControlGroupNo),
		"accessControlGroupRuleList.N.protocolTypeCode": autorest.Encode("query", accessControlGroupRuleListNprotocolTypeCode),
		"responseFormatType":                            autorest.Encode("query", "json"),
		"vpcNo":                                         autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

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

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addAccessControlGroupOutboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addAccessControlGroupOutboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddRuleSender sends the AddRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGOutboundClient) AddRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddRuleResponder handles the response to the AddRule request. The method always
// closes the http.Response Body.
func (client ACGOutboundClient) AddRuleResponder(resp *http.Response) (result AccessControlGroupOutboundRuleAddResponse, err error) {
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
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
func (client ACGOutboundClient) RemoveRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result AccessControlGroupOutboundRuleRemoveResponse, err error) {
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
	req, err := client.RemoveRulePreparer(ctx, vpcNo, accessControlGroupNo, accessControlGroupRuleListNprotocolTypeCode, accessControlGroupRuleListNipBlock, accessControlGroupRuleListNaccessControlGroupSequence, accessControlGroupRuleListNportRange)
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
func (client ACGOutboundClient) RemoveRulePreparer(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"accessControlGroupNo":                          autorest.Encode("query", accessControlGroupNo),
		"accessControlGroupRuleList.N.protocolTypeCode": autorest.Encode("query", accessControlGroupRuleListNprotocolTypeCode),
		"responseFormatType":                            autorest.Encode("query", "json"),
		"vpcNo":                                         autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(accessControlGroupRuleListNipBlock) > 0 {
		queryParameters["accessControlGroupRuleList.N.ipBlock"] = autorest.Encode("query", accessControlGroupRuleListNipBlock)
	}
	if len(accessControlGroupRuleListNaccessControlGroupSequence) > 0 {
		queryParameters["accessControlGroupRuleList.N.accessControlGroupSequence"] = autorest.Encode("query", accessControlGroupRuleListNaccessControlGroupSequence)
	}
	if len(accessControlGroupRuleListNportRange) > 0 {
		queryParameters["accessControlGroupRuleList.N.portRange"] = autorest.Encode("query", accessControlGroupRuleListNportRange)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeAccessControlGroupOutboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeAccessControlGroupOutboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveRuleSender sends the RemoveRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGOutboundClient) RemoveRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveRuleResponder handles the response to the RemoveRule request. The method always
// closes the http.Response Body.
func (client ACGOutboundClient) RemoveRuleResponder(resp *http.Response) (result AccessControlGroupOutboundRuleRemoveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
