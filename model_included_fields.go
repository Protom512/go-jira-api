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

type IncludedFields struct {
	Included []string `json:"included,omitempty"`
	Excluded []string `json:"excluded,omitempty"`
	ActuallyIncluded []string `json:"actuallyIncluded,omitempty"`
}
