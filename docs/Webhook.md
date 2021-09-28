# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the webhook. | [default to null]
**JqlFilter** | **string** | The JQL filter that specifies which issues the webhook is sent for. | [default to null]
**FieldIdsFilter** | **[]string** | A list of field IDs. When the issue changelog contains any of the fields, the webhook &#x60;jira:issue_updated&#x60; is sent. If this parameter is not present, the app is notified about all field updates. | [optional] [default to null]
**IssuePropertyKeysFilter** | **[]string** | A list of issue property keys. A change of those issue properties triggers the &#x60;issue_property_set&#x60; or &#x60;issue_property_deleted&#x60; webhooks. If this parameter is not present, the app is notified about all issue property updates. | [optional] [default to null]
**Events** | **[]string** | The Jira events that trigger the webhook. | [default to null]
**ExpirationDate** | **int64** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

