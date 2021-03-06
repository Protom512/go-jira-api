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

// A group label.
type GroupLabel struct {
	// The group label name.
	Text string `json:"text,omitempty"`
	// The title of the group label.
	Title string `json:"title,omitempty"`
	// The type of the group label.
	Type_ string `json:"type,omitempty"`
}
