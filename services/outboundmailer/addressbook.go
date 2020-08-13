package outboundmailer

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AddressBookClient is the cloud Outbound Mailer Client
type AddressBookClient struct {
	BaseClient
}

// NewAddressBookClient creates an instance of the AddressBookClient client.
func NewAddressBookClient() AddressBookClient {
	return NewAddressBookClientWithBaseURI(DefaultBaseURI)
}

// NewAddressBookClientWithBaseURI creates an instance of the AddressBookClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAddressBookClientWithBaseURI(baseURI string) AddressBookClient {
	return AddressBookClient{NewWithBaseURI(baseURI)}
}

// Create 수신자 주소/그룹 일괄로 입력받아 주소록 생성을 요청
// Parameters:
// requestBody - 주소록 생성 요청
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) Create(ctx context.Context, requestBody AddressBookGenerateRequest, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: requestBody,
			Constraints: []validation.Constraint{{Target: "requestBody.Groups", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("outboundmailer.AddressBookClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, requestBody, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AddressBookClient) CreatePreparer(ctx context.Context, requestBody AddressBookGenerateRequest, xNCPLANG string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book"),
		autorest.WithJSON(requestBody))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AddressBookClient) CreateResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 주소록 초기화
// Parameters:
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) Delete(ctx context.Context, xNCPLANG string) (result AddressBookInitResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AddressBookClient) DeletePreparer(ctx context.Context, xNCPLANG string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book"))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AddressBookClient) DeleteResponder(resp *http.Response) (result AddressBookInitResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAddress 이메일 주소 삭제
// Parameters:
// requestBody - 수신자 주소 삭제 요청
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) DeleteAddress(ctx context.Context, requestBody AddressBookDeleteAddressRequest, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.DeleteAddress")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: requestBody,
			Constraints: []validation.Constraint{{Target: "requestBody.EmailAddresses", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("outboundmailer.AddressBookClient", "DeleteAddress", err.Error())
	}

	req, err := client.DeleteAddressPreparer(ctx, requestBody, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteAddress", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAddressSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteAddress", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAddressResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteAddress", resp, "Failure responding to request")
	}

	return
}

// DeleteAddressPreparer prepares the DeleteAddress request.
func (client AddressBookClient) DeleteAddressPreparer(ctx context.Context, requestBody AddressBookDeleteAddressRequest, xNCPLANG string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book/address"),
		autorest.WithJSON(requestBody))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAddressSender sends the DeleteAddress request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) DeleteAddressSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAddressResponder handles the response to the DeleteAddress request. The method always
// closes the http.Response Body.
func (client AddressBookClient) DeleteAddressResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteEmptyRecipientGroup 수신자 그룹에서 모든 이메일 주소 연관 정보 삭제 (비우기)
// Parameters:
// groupName - 수신자 그룸명
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) DeleteEmptyRecipientGroup(ctx context.Context, groupName string, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.DeleteEmptyRecipientGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteEmptyRecipientGroupPreparer(ctx, groupName, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteEmptyRecipientGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteEmptyRecipientGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteEmptyRecipientGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteEmptyRecipientGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteEmptyRecipientGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteEmptyRecipientGroupPreparer prepares the DeleteEmptyRecipientGroup request.
func (client AddressBookClient) DeleteEmptyRecipientGroupPreparer(ctx context.Context, groupName string, xNCPLANG string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"groupName": autorest.Encode("query", groupName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book/recipient-groups/address/empty"),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteEmptyRecipientGroupSender sends the DeleteEmptyRecipientGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) DeleteEmptyRecipientGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteEmptyRecipientGroupResponder handles the response to the DeleteEmptyRecipientGroup request. The method always
// closes the http.Response Body.
func (client AddressBookClient) DeleteEmptyRecipientGroupResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteRecipientGroup 수신자 그룹 삭제
// Parameters:
// groupName - 수신자 그룸명
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) DeleteRecipientGroup(ctx context.Context, groupName string, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.DeleteRecipientGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteRecipientGroupPreparer(ctx, groupName, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteRecipientGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteRecipientGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteRecipientGroupPreparer prepares the DeleteRecipientGroup request.
func (client AddressBookClient) DeleteRecipientGroupPreparer(ctx context.Context, groupName string, xNCPLANG string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"groupName": autorest.Encode("query", groupName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book/recipient-groups"),
		autorest.WithQueryParameters(queryParameters))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteRecipientGroupSender sends the DeleteRecipientGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) DeleteRecipientGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteRecipientGroupResponder handles the response to the DeleteRecipientGroup request. The method always
// closes the http.Response Body.
func (client AddressBookClient) DeleteRecipientGroupResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteRecipientGroupByAddress 수신자 그룹에서 이메일 주소 삭제
// Parameters:
// requestBody - 수신자 그룹에서 이메일 주소 삭제 요청
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) DeleteRecipientGroupByAddress(ctx context.Context, requestBody AddressBookDeleteRelationRequest, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.DeleteRecipientGroupByAddress")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: requestBody,
			Constraints: []validation.Constraint{{Target: "requestBody.GroupName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "requestBody.EmailAddresses", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("outboundmailer.AddressBookClient", "DeleteRecipientGroupByAddress", err.Error())
	}

	req, err := client.DeleteRecipientGroupByAddressPreparer(ctx, requestBody, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroupByAddress", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteRecipientGroupByAddressSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroupByAddress", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteRecipientGroupByAddressResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "DeleteRecipientGroupByAddress", resp, "Failure responding to request")
	}

	return
}

// DeleteRecipientGroupByAddressPreparer prepares the DeleteRecipientGroupByAddress request.
func (client AddressBookClient) DeleteRecipientGroupByAddressPreparer(ctx context.Context, requestBody AddressBookDeleteRelationRequest, xNCPLANG string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book/recipient-groups/address"),
		autorest.WithJSON(requestBody))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteRecipientGroupByAddressSender sends the DeleteRecipientGroupByAddress request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) DeleteRecipientGroupByAddressSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteRecipientGroupByAddressResponder handles the response to the DeleteRecipientGroupByAddress request. The method always
// closes the http.Response Body.
func (client AddressBookClient) DeleteRecipientGroupByAddressResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get 주소록 현황 조회
// Parameters:
// xNCPLANG - 언어 (ko-KR, en-US, zh-CN), default:en-US
func (client AddressBookClient) Get(ctx context.Context, xNCPLANG string) (result AddressBookResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AddressBookClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, xNCPLANG)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "outboundmailer.AddressBookClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AddressBookClient) GetPreparer(ctx context.Context, xNCPLANG string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/v1/address-book"))
	if len(xNCPLANG) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-NCP-LANG", autorest.String(xNCPLANG)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AddressBookClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AddressBookClient) GetResponder(resp *http.Response) (result AddressBookResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
