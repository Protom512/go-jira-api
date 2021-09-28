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

// Properties that identify a published workflow.
type PublishedWorkflowId struct {
	// The name of the workflow.
	Name string `json:"name"`
	// The entity ID of the workflow.
	EntityId string `json:"entityId,omitempty"`
}