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

// ACGInboundClient is the server Client
type ACGInboundClient struct {
	BaseClient
}

// NewACGInboundClient creates an instance of the ACGInboundClient client.
func NewACGInboundClient() ACGInboundClient {
	return NewACGInboundClientWithBaseURI(DefaultBaseURI)
}

func NewACGInboundClientWithKey(accessKey string, secretKey string) ACGInboundClient {
	return NewACGInboundClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewACGInboundClientWithBaseURI creates an instance of the ACGInboundClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewACGInboundClientWithBaseURI(baseURI string) ACGInboundClient {
	return ACGInboundClient{NewWithBaseURI(baseURI)}
}

func NewACGInboundClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) ACGInboundClient {
	return ACGInboundClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// AddRule ACG Inbound Rule을 추가
// Parameters:
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
// accessControlGroupRuleListNaccessControlGroupRuleDescription - ACG Rule 설명
func (client ACGInboundClient) AddRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (result AccessControlGroupInboundRuleAddResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGInboundClient.AddRule")
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
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "AddRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "AddRule", resp, "Failure sending request")
		return
	}

	result, err = client.AddRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "AddRule", resp, "Failure responding to request")
	}

	return
}

// AddRulePreparer prepares the AddRule request.
func (client ACGInboundClient) AddRulePreparer(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string, accessControlGroupRuleListNaccessControlGroupRuleDescription string) (*http.Request, error) {
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
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/addAccessControlGroupInboundRule")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addAccessControlGroupInboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddRuleSender sends the AddRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGInboundClient) AddRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddRuleResponder handles the response to the AddRule request. The method always
// closes the http.Response Body.
func (client ACGInboundClient) AddRuleResponder(resp *http.Response) (result AccessControlGroupInboundRuleAddResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveRule ACG Inbound Rule을 제거
// Parameters:
// vpcNo - VPC 번호
// accessControlGroupNo - ACG 번호
// accessControlGroupRuleListNprotocolTypeCode - 프로토콜 유형 코드
// accessControlGroupRuleListNipBlock - IP 블록
// accessControlGroupRuleListNaccessControlGroupSequence - 접근 소스 ACG
// accessControlGroupRuleListNportRange - 포트 범위
func (client ACGInboundClient) RemoveRule(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (result AccessControlGroupInboundRuleRemoveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ACGInboundClient.RemoveRule")
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
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "RemoveRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "RemoveRule", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ACGInboundClient", "RemoveRule", resp, "Failure responding to request")
	}

	return
}

// RemoveRulePreparer prepares the RemoveRule request.
func (client ACGInboundClient) RemoveRulePreparer(ctx context.Context, vpcNo string, accessControlGroupNo string, accessControlGroupRuleListNprotocolTypeCode ProtocolTypeCode, accessControlGroupRuleListNipBlock string, accessControlGroupRuleListNaccessControlGroupSequence string, accessControlGroupRuleListNportRange string) (*http.Request, error) {
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
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/removeAccessControlGroupInboundRule")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeAccessControlGroupInboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveRuleSender sends the RemoveRule request. The method will close the
// http.Response Body if it receives an error.
func (client ACGInboundClient) RemoveRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveRuleResponder handles the response to the RemoveRule request. The method always
// closes the http.Response Body.
func (client ACGInboundClient) RemoveRuleResponder(resp *http.Response) (result AccessControlGroupInboundRuleRemoveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
