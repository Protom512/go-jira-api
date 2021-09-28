# EventNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | **string** | Expand options that include additional event notification details in the response. | [optional] [default to null]
**Id** | **int64** | The ID of the notification. | [optional] [default to null]
**NotificationType** | **string** | Identifies the recipients of the notification. | [optional] [default to null]
**Parameter** | **string** | The value of the &#x60;notificationType&#x60;:   *  &#x60;User&#x60; The &#x60;parameter&#x60; is the user account ID.  *  &#x60;Group&#x60; The &#x60;parameter&#x60; is the group name.  *  &#x60;ProjectRole&#x60; The &#x60;parameter&#x60; is the project role ID.  *  &#x60;UserCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field.  *  &#x60;GroupCustomField&#x60; The &#x60;parameter&#x60; is the ID of the custom field. | [optional] [default to null]
**Group** | [***AllOfEventNotificationGroup**](AllOfEventNotificationGroup.md) | The specified group. | [optional] [default to null]
**Field** | [***AllOfEventNotificationField**](AllOfEventNotificationField.md) | The custom user or group field. | [optional] [default to null]
**EmailAddress** | **string** | The email address. | [optional] [default to null]
**ProjectRole** | [***AllOfEventNotificationProjectRole**](AllOfEventNotificationProjectRole.md) | The specified project role. | [optional] [default to null]
**User** | [***AllOfEventNotificationUser**](AllOfEventNotificationUser.md) | The specified user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

