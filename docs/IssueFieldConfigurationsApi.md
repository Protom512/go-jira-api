# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignFieldConfigurationSchemeToProject**](IssueFieldConfigurationsApi.md#AssignFieldConfigurationSchemeToProject) | **Put** /rest/api/3/fieldconfigurationscheme/project | Assign field configuration scheme to project
[**GetAllFieldConfigurationSchemes**](IssueFieldConfigurationsApi.md#GetAllFieldConfigurationSchemes) | **Get** /rest/api/3/fieldconfigurationscheme | Get all field configuration schemes
[**GetAllFieldConfigurations**](IssueFieldConfigurationsApi.md#GetAllFieldConfigurations) | **Get** /rest/api/3/fieldconfiguration | Get all field configurations
[**GetFieldConfigurationItems**](IssueFieldConfigurationsApi.md#GetFieldConfigurationItems) | **Get** /rest/api/3/fieldconfiguration/{id}/fields | Get field configuration items
[**GetFieldConfigurationSchemeMappings**](IssueFieldConfigurationsApi.md#GetFieldConfigurationSchemeMappings) | **Get** /rest/api/3/fieldconfigurationscheme/mapping | Get field configuration issue type items
[**GetFieldConfigurationSchemeProjectMapping**](IssueFieldConfigurationsApi.md#GetFieldConfigurationSchemeProjectMapping) | **Get** /rest/api/3/fieldconfigurationscheme/project | Get field configuration schemes for projects

# **AssignFieldConfigurationSchemeToProject**
> Object AssignFieldConfigurationSchemeToProject(ctx, body)
Assign field configuration scheme to project

Assigns a field configuration scheme to a project. If the field configuration scheme ID is `null`, the operation assigns the default field configuration scheme.  Field configuration schemes can only be assigned to classic projects.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FieldConfigurationSchemeProjectAssociation**](FieldConfigurationSchemeProjectAssociation.md)|  | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFieldConfigurationSchemes**
> PageBeanFieldConfigurationScheme GetAllFieldConfigurationSchemes(ctx, optional)
Get all field configuration schemes

Returns a [paginated](#pagination) list of field configuration schemes.  Only field configuration schemes used in classic projects are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IssueFieldConfigurationsApiGetAllFieldConfigurationSchemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueFieldConfigurationsApiGetAllFieldConfigurationSchemesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]
 **id** | [**optional.Interface of []int64**](int64.md)| The list of field configuration scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 

### Return type

[**PageBeanFieldConfigurationScheme**](PageBeanFieldConfigurationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFieldConfigurations**
> PageBeanFieldConfiguration GetAllFieldConfigurations(ctx, optional)
Get all field configurations

Returns a [paginated](#pagination) list of field configurations. The list can be for all field configurations or a subset determined by any combination of these criteria:   *  a list of field configuration item IDs.  *  whether the field configuration is a default.  *  whether the field configuration name or description contains a query string.  Only field configurations used in classic projects are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IssueFieldConfigurationsApiGetAllFieldConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueFieldConfigurationsApiGetAllFieldConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]
 **id** | [**optional.Interface of []int64**](int64.md)| The list of field configuration IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **isDefault** | **optional.Bool**| If *true* returns default field configurations only. | [default to false]
 **query** | **optional.String**| The query string used to match against field configuration names and descriptions. | 

### Return type

[**PageBeanFieldConfiguration**](PageBeanFieldConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldConfigurationItems**
> PageBeanFieldConfigurationItem GetFieldConfigurationItems(ctx, id, optional)
Get field configuration items

Returns a [paginated](#pagination) list of all fields for a configuration.  Only the fields from configurations used in classic projects are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the field configuration. | 
 **optional** | ***IssueFieldConfigurationsApiGetFieldConfigurationItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueFieldConfigurationsApiGetFieldConfigurationItemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanFieldConfigurationItem**](PageBeanFieldConfigurationItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldConfigurationSchemeMappings**
> PageBeanFieldConfigurationIssueTypeItem GetFieldConfigurationSchemeMappings(ctx, optional)
Get field configuration issue type items

Returns a [paginated](#pagination) list of field configuration issue type items.  Only items used in classic projects are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IssueFieldConfigurationsApiGetFieldConfigurationSchemeMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueFieldConfigurationsApiGetFieldConfigurationSchemeMappingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]
 **fieldConfigurationSchemeId** | [**optional.Interface of []int64**](int64.md)| The list of field configuration scheme IDs. To include multiple field configuration schemes separate IDs with ampersand: &#x60;fieldConfigurationSchemeId&#x3D;10000&amp;fieldConfigurationSchemeId&#x3D;10001&#x60;. | 

### Return type

[**PageBeanFieldConfigurationIssueTypeItem**](PageBeanFieldConfigurationIssueTypeItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFieldConfigurationSchemeProjectMapping**
> PageBeanFieldConfigurationSchemeProjects GetFieldConfigurationSchemeProjectMapping(ctx, projectId, optional)
Get field configuration schemes for projects

Returns a [paginated](#pagination) list of field configuration schemes and, for each scheme, a list of the projects that use it.  The list is sorted by field configuration scheme ID. The first item contains the list of project IDs assigned to the default field configuration scheme.  Only field configuration schemes used in classic projects are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | [**[]int64**](int64.md)| The list of project IDs. To include multiple projects, separate IDs with ampersand: &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 
 **optional** | ***IssueFieldConfigurationsApiGetFieldConfigurationSchemeProjectMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueFieldConfigurationsApiGetFieldConfigurationSchemeProjectMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanFieldConfigurationSchemeProjects**](PageBeanFieldConfigurationSchemeProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

