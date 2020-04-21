package insight

// Copyright (c) Park MinKeun and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

// DataClient is the cloud Insight Client
type DataClient struct {
	BaseClient
}

// NewDataClient creates an instance of the DataClient client.
func NewDataClient() DataClient {
	return NewDataClientWithBaseURI(DefaultBaseURI)
}

// NewDataClientWithBaseURI creates an instance of the DataClient client.
func NewDataClientWithBaseURI(baseURI string) DataClient {
	return DataClient{NewWithBaseURI(baseURI)}
}

// Preview metric information for the data to be retrieved
func (client DataClient) Preview(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.Preview")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PreviewPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Preview", nil, "Failure preparing request")
		return
	}

	resp, err := client.PreviewSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Preview", resp, "Failure sending request")
		return
	}

	result, err = client.PreviewResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Preview", resp, "Failure responding to request")
	}

	return
}

// PreviewPreparer prepares the Preview request.
func (client DataClient) PreviewPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/chart/preview"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PreviewSender sends the Preview request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) PreviewSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PreviewResponder handles the response to the Preview request. The method always
// closes the http.Response Body.
func (client DataClient) PreviewResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Query get widget data preview for dashboard widget
func (client DataClient) Query(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.Query")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Query", nil, "Failure preparing request")
		return
	}

	resp, err := client.QuerySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Query", resp, "Failure sending request")
		return
	}

	result, err = client.QueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "Query", resp, "Failure responding to request")
	}

	return
}

// QueryPreparer prepares the Query request.
func (client DataClient) QueryPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/query"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QuerySender sends the Query request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) QuerySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryResponder handles the response to the Query request. The method always
// closes the http.Response Body.
func (client DataClient) QueryResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// QueryMultiple query multiple metric data for a specific product with specified criteria
func (client DataClient) QueryMultiple(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataClient.QueryMultiple")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.QueryMultiplePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "QueryMultiple", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryMultipleSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insight.DataClient", "QueryMultiple", resp, "Failure sending request")
		return
	}

	result, err = client.QueryMultipleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insight.DataClient", "QueryMultiple", resp, "Failure responding to request")
	}

	return
}

// QueryMultiplePreparer prepares the QueryMultiple request.
func (client DataClient) QueryMultiplePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/cw_fea/real/cw/api/data/query/multiple"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryMultipleSender sends the QueryMultiple request. The method will close the
// http.Response Body if it receives an error.
func (client DataClient) QueryMultipleSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// QueryMultipleResponder handles the response to the QueryMultiple request. The method always
// closes the http.Response Body.
func (client DataClient) QueryMultipleResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
