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

// The scope of the notification scheme.
type AllOfNotificationSchemeScope struct {
	// The type of scope.
	Type_ string `json:"type,omitempty"`
	// The project the item has scope in.
	Project *Object `json:"project,omitempty"`
}
