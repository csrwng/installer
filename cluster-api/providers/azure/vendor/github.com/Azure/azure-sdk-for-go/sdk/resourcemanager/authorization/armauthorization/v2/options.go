//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

// ClassicAdministratorsClientListOptions contains the optional parameters for the ClassicAdministratorsClient.NewListPager
// method.
type ClassicAdministratorsClientListOptions struct {
	// placeholder for future optional parameters
}

// DenyAssignmentsClientGetByIDOptions contains the optional parameters for the DenyAssignmentsClient.GetByID method.
type DenyAssignmentsClientGetByIDOptions struct {
	// placeholder for future optional parameters
}

// DenyAssignmentsClientGetOptions contains the optional parameters for the DenyAssignmentsClient.Get method.
type DenyAssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DenyAssignmentsClientListForResourceGroupOptions contains the optional parameters for the DenyAssignmentsClient.NewListForResourceGroupPager
// method.
type DenyAssignmentsClientListForResourceGroupOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or above the scope. Use $filter=denyAssignmentName
	// eq '{name}' to search deny assignments by name at
	// specified scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for the
	// specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all
	// deny assignments at, above and below the scope for the specified principal. This filter is different from the principalId
	// filter as it returns not only those deny assignments that contain the
	// specified principal is the Principals list but also those deny assignments that contain the specified principal is the
	// ExcludePrincipals list. Additionally, when gdprExportPrincipalId filter is used,
	// only the deny assignment name and description properties are returned.
	Filter *string
}

// DenyAssignmentsClientListForResourceOptions contains the optional parameters for the DenyAssignmentsClient.NewListForResourcePager
// method.
type DenyAssignmentsClientListForResourceOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or above the scope. Use $filter=denyAssignmentName
	// eq '{name}' to search deny assignments by name at
	// specified scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for the
	// specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all
	// deny assignments at, above and below the scope for the specified principal. This filter is different from the principalId
	// filter as it returns not only those deny assignments that contain the
	// specified principal is the Principals list but also those deny assignments that contain the specified principal is the
	// ExcludePrincipals list. Additionally, when gdprExportPrincipalId filter is used,
	// only the deny assignment name and description properties are returned.
	Filter *string
}

// DenyAssignmentsClientListForScopeOptions contains the optional parameters for the DenyAssignmentsClient.NewListForScopePager
// method.
type DenyAssignmentsClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or above the scope. Use $filter=denyAssignmentName
	// eq '{name}' to search deny assignments by name at
	// specified scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for the
	// specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all
	// deny assignments at, above and below the scope for the specified principal. This filter is different from the principalId
	// filter as it returns not only those deny assignments that contain the
	// specified principal is the Principals list but also those deny assignments that contain the specified principal is the
	// ExcludePrincipals list. Additionally, when gdprExportPrincipalId filter is used,
	// only the deny assignment name and description properties are returned.
	Filter *string
}

// DenyAssignmentsClientListOptions contains the optional parameters for the DenyAssignmentsClient.NewListPager method.
type DenyAssignmentsClientListOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or above the scope. Use $filter=denyAssignmentName
	// eq '{name}' to search deny assignments by name at
	// specified scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for the
	// specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all
	// deny assignments at, above and below the scope for the specified principal. This filter is different from the principalId
	// filter as it returns not only those deny assignments that contain the
	// specified principal is the Principals list but also those deny assignments that contain the specified principal is the
	// ExcludePrincipals list. Additionally, when gdprExportPrincipalId filter is used,
	// only the deny assignment name and description properties are returned.
	Filter *string
}

// EligibleChildResourcesClientGetOptions contains the optional parameters for the EligibleChildResourcesClient.NewGetPager
// method.
type EligibleChildResourcesClientGetOptions struct {
	// The filter to apply on the operation. Use $filter=resourceType+eq+'Subscription' to filter on only resource of type = 'Subscription'.
	// Use
	// $filter=resourceType+eq+'subscription'+or+resourceType+eq+'resourcegroup' to filter on resource of type = 'Subscription'
	// or 'ResourceGroup'
	Filter *string
}

// GlobalAdministratorClientElevateAccessOptions contains the optional parameters for the GlobalAdministratorClient.ElevateAccess
// method.
type GlobalAdministratorClientElevateAccessOptions struct {
	// placeholder for future optional parameters
}

// PermissionsClientListForResourceGroupOptions contains the optional parameters for the PermissionsClient.NewListForResourceGroupPager
// method.
type PermissionsClientListForResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// PermissionsClientListForResourceOptions contains the optional parameters for the PermissionsClient.NewListForResourcePager
// method.
type PermissionsClientListForResourceOptions struct {
	// placeholder for future optional parameters
}

// ProviderOperationsMetadataClientGetOptions contains the optional parameters for the ProviderOperationsMetadataClient.Get
// method.
type ProviderOperationsMetadataClientGetOptions struct {
	// Specifies whether to expand the values.
	Expand *string
}

// ProviderOperationsMetadataClientListOptions contains the optional parameters for the ProviderOperationsMetadataClient.NewListPager
// method.
type ProviderOperationsMetadataClientListOptions struct {
	// Specifies whether to expand the values.
	Expand *string
}

// RoleAssignmentScheduleInstancesClientGetOptions contains the optional parameters for the RoleAssignmentScheduleInstancesClient.Get
// method.
type RoleAssignmentScheduleInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentScheduleInstancesClientListForScopeOptions contains the optional parameters for the RoleAssignmentScheduleInstancesClient.NewListForScopePager
// method.
type RoleAssignmentScheduleInstancesClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignment schedules at or above the scope.
	// Use $filter=principalId eq {id} to return all role assignment schedules at,
	// above or below the scope for the specified principal. Use $filter=assignedTo('{userId}') to return all role assignment
	// schedule instances for the user. Use $filter=asTarget() to return all role
	// assignment schedule instances created for the current user.
	Filter *string
}

// RoleAssignmentScheduleRequestsClientCancelOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Cancel
// method.
type RoleAssignmentScheduleRequestsClientCancelOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentScheduleRequestsClientCreateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Create
// method.
type RoleAssignmentScheduleRequestsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentScheduleRequestsClientGetOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Get
// method.
type RoleAssignmentScheduleRequestsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentScheduleRequestsClientListForScopeOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.NewListForScopePager
// method.
type RoleAssignmentScheduleRequestsClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignment schedule requests at or above
	// the scope. Use $filter=principalId eq {id} to return all role assignment
	// schedule requests at, above or below the scope for the specified principal. Use $filter=asRequestor() to return all role
	// assignment schedule requests requested by the current user. Use
	// $filter=asTarget() to return all role assignment schedule requests created for the current user. Use $filter=asApprover()
	// to return all role assignment schedule requests where the current user is an
	// approver.
	Filter *string
}

// RoleAssignmentScheduleRequestsClientValidateOptions contains the optional parameters for the RoleAssignmentScheduleRequestsClient.Validate
// method.
type RoleAssignmentScheduleRequestsClientValidateOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentSchedulesClientGetOptions contains the optional parameters for the RoleAssignmentSchedulesClient.Get method.
type RoleAssignmentSchedulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentSchedulesClientListForScopeOptions contains the optional parameters for the RoleAssignmentSchedulesClient.NewListForScopePager
// method.
type RoleAssignmentSchedulesClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignment schedules at or above the scope.
	// Use $filter=principalId eq {id} to return all role assignment schedules at,
	// above or below the scope for the specified principal. Use $filter=assignedTo('{userId}') to return all role assignment
	// schedules for the current user. Use $filter=asTarget() to return all role
	// assignment schedules created for the current user.
	Filter *string
}

// RoleAssignmentsClientCreateByIDOptions contains the optional parameters for the RoleAssignmentsClient.CreateByID method.
type RoleAssignmentsClientCreateByIDOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentsClientCreateOptions contains the optional parameters for the RoleAssignmentsClient.Create method.
type RoleAssignmentsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// RoleAssignmentsClientDeleteByIDOptions contains the optional parameters for the RoleAssignmentsClient.DeleteByID method.
type RoleAssignmentsClientDeleteByIDOptions struct {
	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientDeleteOptions contains the optional parameters for the RoleAssignmentsClient.Delete method.
type RoleAssignmentsClientDeleteOptions struct {
	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientGetByIDOptions contains the optional parameters for the RoleAssignmentsClient.GetByID method.
type RoleAssignmentsClientGetByIDOptions struct {
	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientGetOptions contains the optional parameters for the RoleAssignmentsClient.Get method.
type RoleAssignmentsClientGetOptions struct {
	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientListForResourceGroupOptions contains the optional parameters for the RoleAssignmentsClient.NewListForResourceGroupPager
// method.
type RoleAssignmentsClientListForResourceGroupOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the
	// scope for the specified principal.
	Filter *string

	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientListForResourceOptions contains the optional parameters for the RoleAssignmentsClient.NewListForResourcePager
// method.
type RoleAssignmentsClientListForResourceOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the
	// scope for the specified principal.
	Filter *string

	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientListForScopeOptions contains the optional parameters for the RoleAssignmentsClient.NewListForScopePager
// method.
type RoleAssignmentsClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the
	// scope for the specified principal.
	Filter *string

	// The skipToken to apply on the operation. Use $skipToken={skiptoken} to return paged role assignments following the skipToken
	// passed. Only supported on provider level calls.
	SkipToken *string

	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleAssignmentsClientListForSubscriptionOptions contains the optional parameters for the RoleAssignmentsClient.NewListForSubscriptionPager
// method.
type RoleAssignmentsClientListForSubscriptionOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the
	// scope for the specified principal.
	Filter *string

	// Tenant ID for cross-tenant request
	TenantID *string
}

// RoleDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the RoleDefinitionsClient.CreateOrUpdate
// method.
type RoleDefinitionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// RoleDefinitionsClientDeleteOptions contains the optional parameters for the RoleDefinitionsClient.Delete method.
type RoleDefinitionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RoleDefinitionsClientGetByIDOptions contains the optional parameters for the RoleDefinitionsClient.GetByID method.
type RoleDefinitionsClientGetByIDOptions struct {
	// placeholder for future optional parameters
}

// RoleDefinitionsClientGetOptions contains the optional parameters for the RoleDefinitionsClient.Get method.
type RoleDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleDefinitionsClientListOptions contains the optional parameters for the RoleDefinitionsClient.NewListPager method.
type RoleDefinitionsClientListOptions struct {
	// The filter to apply on the operation. Use atScopeAndBelow filter to search below the given scope as well.
	Filter *string
}

// RoleEligibilityScheduleInstancesClientGetOptions contains the optional parameters for the RoleEligibilityScheduleInstancesClient.Get
// method.
type RoleEligibilityScheduleInstancesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilityScheduleInstancesClientListForScopeOptions contains the optional parameters for the RoleEligibilityScheduleInstancesClient.NewListForScopePager
// method.
type RoleEligibilityScheduleInstancesClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignment schedules at or above the scope.
	// Use $filter=principalId eq {id} to return all role assignment schedules at,
	// above or below the scope for the specified principal. Use $filter=assignedTo('{userId}') to return all role eligibility
	// schedules for the user. Use $filter=asTarget() to return all role eligibility
	// schedules created for the current user.
	Filter *string
}

// RoleEligibilityScheduleRequestsClientCancelOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Cancel
// method.
type RoleEligibilityScheduleRequestsClientCancelOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilityScheduleRequestsClientCreateOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Create
// method.
type RoleEligibilityScheduleRequestsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilityScheduleRequestsClientGetOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Get
// method.
type RoleEligibilityScheduleRequestsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilityScheduleRequestsClientListForScopeOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.NewListForScopePager
// method.
type RoleEligibilityScheduleRequestsClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role eligibility schedule requests at or above
	// the scope. Use $filter=principalId eq {id} to return all role eligibility
	// schedule requests at, above or below the scope for the specified principal. Use $filter=asRequestor() to return all role
	// eligibility schedule requests requested by the current user. Use
	// $filter=asTarget() to return all role eligibility schedule requests created for the current user. Use $filter=asApprover()
	// to return all role eligibility schedule requests where the current user is an
	// approver.
	Filter *string
}

// RoleEligibilityScheduleRequestsClientValidateOptions contains the optional parameters for the RoleEligibilityScheduleRequestsClient.Validate
// method.
type RoleEligibilityScheduleRequestsClientValidateOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilitySchedulesClientGetOptions contains the optional parameters for the RoleEligibilitySchedulesClient.Get method.
type RoleEligibilitySchedulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleEligibilitySchedulesClientListForScopeOptions contains the optional parameters for the RoleEligibilitySchedulesClient.NewListForScopePager
// method.
type RoleEligibilitySchedulesClientListForScopeOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role eligibility schedules at or above the scope.
	// Use $filter=principalId eq {id} to return all role eligibility schedules at,
	// above or below the scope for the specified principal. Use $filter=assignedTo('{userId}') to return all role eligibility
	// schedules for the user. Use $filter=asTarget() to return all role eligibility
	// schedules created for the current user.
	Filter *string
}

// RoleManagementPoliciesClientDeleteOptions contains the optional parameters for the RoleManagementPoliciesClient.Delete
// method.
type RoleManagementPoliciesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPoliciesClientGetOptions contains the optional parameters for the RoleManagementPoliciesClient.Get method.
type RoleManagementPoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPoliciesClientListForScopeOptions contains the optional parameters for the RoleManagementPoliciesClient.NewListForScopePager
// method.
type RoleManagementPoliciesClientListForScopeOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPoliciesClientUpdateOptions contains the optional parameters for the RoleManagementPoliciesClient.Update
// method.
type RoleManagementPoliciesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPolicyAssignmentsClientCreateOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Create
// method.
type RoleManagementPolicyAssignmentsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPolicyAssignmentsClientDeleteOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Delete
// method.
type RoleManagementPolicyAssignmentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPolicyAssignmentsClientGetOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.Get
// method.
type RoleManagementPolicyAssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// RoleManagementPolicyAssignmentsClientListForScopeOptions contains the optional parameters for the RoleManagementPolicyAssignmentsClient.NewListForScopePager
// method.
type RoleManagementPolicyAssignmentsClientListForScopeOptions struct {
	// placeholder for future optional parameters
}
