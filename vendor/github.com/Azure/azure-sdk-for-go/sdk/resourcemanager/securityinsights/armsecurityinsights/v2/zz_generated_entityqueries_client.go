//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// EntityQueriesClient contains the methods for the EntityQueries group.
// Don't use this type directly, use NewEntityQueriesClient() instead.
type EntityQueriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEntityQueriesClient creates a new instance of EntityQueriesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEntityQueriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EntityQueriesClient, error) {
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
	client := &EntityQueriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the entity query.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// entityQueryID - entity query ID
// entityQuery - The entity query we want to create or update
// options - EntityQueriesClientCreateOrUpdateOptions contains the optional parameters for the EntityQueriesClient.CreateOrUpdate
// method.
func (client *EntityQueriesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, entityQuery CustomEntityQueryClassification, options *EntityQueriesClientCreateOrUpdateOptions) (EntityQueriesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, entityQueryID, entityQuery, options)
	if err != nil {
		return EntityQueriesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EntityQueriesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return EntityQueriesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EntityQueriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, entityQuery CustomEntityQueryClassification, options *EntityQueriesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueries/{entityQueryId}"
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
	if entityQueryID == "" {
		return nil, errors.New("parameter entityQueryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityQueryId}", url.PathEscape(entityQueryID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, entityQuery)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EntityQueriesClient) createOrUpdateHandleResponse(resp *http.Response) (EntityQueriesClientCreateOrUpdateResponse, error) {
	result := EntityQueriesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EntityQueriesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the entity query.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// entityQueryID - entity query ID
// options - EntityQueriesClientDeleteOptions contains the optional parameters for the EntityQueriesClient.Delete method.
func (client *EntityQueriesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, options *EntityQueriesClientDeleteOptions) (EntityQueriesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, entityQueryID, options)
	if err != nil {
		return EntityQueriesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EntityQueriesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return EntityQueriesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return EntityQueriesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EntityQueriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, options *EntityQueriesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueries/{entityQueryId}"
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
	if entityQueryID == "" {
		return nil, errors.New("parameter entityQueryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityQueryId}", url.PathEscape(entityQueryID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an entity query.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// entityQueryID - entity query ID
// options - EntityQueriesClientGetOptions contains the optional parameters for the EntityQueriesClient.Get method.
func (client *EntityQueriesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, options *EntityQueriesClientGetOptions) (EntityQueriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, entityQueryID, options)
	if err != nil {
		return EntityQueriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EntityQueriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EntityQueriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EntityQueriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityQueryID string, options *EntityQueriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueries/{entityQueryId}"
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
	if entityQueryID == "" {
		return nil, errors.New("parameter entityQueryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityQueryId}", url.PathEscape(entityQueryID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EntityQueriesClient) getHandleResponse(resp *http.Response) (EntityQueriesClientGetResponse, error) {
	result := EntityQueriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return EntityQueriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all entity queries.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - EntityQueriesClientListOptions contains the optional parameters for the EntityQueriesClient.List method.
func (client *EntityQueriesClient) NewListPager(resourceGroupName string, workspaceName string, options *EntityQueriesClientListOptions) *runtime.Pager[EntityQueriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntityQueriesClientListResponse]{
		More: func(page EntityQueriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EntityQueriesClientListResponse) (EntityQueriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EntityQueriesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EntityQueriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EntityQueriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EntityQueriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *EntityQueriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entityQueries"
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
	if options != nil && options.Kind != nil {
		reqQP.Set("kind", string(*options.Kind))
	}
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntityQueriesClient) listHandleResponse(resp *http.Response) (EntityQueriesClientListResponse, error) {
	result := EntityQueriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityQueryList); err != nil {
		return EntityQueriesClientListResponse{}, err
	}
	return result, nil
}