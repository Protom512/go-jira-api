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

// Container for a list of registered webhooks. Webhook details are returned in the same order as the request.
type ContainerForRegisteredWebhooks struct {
	// A list of registered webhooks.
	WebhookRegistrationResult []RegisteredWebhook `json:"webhookRegistrationResult,omitempty"`
}