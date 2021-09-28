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

// Details of functions that can be used in advanced searches.
type FunctionReferenceData struct {
	// The function identifier.
	Value string `json:"value,omitempty"`
	// The display name of the function.
	DisplayName string `json:"displayName,omitempty"`
	// Whether the function can take a list of arguments.
	IsList string `json:"isList,omitempty"`
	// The data types returned by the function.
	Types []string `json:"types,omitempty"`
}
