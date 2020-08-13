package registry

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ContainerRegistryImageClient is the container Registry Client
type ContainerRegistryImageClient struct {
	BaseClient
}

// NewContainerRegistryImageClient creates an instance of the ContainerRegistryImageClient client.
func NewContainerRegistryImageClient() ContainerRegistryImageClient {
	return NewContainerRegistryImageClientWithBaseURI(DefaultBaseURI)
}

// NewContainerRegistryImageClientWithBaseURI creates an instance of the ContainerRegistryImageClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewContainerRegistryImageClientWithBaseURI(baseURI string) ContainerRegistryImageClient {
	return ContainerRegistryImageClient{NewWithBaseURI(baseURI)}
}

// GetTagList 레지스트리의 이미지에 등록된 Tag 목록을 반환합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
// page - start page number. It must be greater than zero.
// pagesize - start
func (client ContainerRegistryImageClient) GetTagList(ctx context.Context, registry string, imageName string, page *int32, pagesize *int32) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryImageClient.GetTagList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTagListPreparer(ctx, registry, imageName, page, pagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTagListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagList", resp, "Failure sending request")
		return
	}

	result, err = client.GetTagListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagList", resp, "Failure responding to request")
	}

	return
}

// GetTagListPreparer prepares the GetTagList request.
func (client ContainerRegistryImageClient) GetTagListPreparer(ctx context.Context, registry string, imageName string, page *int32, pagesize *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"registry":  autorest.Encode("path", registry),
	}

	queryParameters := map[string]interface{}{}
	if page != nil {
		queryParameters["page"] = autorest.Encode("query", *page)
	}
	if pagesize != nil {
		queryParameters["pagesize"] = autorest.Encode("query", *pagesize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}/tags", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTagListSender sends the GetTagList request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryImageClient) GetTagListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTagListResponder handles the response to the GetTagList request. The method always
// closes the http.Response Body.
func (client ContainerRegistryImageClient) GetTagListResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTagReference 레지스트리의 이미지에 등록된 특정 Tag의 상세 정보를 반환합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
// reference - 상세 정보를 조회할 태그 이름
func (client ContainerRegistryImageClient) GetTagReference(ctx context.Context, registry string, imageName string, reference string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryImageClient.GetTagReference")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetTagReferencePreparer(ctx, registry, imageName, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagReference", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTagReferenceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagReference", resp, "Failure sending request")
		return
	}

	result, err = client.GetTagReferenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryImageClient", "GetTagReference", resp, "Failure responding to request")
	}

	return
}

// GetTagReferencePreparer prepares the GetTagReference request.
func (client ContainerRegistryImageClient) GetTagReferencePreparer(ctx context.Context, registry string, imageName string, reference string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"reference": autorest.Encode("path", reference),
		"registry":  autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}/tags/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTagReferenceSender sends the GetTagReference request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryImageClient) GetTagReferenceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetTagReferenceResponder handles the response to the GetTagReference request. The method always
// closes the http.Response Body.
func (client ContainerRegistryImageClient) GetTagReferenceResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
