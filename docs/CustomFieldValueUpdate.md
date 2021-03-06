# CustomFieldValueUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueIds** | **[]int64** | The list of issue IDs. | [default to null]
**Value** | [***Object**](.md) | The value for the custom field. The value must be compatible with the [custom field type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#data-types) as follows:   *  &#x60;string&#x60; the value must be a string.  *  &#x60;number&#x60; the value must be a number.  *  &#x60;datetime&#x60; the value must be a string that represents a date in the ISO format, for example &#x60;\&quot;2021-01-18T12:00:00-03:00\&quot;&#x60;.  *  &#x60;user&#x60; the value must be an object that contains the &#x60;accountId&#x60; field.  *  &#x60;group&#x60; the value must be an object that contains the group &#x60;name&#x60; field.  A list of appropriate values must be provided if the field is of the &#x60;list&#x60; [collection type](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/#collection-types). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

