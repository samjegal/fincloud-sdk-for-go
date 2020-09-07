package nas

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// Client is the NAS Client
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

// AddAccessControl NAS 볼륨 접근제어를 추가
// Parameters:
// nasVolumeInstanceNo - NAS 볼륜 인스턴스 번호
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) AddAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result VolumeAccessControlResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.AddAccessControl")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddAccessControlPreparer(ctx, nasVolumeInstanceNo, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "AddAccessControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddAccessControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "AddAccessControl", resp, "Failure sending request")
		return
	}

	result, err = client.AddAccessControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "AddAccessControl", resp, "Failure responding to request")
	}

	return
}

// AddAccessControlPreparer prepares the AddAccessControl request.
func (client Client) AddAccessControlPreparer(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNo":    autorest.Encode("query", nasVolumeInstanceNo),
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addNasVolumeAccessControl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addNasVolumeAccessControl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddAccessControlSender sends the AddAccessControl request. The method will close the
// http.Response Body if it receives an error.
func (client Client) AddAccessControlSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddAccessControlResponder handles the response to the AddAccessControl request. The method always
// closes the http.Response Body.
func (client Client) AddAccessControlResponder(resp *http.Response) (result VolumeAccessControlResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ChangeSize NAS 볼륨 사이즈를 변경
// Parameters:
// nasVolumeInstanceNo - NAS 볼륜 인스턴스 번호
// volumeSize - 볼륨 사이즈
func (client Client) ChangeSize(ctx context.Context, nasVolumeInstanceNo string, volumeSize string) (result VolumeSizeResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ChangeSize")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ChangeSizePreparer(ctx, nasVolumeInstanceNo, volumeSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "ChangeSize", nil, "Failure preparing request")
		return
	}

	resp, err := client.ChangeSizeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "ChangeSize", resp, "Failure sending request")
		return
	}

	result, err = client.ChangeSizeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "ChangeSize", resp, "Failure responding to request")
	}

	return
}

// ChangeSizePreparer prepares the ChangeSize request.
func (client Client) ChangeSizePreparer(ctx context.Context, nasVolumeInstanceNo string, volumeSize string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNo": autorest.Encode("query", nasVolumeInstanceNo),
		"responseFormatType":  autorest.Encode("query", "json"),
		"volumeSize":          autorest.Encode("query", volumeSize),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/changeNasVolumeSize")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/changeNasVolumeSize"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ChangeSizeSender sends the ChangeSize request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ChangeSizeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ChangeSizeResponder handles the response to the ChangeSize request. The method always
// closes the http.Response Body.
func (client Client) ChangeSizeResponder(resp *http.Response) (result VolumeSizeResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create NAS 볼륨 인스턴스를 생성
// Parameters:
// volumeSize - 볼륨 사이즈
// zoneCode - ZONE 코드
// volumeName - 볼륨 이름
// volumeAllotmentProtocolTypeCode - 볼륨 할당 프로토콜 유형 코드
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
// cifsUserName - CIFS 유저 이름
// cifsUserPassword - CIFS 유저 비밀번호
// isEncryptedVolume - 볼륨 암호화 여부
// nasVolumeDescription - NAS 볼륨 설명
func (client Client) Create(ctx context.Context, volumeSize string, zoneCode string, volumeName string, volumeAllotmentProtocolTypeCode VolumeAllotmentProtocolTypeCode, serverInstanceNoListN string, cifsUserName string, cifsUserPassword string, isEncryptedVolume EncryptedVolume, nasVolumeDescription string) (result VolumeInstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, volumeSize, zoneCode, volumeName, volumeAllotmentProtocolTypeCode, serverInstanceNoListN, cifsUserName, cifsUserPassword, isEncryptedVolume, nasVolumeDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, volumeSize string, zoneCode string, volumeName string, volumeAllotmentProtocolTypeCode VolumeAllotmentProtocolTypeCode, serverInstanceNoListN string, cifsUserName string, cifsUserPassword string, isEncryptedVolume EncryptedVolume, nasVolumeDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"volumeSize":         autorest.Encode("query", volumeSize),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}
	if len(volumeName) > 0 {
		queryParameters["volumeName"] = autorest.Encode("query", volumeName)
	}
	if len(string(volumeAllotmentProtocolTypeCode)) > 0 {
		queryParameters["volumeAllotmentProtocolTypeCode"] = autorest.Encode("query", volumeAllotmentProtocolTypeCode)
	}
	if len(serverInstanceNoListN) > 0 {
		queryParameters["serverInstanceNoList.N"] = autorest.Encode("query", serverInstanceNoListN)
	}
	if len(cifsUserName) > 0 {
		queryParameters["cifsUserName"] = autorest.Encode("query", cifsUserName)
	}
	if len(cifsUserPassword) > 0 {
		queryParameters["cifsUserPassword"] = autorest.Encode("query", cifsUserPassword)
	}
	if len(string(isEncryptedVolume)) > 0 {
		queryParameters["isEncryptedVolume"] = autorest.Encode("query", isEncryptedVolume)
	}
	if len(nasVolumeDescription) > 0 {
		queryParameters["nasVolumeDescription"] = autorest.Encode("query", nasVolumeDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createNasVolumeInstance")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createNasVolumeInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result VolumeInstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete NAS 볼륨 인스턴스 삭제
// Parameters:
// nasVolumeInstanceNoListN - NAS 볼륨 인스턴스 번호 리스트
func (client Client) Delete(ctx context.Context, nasVolumeInstanceNoListN string) (result VolumeInstancesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, nasVolumeInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(ctx context.Context, nasVolumeInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNoList.N": autorest.Encode("query", nasVolumeInstanceNoListN),
		"responseFormatType":        autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteNasVolumeInstances")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteNasVolumeInstances"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client Client) DeleteResponder(resp *http.Response) (result VolumeInstancesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail NAS 볼륜 인스턴스 상세 정보를 조회
// Parameters:
// nasVolumeInstanceNo - NAS 볼륜 인스턴스 번호
func (client Client) GetDetail(ctx context.Context, nasVolumeInstanceNo string) (result VolumeInstanceDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, nasVolumeInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client Client) GetDetailPreparer(ctx context.Context, nasVolumeInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNo": autorest.Encode("query", nasVolumeInstanceNo),
		"responseFormatType":  autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNasVolumeInstanceDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNasVolumeInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client Client) GetDetailResponder(resp *http.Response) (result VolumeInstanceDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList NAS 볼륨 인스턴스 리스트를 조회
// Parameters:
// volumeAllotmentProtocolTypeCode - 볼륨 할당 프로토콜 유형 코드
// isEventConfiguration - 이벤트 설정 여부
// isSnapshotConfiguration - 스냅샷 설정 여부
// nasVolumeInstanceNoListN - NAS 볼륨 인스턴스 번호 리스트
// zoneCode - ZONE 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// volumeName - 볼륨 이름
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
func (client Client) GetList(ctx context.Context, volumeAllotmentProtocolTypeCode VolumeAllotmentProtocolTypeCode, isEventConfiguration EventConfiguration, isSnapshotConfiguration SnapshotConfiguration, nasVolumeInstanceNoListN string, zoneCode string, pageNo string, pageSize string, volumeName string, sortedBy SortedBy, sortingOrder SortingOrder) (result VolumeInstanceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, volumeAllotmentProtocolTypeCode, isEventConfiguration, isSnapshotConfiguration, nasVolumeInstanceNoListN, zoneCode, pageNo, pageSize, volumeName, sortedBy, sortingOrder)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client Client) GetListPreparer(ctx context.Context, volumeAllotmentProtocolTypeCode VolumeAllotmentProtocolTypeCode, isEventConfiguration EventConfiguration, isSnapshotConfiguration SnapshotConfiguration, nasVolumeInstanceNoListN string, zoneCode string, pageNo string, pageSize string, volumeName string, sortedBy SortedBy, sortingOrder SortingOrder) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(string(volumeAllotmentProtocolTypeCode)) > 0 {
		queryParameters["volumeAllotmentProtocolTypeCode"] = autorest.Encode("query", volumeAllotmentProtocolTypeCode)
	}
	if len(string(isEventConfiguration)) > 0 {
		queryParameters["isEventConfiguration"] = autorest.Encode("query", isEventConfiguration)
	}
	if len(string(isSnapshotConfiguration)) > 0 {
		queryParameters["isSnapshotConfiguration"] = autorest.Encode("query", isSnapshotConfiguration)
	}
	if len(nasVolumeInstanceNoListN) > 0 {
		queryParameters["nasVolumeInstanceNoList.N"] = autorest.Encode("query", nasVolumeInstanceNoListN)
	}
	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(volumeName) > 0 {
		queryParameters["volumeName"] = autorest.Encode("query", volumeName)
	}
	if len(string(sortedBy)) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(string(sortingOrder)) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNasVolumeInstanceList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNasVolumeInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client Client) GetListResponder(resp *http.Response) (result VolumeInstanceListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveAccessControl NAS 볼륨 접근제어를 제거
// Parameters:
// nasVolumeInstanceNo - NAS 볼륜 인스턴스 번호
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) RemoveAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result VolumeAccessControlResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.RemoveAccessControl")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveAccessControlPreparer(ctx, nasVolumeInstanceNo, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "RemoveAccessControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveAccessControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "RemoveAccessControl", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveAccessControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "RemoveAccessControl", resp, "Failure responding to request")
	}

	return
}

// RemoveAccessControlPreparer prepares the RemoveAccessControl request.
func (client Client) RemoveAccessControlPreparer(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNo":    autorest.Encode("query", nasVolumeInstanceNo),
		"responseFormatType":     autorest.Encode("query", "json"),
		"serverInstanceNoList.N": autorest.Encode("query", serverInstanceNoListN),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeNasVolumeAccessControl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeNasVolumeAccessControl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveAccessControlSender sends the RemoveAccessControl request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RemoveAccessControlSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveAccessControlResponder handles the response to the RemoveAccessControl request. The method always
// closes the http.Response Body.
func (client Client) RemoveAccessControlResponder(resp *http.Response) (result VolumeAccessControlResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetAccessControl NAS 볼륨 접근제어를 설정
// Parameters:
// nasVolumeInstanceNo - NAS 볼륜 인스턴스 번호
// serverInstanceNoListN - 서버 인스턴스 번호 리스트
func (client Client) SetAccessControl(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (result VolumeAccessControlResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.SetAccessControl")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SetAccessControlPreparer(ctx, nasVolumeInstanceNo, serverInstanceNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "SetAccessControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetAccessControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "nas.Client", "SetAccessControl", resp, "Failure sending request")
		return
	}

	result, err = client.SetAccessControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "nas.Client", "SetAccessControl", resp, "Failure responding to request")
	}

	return
}

// SetAccessControlPreparer prepares the SetAccessControl request.
func (client Client) SetAccessControlPreparer(ctx context.Context, nasVolumeInstanceNo string, serverInstanceNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"nasVolumeInstanceNo": autorest.Encode("query", nasVolumeInstanceNo),
		"responseFormatType":  autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(serverInstanceNoListN) > 0 {
		queryParameters["serverInstanceNoList.N"] = autorest.Encode("query", serverInstanceNoListN)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/setNasVolumeAccessControl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/setNasVolumeAccessControl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetAccessControlSender sends the SetAccessControl request. The method will close the
// http.Response Body if it receives an error.
func (client Client) SetAccessControlSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SetAccessControlResponder handles the response to the SetAccessControl request. The method always
// closes the http.Response Body.
func (client Client) SetAccessControlResponder(resp *http.Response) (result VolumeAccessControlResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
