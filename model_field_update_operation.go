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

// Details of an operation to perform on a field.
type FieldUpdateOperation struct {
	// The value to add to the field.
	Add *Object `json:"add,omitempty"`
	// The value to set in the field.
	Set *Object `json:"set,omitempty"`
	// The value to removed from the field.
	Remove *Object `json:"remove,omitempty"`
	// The value to edit in the field.
	Edit *Object `json:"edit,omitempty"`
}
