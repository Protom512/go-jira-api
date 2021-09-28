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

// Details of an entity property.
type JqlQueryFieldEntityProperty struct {
	// The object on which the property is set.
	Entity string `json:"entity"`
	// The key of the property.
	Key string `json:"key"`
	// The path in the property value to query.
	Path string `json:"path"`
	// The type of the property value extraction. Not available if the extraction for the property is not registered on the instance with the [Entity property](https://developer.atlassian.com/cloud/jira/platform/modules/entity-property/) module.
	Type_ string `json:"type,omitempty"`
}
