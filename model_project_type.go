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

// Details about a project type.
type ProjectType struct {
	// The key of the project type.
	Key string `json:"key,omitempty"`
	// The formatted key of the project type.
	FormattedKey string `json:"formattedKey,omitempty"`
	// The key of the project type's description.
	DescriptionI18nKey string `json:"descriptionI18nKey,omitempty"`
	// The icon of the project type.
	Icon string `json:"icon,omitempty"`
	// The color of the project type.
	Color string `json:"color,omitempty"`
}
