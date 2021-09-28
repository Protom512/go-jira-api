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

// Details of a changed worklog.
type ChangedWorklog struct {
	// The ID of the worklog.
	WorklogId int64 `json:"worklogId,omitempty"`
	// The datetime of the change.
	UpdatedTime int64 `json:"updatedTime,omitempty"`
	// Details of properties associated with the change.
	Properties []EntityProperty `json:"properties,omitempty"`
}
