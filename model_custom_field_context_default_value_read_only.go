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

// The default text for a read only custom field.
type CustomFieldContextDefaultValueReadOnly struct {
	// The default text. The maximum length is 255 characters.
	Text string `json:"text,omitempty"`
	Type_ string `json:"type"`
}
