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

// Lists of JQL reference data.
type JqlReferenceData struct {
	// List of fields usable in JQL queries.
	VisibleFieldNames []FieldReferenceData `json:"visibleFieldNames,omitempty"`
	// List of functions usable in JQL queries.
	VisibleFunctionNames []FunctionReferenceData `json:"visibleFunctionNames,omitempty"`
	// List of JQL query reserved words.
	JqlReservedWords []string `json:"jqlReservedWords,omitempty"`
}