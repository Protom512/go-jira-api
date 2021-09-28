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

// List of permission grants.
type PermissionGrants struct {
	// Permission grants list.
	Permissions []PermissionGrant `json:"permissions,omitempty"`
	// Expand options that include additional permission grant details in the response.
	Expand string `json:"expand,omitempty"`
}
