package cloudinsight

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SchemaClient is the cloud Insight Client
type SchemaClient struct {
	BaseClient
}

// NewSchemaClient creates an instance of the SchemaClient client.
func NewSchemaClient() SchemaClient {
	return NewSchemaClientWithBaseURI(DefaultBaseURI)
}

// NewSchemaClientWithBaseURI creates an instance of the SchemaClient client.
func NewSchemaClientWithBaseURI(baseURI string) SchemaClient {
	return SchemaClient{NewWithBaseURI(baseURI)}
}

// Create create schema for user application. If product does not exist, it will registered automatically
// Parameters:
// parameters - cloud Insight Custom 메트릭 생성 데이터
// prodName - product 이름
func (client SchemaClient) Create(ctx context.Context, parameters SchemaParameter, prodName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, parameters, prodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SchemaClient) CreatePreparer(ctx context.Context, parameters SchemaParameter, prodName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(prodName) > 0 {
		queryParameters["prodName"] = autorest.Encode("query", prodName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SchemaClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete schema for an application
// Parameters:
// prodName - product 이름
func (client SchemaClient) Delete(ctx context.Context, prodName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, prodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SchemaClient) DeletePreparer(ctx context.Context, prodName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(prodName) > 0 {
		queryParameters["prodName"] = autorest.Encode("query", prodName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SchemaClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disable disable the extended metrics for a product
// Parameters:
// cwKey - product key
// instanceIds - target instance id, string separated by commas
func (client SchemaClient) Disable(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Disable")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisablePreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Disable", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Disable", resp, "Failure sending request")
		return
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client SchemaClient) DisablePreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(cwKey) > 0 {
		queryParameters["cw_key"] = autorest.Encode("query", cwKey)
	}
	if len(instanceIds) > 0 {
		queryParameters["instanceIds"] = autorest.Encode("query", instanceIds)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/disable"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) DisableSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client SchemaClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable enable the extended metrics for a product
// Parameters:
// cwKey - product key
// instanceIds - target instance id, string separated by commas
func (client SchemaClient) Enable(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Enable")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnablePreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client SchemaClient) EnablePreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(cwKey) > 0 {
		queryParameters["cw_key"] = autorest.Encode("query", cwKey)
	}
	if len(instanceIds) > 0 {
		queryParameters["instanceIds"] = autorest.Encode("query", instanceIds)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/enable"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) EnableSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client SchemaClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get schema information of the specific product
// Parameters:
// prodName - product 이름
func (client SchemaClient) Get(ctx context.Context, prodName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Get")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, prodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SchemaClient) GetPreparer(ctx context.Context, prodName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(prodName) > 0 {
		queryParameters["prodName"] = autorest.Encode("query", prodName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SchemaClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetList get schema list
func (client SchemaClient) GetList(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.GetList")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client SchemaClient) GetListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/list"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) GetListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client SchemaClient) GetListResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// QueryStatus query the status of extended metrics of those product
// Parameters:
// cwKey - product key
// instanceIds - target instance id, string separated by commas
func (client SchemaClient) QueryStatus(ctx context.Context, cwKey string, instanceIds string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.QueryStatus")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryStatusPreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "QueryStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryStatusSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "QueryStatus", resp, "Failure sending request")
		return
	}

	result, err = client.QueryStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "QueryStatus", resp, "Failure responding to request")
	}

	return
}

// QueryStatusPreparer prepares the QueryStatus request.
func (client SchemaClient) QueryStatusPreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(cwKey) > 0 {
		queryParameters["cw_key"] = autorest.Encode("query", cwKey)
	}
	if len(instanceIds) > 0 {
		queryParameters["instanceIds"] = autorest.Encode("query", instanceIds)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/status"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryStatusSender sends the QueryStatus request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) QueryStatusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryStatusResponder handles the response to the QueryStatus request. The method always
// closes the http.Response Body.
func (client SchemaClient) QueryStatusResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update schema for an existing product
// Parameters:
// prodName - product 이름
func (client SchemaClient) Update(ctx context.Context, prodName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SchemaClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, prodName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.SchemaClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SchemaClient) UpdatePreparer(ctx context.Context, prodName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if len(prodName) > 0 {
		queryParameters["prodName"] = autorest.Encode("query", prodName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SchemaClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SchemaClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
