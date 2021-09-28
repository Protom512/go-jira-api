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

// A page of items.
type PageBeanWebhook struct {
	// The URL of the page.
	Self string `json:"self,omitempty"`
	// If there is another page of results, the URL of the next page.
	NextPage string `json:"nextPage,omitempty"`
	// The maximum number of items that could be returned.
	MaxResults int32 `json:"maxResults,omitempty"`
	// The index of the first item returned.
	StartAt int64 `json:"startAt,omitempty"`
	// The number of items returned.
	Total int64 `json:"total,omitempty"`
	// Whether this is the last page.
	IsLast bool `json:"isLast,omitempty"`
	// The list of items.
	Values []Webhook `json:"values,omitempty"`
}
