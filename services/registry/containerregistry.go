package registry

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ContainerRegistryClient is the container Registry Client
type ContainerRegistryClient struct {
	BaseClient
}

// NewContainerRegistryClient creates an instance of the ContainerRegistryClient client.
func NewContainerRegistryClient() ContainerRegistryClient {
	return NewContainerRegistryClientWithBaseURI(DefaultBaseURI)
}

// NewContainerRegistryClientWithBaseURI creates an instance of the ContainerRegistryClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewContainerRegistryClientWithBaseURI(baseURI string) ContainerRegistryClient {
	return ContainerRegistryClient{NewWithBaseURI(baseURI)}
}

// Create object Storage의 Bucket과 연동된 Registry를 생성합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// parameters - 요청 바디
func (client ContainerRegistryClient) Create(ctx context.Context, registry string, parameters RepositoryRequest) (result MessageResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Bucket", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("registry.ContainerRegistryClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, registry, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ContainerRegistryClient) CreatePreparer(ctx context.Context, registry string, parameters RepositoryRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registry": autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) CreateResponder(resp *http.Response) (result MessageResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusConflict, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 특정 Registry를 삭제합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
func (client ContainerRegistryClient) Delete(ctx context.Context, registry string) (result MessageResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, registry)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContainerRegistryClient) DeletePreparer(ctx context.Context, registry string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registry": autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) DeleteResponder(resp *http.Response) (result MessageResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteImage 레지스트리에 등록된 이미지를 삭제합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
func (client ContainerRegistryClient) DeleteImage(ctx context.Context, registry string, imageName string) (result MessageResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.DeleteImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteImagePreparer(ctx, registry, imageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteImage", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteImage", resp, "Failure responding to request")
	}

	return
}

// DeleteImagePreparer prepares the DeleteImage request.
func (client ContainerRegistryClient) DeleteImagePreparer(ctx context.Context, registry string, imageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"registry":  autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteImageSender sends the DeleteImage request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) DeleteImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteImageResponder handles the response to the DeleteImage request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) DeleteImageResponder(resp *http.Response) (result MessageResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteTagReference 레지스트리의 이미지에 등록된 특정 Tag를 제거합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
// reference - 상세 정보를 조회할 태그 이름
func (client ContainerRegistryClient) DeleteTagReference(ctx context.Context, registry string, imageName string, reference string) (result MessageResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.DeleteTagReference")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteTagReferencePreparer(ctx, registry, imageName, reference)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteTagReference", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteTagReferenceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteTagReference", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteTagReferenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "DeleteTagReference", resp, "Failure responding to request")
	}

	return
}

// DeleteTagReferencePreparer prepares the DeleteTagReference request.
func (client ContainerRegistryClient) DeleteTagReferencePreparer(ctx context.Context, registry string, imageName string, reference string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"reference": autorest.Encode("path", reference),
		"registry":  autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}/tags/{reference}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteTagReferenceSender sends the DeleteTagReference request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) DeleteTagReferenceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteTagReferenceResponder handles the response to the DeleteTagReference request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) DeleteTagReferenceResponder(resp *http.Response) (result MessageResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetImageDetail 레지스트리 내 등록된 Image의 상세 정보를 반환합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
func (client ContainerRegistryClient) GetImageDetail(ctx context.Context, registry string, imageName string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.GetImageDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetImageDetailPreparer(ctx, registry, imageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetImageDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetImageDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageDetail", resp, "Failure responding to request")
	}

	return
}

// GetImageDetailPreparer prepares the GetImageDetail request.
func (client ContainerRegistryClient) GetImageDetailPreparer(ctx context.Context, registry string, imageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"registry":  autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetImageDetailSender sends the GetImageDetail request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) GetImageDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetImageDetailResponder handles the response to the GetImageDetail request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) GetImageDetailResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetImageList 레지스트리 내 등록된 Image들의 목록을 반환합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// page - start page number. It must be greater than zero.
// pagesize - start
func (client ContainerRegistryClient) GetImageList(ctx context.Context, registry string, page *int32, pagesize *int32) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.GetImageList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetImageListPreparer(ctx, registry, page, pagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetImageListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageList", resp, "Failure sending request")
		return
	}

	result, err = client.GetImageListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetImageList", resp, "Failure responding to request")
	}

	return
}

// GetImageListPreparer prepares the GetImageList request.
func (client ContainerRegistryClient) GetImageListPreparer(ctx context.Context, registry string, page *int32, pagesize *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registry": autorest.Encode("path", registry),
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
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetImageListSender sends the GetImageList request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) GetImageListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetImageListResponder handles the response to the GetImageList request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) GetImageListResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList container Registry에 등록된 Registry 목록을 반환합니다.
// Parameters:
// page - start page number. It must be greater than zero.
// pagesize - start
func (client ContainerRegistryClient) GetList(ctx context.Context, page *int32, pagesize *int32) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, page, pagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client ContainerRegistryClient) GetListPreparer(ctx context.Context, page *int32, pagesize *int32) (*http.Request, error) {
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
		autorest.WithPath("/ncr/api/v2/repositories"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) GetListResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateImage 레지스트리에 등록된 이미지에 대한 description, full_description을 업데이트 합니다.
// Parameters:
// registry - 이미지 목록을 조회할 레지스트리 이름
// imageName - 상세 정보 조회 대상 이미지 이름
// parameters - 요청 바디
func (client ContainerRegistryClient) UpdateImage(ctx context.Context, registry string, imageName string, parameters UpdateRepositoryImageRequest) (result MessageResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerRegistryClient.UpdateImage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateImagePreparer(ctx, registry, imageName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "UpdateImage", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateImageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "UpdateImage", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateImageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registry.ContainerRegistryClient", "UpdateImage", resp, "Failure responding to request")
	}

	return
}

// UpdateImagePreparer prepares the UpdateImage request.
func (client ContainerRegistryClient) UpdateImagePreparer(ctx context.Context, registry string, imageName string, parameters UpdateRepositoryImageRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"imageName": autorest.Encode("path", imageName),
		"registry":  autorest.Encode("path", registry),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/ncr/api/v2/repositories/{registry}/{imageName}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateImageSender sends the UpdateImage request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerRegistryClient) UpdateImageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateImageResponder handles the response to the UpdateImage request. The method always
// closes the http.Response Body.
func (client ContainerRegistryClient) UpdateImageResponder(resp *http.Response) (result MessageResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
