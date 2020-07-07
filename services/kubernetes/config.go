package kubernetes

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ConfigClient is the kubernetes Client
type ConfigClient struct {
	BaseClient
}

// NewConfigClient creates an instance of the ConfigClient client.
func NewConfigClient() ConfigClient {
	return NewConfigClientWithBaseURI(DefaultBaseURI)
}

// NewConfigClientWithBaseURI creates an instance of the ConfigClient client.
func NewConfigClientWithBaseURI(baseURI string) ConfigClient {
	return ConfigClient{NewWithBaseURI(baseURI)}
}

// Get kubeConfig 조회
// Parameters:
// UUID - 클러스터 UUID
func (client ConfigClient) Get(ctx context.Context, UUID string) (result ConfigParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, UUID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConfigClient) GetPreparer(ctx context.Context, UUID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"uuid": autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/clusters/{uuid}/kubeconfig", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConfigClient) GetResponder(resp *http.Response) (result ConfigParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Reset kubeConfig 재설정
// Parameters:
// UUID - 클러스터 UUID
func (client ConfigClient) Reset(ctx context.Context, UUID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigClient.Reset")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetPreparer(ctx, UUID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Reset", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Reset", resp, "Failure sending request")
		return
	}

	result, err = client.ResetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.ConfigClient", "Reset", resp, "Failure responding to request")
	}

	return
}

// ResetPreparer prepares the Reset request.
func (client ConfigClient) ResetPreparer(ctx context.Context, UUID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"uuid": autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/clusters/{uuid}/kubeconfig/reset", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigClient) ResetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client ConfigClient) ResetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
