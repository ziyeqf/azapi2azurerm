//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SQLPoolOperationsClient contains the methods for the SQLPoolOperations group.
// Don't use this type directly, use NewSQLPoolOperationsClient() instead.
type SQLPoolOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolOperationsClient creates a new instance of SQLPoolOperationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLPoolOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolOperationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SQLPoolOperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Gets a list of operations performed on the SQL pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// options - SQLPoolOperationsClientListOptions contains the optional parameters for the SQLPoolOperationsClient.List method.
func (client *SQLPoolOperationsClient) NewListPager(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolOperationsClientListOptions) *runtime.Pager[SQLPoolOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLPoolOperationsClientListResponse]{
		More: func(page SQLPoolOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLPoolOperationsClientListResponse) (SQLPoolOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLPoolOperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLPoolOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLPoolOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLPoolOperationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolOperationsClient) listHandleResponse(resp *http.Response) (SQLPoolOperationsClientListResponse, error) {
	result := SQLPoolOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolBlobAuditingPolicySQLPoolOperationListResult); err != nil {
		return SQLPoolOperationsClientListResponse{}, err
	}
	return result, nil
}
