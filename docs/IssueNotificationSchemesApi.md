# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationScheme**](IssueNotificationSchemesApi.md#GetNotificationScheme) | **Get** /rest/api/3/notificationscheme/{id} | Get notification scheme
[**GetNotificationSchemes**](IssueNotificationSchemesApi.md#GetNotificationSchemes) | **Get** /rest/api/3/notificationscheme | Get notification schemes paginated

# **GetNotificationScheme**
> NotificationScheme GetNotificationScheme(ctx, id, optional)
Get notification scheme

Returns a [notification scheme](https://confluence.atlassian.com/x/8YdKLg), including the list of events and the recipients who will receive notifications for those events.  **[Permissions](#permissions) required:** Permission to access Jira, however the user must have permission to administer at least one project associated with the notification scheme.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the notification scheme. Use [Get notification schemes paginated](#api-rest-api-3-notificationscheme-get) to get a list of notification scheme IDs. | 
 **optional** | ***IssueNotificationSchemesApiGetNotificationSchemeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueNotificationSchemesApiGetNotificationSchemeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **optional.String**| Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about any custom fields assigned to receive an event.  *  &#x60;group&#x60; Returns information about any groups assigned to receive an event.  *  &#x60;notificationSchemeEvents&#x60; Returns a list of event associations. This list is returned for all expandable information.  *  &#x60;projectRole&#x60; Returns information about any project roles assigned to receive an event.  *  &#x60;user&#x60; Returns information about any users assigned to receive an event. | 

### Return type

[**NotificationScheme**](NotificationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationSchemes**
> PageBeanNotificationScheme GetNotificationSchemes(ctx, optional)
Get notification schemes paginated

Returns a [paginated](#pagination) list of [notification schemes](https://confluence.atlassian.com/x/8YdKLg) ordered by display name.  ### About notification schemes ###  A notification scheme is a list of events and recipients who will receive notifications for those events. The list is contained within the `notificationSchemeEvents` object and contains pairs of `events` and `notifications`:   *  `event` Identifies the type of event. The events can be [Jira system events](https://confluence.atlassian.com/x/8YdKLg#Creatinganotificationscheme-eventsEvents) or [custom events](https://confluence.atlassian.com/x/AIlKLg).  *  `notifications` Identifies the [recipients](https://confluence.atlassian.com/x/8YdKLg#Creatinganotificationscheme-recipientsRecipients) of notifications for each event. Recipients can be any of the following types:           *  `CurrentAssignee`      *  `Reporter`      *  `CurrentUser`      *  `ProjectLead`      *  `ComponentLead`      *  `User` (the `parameter` is the user key)      *  `Group` (the `parameter` is the group name)      *  `ProjectRole` (the `parameter` is the project role ID)      *  `EmailAddress`      *  `AllWatchers`      *  `UserCustomField` (the `parameter` is the ID of the custom field)      *  `GroupCustomField`(the `parameter` is the ID of the custom field)  *Note that you should allow for events without recipients to appear in responses.*  **[Permissions](#permissions) required:** Permission to access Jira, however the user must have permission to administer at least one project associated with a notification scheme for it to be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IssueNotificationSchemesApiGetNotificationSchemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueNotificationSchemesApiGetNotificationSchemesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 50]
 **expand** | **optional.String**| Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about any custom fields assigned to receive an event.  *  &#x60;group&#x60; Returns information about any groups assigned to receive an event.  *  &#x60;notificationSchemeEvents&#x60; Returns a list of event associations. This list is returned for all expandable information.  *  &#x60;projectRole&#x60; Returns information about any project roles assigned to receive an event.  *  &#x60;user&#x60; Returns information about any users assigned to receive an event. | 

### Return type

[**PageBeanNotificationScheme**](PageBeanNotificationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

