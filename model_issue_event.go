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

// Details about an issue event.
type IssueEvent struct {
	// The ID of the event.
	Id int64 `json:"id,omitempty"`
	// The name of the event.
	Name string `json:"name,omitempty"`
}
