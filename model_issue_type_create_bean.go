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

type IssueTypeCreateBean struct {
	// The unique name for the issue type. The maximum length is 60 characters.
	Name string `json:"name"`
	// The description of the issue type.
	Description string `json:"description,omitempty"`
	// Deprecated. Use `hierarchyLevel` instead.  Whether the issue type is `subtype` or `standard`. Defaults to `standard`.
	Type_ string `json:"type,omitempty"`
	// The hierarchy level of the issue type. Use:   *  `-1` for Subtask.  *  `0` for Base.  Defaults to `0`.
	HierarchyLevel int32 `json:"hierarchyLevel,omitempty"`
}
