# IssueTypeCreateBean

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name for the issue type. The maximum length is 60 characters. | [default to null]
**Description** | **string** | The description of the issue type. | [optional] [default to null]
**Type_** | **string** | Deprecated. Use &#x60;hierarchyLevel&#x60; instead.  Whether the issue type is &#x60;subtype&#x60; or &#x60;standard&#x60;. Defaults to &#x60;standard&#x60;. | [optional] [default to null]
**HierarchyLevel** | **int32** | The hierarchy level of the issue type. Use:   *  &#x60;-1&#x60; for Subtask.  *  &#x60;0&#x60; for Base.  Defaults to &#x60;0&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

