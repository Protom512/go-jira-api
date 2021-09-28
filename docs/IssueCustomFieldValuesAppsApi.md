# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCustomFieldValue**](IssueCustomFieldValuesAppsApi.md#UpdateCustomFieldValue) | **Put** /rest/api/3/app/field/{fieldIdOrKey}/value | Update custom field value

# **UpdateCustomFieldValue**
> Object UpdateCustomFieldValue(ctx, body, fieldIdOrKey, optional)
Update custom field value

Updates the value of a custom field on one or more issues. Custom fields can only be updated by the Forge app that created them.  **[Permissions](#permissions) required:** Only the app that created the custom field can update its values with this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomFieldValueUpdateRequest**](CustomFieldValueUpdateRequest.md)|  | 
  **fieldIdOrKey** | **string**| The ID or key of the custom field. For example, &#x60;customfield_10010&#x60;. | 
 **optional** | ***IssueCustomFieldValuesAppsApiUpdateCustomFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueCustomFieldValuesAppsApiUpdateCustomFieldValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **generateChangelog** | **optional.**| Whether to generate a changelog for this update. Default: true. | [default to true]

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

