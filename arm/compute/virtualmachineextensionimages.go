package compute

// Copyright (c) Microsoft and contributors.  All rights reserved.
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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// VirtualMachineExtensionImagesClient is the client for the
// VirtualMachineExtensionImages methods of the Compute service.
type VirtualMachineExtensionImagesClient struct {
	ManagementClient
}

// NewVirtualMachineExtensionImagesClient creates an instance of the
// VirtualMachineExtensionImagesClient client.
func NewVirtualMachineExtensionImagesClient(subscriptionID string) VirtualMachineExtensionImagesClient {
	return NewVirtualMachineExtensionImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineExtensionImagesClientWithBaseURI creates an instance of
// the VirtualMachineExtensionImagesClient client.
func NewVirtualMachineExtensionImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineExtensionImagesClient {
	return VirtualMachineExtensionImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a virtual machine extension image.
//
func (client VirtualMachineExtensionImagesClient) Get(location string, publisherName string, typeParameter string, version string) (result VirtualMachineExtensionImage, ae error) {
	req, err := client.GetPreparer(location, publisherName, typeParameter, version)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineExtensionImagesClient) GetPreparer(location string, publisherName string, typeParameter string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"publisherName":  url.QueryEscape(publisherName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"type":           url.QueryEscape(typeParameter),
		"version":        url.QueryEscape(version),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) GetResponder(resp *http.Response) (result VirtualMachineExtensionImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTypes gets a list of virtual machine extension image types.
//
func (client VirtualMachineExtensionImagesClient) ListTypes(location string, publisherName string) (result VirtualMachineImageResourceList, ae error) {
	req, err := client.ListTypesPreparer(location, publisherName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListTypes", "Failure preparing request")
	}

	resp, err := client.ListTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListTypes", "Failure sending request")
	}

	result, err = client.ListTypesResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListTypes", "Failure responding to request")
	}

	return
}

// ListTypesPreparer prepares the ListTypes request.
func (client VirtualMachineExtensionImagesClient) ListTypesPreparer(location string, publisherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"publisherName":  url.QueryEscape(publisherName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListTypesSender sends the ListTypes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) ListTypesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListTypesResponder handles the response to the ListTypes request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) ListTypesResponder(resp *http.Response) (result VirtualMachineImageResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListVersions gets a list of virtual machine extension image versions.
//
// filter is the filter to apply on the operation.
func (client VirtualMachineExtensionImagesClient) ListVersions(location string, publisherName string, typeParameter string, filter string, top int, orderby string) (result VirtualMachineImageResourceList, ae error) {
	req, err := client.ListVersionsPreparer(location, publisherName, typeParameter, filter, top, orderby)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListVersions", "Failure preparing request")
	}

	resp, err := client.ListVersionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListVersions", "Failure sending request")
	}

	result, err = client.ListVersionsResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "compute/VirtualMachineExtensionImagesClient", "ListVersions", "Failure responding to request")
	}

	return
}

// ListVersionsPreparer prepares the ListVersions request.
func (client VirtualMachineExtensionImagesClient) ListVersionsPreparer(location string, publisherName string, typeParameter string, filter string, top int, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       url.QueryEscape(location),
		"publisherName":  url.QueryEscape(publisherName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
		"type":           url.QueryEscape(typeParameter),
	}

	queryParameters := map[string]interface{}{
		"$filter":     filter,
		"$orderby":    orderby,
		"$top":        top,
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListVersionsSender sends the ListVersions request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineExtensionImagesClient) ListVersionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListVersionsResponder handles the response to the ListVersions request. The method always
// closes the http.Response Body.
func (client VirtualMachineExtensionImagesClient) ListVersionsResponder(resp *http.Response) (result VirtualMachineImageResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
