package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"io"
	"net/http"
)

// FileClient is the cloud Outbound Mailer Client
type FileClient struct {
	BaseClient
}

// NewFileClient creates an instance of the FileClient client.
func NewFileClient() FileClient {
	return NewFileClientWithBaseURI(DefaultBaseURI)
}

// NewFileClientWithBaseURI creates an instance of the FileClient client.
func NewFileClientWithBaseURI(baseURI string) FileClient {
	return FileClient{NewWithBaseURI(baseURI)}
}

// Create 파일을 업로드 할 수 있는 기능을 제공
// Parameters:
// fileList - 파일 업로드
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client FileClient) Create(ctx context.Context, fileList io.ReadCloser, xNCPLANG string) (result FileUploadResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, fileList, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client FileClient) CreatePreparer(ctx context.Context, fileList io.ReadCloser, xNCPLANG string) (*http.Request, error) {
	formDataParameters := map[string]interface{}{
		"fileList": fileList,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/files"),
		autorest.WithMultiPartFormData(formDataParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client FileClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client FileClient) CreateResponder(resp *http.Response) (result FileUploadResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 저장된 파일 삭제
// Parameters:
// tempRequestID - tempRequestId
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client FileClient) Delete(ctx context.Context, tempRequestID string, xNCPLANG string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, tempRequestID, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FileClient) DeletePreparer(ctx context.Context, tempRequestID string, xNCPLANG string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tempRequestId": autorest.Encode("path", tempRequestID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/files/{tempRequestId}", pathParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FileClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FileClient) DeleteResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get 저장된 파일 조회
// Parameters:
// tempRequestID - tempRequestId
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client FileClient) Get(ctx context.Context, tempRequestID string, xNCPLANG string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FileClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, tempRequestID, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.FileClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FileClient) GetPreparer(ctx context.Context, tempRequestID string, xNCPLANG string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tempRequestId": autorest.Encode("path", tempRequestID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/files/{tempRequestId}", pathParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FileClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FileClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusForbidden, http.StatusMethodNotAllowed, http.StatusUnsupportedMediaType, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
