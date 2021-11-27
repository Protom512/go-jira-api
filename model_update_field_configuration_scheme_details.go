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

// The details of the field configuration scheme.
type UpdateFieldConfigurationSchemeDetails struct {
	// The name of the field configuration scheme. The name must be unique.
	Name string `json:"name"`
	// The description of the field configuration scheme.
	Description string `json:"description,omitempty"`
}
