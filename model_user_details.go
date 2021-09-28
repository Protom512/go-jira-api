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

// User details permitted by the user's Atlassian Account privacy settings. However, be aware of these exceptions:   *  User record deleted from Atlassian: This occurs as the result of a right to be forgotten request. In this case, `displayName` provides an indication and other parameters have default values or are blank (for example, email is blank).  *  User record corrupted: This occurs as a results of events such as a server import and can only happen to deleted users. In this case, `accountId` returns *unknown* and all other parameters have fallback values.  *  User record unavailable: This usually occurs due to an internal service outage. In this case, all parameters have fallback values.
type UserDetails struct {
	// The URL of the user.
	Self string `json:"self,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Name string `json:"name,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Key string `json:"key,omitempty"`
	// The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	AccountId string `json:"accountId,omitempty"`
	// The email address of the user. Depending on the user’s privacy settings, this may be returned as null.
	EmailAddress string `json:"emailAddress,omitempty"`
	// The avatars of the user.
	AvatarUrls *AllOfUserDetailsAvatarUrls `json:"avatarUrls,omitempty"`
	// The display name of the user. Depending on the user’s privacy settings, this may return an alternative value.
	DisplayName string `json:"displayName,omitempty"`
	// Whether the user is active.
	Active bool `json:"active,omitempty"`
	// The time zone specified in the user's profile. Depending on the user’s privacy settings, this may be returned as null.
	TimeZone string `json:"timeZone,omitempty"`
	// The type of account represented by this user. This will be one of 'atlassian' (normal users), 'app' (application user) or 'customer' (Jira Service Desk customer user)
	AccountType string `json:"accountType,omitempty"`
}
