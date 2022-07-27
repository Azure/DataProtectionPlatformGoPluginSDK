//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.0, generator: @autorest/go@4.0.0-preview.42)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dataprotectiondatasourceplugin

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type StopProtectionClient struct {
	pl runtime.Pipeline
}

// NewStopProtectionClient creates a new instance of StopProtectionClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewStopProtectionClient(pl runtime.Pipeline) *StopProtectionClient {
	client := &StopProtectionClient{
		pl: pl,
	}
	return client
}

// Cancel - Cancel the operation. Poll the LRO to get the final status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - StopProtectionClientCancelOptions contains the optional parameters for the StopProtectionClient.Cancel method.
func (client *StopProtectionClient) Cancel(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *StopProtectionClientCancelOptions) (StopProtectionClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return StopProtectionClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StopProtectionClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StopProtectionClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *StopProtectionClient) cancelCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *StopProtectionClientCancelOptions) (*policy.Request, error) {
	urlPath := "/plugin/stopProtectionOperations/{operationId}:cancel"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// cancelHandleResponse handles the Cancel response.
func (client *StopProtectionClient) cancelHandleResponse(resp *http.Response) (StopProtectionClientCancelResponse, error) {
	result := StopProtectionClientCancelResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return StopProtectionClientCancelResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StopProtectionResponse); err != nil {
		return StopProtectionClientCancelResponse{}, err
	}
	return result, nil
}

// Get - Gets the status of a stopProtection LRO.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - Unique id of the resource
// taskID - Unique id of the current task
// xmsClientRequestID - Correlation request Id for tracking a particular request.
// options - StopProtectionClientGetOptions contains the optional parameters for the StopProtectionClient.Get method.
func (client *StopProtectionClient) Get(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *StopProtectionClientGetOptions) (StopProtectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, options)
	if err != nil {
		return StopProtectionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StopProtectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StopProtectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StopProtectionClient) getCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *StopProtectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/plugin/stopProtectionOperations/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StopProtectionClient) getHandleResponse(resp *http.Response) (StopProtectionClientGetResponse, error) {
	result := StopProtectionClientGetResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return StopProtectionClientGetResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StopProtectionResponse); err != nil {
		return StopProtectionClientGetResponse{}, err
	}
	return result, nil
}

// RefreshTokens - Refresh tokens for a given operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - StopProtectionClientRefreshTokensOptions contains the optional parameters for the StopProtectionClient.RefreshTokens
// method.
func (client *StopProtectionClient) RefreshTokens(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *StopProtectionClientRefreshTokensOptions) (StopProtectionClientRefreshTokensResponse, error) {
	req, err := client.refreshTokensCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return StopProtectionClientRefreshTokensResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StopProtectionClientRefreshTokensResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StopProtectionClientRefreshTokensResponse{}, runtime.NewResponseError(resp)
	}
	return client.refreshTokensHandleResponse(resp)
}

// refreshTokensCreateRequest creates the RefreshTokens request.
func (client *StopProtectionClient) refreshTokensCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *StopProtectionClientRefreshTokensOptions) (*policy.Request, error) {
	urlPath := "/plugin/stopProtectionOperations/{operationId}:refreshTokens"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// refreshTokensHandleResponse handles the RefreshTokens response.
func (client *StopProtectionClient) refreshTokensHandleResponse(resp *http.Response) (StopProtectionClientRefreshTokensResponse, error) {
	result := StopProtectionClientRefreshTokensResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return StopProtectionClientRefreshTokensResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.StopProtectionResponse); err != nil {
		return StopProtectionClientRefreshTokensResponse{}, err
	}
	return result, nil
}

