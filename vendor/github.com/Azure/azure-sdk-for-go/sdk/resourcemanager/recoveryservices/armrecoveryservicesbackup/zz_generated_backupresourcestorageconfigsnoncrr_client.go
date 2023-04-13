//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// BackupResourceStorageConfigsNonCRRClient contains the methods for the BackupResourceStorageConfigsNonCRR group.
// Don't use this type directly, use NewBackupResourceStorageConfigsNonCRRClient() instead.
type BackupResourceStorageConfigsNonCRRClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupResourceStorageConfigsNonCRRClient creates a new instance of BackupResourceStorageConfigsNonCRRClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupResourceStorageConfigsNonCRRClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupResourceStorageConfigsNonCRRClient, error) {
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
	client := &BackupResourceStorageConfigsNonCRRClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Fetches resource storage config.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// options - BackupResourceStorageConfigsNonCRRClientGetOptions contains the optional parameters for the BackupResourceStorageConfigsNonCRRClient.Get
// method.
func (client *BackupResourceStorageConfigsNonCRRClient) Get(ctx context.Context, vaultName string, resourceGroupName string, options *BackupResourceStorageConfigsNonCRRClientGetOptions) (BackupResourceStorageConfigsNonCRRClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, options)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupResourceStorageConfigsNonCRRClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupResourceStorageConfigsNonCRRClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupResourceStorageConfigsNonCRRClientGetOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupResourceStorageConfigsNonCRRClient) getHandleResponse(resp *http.Response) (BackupResourceStorageConfigsNonCRRClientGetResponse, error) {
	result := BackupResourceStorageConfigsNonCRRClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResourceConfigResource); err != nil {
		return BackupResourceStorageConfigsNonCRRClientGetResponse{}, err
	}
	return result, nil
}

// Patch - Updates vault storage model type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - Vault storage config request
// options - BackupResourceStorageConfigsNonCRRClientPatchOptions contains the optional parameters for the BackupResourceStorageConfigsNonCRRClient.Patch
// method.
func (client *BackupResourceStorageConfigsNonCRRClient) Patch(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceConfigResource, options *BackupResourceStorageConfigsNonCRRClientPatchOptions) (BackupResourceStorageConfigsNonCRRClientPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return BackupResourceStorageConfigsNonCRRClientPatchResponse{}, runtime.NewResponseError(resp)
	}
	return BackupResourceStorageConfigsNonCRRClientPatchResponse{}, nil
}

// patchCreateRequest creates the Patch request.
func (client *BackupResourceStorageConfigsNonCRRClient) patchCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceConfigResource, options *BackupResourceStorageConfigsNonCRRClientPatchOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Update - Updates vault storage model type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// parameters - Vault storage config request
// options - BackupResourceStorageConfigsNonCRRClientUpdateOptions contains the optional parameters for the BackupResourceStorageConfigsNonCRRClient.Update
// method.
func (client *BackupResourceStorageConfigsNonCRRClient) Update(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceConfigResource, options *BackupResourceStorageConfigsNonCRRClientUpdateOptions) (BackupResourceStorageConfigsNonCRRClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupResourceStorageConfigsNonCRRClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupResourceStorageConfigsNonCRRClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *BackupResourceStorageConfigsNonCRRClient) updateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceConfigResource, options *BackupResourceStorageConfigsNonCRRClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *BackupResourceStorageConfigsNonCRRClient) updateHandleResponse(resp *http.Response) (BackupResourceStorageConfigsNonCRRClientUpdateResponse, error) {
	result := BackupResourceStorageConfigsNonCRRClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResourceConfigResource); err != nil {
		return BackupResourceStorageConfigsNonCRRClientUpdateResponse{}, err
	}
	return result, nil
}
