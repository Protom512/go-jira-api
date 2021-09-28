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

// The details of a workflow transition rules.
type CreateWorkflowTransitionRulesDetails struct {
	// The workflow conditions.
	Conditions *AllOfCreateWorkflowTransitionRulesDetailsConditions `json:"conditions,omitempty"`
	// The workflow validators.
	Validators []CreateWorkflowTransitionRule `json:"validators,omitempty"`
	// The workflow post functions.
	PostFunctions []CreateWorkflowTransitionRule `json:"postFunctions,omitempty"`
}