//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.3, generator: @autorest/go@4.0.0-preview.43)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package dataprotectiondatasourceplugin

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

type PluginClient struct {
	pl runtime.Pipeline
}

// NewPluginClient creates a new instance of PluginClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewPluginClient(pl runtime.Pipeline) *PluginClient {
	client := &PluginClient{
		pl: pl,
	}
	return client
}

// Backup - Start the Backup operation
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientBackupOptions contains the optional parameters for the PluginClient.Backup method.
func (client *PluginClient) Backup(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters BackupRequest, options *PluginClientBackupOptions) (PluginClientBackupResponse, error) {
	req, err := client.backupCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientBackupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientBackupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientBackupResponse{}, runtime.NewResponseError(resp)
	}
	return client.backupHandleResponse(resp)
}

// backupCreateRequest creates the Backup request.
func (client *PluginClient) backupCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters BackupRequest, options *PluginClientBackupOptions) (*policy.Request, error) {
	urlPath := "/plugin:backup"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// backupHandleResponse handles the Backup response.
func (client *PluginClient) backupHandleResponse(resp *http.Response) (PluginClientBackupResponse, error) {
	result := PluginClientBackupResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientBackupResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResponse); err != nil {
		return PluginClientBackupResponse{}, err
	}
	return result, nil
}

// CommitOrRollbackBackup - Start the commitOrRollbackBackup operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientCommitOrRollbackBackupOptions contains the optional parameters for the PluginClient.CommitOrRollbackBackup
// method.
func (client *PluginClient) CommitOrRollbackBackup(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CommitOrRollbackBackupRequest, options *PluginClientCommitOrRollbackBackupOptions) (PluginClientCommitOrRollbackBackupResponse, error) {
	req, err := client.commitOrRollbackBackupCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientCommitOrRollbackBackupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientCommitOrRollbackBackupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientCommitOrRollbackBackupResponse{}, runtime.NewResponseError(resp)
	}
	return client.commitOrRollbackBackupHandleResponse(resp)
}

// commitOrRollbackBackupCreateRequest creates the CommitOrRollbackBackup request.
func (client *PluginClient) commitOrRollbackBackupCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CommitOrRollbackBackupRequest, options *PluginClientCommitOrRollbackBackupOptions) (*policy.Request, error) {
	urlPath := "/plugin:commitOrRollbackBackup"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// commitOrRollbackBackupHandleResponse handles the CommitOrRollbackBackup response.
func (client *PluginClient) commitOrRollbackBackupHandleResponse(resp *http.Response) (PluginClientCommitOrRollbackBackupResponse, error) {
	result := PluginClientCommitOrRollbackBackupResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientCommitOrRollbackBackupResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitOrRollbackBackupResponse); err != nil {
		return PluginClientCommitOrRollbackBackupResponse{}, err
	}
	return result, nil
}

// CommitOrRollbackRestore - Start the commitOrRollbackRestore operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientCommitOrRollbackRestoreOptions contains the optional parameters for the PluginClient.CommitOrRollbackRestore
// method.
func (client *PluginClient) CommitOrRollbackRestore(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CommitOrRollbackRestoreRequest, options *PluginClientCommitOrRollbackRestoreOptions) (PluginClientCommitOrRollbackRestoreResponse, error) {
	req, err := client.commitOrRollbackRestoreCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientCommitOrRollbackRestoreResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientCommitOrRollbackRestoreResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientCommitOrRollbackRestoreResponse{}, runtime.NewResponseError(resp)
	}
	return client.commitOrRollbackRestoreHandleResponse(resp)
}

// commitOrRollbackRestoreCreateRequest creates the CommitOrRollbackRestore request.
func (client *PluginClient) commitOrRollbackRestoreCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CommitOrRollbackRestoreRequest, options *PluginClientCommitOrRollbackRestoreOptions) (*policy.Request, error) {
	urlPath := "/plugin:commitOrRollbackRestore"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// commitOrRollbackRestoreHandleResponse handles the CommitOrRollbackRestore response.
func (client *PluginClient) commitOrRollbackRestoreHandleResponse(resp *http.Response) (PluginClientCommitOrRollbackRestoreResponse, error) {
	result := PluginClientCommitOrRollbackRestoreResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientCommitOrRollbackRestoreResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommitOrRollbackRestoreResponse); err != nil {
		return PluginClientCommitOrRollbackRestoreResponse{}, err
	}
	return result, nil
}

// Restore - Start the restore operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientRestoreOptions contains the optional parameters for the PluginClient.Restore method.
func (client *PluginClient) Restore(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RestoreRequest, options *PluginClientRestoreOptions) (PluginClientRestoreResponse, error) {
	req, err := client.restoreCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientRestoreResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientRestoreResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientRestoreResponse{}, runtime.NewResponseError(resp)
	}
	return client.restoreHandleResponse(resp)
}

// restoreCreateRequest creates the Restore request.
func (client *PluginClient) restoreCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RestoreRequest, options *PluginClientRestoreOptions) (*policy.Request, error) {
	urlPath := "/plugin:restore"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// restoreHandleResponse handles the Restore response.
func (client *PluginClient) restoreHandleResponse(resp *http.Response) (PluginClientRestoreResponse, error) {
	result := PluginClientRestoreResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientRestoreResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestoreResponse); err != nil {
		return PluginClientRestoreResponse{}, err
	}
	return result, nil
}

// StartProtection - Start the startProtection operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientStartProtectionOptions contains the optional parameters for the PluginClient.StartProtection method.
func (client *PluginClient) StartProtection(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters StartProtectionRequest, options *PluginClientStartProtectionOptions) (PluginClientStartProtectionResponse, error) {
	req, err := client.startProtectionCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientStartProtectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientStartProtectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientStartProtectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.startProtectionHandleResponse(resp)
}

// startProtectionCreateRequest creates the StartProtection request.
func (client *PluginClient) startProtectionCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters StartProtectionRequest, options *PluginClientStartProtectionOptions) (*policy.Request, error) {
	urlPath := "/plugin:startProtection"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// startProtectionHandleResponse handles the StartProtection response.
func (client *PluginClient) startProtectionHandleResponse(resp *http.Response) (PluginClientStartProtectionResponse, error) {
	result := PluginClientStartProtectionResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientStartProtectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StartProtectionResponse); err != nil {
		return PluginClientStartProtectionResponse{}, err
	}
	return result, nil
}

// StopProtection - Start the stopProtection operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientStopProtectionOptions contains the optional parameters for the PluginClient.StopProtection method.
func (client *PluginClient) StopProtection(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters StopProtectionRequest, options *PluginClientStopProtectionOptions) (PluginClientStopProtectionResponse, error) {
	req, err := client.stopProtectionCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientStopProtectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientStopProtectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientStopProtectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.stopProtectionHandleResponse(resp)
}

// stopProtectionCreateRequest creates the StopProtection request.
func (client *PluginClient) stopProtectionCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters StopProtectionRequest, options *PluginClientStopProtectionOptions) (*policy.Request, error) {
	urlPath := "/plugin:stopProtection"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// stopProtectionHandleResponse handles the StopProtection response.
func (client *PluginClient) stopProtectionHandleResponse(resp *http.Response) (PluginClientStopProtectionResponse, error) {
	result := PluginClientStopProtectionResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientStopProtectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StopProtectionResponse); err != nil {
		return PluginClientStopProtectionResponse{}, err
	}
	return result, nil
}

// UpdateProtection - Start the updateProtection operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientUpdateProtectionOptions contains the optional parameters for the PluginClient.UpdateProtection method.
func (client *PluginClient) UpdateProtection(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters UpdateProtectionRequest, options *PluginClientUpdateProtectionOptions) (PluginClientUpdateProtectionResponse, error) {
	req, err := client.updateProtectionCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientUpdateProtectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientUpdateProtectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientUpdateProtectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateProtectionHandleResponse(resp)
}

// updateProtectionCreateRequest creates the UpdateProtection request.
func (client *PluginClient) updateProtectionCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters UpdateProtectionRequest, options *PluginClientUpdateProtectionOptions) (*policy.Request, error) {
	urlPath := "/plugin:updateProtection"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateProtectionHandleResponse handles the UpdateProtection response.
func (client *PluginClient) updateProtectionHandleResponse(resp *http.Response) (PluginClientUpdateProtectionResponse, error) {
	result := PluginClientUpdateProtectionResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientUpdateProtectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateProtectionResponse); err != nil {
		return PluginClientUpdateProtectionResponse{}, err
	}
	return result, nil
}

// ValidateForBackup - Start the validateForBackup operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientValidateForBackupOptions contains the optional parameters for the PluginClient.ValidateForBackup
// method.
func (client *PluginClient) ValidateForBackup(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForBackupRequest, options *PluginClientValidateForBackupOptions) (PluginClientValidateForBackupResponse, error) {
	req, err := client.validateForBackupCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientValidateForBackupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientValidateForBackupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientValidateForBackupResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateForBackupHandleResponse(resp)
}

// validateForBackupCreateRequest creates the ValidateForBackup request.
func (client *PluginClient) validateForBackupCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForBackupRequest, options *PluginClientValidateForBackupOptions) (*policy.Request, error) {
	urlPath := "/plugin:validateForBackup"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateForBackupHandleResponse handles the ValidateForBackup response.
func (client *PluginClient) validateForBackupHandleResponse(resp *http.Response) (PluginClientValidateForBackupResponse, error) {
	result := PluginClientValidateForBackupResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientValidateForBackupResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateForBackupResponse); err != nil {
		return PluginClientValidateForBackupResponse{}, err
	}
	return result, nil
}

// ValidateForProtection - Start the validateForProtection operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientValidateForProtectionOptions contains the optional parameters for the PluginClient.ValidateForProtection
// method.
func (client *PluginClient) ValidateForProtection(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForProtectionRequest, options *PluginClientValidateForProtectionOptions) (PluginClientValidateForProtectionResponse, error) {
	req, err := client.validateForProtectionCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientValidateForProtectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientValidateForProtectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientValidateForProtectionResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateForProtectionHandleResponse(resp)
}

// validateForProtectionCreateRequest creates the ValidateForProtection request.
func (client *PluginClient) validateForProtectionCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForProtectionRequest, options *PluginClientValidateForProtectionOptions) (*policy.Request, error) {
	urlPath := "/plugin:validateForProtection"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateForProtectionHandleResponse handles the ValidateForProtection response.
func (client *PluginClient) validateForProtectionHandleResponse(resp *http.Response) (PluginClientValidateForProtectionResponse, error) {
	result := PluginClientValidateForProtectionResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientValidateForProtectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateForProtectionResponse); err != nil {
		return PluginClientValidateForProtectionResponse{}, err
	}
	return result, nil
}

// ValidateForRestore - Start the validateForRestore operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique id of the LRO
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - PluginClientValidateForRestoreOptions contains the optional parameters for the PluginClient.ValidateForRestore
// method.
func (client *PluginClient) ValidateForRestore(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForRestoreRequest, options *PluginClientValidateForRestoreOptions) (PluginClientValidateForRestoreResponse, error) {
	req, err := client.validateForRestoreCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return PluginClientValidateForRestoreResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PluginClientValidateForRestoreResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return PluginClientValidateForRestoreResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateForRestoreHandleResponse(resp)
}

// validateForRestoreCreateRequest creates the ValidateForRestore request.
func (client *PluginClient) validateForRestoreCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters ValidateForRestoreRequest, options *PluginClientValidateForRestoreOptions) (*policy.Request, error) {
	urlPath := "/plugin:validateForRestore"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["operation-Id"] = []string{operationID}
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// validateForRestoreHandleResponse handles the ValidateForRestore response.
func (client *PluginClient) validateForRestoreHandleResponse(resp *http.Response) (PluginClientValidateForRestoreResponse, error) {
	result := PluginClientValidateForRestoreResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PluginClientValidateForRestoreResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("Operation-Location"); val != "" {
		result.OperationLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateForRestoreResponse); err != nil {
		return PluginClientValidateForRestoreResponse{}, err
	}
	return result, nil
}

