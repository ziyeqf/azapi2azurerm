//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// PreconfiguredEndpointsClient contains the methods for the PreconfiguredEndpoints group.
// Don't use this type directly, use NewPreconfiguredEndpointsClient() instead.
type PreconfiguredEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPreconfiguredEndpointsClient creates a new instance of PreconfiguredEndpointsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPreconfiguredEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PreconfiguredEndpointsClient, error) {
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
	client := &PreconfiguredEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Gets a list of Preconfigured Endpoints
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - The Profile identifier associated with the Tenant and Partner
// options - PreconfiguredEndpointsClientListOptions contains the optional parameters for the PreconfiguredEndpointsClient.List
// method.
func (client *PreconfiguredEndpointsClient) NewListPager(resourceGroupName string, profileName string, options *PreconfiguredEndpointsClientListOptions) *runtime.Pager[PreconfiguredEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PreconfiguredEndpointsClientListResponse]{
		More: func(page PreconfiguredEndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PreconfiguredEndpointsClientListResponse) (PreconfiguredEndpointsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, profileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PreconfiguredEndpointsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PreconfiguredEndpointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PreconfiguredEndpointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PreconfiguredEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *PreconfiguredEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/PreconfiguredEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PreconfiguredEndpointsClient) listHandleResponse(resp *http.Response) (PreconfiguredEndpointsClientListResponse, error) {
	result := PreconfiguredEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PreconfiguredEndpointList); err != nil {
		return PreconfiguredEndpointsClientListResponse{}, err
	}
	return result, nil
}