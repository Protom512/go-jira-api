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

type ActorInputBean struct {
	// The account IDs of the users to add as default actors. This parameter accepts a comma-separated list. For example, `\"user\":[\"5b10a2844c20165700ede21g\", \"5b109f2e9729b51b54dc274d\"]`.
	User []string `json:"user,omitempty"`
	// The name of the group to add as a default actor. This parameter accepts a comma-separated list. For example, `\"group\":[\"project-admin\", \"jira-developers\"]`.
	Group []string `json:"group,omitempty"`
}
