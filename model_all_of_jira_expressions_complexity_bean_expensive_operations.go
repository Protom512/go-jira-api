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

// The number of expensive operations executed while evaluating the expression. Expensive operations are those that load additional data, such as entity properties, comments, or custom fields.
type AllOfJiraExpressionsComplexityBeanExpensiveOperations struct {
	// The complexity value of the current expression.
	Value int32 `json:"value"`
	// The maximum allowed complexity. The evaluation will fail if this value is exceeded.
	Limit int32 `json:"limit"`
}
