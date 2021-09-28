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

// A screen scheme.
type ScreenScheme struct {
	// The ID of the screen scheme.
	Id int64 `json:"id,omitempty"`
	// The name of the screen scheme.
	Name string `json:"name,omitempty"`
	// The description of the screen scheme.
	Description string `json:"description,omitempty"`
	// The IDs of the screens for the screen types of the screen scheme.
	Screens *AllOfScreenSchemeScreens `json:"screens,omitempty"`
	// Details of the issue type screen schemes associated with the screen scheme.
	IssueTypeScreenSchemes *AllOfScreenSchemeIssueTypeScreenSchemes `json:"issueTypeScreenSchemes,omitempty"`
}
