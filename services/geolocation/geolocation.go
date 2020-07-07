package geolocation

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the geoLocation Client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// Get 위치정보 조회
// Parameters:
// IP - 쿼리 IP 주소 (x.x.x.x)
// enc - 인코딩 타입 utf8(기본값) 또는 euckr
// ext - 추가 정보 포함 여부 t 또는 f(기본값)
// responseFormatType - 응답 결과 포맷 xml 또는 json(기본값)
func (client Client) Get(ctx context.Context, IP string, enc string, ext string, responseFormatType string) (result ResultParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, IP, enc, ext, responseFormatType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "geolocation.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "geolocation.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, IP string, enc string, ext string, responseFormatType string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"ip": autorest.Encode("query", IP),
	}
	if len(enc) > 0 {
		queryParameters["enc"] = autorest.Encode("query", enc)
	} else {
		queryParameters["enc"] = autorest.Encode("query", "utf8")
	}
	if len(ext) > 0 {
		queryParameters["ext"] = autorest.Encode("query", ext)
	} else {
		queryParameters["ext"] = autorest.Encode("query", "f")
	}
	if len(responseFormatType) > 0 {
		queryParameters["responseFormatType"] = autorest.Encode("query", responseFormatType)
	} else {
		queryParameters["responseFormatType"] = autorest.Encode("query", "json")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/geoLocation"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result ResultParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
