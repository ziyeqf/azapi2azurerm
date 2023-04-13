//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// BuildServiceBuilderClient contains the methods for the BuildServiceBuilder group.
// Don't use this type directly, use NewBuildServiceBuilderClient() instead.
type BuildServiceBuilderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBuildServiceBuilderClient creates a new instance of BuildServiceBuilderClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBuildServiceBuilderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BuildServiceBuilderClient, error) {
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
	client := &BuildServiceBuilderClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a KPack builder.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// builderResource - The target builder for the create or update operation
// options - BuildServiceBuilderClientBeginCreateOrUpdateOptions contains the optional parameters for the BuildServiceBuilderClient.BeginCreateOrUpdate
// method.
func (client *BuildServiceBuilderClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource BuilderResource, options *BuildServiceBuilderClientBeginCreateOrUpdateOptions) (*runtime.Poller[BuildServiceBuilderClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, buildServiceName, builderName, builderResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BuildServiceBuilderClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BuildServiceBuilderClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a KPack builder.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *BuildServiceBuilderClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource BuilderResource, options *BuildServiceBuilderClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, builderResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BuildServiceBuilderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource BuilderResource, options *BuildServiceBuilderClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, builderResource)
}

// BeginDelete - Delete a KPack builder.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// options - BuildServiceBuilderClientBeginDeleteOptions contains the optional parameters for the BuildServiceBuilderClient.BeginDelete
// method.
func (client *BuildServiceBuilderClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildServiceBuilderClientBeginDeleteOptions) (*runtime.Poller[BuildServiceBuilderClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, buildServiceName, builderName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[BuildServiceBuilderClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[BuildServiceBuilderClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a KPack builder.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *BuildServiceBuilderClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildServiceBuilderClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BuildServiceBuilderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildServiceBuilderClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
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

// Get - Get a KPack builder.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// options - BuildServiceBuilderClientGetOptions contains the optional parameters for the BuildServiceBuilderClient.Get method.
func (client *BuildServiceBuilderClient) Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildServiceBuilderClientGetOptions) (BuildServiceBuilderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, options)
	if err != nil {
		return BuildServiceBuilderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BuildServiceBuilderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BuildServiceBuilderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BuildServiceBuilderClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildServiceBuilderClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
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
func (client *BuildServiceBuilderClient) getHandleResponse(resp *http.Response) (BuildServiceBuilderClientGetResponse, error) {
	result := BuildServiceBuilderClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuilderResource); err != nil {
		return BuildServiceBuilderClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List KPack builders result.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// options - BuildServiceBuilderClientListOptions contains the optional parameters for the BuildServiceBuilderClient.List
// method.
func (client *BuildServiceBuilderClient) NewListPager(resourceGroupName string, serviceName string, buildServiceName string, options *BuildServiceBuilderClientListOptions) *runtime.Pager[BuildServiceBuilderClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BuildServiceBuilderClientListResponse]{
		More: func(page BuildServiceBuilderClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BuildServiceBuilderClientListResponse) (BuildServiceBuilderClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BuildServiceBuilderClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BuildServiceBuilderClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BuildServiceBuilderClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BuildServiceBuilderClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, options *BuildServiceBuilderClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
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

// listHandleResponse handles the List response.
func (client *BuildServiceBuilderClient) listHandleResponse(resp *http.Response) (BuildServiceBuilderClientListResponse, error) {
	result := BuildServiceBuilderClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuilderResourceCollection); err != nil {
		return BuildServiceBuilderClientListResponse{}, err
	}
	return result, nil
}
