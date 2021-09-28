/*
 * The Jira Cloud platform REST API
 *
 * Jira Cloud platform REST API documentation
 *
 * API version: 1001.0.0-SNAPSHOT
 * Contact: ecosystem@atlassian.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The groups that the user belongs to.
type AllOfUserGroups struct {
	Size int32 `json:"size,omitempty"`
	Items []GroupName `json:"items,omitempty"`
	PagingCallback *ListWrapperCallbackGroupName `json:"pagingCallback,omitempty"`
	Callback *ListWrapperCallbackGroupName `json:"callback,omitempty"`
	MaxResults int32 `json:"max-results,omitempty"`
}
