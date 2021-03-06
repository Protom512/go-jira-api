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

// The JQL queries to be converted.
type JqlPersonalDataMigrationRequest struct {
	// A list of queries with user identifiers. Maximum of 100 queries.
	QueryStrings []string `json:"queryStrings,omitempty"`
}
