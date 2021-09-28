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

// A paginated list of users sharing the filter. This includes users that are members of the groups or can browse the projects that the filter is shared with.
type UserList struct {
	// The number of items on the page.
	Size int32 `json:"size,omitempty"`
	// The list of items.
	Items []User `json:"items,omitempty"`
	// The maximum number of results that could be on the page.
	MaxResults int32 `json:"max-results,omitempty"`
	// The index of the first item returned on the page.
	StartIndex int32 `json:"start-index,omitempty"`
	// The index of the last item returned on the page.
	EndIndex int32 `json:"end-index,omitempty"`
}
