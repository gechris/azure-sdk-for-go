package mysql

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// DatabasesClient is the the Microsoft Azure management API provides create, read, update, and delete functionality
// for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log files and configurations
// with new business model.
type DatabasesClient struct {
    BaseClient
}
// NewDatabasesClient creates an instance of the DatabasesClient client.
func NewDatabasesClient(subscriptionID string) DatabasesClient {
    return NewDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabasesClientWithBaseURI creates an instance of the DatabasesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
        return DatabasesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates a new database or updates an existing database.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // serverName - the name of the server.
        // databaseName - the name of the database.
        // parameters - the required parameters for creating or updating a database.
func (client DatabasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters Database) (result DatabasesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("mysql.DatabasesClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client DatabasesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters Database) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "databaseName": autorest.Encode("path",databaseName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-07-01-privatepreview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/databases/{databaseName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) CreateOrUpdateSender(req *http.Request) (future DatabasesCreateOrUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client DatabasesClient) CreateOrUpdateResponder(resp *http.Response) (result Database, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes a database.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // serverName - the name of the server.
        // databaseName - the name of the database.
func (client DatabasesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DatabasesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.Delete")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("mysql.DatabasesClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client DatabasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "databaseName": autorest.Encode("path",databaseName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-07-01-privatepreview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/databases/{databaseName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) DeleteSender(req *http.Request) (future DatabasesDeleteFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client DatabasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets information about a database.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // serverName - the name of the server.
        // databaseName - the name of the database.
func (client DatabasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result Database, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("mysql.DatabasesClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client DatabasesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "databaseName": autorest.Encode("path",databaseName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-07-01-privatepreview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/databases/{databaseName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client DatabasesClient) GetResponder(resp *http.Response) (result Database, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByServer list all the databases in a given server.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // serverName - the name of the server.
func (client DatabasesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result DatabaseListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.ListByServer")
        defer func() {
            sc := -1
        if result.dlr.Response.Response != nil {
        sc = result.dlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("mysql.DatabasesClient", "ListByServer", err.Error())
        }

            result.fn = client.listByServerNextResults
    req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "ListByServer", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByServerSender(req)
        if err != nil {
        result.dlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "ListByServer", resp, "Failure sending request")
        return
        }

        result.dlr, err = client.ListByServerResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "ListByServer", resp, "Failure responding to request")
        }
            if result.dlr.hasNextLink() && result.dlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListByServerPreparer prepares the ListByServer request.
    func (client DatabasesClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-07-01-privatepreview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/databases",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServerSender sends the ListByServer request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByServerResponder handles the response to the ListByServer request. The method always
    // closes the http.Response Body.
    func (client DatabasesClient) ListByServerResponder(resp *http.Response) (result DatabaseListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByServerNextResults retrieves the next set of results, if any.
            func (client DatabasesClient) listByServerNextResults(ctx context.Context, lastResults DatabaseListResult) (result DatabaseListResult, err error) {
            req, err := lastResults.databaseListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "mysql.DatabasesClient", "listByServerNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServerSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "mysql.DatabasesClient", "listByServerNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServerResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "mysql.DatabasesClient", "listByServerNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
            func (client DatabasesClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result DatabaseListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.ListByServer")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
                            return
            }

