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

type SimpleListWrapperApplicationRole struct {
	Size int32 `json:"size,omitempty"`
	Items []ApplicationRole `json:"items,omitempty"`
	PagingCallback *ListWrapperCallbackApplicationRole `json:"pagingCallback,omitempty"`
	Callback *ListWrapperCallbackApplicationRole `json:"callback,omitempty"`
	MaxResults int32 `json:"max-results,omitempty"`
}
