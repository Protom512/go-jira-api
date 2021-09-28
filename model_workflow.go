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

// Details about a workflow.
type Workflow struct {
	Id *PublishedWorkflowId `json:"id"`
	// The description of the workflow.
	Description string `json:"description"`
	// The transitions of the workflow.
	Transitions []Transition `json:"transitions,omitempty"`
	// The statuses of the workflow.
	Statuses []WorkflowStatus `json:"statuses,omitempty"`
	// Whether this is the default workflow.
	IsDefault bool `json:"isDefault,omitempty"`
}