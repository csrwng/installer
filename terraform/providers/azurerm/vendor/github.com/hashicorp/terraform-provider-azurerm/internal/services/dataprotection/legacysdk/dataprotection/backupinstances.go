package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// BackupInstancesClient is the open API 2.0 Specs for Azure Data Protection service
type BackupInstancesClient struct {
	BaseClient
}

// NewBackupInstancesClient creates an instance of the BackupInstancesClient client.
func NewBackupInstancesClient(subscriptionID string) BackupInstancesClient {
	return NewBackupInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupInstancesClientWithBaseURI creates an instance of the BackupInstancesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBackupInstancesClientWithBaseURI(baseURI string, subscriptionID string) BackupInstancesClient {
	return BackupInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// AdhocBackup trigger adhoc backup
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
// parameters - request body for operation
func (client BackupInstancesClient) AdhocBackup(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters TriggerBackupRequest) (result BackupInstancesAdhocBackupFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.AdhocBackup")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{
			TargetValue: parameters,
			Constraints: []validation.Constraint{{
				Target: "parameters.BackupRuleOptions", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{
					{Target: "parameters.BackupRuleOptions.RuleName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BackupRuleOptions.TriggerOption", Name: validation.Null, Rule: true, Chain: nil},
				},
			}},
		},
	}); err != nil {
		return result, validation.NewError("dataprotection.BackupInstancesClient", "AdhocBackup", err.Error())
	}

	req, err := client.AdhocBackupPreparer(ctx, vaultName, resourceGroupName, backupInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "AdhocBackup", nil, "Failure preparing request")
		return
	}

	result, err = client.AdhocBackupSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "AdhocBackup", nil, "Failure sending request")
		return
	}

	return
}

// AdhocBackupPreparer prepares the AdhocBackup request.
func (client BackupInstancesClient) AdhocBackupPreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters TriggerBackupRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/backup", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AdhocBackupSender sends the AdhocBackup request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) AdhocBackupSender(req *http.Request) (future BackupInstancesAdhocBackupFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// AdhocBackupResponder handles the response to the AdhocBackup request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) AdhocBackupResponder(resp *http.Response) (result OperationJobExtendedInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate create or update a backup instance in a backup vault
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
// parameters - request body for operation
func (client BackupInstancesClient) CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters BackupInstanceResource) (result BackupInstancesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{
			TargetValue: parameters,
			Constraints: []validation.Constraint{{
				Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{
					{
						Target: "parameters.Properties.DataSourceInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.Properties.DataSourceInfo.ResourceID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.Properties.DataSourceSetInfo", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.DataSourceSetInfo.ResourceID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.Properties.PolicyInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.Properties.PolicyInfo.PolicyID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.Properties.ProtectionStatus", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{
							{
								Target: "parameters.Properties.ProtectionStatus.ErrorDetails", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{
									{
										Target: "parameters.Properties.ProtectionStatus.ErrorDetails.InnerError", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.Properties.ProtectionStatus.ErrorDetails.InnerError.EmbeddedInnerError", Name: validation.Null, Rule: false, Chain: nil}},
									},
								},
							},
						},
					},
					{
						Target: "parameters.Properties.ProtectionErrorDetails", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{
							{
								Target: "parameters.Properties.ProtectionErrorDetails.InnerError", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Properties.ProtectionErrorDetails.InnerError.EmbeddedInnerError", Name: validation.Null, Rule: false, Chain: nil}},
							},
						},
					},
					{Target: "parameters.Properties.ObjectType", Name: validation.Null, Rule: true, Chain: nil},
				},
			}},
		},
	}); err != nil {
		return result, validation.NewError("dataprotection.BackupInstancesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, vaultName, resourceGroupName, backupInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupInstancesClient) CreateOrUpdatePreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters BackupInstanceResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) CreateOrUpdateSender(req *http.Request) (future BackupInstancesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) CreateOrUpdateResponder(resp *http.Response) (result BackupInstanceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a backup instance in a backup vault
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
func (client BackupInstancesClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result BackupInstancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vaultName, resourceGroupName, backupInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupInstancesClient) DeletePreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) DeleteSender(req *http.Request) (future BackupInstancesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a backup instance with name in a backup vault
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
func (client BackupInstancesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result BackupInstanceResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, backupInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupInstancesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) GetResponder(resp *http.Response) (result BackupInstanceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a backup instances belonging to a backup vault
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
func (client BackupInstancesClient) List(ctx context.Context, vaultName string, resourceGroupName string) (result BackupInstanceResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.List")
		defer func() {
			sc := -1
			if result.birl.Response.Response != nil {
				sc = result.birl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.birl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.birl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.birl.hasNextLink() && result.birl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BackupInstancesClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) ListResponder(resp *http.Response) (result BackupInstanceResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BackupInstancesClient) listNextResults(ctx context.Context, lastResults BackupInstanceResourceList) (result BackupInstanceResourceList, err error) {
	req, err := lastResults.backupInstanceResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BackupInstancesClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string) (result BackupInstanceResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName)
	return
}

// TriggerRehydrate rehydrate recovery point for restore for a BackupInstance
// Parameters:
// resourceGroupName - the name of the resource group where the backup vault is present.
// vaultName - the name of the backup vault.
// parameters - request body for operation
func (client BackupInstancesClient) TriggerRehydrate(ctx context.Context, resourceGroupName string, vaultName string, parameters AzureBackupRehydrationRequest, backupInstanceName string) (result BackupInstancesTriggerRehydrateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.TriggerRehydrate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{
			TargetValue: parameters,
			Constraints: []validation.Constraint{
				{Target: "parameters.RecoveryPointID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.RehydrationRetentionDuration", Name: validation.Null, Rule: true, Chain: nil},
			},
		},
	}); err != nil {
		return result, validation.NewError("dataprotection.BackupInstancesClient", "TriggerRehydrate", err.Error())
	}

	req, err := client.TriggerRehydratePreparer(ctx, resourceGroupName, vaultName, parameters, backupInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "TriggerRehydrate", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerRehydrateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "TriggerRehydrate", nil, "Failure sending request")
		return
	}

	return
}

// TriggerRehydratePreparer prepares the TriggerRehydrate request.
func (client BackupInstancesClient) TriggerRehydratePreparer(ctx context.Context, resourceGroupName string, vaultName string, parameters AzureBackupRehydrationRequest, backupInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/rehydrate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerRehydrateSender sends the TriggerRehydrate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) TriggerRehydrateSender(req *http.Request) (future BackupInstancesTriggerRehydrateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// TriggerRehydrateResponder handles the response to the TriggerRehydrate request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) TriggerRehydrateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// TriggerRestore triggers restore for a BackupInstance
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
// parameters - request body for operation
func (client BackupInstancesClient) TriggerRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters BasicAzureBackupRestoreRequest) (result BackupInstancesTriggerRestoreFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.TriggerRestore")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerRestorePreparer(ctx, vaultName, resourceGroupName, backupInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "TriggerRestore", nil, "Failure preparing request")
		return
	}

	result, err = client.TriggerRestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "TriggerRestore", nil, "Failure sending request")
		return
	}

	return
}

// TriggerRestorePreparer prepares the TriggerRestore request.
func (client BackupInstancesClient) TriggerRestorePreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters BasicAzureBackupRestoreRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/restore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerRestoreSender sends the TriggerRestore request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) TriggerRestoreSender(req *http.Request) (future BackupInstancesTriggerRestoreFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// TriggerRestoreResponder handles the response to the TriggerRestore request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) TriggerRestoreResponder(resp *http.Response) (result OperationJobExtendedInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateForBackup validate whether adhoc backup will be successful or not
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// parameters - request body for operation
func (client BackupInstancesClient) ValidateForBackup(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateForBackupRequest) (result BackupInstancesValidateForBackupFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.ValidateForBackup")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{
			TargetValue: parameters,
			Constraints: []validation.Constraint{{
				Target: "parameters.BackupInstance", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{
					{
						Target: "parameters.BackupInstance.DataSourceInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.BackupInstance.DataSourceInfo.ResourceID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.BackupInstance.DataSourceSetInfo", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.BackupInstance.DataSourceSetInfo.ResourceID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.BackupInstance.PolicyInfo", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.BackupInstance.PolicyInfo.PolicyID", Name: validation.Null, Rule: true, Chain: nil}},
					},
					{
						Target: "parameters.BackupInstance.ProtectionStatus", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{
							{
								Target: "parameters.BackupInstance.ProtectionStatus.ErrorDetails", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{
									{
										Target: "parameters.BackupInstance.ProtectionStatus.ErrorDetails.InnerError", Name: validation.Null, Rule: false,
										Chain: []validation.Constraint{{Target: "parameters.BackupInstance.ProtectionStatus.ErrorDetails.InnerError.EmbeddedInnerError", Name: validation.Null, Rule: false, Chain: nil}},
									},
								},
							},
						},
					},
					{
						Target: "parameters.BackupInstance.ProtectionErrorDetails", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{
							{
								Target: "parameters.BackupInstance.ProtectionErrorDetails.InnerError", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.BackupInstance.ProtectionErrorDetails.InnerError.EmbeddedInnerError", Name: validation.Null, Rule: false, Chain: nil}},
							},
						},
					},
					{Target: "parameters.BackupInstance.ObjectType", Name: validation.Null, Rule: true, Chain: nil},
				},
			}},
		},
	}); err != nil {
		return result, validation.NewError("dataprotection.BackupInstancesClient", "ValidateForBackup", err.Error())
	}

	req, err := client.ValidateForBackupPreparer(ctx, vaultName, resourceGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "ValidateForBackup", nil, "Failure preparing request")
		return
	}

	result, err = client.ValidateForBackupSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "ValidateForBackup", nil, "Failure sending request")
		return
	}

	return
}

// ValidateForBackupPreparer prepares the ValidateForBackup request.
func (client BackupInstancesClient) ValidateForBackupPreparer(ctx context.Context, vaultName string, resourceGroupName string, parameters ValidateForBackupRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/validateForBackup", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateForBackupSender sends the ValidateForBackup request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) ValidateForBackupSender(req *http.Request) (future BackupInstancesValidateForBackupFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ValidateForBackupResponder handles the response to the ValidateForBackup request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) ValidateForBackupResponder(resp *http.Response) (result OperationJobExtendedInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateForRestore validates if Restore can be triggered for a DataSource
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
// parameters - request body for operation
func (client BackupInstancesClient) ValidateForRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters ValidateRestoreRequestObject) (result BackupInstancesValidateForRestoreFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupInstancesClient.ValidateForRestore")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidateForRestorePreparer(ctx, vaultName, resourceGroupName, backupInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "ValidateForRestore", nil, "Failure preparing request")
		return
	}

	result, err = client.ValidateForRestoreSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.BackupInstancesClient", "ValidateForRestore", nil, "Failure sending request")
		return
	}

	return
}

// ValidateForRestorePreparer prepares the ValidateForRestore request.
func (client BackupInstancesClient) ValidateForRestorePreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters ValidateRestoreRequestObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/validateRestore", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateForRestoreSender sends the ValidateForRestore request. The method will close the
// http.Response Body if it receives an error.
func (client BackupInstancesClient) ValidateForRestoreSender(req *http.Request) (future BackupInstancesValidateForRestoreFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ValidateForRestoreResponder handles the response to the ValidateForRestore request. The method always
// closes the http.Response Body.
func (client BackupInstancesClient) ValidateForRestoreResponder(resp *http.Response) (result OperationJobExtendedInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
