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

// Details about a filter.
type Filter struct {
	// The URL of the filter.
	Self string `json:"self,omitempty"`
	// The unique identifier for the filter.
	Id string `json:"id,omitempty"`
	// The name of the filter. Must be unique.
	Name string `json:"name"`
	// A description of the filter.
	Description string `json:"description,omitempty"`
	// The user who owns the filter. This is defaulted to the creator of the filter, however Jira administrators can change the owner of a shared filter in the admin settings.
	Owner *AllOfFilterOwner `json:"owner,omitempty"`
	// The JQL query for the filter. For example, *project = SSP AND issuetype = Bug*.
	Jql string `json:"jql,omitempty"`
	// A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter=10100*.
	ViewUrl string `json:"viewUrl,omitempty"`
	// A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter's JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql=project+%3D+SSP+AND+issuetype+%3D+Bug*.
	SearchUrl string `json:"searchUrl,omitempty"`
	// Whether the filter is selected as a favorite.
	Favourite bool `json:"favourite,omitempty"`
	// The count of how many users have selected this filter as a favorite, including the filter owner.
	FavouritedCount int64 `json:"favouritedCount,omitempty"`
	// The groups and projects that the filter is shared with.
	SharePermissions []SharePermission `json:"sharePermissions,omitempty"`
	// A paginated list of the users that the filter is shared with. This includes users that are members of the groups or can browse the projects that the filter is shared with.
	SharedUsers *AllOfFilterSharedUsers `json:"sharedUsers,omitempty"`
	// A paginated list of the users that are subscribed to the filter.
	Subscriptions *AllOfFilterSubscriptions `json:"subscriptions,omitempty"`
}
