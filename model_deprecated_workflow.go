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
type DeprecatedWorkflow struct {
	// The name of the workflow.
	Name string `json:"name,omitempty"`
	// The description of the workflow.
	Description string `json:"description,omitempty"`
	// The datetime the workflow was last modified.
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	LastModifiedUser string `json:"lastModifiedUser,omitempty"`
	// The account ID of the user that last modified the workflow.
	LastModifiedUserAccountId string `json:"lastModifiedUserAccountId,omitempty"`
	// The number of steps included in the workflow.
	Steps int32 `json:"steps,omitempty"`
	// The scope where this workflow applies
	Scope *AllOfDeprecatedWorkflowScope `json:"scope,omitempty"`
	Default_ bool `json:"default,omitempty"`
}
