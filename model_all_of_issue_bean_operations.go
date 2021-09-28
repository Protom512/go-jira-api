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

// The operations that can be performed on the issue.
type AllOfIssueBeanOperations struct {
	// Details of the link groups defining issue operations.
	LinkGroups []LinkGroup `json:"linkGroups,omitempty"`
}