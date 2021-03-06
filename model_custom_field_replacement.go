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

// Details about the replacement for a deleted version.
type CustomFieldReplacement struct {
	// The ID of the custom field in which to replace the version number.
	CustomFieldId int64 `json:"customFieldId,omitempty"`
	// The version number to use as a replacement for the deleted version.
	MoveTo int64 `json:"moveTo,omitempty"`
}
