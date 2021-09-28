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

// Details of workflow transition rules.
type WorkflowRulesSearchDetails struct {
	// The workflow ID.
	WorkflowEntityId string `json:"workflowEntityId,omitempty"`
	// List of workflow rule IDs that do not belong to the workflow or can not be found.
	InvalidRules []string `json:"invalidRules,omitempty"`
	// List of valid workflow transition rules.
	ValidRules []WorkflowTransitionRules `json:"validRules,omitempty"`
}