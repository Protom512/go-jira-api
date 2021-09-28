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

// The avatars of the user.
type AllOfUserDetailsAvatarUrls struct {
	// The URL of the item's 16x16 pixel avatar.
	Var16x16 string `json:"16x16,omitempty"`
	// The URL of the item's 24x24 pixel avatar.
	Var24x24 string `json:"24x24,omitempty"`
	// The URL of the item's 32x32 pixel avatar.
	Var32x32 string `json:"32x32,omitempty"`
	// The URL of the item's 48x48 pixel avatar.
	Var48x48 string `json:"48x48,omitempty"`
}
