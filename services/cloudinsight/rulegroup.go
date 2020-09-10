package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/fincloud-sdk-for-go/common"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// RuleGroupClient is the cloud Insight Client
type RuleGroupClient struct {
	BaseClient
}

// NewRuleGroupClient creates an instance of the RuleGroupClient client.
func NewRuleGroupClient() RuleGroupClient {
	return NewRuleGroupClientWithBaseURI(DefaultBaseURI)
}

// NewRuleGroupClientWithBaseURI creates an instance of the RuleGroupClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRuleGroupClientWithBaseURI(baseURI string) RuleGroupClient {
	return RuleGroupClient{NewWithBaseURI(baseURI)}
}

// Copy event Rule을 복제합니다.
// Parameters:
// ID - 감시 항목 그룹(템플릿)의 id
func (client RuleGroupClient) Copy(ctx context.Context, ID string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.Copy")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CopyPreparer(ctx, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Copy", nil, "Failure preparing request")
		return
	}

	resp, err := client.CopySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Copy", resp, "Failure sending request")
		return
	}

	result, err = client.CopyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Copy", resp, "Failure responding to request")
	}

	return
}

// CopyPreparer prepares the Copy request.
func (client RuleGroupClient) CopyPreparer(ctx context.Context, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id": autorest.Encode("path", ID),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("PUT", common.GetPathParameters(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/copy/{id}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/cw_fea/real/cw/api/rule/group/ruleGrp/copy/{id}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CopySender sends the Copy request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) CopySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CopyResponder handles the response to the Copy request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) CopyResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CopyAsgGroup copy new rules according to proto asgGroupNo
// Parameters:
// parameters - copy rule for asgGroup
func (client RuleGroupClient) CopyAsgGroup(ctx context.Context, parameters RuleGroupCopyForAsgGroupRequest) (result ListSetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.CopyAsgGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.NewAsgGroupNo", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.NewAsgPolicyNo", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProtoAsgGroupNo", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProtoAsgPolicyNo", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "CopyAsgGroup", err.Error())
	}

	req, err := client.CopyAsgGroupPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopyAsgGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CopyAsgGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopyAsgGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CopyAsgGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopyAsgGroup", resp, "Failure responding to request")
	}

	return
}

// CopyAsgGroupPreparer prepares the CopyAsgGroup request.
func (client RuleGroupClient) CopyAsgGroupPreparer(ctx context.Context, parameters RuleGroupCopyForAsgGroupRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/asgCopy"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/asgCopy"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CopyAsgGroupSender sends the CopyAsgGroup request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) CopyAsgGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CopyAsgGroupResponder handles the response to the CopyAsgGroup request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) CopyAsgGroupResponder(resp *http.Response) (result ListSetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CopySettings event Rule의 설정을 복사하여 지정한 resource에 대한 새로운 Event Rule을 생성합니다.
// Parameters:
// parameters - copy given rule's settings into the resourceId
func (client RuleGroupClient) CopySettings(ctx context.Context, parameters RuleGroupCopySettingRequest) (result ListInt64, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.CopySettings")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ProdKey", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ResourceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.RuleGroupIds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "CopySettings", err.Error())
	}

	req, err := client.CopySettingsPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopySettings", nil, "Failure preparing request")
		return
	}

	resp, err := client.CopySettingsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopySettings", resp, "Failure sending request")
		return
	}

	result, err = client.CopySettingsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CopySettings", resp, "Failure responding to request")
	}

	return
}

// CopySettingsPreparer prepares the CopySettings request.
func (client RuleGroupClient) CopySettingsPreparer(ctx context.Context, parameters RuleGroupCopySettingRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/copySettings"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/copySettings"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CopySettingsSender sends the CopySettings request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) CopySettingsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CopySettingsResponder handles the response to the CopySettings request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) CopySettingsResponder(resp *http.Response) (result ListInt64, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create event Rule을 생성합니다.
// Parameters:
// parameters - 이벤트 룰 그룹 요청
func (client RuleGroupClient) Create(ctx context.Context, parameters RuleGroupRequest) (result Int64, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.GroupName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.MetricsGroupKey", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.MonitorGroupKey", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProdKey", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RuleGroupClient) CreatePreparer(ctx context.Context, parameters RuleGroupRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) CreateResponder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateDirectly 감시 대상 그룹과 감시 항목 그룹 생성 없이 감시 대상과 감시 항목을 지정하여 바로 Event Rule을 생성합니다.
// Parameters:
// parameters - 감시 대상 그룹, 감시 항목 그룹 생성 없이 Event Rule 생성
func (client RuleGroupClient) CreateDirectly(ctx context.Context, parameters DirectRuleGroupCreateRequest) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.CreateDirectly")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.MetricsGroup", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.MetricsGroup.MetricsGroupItems", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.MetricsGroup.ProdKey", Name: validation.Null, Rule: true, Chain: nil},
				}},
				{Target: "parameters.MonitorGroup", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.MonitorGroup.MonitorGroupItemList", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.MonitorGroup.ProdKey", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "CreateDirectly", err.Error())
	}

	req, err := client.CreateDirectlyPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CreateDirectly", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateDirectlySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CreateDirectly", resp, "Failure sending request")
		return
	}

	result, err = client.CreateDirectlyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "CreateDirectly", resp, "Failure responding to request")
	}

	return
}

// CreateDirectlyPreparer prepares the CreateDirectly request.
func (client RuleGroupClient) CreateDirectlyPreparer(ctx context.Context, parameters DirectRuleGroupCreateRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/createDirectly"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/createDirectly"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateDirectlySender sends the CreateDirectly request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) CreateDirectlySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateDirectlyResponder handles the response to the CreateDirectly request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) CreateDirectlyResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete event Rule을 삭제합니다.
// Parameters:
// parameters - 삭제하고자 하는 Event Rule id, 여러개 입력 가능
func (client RuleGroupClient) Delete(ctx context.Context, parameters RuleGroupDeleteRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Items", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RuleGroupClient) DeletePreparer(ctx context.Context, parameters RuleGroupDeleteRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/del"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/del"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByProdKeyAndID event Rule id로 Event Rule을 삭제합니다.
// Parameters:
// prodKey - 상품의 cw_key
// ID - 감시 항목 그룹(템플릿)의 id
func (client RuleGroupClient) DeleteByProdKeyAndID(ctx context.Context, prodKey string, ID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.DeleteByProdKeyAndID")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByProdKeyAndIDPreparer(ctx, prodKey, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "DeleteByProdKeyAndID", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByProdKeyAndIDSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "DeleteByProdKeyAndID", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByProdKeyAndIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "DeleteByProdKeyAndID", resp, "Failure responding to request")
	}

	return
}

// DeleteByProdKeyAndIDPreparer prepares the DeleteByProdKeyAndID request.
func (client RuleGroupClient) DeleteByProdKeyAndIDPreparer(ctx context.Context, prodKey string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":      autorest.Encode("path", ID),
		"prodKey": autorest.Encode("path", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPathParameters(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/del/{prodKey}/{id}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/cw_fea/real/cw/api/rule/group/ruleGrp/del/{prodKey}/{id}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByProdKeyAndIDSender sends the DeleteByProdKeyAndID request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) DeleteByProdKeyAndIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteByProdKeyAndIDResponder handles the response to the DeleteByProdKeyAndID request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) DeleteByProdKeyAndIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetByMonitorGroupIds 감시 대상 그룹과 관련된 Event Rule을 조회합니다.
// Parameters:
// prodKey - 상품의 cw_key
// parameters - 조회하려는 감시 대상 그룹의 Id이며, 여러개의 감시 대상 그룹을 한번에 조회할 수 있습니다.
func (client RuleGroupClient) GetByMonitorGroupIds(ctx context.Context, prodKey string, parameters []string) (result ListRuleGroupItemListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.GetByMonitorGroupIds")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "GetByMonitorGroupIds", err.Error())
	}

	req, err := client.GetByMonitorGroupIdsPreparer(ctx, prodKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "GetByMonitorGroupIds", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByMonitorGroupIdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "GetByMonitorGroupIds", resp, "Failure sending request")
		return
	}

	result, err = client.GetByMonitorGroupIdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "GetByMonitorGroupIds", resp, "Failure responding to request")
	}

	return
}

// GetByMonitorGroupIdsPreparer prepares the GetByMonitorGroupIds request.
func (client RuleGroupClient) GetByMonitorGroupIdsPreparer(ctx context.Context, prodKey string, parameters []string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"prodKey": autorest.Encode("query", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/metric/group/related")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/metric/group/related"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByMonitorGroupIdsSender sends the GetByMonitorGroupIds request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) GetByMonitorGroupIdsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByMonitorGroupIdsResponder handles the response to the GetByMonitorGroupIds request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) GetByMonitorGroupIdsResponder(resp *http.Response) (result ListRuleGroupItemListParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Query event Rule을 조회합니다.
// Parameters:
// parameters - rule group list query item
func (client RuleGroupClient) Query(ctx context.Context, parameters RuleGroupListQueryRequest) (result RuleGroupListQueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.Query")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Query", nil, "Failure preparing request")
		return
	}

	resp, err := client.QuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Query", resp, "Failure sending request")
		return
	}

	result, err = client.QueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Query", resp, "Failure responding to request")
	}

	return
}

// QueryPreparer prepares the Query request.
func (client RuleGroupClient) QueryPreparer(ctx context.Context, parameters RuleGroupListQueryRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/query"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/query"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QuerySender sends the Query request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) QuerySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryResponder handles the response to the Query request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) QueryResponder(resp *http.Response) (result RuleGroupListQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// QueryByProdKeyAndID event Rule을 조회합니다.
// Parameters:
// prodKey - 상품의 cw_key
// ID - 감시 항목 그룹(템플릿)의 id
func (client RuleGroupClient) QueryByProdKeyAndID(ctx context.Context, prodKey string, ID string) (result RuleGroupParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.QueryByProdKeyAndID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryByProdKeyAndIDPreparer(ctx, prodKey, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "QueryByProdKeyAndID", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryByProdKeyAndIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "QueryByProdKeyAndID", resp, "Failure sending request")
		return
	}

	result, err = client.QueryByProdKeyAndIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "QueryByProdKeyAndID", resp, "Failure responding to request")
	}

	return
}

// QueryByProdKeyAndIDPreparer prepares the QueryByProdKeyAndID request.
func (client RuleGroupClient) QueryByProdKeyAndIDPreparer(ctx context.Context, prodKey string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":      autorest.Encode("path", ID),
		"prodKey": autorest.Encode("path", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/query/{prodKey}/{id}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/cw_fea/real/cw/api/rule/group/ruleGrp/query/{prodKey}/{id}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryByProdKeyAndIDSender sends the QueryByProdKeyAndID request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) QueryByProdKeyAndIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryByProdKeyAndIDResponder handles the response to the QueryByProdKeyAndID request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) QueryByProdKeyAndIDResponder(resp *http.Response) (result RuleGroupParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update event Rule을 수정합니다.
// Parameters:
// parameters - event Rule 생성/수정 요청
func (client RuleGroupClient) Update(ctx context.Context, parameters RuleGroupRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleGroupClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.GroupName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.MetricsGroupKey", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.MonitorGroupKey", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProdKey", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.RuleGroupClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.RuleGroupClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RuleGroupClient) UpdatePreparer(ctx context.Context, parameters RuleGroupRequest) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/ruleGrp/update"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp/update"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RuleGroupClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RuleGroupClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
