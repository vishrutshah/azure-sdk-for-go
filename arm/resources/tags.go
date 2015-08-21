package resources

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
	"github.com/azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// Tags Client
type TagsClient struct {
	ResourceManagementClient
}

func NewTagsClient(subscriptionId string) TagsClient {
	return NewTagsClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewTagsClientWithBaseUri(baseUri string, subscriptionId string) TagsClient {
	return TagsClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// DeleteValue delete a subscription resource tag value.
//
// tagName is the name of the tag. tagValue is the value of the tag.
func (client TagsClient) DeleteValue(tagName string, tagValue string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteValueRequest(tagName, tagValue)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.TagsClient", "DeleteValue", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the DeleteValue request.
func (client TagsClient) NewDeleteValueRequest(tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
		"tagName":        url.QueryEscape(tagName),
		"tagValue":       url.QueryEscape(tagValue),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteValueRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the DeleteValue request.
func (client TagsClient) DeleteValueRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"))
}

// CreateOrUpdateValue create a subscription resource tag value.
//
// tagName is the name of the tag. tagValue is the value of the tag.
func (client TagsClient) CreateOrUpdateValue(tagName string, tagValue string) (result TagValue, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateValueRequest(tagName, tagValue)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdateValue", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdateValue request.
func (client TagsClient) NewCreateOrUpdateValueRequest(tagName string, tagValue string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
		"tagName":        url.QueryEscape(tagName),
		"tagValue":       url.QueryEscape(tagValue),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateValueRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdateValue request.
func (client TagsClient) CreateOrUpdateValueRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"))
}

// CreateOrUpdate create a subscription resource tag.
//
// tagName is the name of the tag.
func (client TagsClient) CreateOrUpdate(tagName string) (result TagDetails, ae autorest.Error) {
	req, err := client.NewCreateOrUpdateRequest(tagName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.TagsClient", "CreateOrUpdate", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the CreateOrUpdate request.
func (client TagsClient) NewCreateOrUpdateRequest(tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
		"tagName":        url.QueryEscape(tagName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.CreateOrUpdateRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the CreateOrUpdate request.
func (client TagsClient) CreateOrUpdateRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}"))
}

// Delete delete a subscription resource tag.
//
// tagName is the name of the tag.
func (client TagsClient) Delete(tagName string) (result autorest.Response, ae autorest.Error) {
	req, err := client.NewDeleteRequest(tagName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK())
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.TagsClient", "Delete", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = resp

	return
}

// Create the Delete request.
func (client TagsClient) NewDeleteRequest(tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
		"tagName":        url.QueryEscape(tagName),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.DeleteRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the Delete request.
func (client TagsClient) DeleteRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames/{tagName}"))
}

// List get a list of subscription resource tags.
func (client TagsClient) List() (result TagsListResult, ae autorest.Error) {
	req, err := client.NewListRequest()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.TagsClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "resources.TagsClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "resources.TagsClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client TagsClient) NewListRequest() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client TagsClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/tagNames"))
}