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

// The field configuration for an issue type.
type FieldConfigurationIssueTypeItem struct {
	// The ID of the field configuration scheme.
	FieldConfigurationSchemeId string `json:"fieldConfigurationSchemeId"`
	// The ID of the issue type or *default*. When set to *default* this field configuration issue type item applies to all issue types without a field configuration.
	IssueTypeId string `json:"issueTypeId"`
	// The ID of the field configuration.
	FieldConfigurationId string `json:"fieldConfigurationId"`
}
