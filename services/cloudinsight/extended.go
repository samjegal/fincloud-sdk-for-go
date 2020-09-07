package cloudinsight

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

// ExtendedClient is the cloud Insight Client
type ExtendedClient struct {
	BaseClient
}

// NewExtendedClient creates an instance of the ExtendedClient client.
func NewExtendedClient() ExtendedClient {
	return NewExtendedClientWithBaseURI(DefaultBaseURI)
}

// NewExtendedClientWithBaseURI creates an instance of the ExtendedClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExtendedClientWithBaseURI(baseURI string) ExtendedClient {
	return ExtendedClient{NewWithBaseURI(baseURI)}
}

// Disable NCP 상품의 extended 설정을 해제합니다.
// Parameters:
// cwKey - extended 설정을 해지하려는 상품의 cw_key
// instanceIds - extended 설정을 해지하려는 instanceId
func (client ExtendedClient) Disable(ctx context.Context, cwKey string, instanceIds string) (result SchemaExtendedDisableResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedClient.Disable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisablePreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Disable", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Disable", resp, "Failure sending request")
		return
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client ExtendedClient) DisablePreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"cw_key":      autorest.Encode("query", cwKey),
		"instanceIds": autorest.Encode("query", instanceIds),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("PUT", autorest.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/schema/extended/disable")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/disable"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedClient) DisableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client ExtendedClient) DisableResponder(resp *http.Response) (result SchemaExtendedDisableResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Enable NCP 상품의 extended 설정을 합니다.
// Parameters:
// cwKey - extended 설정을 하려는 상품의 cw_key
// instanceIds - extended 설정을 하려는 instanceId
func (client ExtendedClient) Enable(ctx context.Context, cwKey string, instanceIds string) (result SchemaExtendedEnableResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedClient.Enable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnablePreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client ExtendedClient) EnablePreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"cw_key":      autorest.Encode("query", cwKey),
		"instanceIds": autorest.Encode("query", instanceIds),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("PUT", autorest.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/schema/extended/enable")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/enable"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedClient) EnableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client ExtendedClient) EnableResponder(resp *http.Response) (result SchemaExtendedEnableResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Status instance의 Extended 설정 여부를 조회합니다.
// Parameters:
// cwKey - 상품의 cw_key
// instanceIds - extended 설정 여부를 조회하고자하는 instanceId
func (client ExtendedClient) Status(ctx context.Context, cwKey string, instanceIds string) (result ListSchemaExtendedStatusParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExtendedClient.Status")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StatusPreparer(ctx, cwKey, instanceIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Status", nil, "Failure preparing request")
		return
	}

	resp, err := client.StatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Status", resp, "Failure sending request")
		return
	}

	result, err = client.StatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudinsight.ExtendedClient", "Status", resp, "Failure responding to request")
	}

	return
}

// StatusPreparer prepares the Status request.
func (client ExtendedClient) StatusPreparer(ctx context.Context, cwKey string, instanceIds string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"cw_key":      autorest.Encode("query", cwKey),
		"instanceIds": autorest.Encode("query", instanceIds),
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/cw_fea/real/cw/api/schema/extended/status")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/schema/extended/status"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-api-key", client.Client.APIGatewayAPIKey),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StatusSender sends the Status request. The method will close the
// http.Response Body if it receives an error.
func (client ExtendedClient) StatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StatusResponder handles the response to the Status request. The method always
// closes the http.Response Body.
func (client ExtendedClient) StatusResponder(resp *http.Response) (result ListSchemaExtendedStatusParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
