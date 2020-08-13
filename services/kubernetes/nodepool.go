package kubernetes

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NodePoolClient is the kubernetes Client
type NodePoolClient struct {
	BaseClient
}

// NewNodePoolClient creates an instance of the NodePoolClient client.
func NewNodePoolClient() NodePoolClient {
	return NewNodePoolClientWithBaseURI(DefaultBaseURI)
}

// NewNodePoolClientWithBaseURI creates an instance of the NodePoolClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNodePoolClientWithBaseURI(baseURI string) NodePoolClient {
	return NodePoolClient{NewWithBaseURI(baseURI)}
}

// Create 노드풀 생성
// Parameters:
// UUID - 클러스터 UUID
// parameters - 노드풀 생성 파라미터
func (client NodePoolClient) Create(ctx context.Context, UUID string, parameters NodePoolRequestParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodePoolClient.Create")
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
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.NodeCount", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProductCode", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kubernetes.NodePoolClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, UUID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client NodePoolClient) CreatePreparer(ctx context.Context, UUID string, parameters NodePoolRequestParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"uuid": autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/nks/v2/clusters/{uuid}/node-pool", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client NodePoolClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client NodePoolClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete 노드푹 삭제
// Parameters:
// UUID - 클러스터 UUID
// instanceNo - 인스턴스 번호
func (client NodePoolClient) Delete(ctx context.Context, UUID string, instanceNo string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodePoolClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, UUID, instanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NodePoolClient) DeletePreparer(ctx context.Context, UUID string, instanceNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
		"uuid":       autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/nks/v2/clusters/{uuid}/node-pool/{instanceNo}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NodePoolClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NodePoolClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get 노드풀 조회
// Parameters:
// UUID - 클러스터 UUID
func (client NodePoolClient) Get(ctx context.Context, UUID string) (result NodePoolResponseParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodePoolClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, UUID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NodePoolClient) GetPreparer(ctx context.Context, UUID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"uuid": autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/nks/v2/clusters/{uuid}/node-pool", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NodePoolClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NodePoolClient) GetResponder(resp *http.Response) (result NodePoolResponseParameter, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update 노드풀 수정
// Parameters:
// UUID - 클러스터 UUID
// instanceNo - 인스턴스 번호
// parameters - 노드풀 생성 파라미터
func (client NodePoolClient) Update(ctx context.Context, UUID string, instanceNo string, parameters NodePoolUpdateParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NodePoolClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, UUID, instanceNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubernetes.NodePoolClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client NodePoolClient) UpdatePreparer(ctx context.Context, UUID string, instanceNo string, parameters NodePoolUpdateParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
		"uuid":       autorest.Encode("path", UUID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/nks/v2/clusters/{uuid}/node-pool/{instanceNo}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client NodePoolClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client NodePoolClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError),
		autorest.ByClosing())
	result.Response = resp
	return
}
