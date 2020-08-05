package server

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ImageProductClient is the server Client
type ImageProductClient struct {
	BaseClient
}

// NewImageProductClient creates an instance of the ImageProductClient client.
func NewImageProductClient() ImageProductClient {
	return NewImageProductClientWithBaseURI(DefaultBaseURI)
}

// NewImageProductClientWithBaseURI creates an instance of the ImageProductClient client.
func NewImageProductClientWithBaseURI(baseURI string) ImageProductClient {
	return ImageProductClient{NewWithBaseURI(baseURI)}
}

// GetList 서버 이미지 상품 리스트를 조회
// Parameters:
// responseFormatType - 반환 데이터 포맷 타입
// regionCode - REGION 코드
// blockStorageSize - 블록스토리지 사이즈
// exclusionProductCode - 제외할 상품 코드
// productCode - 조회할 상품 코드
// platformTypeCodeListN - 플랫폼 유형 코드 리스트
func (client ImageProductClient) GetList(ctx context.Context, responseFormatType string, regionCode string, blockStorageSize string, exclusionProductCode string, productCode string, platformTypeCodeListN PlatformTypeCode) (result ImageProductListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ImageProductClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, responseFormatType, regionCode, blockStorageSize, exclusionProductCode, productCode, platformTypeCodeListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageProductClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "server.ImageProductClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "server.ImageProductClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ImageProductClient) GetListPreparer(ctx context.Context, responseFormatType string, regionCode string, blockStorageSize string, exclusionProductCode string, productCode string, platformTypeCodeListN PlatformTypeCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(blockStorageSize) > 0 {
		queryParameters["blockStorageSize"] = autorest.Encode("query", blockStorageSize)
	}
	if len(exclusionProductCode) > 0 {
		queryParameters["exclusionProductCode"] = autorest.Encode("query", exclusionProductCode)
	}
	if len(productCode) > 0 {
		queryParameters["productCode"] = autorest.Encode("query", productCode)
	}
	if len(string(platformTypeCodeListN)) > 0 {
		queryParameters["platformTypeCodeList.N"] = autorest.Encode("query", platformTypeCodeListN)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getServerImageProductList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ImageProductClient) GetListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ImageProductClient) GetListResponder(resp *http.Response) (result ImageProductListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}