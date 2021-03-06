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

// Details about a licensed Jira application.
type LicensedApplication struct {
	// The ID of the application.
	Id string `json:"id"`
	// The licensing plan.
	Plan string `json:"plan"`
}
