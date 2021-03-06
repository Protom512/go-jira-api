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

// The results from a JQL query.
type AutoCompleteSuggestions struct {
	// The list of suggested item.
	Results []AutoCompleteSuggestion `json:"results,omitempty"`
}
