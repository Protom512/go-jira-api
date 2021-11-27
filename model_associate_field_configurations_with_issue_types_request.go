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

// Details of a field configuration to issue type mappings.
type AssociateFieldConfigurationsWithIssueTypesRequest struct {
	// Field configuration to issue type mappings.
	Mappings []FieldConfigurationToIssueTypeMapping `json:"mappings"`
}
