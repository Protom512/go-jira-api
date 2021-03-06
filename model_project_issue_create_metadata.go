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

// Details of the issue creation metadata for a project.
type ProjectIssueCreateMetadata struct {
	// Expand options that include additional project issue create metadata details in the response.
	Expand string `json:"expand,omitempty"`
	// The URL of the project.
	Self string `json:"self,omitempty"`
	// The ID of the project.
	Id string `json:"id,omitempty"`
	// The key of the project.
	Key string `json:"key,omitempty"`
	// The name of the project.
	Name string `json:"name,omitempty"`
	// List of the project's avatars, returning the avatar size and associated URL.
	AvatarUrls *AllOfProjectIssueCreateMetadataAvatarUrls `json:"avatarUrls,omitempty"`
	// List of the issue types supported by the project.
	Issuetypes []IssueTypeIssueCreateMetadata `json:"issuetypes,omitempty"`
}
