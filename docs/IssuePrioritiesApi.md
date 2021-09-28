# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriorities**](IssuePrioritiesApi.md#GetPriorities) | **Get** /rest/api/3/priority | Get priorities
[**GetPriority**](IssuePrioritiesApi.md#GetPriority) | **Get** /rest/api/3/priority/{id} | Get priority

# **GetPriorities**
> []ModelMap GetPriorities(ctx, )
Get priorities

Returns the list of all issue priorities.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPriority**
> ModelMap GetPriority(ctx, id)
Get priority

Returns an issue priority.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the issue priority. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

