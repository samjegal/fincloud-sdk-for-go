package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProductClient is the server Client
type ProductClient struct {
	BaseClient
}

// NewProductClient creates an instance of the ProductClient client.
func NewProductClient() ProductClient {
	return NewProductClientWithBaseURI(DefaultBaseURI)
}

// NewProductClientWithBaseURI creates an instance of the ProductClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProductClientWithBaseURI(baseURI string) ProductClient {
	return ProductClient{NewWithBaseURI(baseURI)}
}

// GetList 서버 상품 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// serverImageProductCode - 서버 이미지 상품 코드
// regionCode - REGION 코드
// zoneCode - ZONE 코드
// exclusionProductCode - 제외할 상품 코드
// productCode - 조회할 상품 코드
// generationCode - 세대 코드
func (client ProductClient) GetList(ctx context.Context, responseFormatType string, serverImageProductCode string, regionCode string, zoneCode string, exclusionProductCode string, productCode string, generationCode string) (result ProductListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, serverImageProductCode, regionCode, zoneCode, exclusionProductCode, productCode, generationCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ProductClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ProductClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ProductClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ProductClient) GetListPreparer(ctx context.Context, responseFormatType string, serverImageProductCode string, regionCode string, zoneCode string, exclusionProductCode string, productCode string, generationCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":     autorest.Encode("query", responseFormatType),
		"serverImageProductCode": autorest.Encode("query", serverImageProductCode),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}
	if len(exclusionProductCode) > 0 {
		queryParameters["exclusionProductCode"] = autorest.Encode("query", exclusionProductCode)
	}
	if len(productCode) > 0 {
		queryParameters["productCode"] = autorest.Encode("query", productCode)
	}
	if len(generationCode) > 0 {
		queryParameters["generationCode"] = autorest.Encode("query", generationCode)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vserver/v2/getServerProductList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ProductClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ProductClient) GetListResponder(resp *http.Response) (result ProductListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
