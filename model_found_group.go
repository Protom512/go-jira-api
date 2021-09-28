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

// A group found in a search.
type FoundGroup struct {
	// The name of the group.
	Name string `json:"name,omitempty"`
	// The group name with the matched query string highlighted with the HTML bold tag.
	Html string `json:"html,omitempty"`
	Labels []GroupLabel `json:"labels,omitempty"`
	// The ID of the group, if available, which uniquely identifies the group across all Atlassian products. For example, *952d12c3-5b5b-4d04-bb32-44d383afc4b2*.
	GroupId string `json:"groupId,omitempty"`
}