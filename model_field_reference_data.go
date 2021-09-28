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

// Details of a field that can be used in advanced searches.
type FieldReferenceData struct {
	// The field identifier.
	Value string `json:"value,omitempty"`
	// The display name contains the following:   *  for system fields, the field name. For example, `Summary`.  *  for collapsed custom fields, the field name followed by a hyphen and then the field name and field type. For example, `Component - Component[Dropdown]`.  *  for other custom fields, the field name followed by a hyphen and then the custom field ID. For example, `Component - cf[10061]`.
	DisplayName string `json:"displayName,omitempty"`
	// Whether the field can be used in a query's `ORDER BY` clause.
	Orderable string `json:"orderable,omitempty"`
	// Whether the content of this field can be searched.
	Searchable string `json:"searchable,omitempty"`
	// Whether the field provide auto-complete suggestions.
	Auto string `json:"auto,omitempty"`
	// If the item is a custom field, the ID of the custom field.
	Cfid string `json:"cfid,omitempty"`
	// The valid search operators for the field.
	Operators []string `json:"operators,omitempty"`
	// The data types of items in the field.
	Types []string `json:"types,omitempty"`
}
