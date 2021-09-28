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

type ProjectLandingPageInfo struct {
	Url string `json:"url,omitempty"`
	ProjectKey string `json:"projectKey,omitempty"`
	ProjectType string `json:"projectType,omitempty"`
	BoardId int64 `json:"boardId,omitempty"`
	Simplified bool `json:"simplified,omitempty"`
}