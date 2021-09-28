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

type Group struct {
	// The name of group.
	Name string `json:"name,omitempty"`
	// The URL for these group details.
	Self string `json:"self,omitempty"`
	// A paginated list of the users that are members of the group. A maximum of 50 users is returned in the list, to access additional users append `[start-index:end-index]` to the expand request. For example, to access the next 50 users, use`?expand=users[51:100]`.
	Users *AllOfGroupUsers `json:"users,omitempty"`
	// Expand options that include additional group details in the response.
	Expand string `json:"expand,omitempty"`
}
