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

// Details a link group, which defines issue operations.
type LinkGroup struct {
	Id string `json:"id,omitempty"`
	StyleClass string `json:"styleClass,omitempty"`
	Header *SimpleLink `json:"header,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	Links []SimpleLink `json:"links,omitempty"`
	Groups []LinkGroup `json:"groups,omitempty"`
}
