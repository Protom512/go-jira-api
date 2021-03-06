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

type EntityPropertyDetails struct {
	// The entity property ID.
	EntityId float64 `json:"entityId"`
	// The entity property key.
	Key string `json:"key"`
	// The new value of the entity property.
	Value string `json:"value"`
}
