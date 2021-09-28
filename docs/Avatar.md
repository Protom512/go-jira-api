# Avatar

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the avatar. | [default to null]
**Owner** | **string** | The owner of the avatar. For a system avatar the owner is null (and nothing is returned). For non-system avatars this is the appropriate identifier, such as the ID for a project or the account ID for a user. | [optional] [default to null]
**IsSystemAvatar** | **bool** | Whether the avatar is a system avatar. | [optional] [default to null]
**IsSelected** | **bool** | Whether the avatar is used in Jira. For example, shown as a project&#x27;s avatar. | [optional] [default to null]
**IsDeletable** | **bool** | Whether the avatar can be deleted. | [optional] [default to null]
**FileName** | **string** | The file name of the avatar icon. Returned for system avatars. | [optional] [default to null]
**Urls** | **map[string]string** | The list of avatar icon URLs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

