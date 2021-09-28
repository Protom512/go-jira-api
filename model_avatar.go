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

// Details of an avatar.
type Avatar struct {
	// The ID of the avatar.
	Id string `json:"id"`
	// The owner of the avatar. For a system avatar the owner is null (and nothing is returned). For non-system avatars this is the appropriate identifier, such as the ID for a project or the account ID for a user.
	Owner string `json:"owner,omitempty"`
	// Whether the avatar is a system avatar.
	IsSystemAvatar bool `json:"isSystemAvatar,omitempty"`
	// Whether the avatar is used in Jira. For example, shown as a project's avatar.
	IsSelected bool `json:"isSelected,omitempty"`
	// Whether the avatar can be deleted.
	IsDeletable bool `json:"isDeletable,omitempty"`
	// The file name of the avatar icon. Returned for system avatars.
	FileName string `json:"fileName,omitempty"`
	// The list of avatar icon URLs.
	Urls map[string]string `json:"urls,omitempty"`
}
