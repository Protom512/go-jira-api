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

type PaginatedResponseComment struct {
	Total int64 `json:"total,omitempty"`
	MaxResults int32 `json:"maxResults,omitempty"`
	StartAt int64 `json:"startAt,omitempty"`
	Results []ModelMap `json:"results,omitempty"`
}
