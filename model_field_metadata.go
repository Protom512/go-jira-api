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

// The metadata describing an issue field.
type FieldMetadata struct {
	// Whether the field is required.
	Required bool `json:"required"`
	// The data type of the field.
	Schema *AllOfFieldMetadataSchema `json:"schema"`
	// The name of the field.
	Name string `json:"name"`
	// The key of the field.
	Key string `json:"key"`
	// The URL that can be used to automatically complete the field.
	AutoCompleteUrl string `json:"autoCompleteUrl,omitempty"`
	// Whether the field has a default value.
	HasDefaultValue bool `json:"hasDefaultValue,omitempty"`
	// The list of operations that can be performed on the field.
	Operations []string `json:"operations"`
	// The list of values allowed in the field.
	AllowedValues []Object `json:"allowedValues,omitempty"`
	// The default value of the field.
	DefaultValue *Object `json:"defaultValue,omitempty"`
}
