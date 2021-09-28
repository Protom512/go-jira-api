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

type BulkOperationErrorResult struct {
	Status int32 `json:"status,omitempty"`
	ElementErrors *ErrorCollection `json:"elementErrors,omitempty"`
	FailedElementNumber int32 `json:"failedElementNumber,omitempty"`
}
