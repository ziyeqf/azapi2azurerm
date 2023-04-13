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

// WorkspaceManagedSQLServerEncryptionProtectorClient contains the methods for the WorkspaceManagedSQLServerEncryptionProtector group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerEncryptionProtectorClient() instead.
type WorkspaceManagedSQLServerEncryptionProtectorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerEncryptionProtectorClient creates a new instance of WorkspaceManagedSQLServerEncryptionProtectorClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerEncryptionProtectorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerEncryptionProtectorClient, error) {
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
	client := &WorkspaceManagedSQLServerEncryptionProtectorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Updates workspace managed sql server's encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// encryptionProtectorName - The name of the encryption protector.
// parameters - The requested encryption protector resource state.
// options - WorkspaceManagedSQLServerEncryptionProtectorClientBeginCreateOrUpdateOptions contains the optional parameters
// for the WorkspaceManagedSQLServerEncryptionProtectorClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceManagedSQLServerEncryptionProtectorClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, encryptionProtectorName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkspaceManagedSQLServerEncryptionProtectorClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerEncryptionProtectorClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Updates workspace managed sql server's encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, parameters EncryptionProtector, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}"
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
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get workspace managed sql server's encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// encryptionProtectorName - The name of the encryption protector.
// options - WorkspaceManagedSQLServerEncryptionProtectorClientGetOptions contains the optional parameters for the WorkspaceManagedSQLServerEncryptionProtectorClient.Get
// method.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorClientGetOptions) (WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}"
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
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
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

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse, error) {
	result := WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtector); err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get list of encryption protectors for workspace managed sql server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceManagedSQLServerEncryptionProtectorClientListOptions contains the optional parameters for the WorkspaceManagedSQLServerEncryptionProtectorClient.List
// method.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerEncryptionProtectorClientListOptions) *runtime.Pager[WorkspaceManagedSQLServerEncryptionProtectorClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerEncryptionProtectorClientListResponse]{
		More: func(page WorkspaceManagedSQLServerEncryptionProtectorClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerEncryptionProtectorClientListResponse) (WorkspaceManagedSQLServerEncryptionProtectorClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerEncryptionProtectorClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector"
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
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerEncryptionProtectorClientListResponse, error) {
	result := WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EncryptionProtectorListResult); err != nil {
		return WorkspaceManagedSQLServerEncryptionProtectorClientListResponse{}, err
	}
	return result, nil
}

// BeginRevalidate - Revalidates workspace managed sql server's existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// encryptionProtectorName - The name of the encryption protector.
// options - WorkspaceManagedSQLServerEncryptionProtectorClientBeginRevalidateOptions contains the optional parameters for
// the WorkspaceManagedSQLServerEncryptionProtectorClient.BeginRevalidate method.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) BeginRevalidate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginRevalidateOptions) (*runtime.Poller[WorkspaceManagedSQLServerEncryptionProtectorClientRevalidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revalidate(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkspaceManagedSQLServerEncryptionProtectorClientRevalidateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerEncryptionProtectorClientRevalidateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Revalidate - Revalidates workspace managed sql server's existing encryption protector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) revalidate(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginRevalidateOptions) (*http.Response, error) {
	req, err := client.revalidateCreateRequest(ctx, resourceGroupName, workspaceName, encryptionProtectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revalidateCreateRequest creates the Revalidate request.
func (client *WorkspaceManagedSQLServerEncryptionProtectorClient) revalidateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, encryptionProtectorName EncryptionProtectorName, options *WorkspaceManagedSQLServerEncryptionProtectorClientBeginRevalidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}/revalidate"
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
	if encryptionProtectorName == "" {
		return nil, errors.New("parameter encryptionProtectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{encryptionProtectorName}", url.PathEscape(string(encryptionProtectorName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}