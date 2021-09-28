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

// A workflow transition.
type WorkflowTransition struct {
	// The transition ID.
	Id int32 `json:"id"`
	// The transition name.
	Name string `json:"name"`
}
