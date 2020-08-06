package kms

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the KMS Client
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

// CreateCustomKey 임의의 high-entropy 키를 주어진 bit로 생성합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) CreateCustomKey(ctx context.Context, keyTag string, parameters CreateCustomKeyRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CreateCustomKey")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.RequestPlainKey", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "CreateCustomKey", err.Error())
	}

	req, err := client.CreateCustomKeyPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "CreateCustomKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateCustomKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "CreateCustomKey", resp, "Failure sending request")
		return
	}

	result, err = client.CreateCustomKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "CreateCustomKey", resp, "Failure responding to request")
	}

	return
}

// CreateCustomKeyPreparer prepares the CreateCustomKey request.
func (client Client) CreateCustomKeyPreparer(ctx context.Context, keyTag string, parameters CreateCustomKeyRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/createCustomKey", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateCustomKeySender sends the CreateCustomKey request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateCustomKeySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateCustomKeyResponder handles the response to the CreateCustomKey request. The method always
// closes the http.Response Body.
func (client Client) CreateCustomKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Descrypt 지정된 마스터키로 암호문을 복호화합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) Descrypt(ctx context.Context, keyTag string, parameters DecryptRequestKey) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Descrypt")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Ciphertext", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "Descrypt", err.Error())
	}

	req, err := client.DescryptPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Descrypt", nil, "Failure preparing request")
		return
	}

	resp, err := client.DescryptSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "Descrypt", resp, "Failure sending request")
		return
	}

	result, err = client.DescryptResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Descrypt", resp, "Failure responding to request")
	}

	return
}

// DescryptPreparer prepares the Descrypt request.
func (client Client) DescryptPreparer(ctx context.Context, keyTag string, parameters DecryptRequestKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/decrypt", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DescryptSender sends the Descrypt request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DescryptSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DescryptResponder handles the response to the Descrypt request. The method always
// closes the http.Response Body.
func (client Client) DescryptResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Encrypt 지정된 마스터키의 현재 버전으로 최대 32KB 데이터를 암호화합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) Encrypt(ctx context.Context, keyTag string, parameters EncryptRequestKey) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Encrypt")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Plaintext", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "Encrypt", err.Error())
	}

	req, err := client.EncryptPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Encrypt", nil, "Failure preparing request")
		return
	}

	resp, err := client.EncryptSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "Encrypt", resp, "Failure sending request")
		return
	}

	result, err = client.EncryptResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Encrypt", resp, "Failure responding to request")
	}

	return
}

// EncryptPreparer prepares the Encrypt request.
func (client Client) EncryptPreparer(ctx context.Context, keyTag string, parameters EncryptRequestKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/encrypt", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EncryptSender sends the Encrypt request. The method will close the
// http.Response Body if it receives an error.
func (client Client) EncryptSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// EncryptResponder handles the response to the Encrypt request. The method always
// closes the http.Response Body.
func (client Client) EncryptResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Reencrypt 암호문을 지정된 마스터키의 가장 최신 버전으로 재암호화합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) Reencrypt(ctx context.Context, keyTag string, parameters ReencryptRequestKey) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Reencrypt")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Ciphertext", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "Reencrypt", err.Error())
	}

	req, err := client.ReencryptPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Reencrypt", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReencryptSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "Reencrypt", resp, "Failure sending request")
		return
	}

	result, err = client.ReencryptResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Reencrypt", resp, "Failure responding to request")
	}

	return
}

// ReencryptPreparer prepares the Reencrypt request.
func (client Client) ReencryptPreparer(ctx context.Context, keyTag string, parameters ReencryptRequestKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/reencrypt", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReencryptSender sends the Reencrypt request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ReencryptSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ReencryptResponder handles the response to the Reencrypt request. The method always
// closes the http.Response Body.
func (client Client) ReencryptResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Sign 지정된 마스터키(비밀키)로 최대 8KB 데이터의 서명값을 생성합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) Sign(ctx context.Context, keyTag string, parameters SignRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Sign")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Data", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "Sign", err.Error())
	}

	req, err := client.SignPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Sign", nil, "Failure preparing request")
		return
	}

	resp, err := client.SignSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "Sign", resp, "Failure sending request")
		return
	}

	result, err = client.SignResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Sign", resp, "Failure responding to request")
	}

	return
}

// SignPreparer prepares the Sign request.
func (client Client) SignPreparer(ctx context.Context, keyTag string, parameters SignRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/sign", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SignSender sends the Sign request. The method will close the
// http.Response Body if it receives an error.
func (client Client) SignSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// SignResponder handles the response to the Sign request. The method always
// closes the http.Response Body.
func (client Client) SignResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Verify 지정된 마스터키(공개키)로 최대 8KB 데이터와 서명값을 비교하여 일치 여부를 반환합니다.
// Parameters:
// keyTag - 생성된 키 봉인에 사용할 마스터키의 tag 값
// parameters - 요청 바디
func (client Client) Verify(ctx context.Context, keyTag string, parameters VerifyRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Verify")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Data", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Signature", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kms.Client", "Verify", err.Error())
	}

	req, err := client.VerifyPreparer(ctx, keyTag, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Verify", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kms.Client", "Verify", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kms.Client", "Verify", resp, "Failure responding to request")
	}

	return
}

// VerifyPreparer prepares the Verify request.
func (client Client) VerifyPreparer(ctx context.Context, keyTag string, parameters VerifyRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"keyTag": autorest.Encode("path", keyTag),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{keyTag}/verify", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VerifySender sends the Verify request. The method will close the
// http.Response Body if it receives an error.
func (client Client) VerifySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// VerifyResponder handles the response to the Verify request. The method always
// closes the http.Response Body.
func (client Client) VerifyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
