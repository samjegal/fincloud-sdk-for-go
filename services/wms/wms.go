package wms

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the WMS Client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// Detail wMS에 등록되어 있는 URL 및 SCENARIO의 세부 정보를 조회
// Parameters:
// scriptID - 스크립트 아이디
func (client Client) Detail(ctx context.Context, scriptID string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Detail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetailPreparer(ctx, scriptID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Detail", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "wms.Client", "Detail", resp, "Failure sending request")
		return
	}

	result, err = client.DetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Detail", resp, "Failure responding to request")
	}

	return
}

// DetailPreparer prepares the Detail request.
func (client Client) DetailPreparer(ctx context.Context, scriptID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scriptId": autorest.Encode("path", scriptID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/sms/v1/scripts/{scriptId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetailSender sends the Detail request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetailResponder handles the response to the Detail request. The method always
// closes the http.Response Body.
func (client Client) DetailResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List wMS에 등록되어 있는 URL 및 SCENARIO 전체 목록을 조회
func (client Client) List(ctx context.Context) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "wms.Client", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/sms/v1/scripts"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Result wMS에 등록되어 있는 URL 및 SCENARIO의 모니터링 결과를 조회
// Parameters:
// scriptID - 스크립트 아이디
// timestamp - 조회 시간 기준
func (client Client) Result(ctx context.Context, scriptID string, timestamp string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Result")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResultPreparer(ctx, scriptID, timestamp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Result", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResultSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "wms.Client", "Result", resp, "Failure sending request")
		return
	}

	result, err = client.ResultResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Result", resp, "Failure responding to request")
	}

	return
}

// ResultPreparer prepares the Result request.
func (client Client) ResultPreparer(ctx context.Context, scriptID string, timestamp string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scriptId": autorest.Encode("path", scriptID),
	}

	queryParameters := map[string]interface{}{}
	if len(timestamp) > 0 {
		queryParameters["timestamp"] = autorest.Encode("query", timestamp)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/sms/v1/scripts/{scriptId}/result", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResultSender sends the Result request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ResultSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResultResponder handles the response to the Result request. The method always
// closes the http.Response Body.
func (client Client) ResultResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start wMS에 등록되어 있는 URL 및 SCENARIO의 상태를 “시작”으로 변경
// Parameters:
// scriptID - 스크립트 아이디
func (client Client) Start(ctx context.Context, scriptID string) (result ReturnParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Start")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, scriptID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Start", nil, "Failure preparing request")
		return
	}

	resp, err := client.StartSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "wms.Client", "Start", resp, "Failure sending request")
		return
	}

	result, err = client.StartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Start", resp, "Failure responding to request")
	}

	return
}

// StartPreparer prepares the Start request.
func (client Client) StartPreparer(ctx context.Context, scriptID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scriptId": autorest.Encode("path", scriptID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/sms/v1/scripts/{scriptId}/setting/status/start", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client Client) StartSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client Client) StartResponder(resp *http.Response) (result ReturnParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop wMS에 등록되어 있는 URL 및 SCENARIO의 상태를 “정지”로 변경
// Parameters:
// scriptID - 스크립트 아이디
func (client Client) Stop(ctx context.Context, scriptID string) (result ReturnParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Stop")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, scriptID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "wms.Client", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "wms.Client", "Stop", resp, "Failure responding to request")
	}

	return
}

// StopPreparer prepares the Stop request.
func (client Client) StopPreparer(ctx context.Context, scriptID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scriptId": autorest.Encode("path", scriptID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/sms/v1/scripts/{scriptId}/setting/status/stop", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client Client) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client Client) StopResponder(resp *http.Response) (result ReturnParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusUnauthorized, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
