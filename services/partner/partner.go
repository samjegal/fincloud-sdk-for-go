package partner

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the partner Client
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

// GetContractProductDemandList 계약 상품 청구 리스트
// Parameters:
// responseFormatType - 리턴 타입
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// startMonth - 시작 년월
// endMonth - 마지막 년월
// loginID - 로그인 아이디
func (client Client) GetContractProductDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result ContractProductDemandListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetContractProductDemandList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetContractProductDemandListPreparer(ctx, responseFormatType, pageNo, pageSize, startMonth, endMonth, loginID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetContractProductDemandList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContractProductDemandListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "partner.Client", "GetContractProductDemandList", resp, "Failure sending request")
		return
	}

	result, err = client.GetContractProductDemandListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetContractProductDemandList", resp, "Failure responding to request")
	}

	return
}

// GetContractProductDemandListPreparer prepares the GetContractProductDemandList request.
func (client Client) GetContractProductDemandListPreparer(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if pageNo != nil {
		queryParameters["pageNo"] = autorest.Encode("query", *pageNo)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}
	if len(startMonth) > 0 {
		queryParameters["startMonth"] = autorest.Encode("query", startMonth)
	}
	if len(endMonth) > 0 {
		queryParameters["endMonth"] = autorest.Encode("query", endMonth)
	}
	if len(loginID) > 0 {
		queryParameters["loginId"] = autorest.Encode("query", loginID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/billing/v2/getContractProductDemandList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContractProductDemandListSender sends the GetContractProductDemandList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetContractProductDemandListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetContractProductDemandListResponder handles the response to the GetContractProductDemandList request. The method always
// closes the http.Response Body.
func (client Client) GetContractProductDemandListResponder(resp *http.Response) (result ContractProductDemandListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDemandList 청구 리스트
// Parameters:
// responseFormatType - 리턴 타입
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// startMonth - 시작 년월
// endMonth - 마지막 년월
// loginID - 로그인 아이디
func (client Client) GetDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result DemandListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDemandList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDemandListPreparer(ctx, responseFormatType, pageNo, pageSize, startMonth, endMonth, loginID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDemandListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandList", resp, "Failure sending request")
		return
	}

	result, err = client.GetDemandListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandList", resp, "Failure responding to request")
	}

	return
}

// GetDemandListPreparer prepares the GetDemandList request.
func (client Client) GetDemandListPreparer(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if pageNo != nil {
		queryParameters["pageNo"] = autorest.Encode("query", *pageNo)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}
	if len(startMonth) > 0 {
		queryParameters["startMonth"] = autorest.Encode("query", startMonth)
	}
	if len(endMonth) > 0 {
		queryParameters["endMonth"] = autorest.Encode("query", endMonth)
	}
	if len(loginID) > 0 {
		queryParameters["loginId"] = autorest.Encode("query", loginID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/billing/v2/getDemandList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDemandListSender sends the GetDemandList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDemandListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDemandListResponder handles the response to the GetDemandList request. The method always
// closes the http.Response Body.
func (client Client) GetDemandListResponder(resp *http.Response) (result DemandListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDemandProductList 청구 상품 리스트
// Parameters:
// responseFormatType - 리턴 타입
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// startMonth - 시작 년월
// endMonth - 마지막 년월
// loginID - 로그인 아이디
func (client Client) GetDemandProductList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result DemandProductListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDemandProductList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDemandProductListPreparer(ctx, responseFormatType, pageNo, pageSize, startMonth, endMonth, loginID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandProductList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDemandProductListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandProductList", resp, "Failure sending request")
		return
	}

	result, err = client.GetDemandProductListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetDemandProductList", resp, "Failure responding to request")
	}

	return
}

// GetDemandProductListPreparer prepares the GetDemandProductList request.
func (client Client) GetDemandProductListPreparer(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if pageNo != nil {
		queryParameters["pageNo"] = autorest.Encode("query", *pageNo)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}
	if len(startMonth) > 0 {
		queryParameters["startMonth"] = autorest.Encode("query", startMonth)
	}
	if len(endMonth) > 0 {
		queryParameters["endMonth"] = autorest.Encode("query", endMonth)
	}
	if len(loginID) > 0 {
		queryParameters["loginId"] = autorest.Encode("query", loginID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/billing/v2/getDemandProductList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDemandProductListSender sends the GetDemandProductList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDemandProductListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDemandProductListResponder handles the response to the GetDemandProductList request. The method always
// closes the http.Response Body.
func (client Client) GetDemandProductListResponder(resp *http.Response) (result DemandProductListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPartnerDemandList 파트너 청구 리스트
// Parameters:
// responseFormatType - 리턴 타입
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// startMonth - 시작 년월
// endMonth - 마지막 년월
// loginID - 로그인 아이디
func (client Client) GetPartnerDemandList(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (result DemandListResponseType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetPartnerDemandList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPartnerDemandListPreparer(ctx, responseFormatType, pageNo, pageSize, startMonth, endMonth, loginID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetPartnerDemandList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPartnerDemandListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "partner.Client", "GetPartnerDemandList", resp, "Failure sending request")
		return
	}

	result, err = client.GetPartnerDemandListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "partner.Client", "GetPartnerDemandList", resp, "Failure responding to request")
	}

	return
}

// GetPartnerDemandListPreparer prepares the GetPartnerDemandList request.
func (client Client) GetPartnerDemandListPreparer(ctx context.Context, responseFormatType string, pageNo *int32, pageSize *int32, startMonth string, endMonth string, loginID string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", responseFormatType),
	}
	if pageNo != nil {
		queryParameters["pageNo"] = autorest.Encode("query", *pageNo)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}
	if len(startMonth) > 0 {
		queryParameters["startMonth"] = autorest.Encode("query", startMonth)
	}
	if len(endMonth) > 0 {
		queryParameters["endMonth"] = autorest.Encode("query", endMonth)
	}
	if len(loginID) > 0 {
		queryParameters["loginId"] = autorest.Encode("query", loginID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/billing/v2/getPartnerDemandList"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPartnerDemandListSender sends the GetPartnerDemandList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetPartnerDemandListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetPartnerDemandListResponder handles the response to the GetPartnerDemandList request. The method always
// closes the http.Response Body.
func (client Client) GetPartnerDemandListResponder(resp *http.Response) (result DemandListResponseType, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
