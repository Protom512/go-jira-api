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

// Count of a version's unresolved issues.
type VersionUnresolvedIssuesCount struct {
	// The URL of these count details.
	Self string `json:"self,omitempty"`
	// Count of unresolved issues.
	IssuesUnresolvedCount int64 `json:"issuesUnresolvedCount,omitempty"`
	// Count of issues.
	IssuesCount int64 `json:"issuesCount,omitempty"`
}