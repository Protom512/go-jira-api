# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResolution**](IssueResolutionsApi.md#GetResolution) | **Get** /rest/api/3/resolution/{id} | Get resolution
[**GetResolutions**](IssueResolutionsApi.md#GetResolutions) | **Get** /rest/api/3/resolution | Get resolutions

# **GetResolution**
> Resolution GetResolution(ctx, id)
Get resolution

Returns an issue resolution value.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the issue resolution value. | 

### Return type

[**Resolution**](Resolution.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResolutions**
> []Resolution GetResolutions(ctx, )
Get resolutions

Returns a list of all issue resolution values.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Resolution**](Resolution.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

