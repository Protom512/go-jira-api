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

// Details about a workflow scheme.
type WorkflowScheme struct {
	// The ID of the workflow scheme.
	Id int64 `json:"id,omitempty"`
	// The name of the workflow scheme. The name must be unique. The maximum length is 255 characters. Required when creating a workflow scheme.
	Name string `json:"name,omitempty"`
	// The description of the workflow scheme.
	Description string `json:"description,omitempty"`
	// The name of the default workflow for the workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira. If `defaultWorkflow` is not specified when creating a workflow scheme, it is set to *Jira Workflow (jira)*.
	DefaultWorkflow string `json:"defaultWorkflow,omitempty"`
	// The issue type to workflow mappings, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme.
	IssueTypeMappings map[string]string `json:"issueTypeMappings,omitempty"`
	// For draft workflow schemes, this property is the name of the default workflow for the original workflow scheme. The default workflow has *All Unassigned Issue Types* assigned to it in Jira.
	OriginalDefaultWorkflow string `json:"originalDefaultWorkflow,omitempty"`
	// For draft workflow schemes, this property is the issue type to workflow mappings for the original workflow scheme, where each mapping is an issue type ID and workflow name pair. Note that an issue type can only be mapped to one workflow in a workflow scheme.
	OriginalIssueTypeMappings map[string]string `json:"originalIssueTypeMappings,omitempty"`
	// Whether the workflow scheme is a draft or not.
	Draft bool `json:"draft,omitempty"`
	// The user that last modified the draft workflow scheme. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows.
	LastModifiedUser *AllOfWorkflowSchemeLastModifiedUser `json:"lastModifiedUser,omitempty"`
	// The date-time that the draft workflow scheme was last modified. A modification is a change to the issue type-project mappings only. This property does not apply to non-draft workflows.
	LastModified string `json:"lastModified,omitempty"`
	Self string `json:"self,omitempty"`
	// Whether to create or update a draft workflow scheme when updating an active workflow scheme. An active workflow scheme is a workflow scheme that is used by at least one project. The following examples show how this property works:   *  Update an active workflow scheme with `updateDraftIfNeeded` set to `true`: If a draft workflow scheme exists, it is updated. Otherwise, a draft workflow scheme is created.  *  Update an active workflow scheme with `updateDraftIfNeeded` set to `false`: An error is returned, as active workflow schemes cannot be updated.  *  Update an inactive workflow scheme with `updateDraftIfNeeded` set to `true`: The workflow scheme is updated, as inactive workflow schemes do not require drafts to update.  Defaults to `false`.
	UpdateDraftIfNeeded bool `json:"updateDraftIfNeeded,omitempty"`
	// The issue types available in Jira.
	IssueTypes map[string]IssueTypeDetails `json:"issueTypes,omitempty"`
}
