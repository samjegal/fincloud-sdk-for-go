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

// MonitorGroupClient is the cloud Insight Client
type MonitorGroupClient struct {
	BaseClient
}

// NewMonitorGroupClient creates an instance of the MonitorGroupClient client.
func NewMonitorGroupClient() MonitorGroupClient {
	return NewMonitorGroupClientWithBaseURI(DefaultBaseURI)
}

func NewMonitorGroupClientWithKey(accessKey string, secretKey string, apiGatewayKey string) MonitorGroupClient {
	return NewMonitorGroupClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey, apiGatewayKey)
}

// NewMonitorGroupClientWithBaseURI creates an instance of the MonitorGroupClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMonitorGroupClientWithBaseURI(baseURI string) MonitorGroupClient {
	return MonitorGroupClient{NewWithBaseURI(baseURI)}
}

func NewMonitorGroupClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string, apiGatewayKey string) MonitorGroupClient {
	return MonitorGroupClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey, apiGatewayKey)}
}

// DeleteForce 감시 대상 그룹과 관련된 모든 Event Rule을 삭제합니다.
// Parameters:
// prodKey - 상품의 cw_key
// parameters - rule group for deletion
func (client MonitorGroupClient) DeleteForce(ctx context.Context, prodKey string, parameters []TypeGroupRelatedRuleParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.DeleteForce")
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
			Constraints: []validation.Constraint{{Target: "parameters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cloudinsight.MonitorGroupClient", "DeleteForce", err.Error())
	}

	req, err := client.DeleteForcePreparer(ctx, prodKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteForceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteForceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce", resp, "Failure responding to request")
	}

	return
}

// DeleteForcePreparer prepares the DeleteForce request.
func (client MonitorGroupClient) DeleteForcePreparer(ctx context.Context, prodKey string, parameters []TypeGroupRelatedRuleParameter) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"prodKey": autorest.Encode("query", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/metric/groups")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/metric/groups"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteForceSender sends the DeleteForce request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorGroupClient) DeleteForceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteForceResponder handles the response to the DeleteForce request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) DeleteForceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteForce1 감시 대상 그룹과 관련된 모든 Event Rule을 삭제합니다.
// Parameters:
// prodKey - 상품의 cw_key
// parameters - 감시 대상 그룹 혹은 감시 항목 그룹과 관련된 모든 Event Rule 리스트
func (client MonitorGroupClient) DeleteForce1(ctx context.Context, prodKey string, parameters []TypeGroupRelatedRuleParameter) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.DeleteForce1")
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
		return result, validation.NewError("cloudinsight.MonitorGroupClient", "DeleteForce1", err.Error())
	}

	req, err := client.DeleteForce1Preparer(ctx, prodKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce1", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteForce1Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce1", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteForce1Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "DeleteForce1", resp, "Failure responding to request")
	}

	return
}

// DeleteForce1Preparer prepares the DeleteForce1 request.
func (client MonitorGroupClient) DeleteForce1Preparer(ctx context.Context, prodKey string, parameters []TypeGroupRelatedRuleParameter) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"prodKey": autorest.Encode("query", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("DELETE", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/monitor/groups")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/monitor/groups"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteForce1Sender sends the DeleteForce1 request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorGroupClient) DeleteForce1Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteForce1Responder handles the response to the DeleteForce1 request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) DeleteForce1Responder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get 감시 항목 그룹(템플릿) id로 감시 항목 그룹에 대한 상세정보를 조회합니다.
// Parameters:
// prodKey - 상품의 cw_key
// ID - 감시 항목 그룹(템플릿)의 id
func (client MonitorGroupClient) Get(ctx context.Context, prodKey string, ID string) (result MonitorGroupParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, prodKey, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitorGroupClient) GetPreparer(ctx context.Context, prodKey string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":      autorest.Encode("path", ID),
		"prodKey": autorest.Encode("path", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/monitor/{prodKey}/{id}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/cw_fea/real/cw/api/rule/group/monitor/{prodKey}/{id}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) GetResponder(resp *http.Response) (result MonitorGroupParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByMonitorGroupIds 감시 대상 그룹과 관련된 Event Rule을 조회합니다.
// Parameters:
// prodKey - 상품의 cw_key
// parameters - 조회하려는 감시 대상 그룹의 Id이며, 여러개의 감시 대상 그룹을 한번에 조회할 수 있습니다.
func (client MonitorGroupClient) GetByMonitorGroupIds(ctx context.Context, prodKey string, parameters []string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.GetByMonitorGroupIds")
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
		return result, validation.NewError("cloudinsight.MonitorGroupClient", "GetByMonitorGroupIds", err.Error())
	}

	req, err := client.GetByMonitorGroupIdsPreparer(ctx, prodKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "GetByMonitorGroupIds", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByMonitorGroupIdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "GetByMonitorGroupIds", resp, "Failure sending request")
		return
	}

	result, err = client.GetByMonitorGroupIdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "GetByMonitorGroupIds", resp, "Failure responding to request")
	}

	return
}

// GetByMonitorGroupIdsPreparer prepares the GetByMonitorGroupIds request.
func (client MonitorGroupClient) GetByMonitorGroupIdsPreparer(ctx context.Context, prodKey string, parameters []string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"prodKey": autorest.Encode("query", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/monitor/group/related")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/monitor/group/related"),
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
func (client MonitorGroupClient) GetByMonitorGroupIdsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByMonitorGroupIdsResponder handles the response to the GetByMonitorGroupIds request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) GetByMonitorGroupIdsResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List 해당 상품에 대한 모든 감시 항목 그룹(템플릿)을 조회합니다.
// Parameters:
// prodKey - 상품의 cw_key
func (client MonitorGroupClient) List(ctx context.Context, prodKey string) (result ListMonitorGroupParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, prodKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MonitorGroupClient) ListPreparer(ctx context.Context, prodKey string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"prodKey": autorest.Encode("path", prodKey),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPathParameters(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/monitor/{prodKey}", pathParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/cw_fea/real/cw/api/rule/group/monitor/{prodKey}", pathParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) ListResponder(resp *http.Response) (result ListMonitorGroupParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveResourceFromRules event Rule의 감시 대상 그룹에서 특정 감시 대상을 삭제합니다.
// Parameters:
// parameters - remove info
func (client MonitorGroupClient) RemoveResourceFromRules(ctx context.Context, parameters RemoveResourceFromRulesParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorGroupClient.RemoveResourceFromRules")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveResourceFromRulesPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "RemoveResourceFromRules", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveResourceFromRulesSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "RemoveResourceFromRules", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResourceFromRulesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.MonitorGroupClient", "RemoveResourceFromRules", resp, "Failure responding to request")
	}

	return
}

// RemoveResourceFromRulesPreparer prepares the RemoveResourceFromRules request.
func (client MonitorGroupClient) RemoveResourceFromRulesPreparer(ctx context.Context, parameters RemoveResourceFromRulesParameter) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/rule/group/monitor/removeResourceFromRules"), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/monitor/removeResourceFromRules"),
		autorest.WithJSON(parameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveResourceFromRulesSender sends the RemoveResourceFromRules request. The method will close the
// http.Response Body if it receives an error.
func (client MonitorGroupClient) RemoveResourceFromRulesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResourceFromRulesResponder handles the response to the RemoveResourceFromRules request. The method always
// closes the http.Response Body.
func (client MonitorGroupClient) RemoveResourceFromRulesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
