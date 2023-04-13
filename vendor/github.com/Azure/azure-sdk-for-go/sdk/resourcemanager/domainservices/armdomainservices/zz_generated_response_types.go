//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	DomainService
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	DomainService
}

// ClientListByResourceGroupResponse contains the response from method Client.ListByResourceGroup.
type ClientListByResourceGroupResponse struct {
	DomainServiceListResult
}

// ClientListResponse contains the response from method Client.List.
type ClientListResponse struct {
	DomainServiceListResult
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	DomainService
}

// DomainServiceOperationsClientListResponse contains the response from method DomainServiceOperationsClient.List.
type DomainServiceOperationsClientListResponse struct {
	OperationEntityListResult
}

// OuContainerClientCreateResponse contains the response from method OuContainerClient.Create.
type OuContainerClientCreateResponse struct {
	OuContainer
}

// OuContainerClientDeleteResponse contains the response from method OuContainerClient.Delete.
type OuContainerClientDeleteResponse struct {
	// placeholder for future response values
}

// OuContainerClientGetResponse contains the response from method OuContainerClient.Get.
type OuContainerClientGetResponse struct {
	OuContainer
}

// OuContainerClientListResponse contains the response from method OuContainerClient.List.
type OuContainerClientListResponse struct {
	OuContainerListResult
}

// OuContainerClientUpdateResponse contains the response from method OuContainerClient.Update.
type OuContainerClientUpdateResponse struct {
	OuContainer
}

// OuContainerOperationsClientListResponse contains the response from method OuContainerOperationsClient.List.
type OuContainerOperationsClientListResponse struct {
	OperationEntityListResult
}
