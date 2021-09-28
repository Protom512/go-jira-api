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

// An associated workflow scheme and project.
type WorkflowSchemeProjectAssociation struct {
	// The ID of the workflow scheme. If the workflow scheme ID is `null`, the operation assigns the default workflow scheme.
	WorkflowSchemeId string `json:"workflowSchemeId,omitempty"`
	// The ID of the project.
	ProjectId string `json:"projectId"`
}
