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

// The default value for a single select custom field.
type CustomFieldContextDefaultValueSingleOption struct {
	// The ID of the context.
	ContextId string `json:"contextId"`
	// The ID of the default option.
	OptionId string `json:"optionId"`
	Type_ string `json:"type"`
}
