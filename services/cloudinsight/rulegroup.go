package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
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
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/rule/group/ruleGrp"),
		autorest.WithJSON(parameters))
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
