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

// Details about an issue type.
type IssueTypeDetails struct {
	// The URL of these issue type details.
	Self string `json:"self,omitempty"`
	// The ID of the issue type.
	Id string `json:"id,omitempty"`
	// The description of the issue type.
	Description string `json:"description,omitempty"`
	// The URL of the issue type's avatar.
	IconUrl string `json:"iconUrl,omitempty"`
	// The name of the issue type.
	Name string `json:"name,omitempty"`
	// Whether this issue type is used to create subtasks.
	Subtask bool `json:"subtask,omitempty"`
	// The ID of the issue type's avatar.
	AvatarId int64 `json:"avatarId,omitempty"`
	// Unique ID for next-gen projects.
	EntityId string `json:"entityId,omitempty"`
	// Hierarchy level of the issue type.
	HierarchyLevel int32 `json:"hierarchyLevel,omitempty"`
	// Details of the next-gen projects the issue type is available in.
	Scope *AllOfIssueTypeDetailsScope `json:"scope,omitempty"`
}
