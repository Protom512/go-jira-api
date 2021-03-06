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

type JiraExpressionEvaluationMetaDataBean struct {
	// Contains information about the expression complexity. For example, the number of steps it took to evaluate the expression.
	Complexity *AllOfJiraExpressionEvaluationMetaDataBeanComplexity `json:"complexity,omitempty"`
	// Contains information about the `issues` variable in the context. For example, is the issues were loaded with JQL, information about the page will be included here.
	Issues *AllOfJiraExpressionEvaluationMetaDataBeanIssues `json:"issues,omitempty"`
}
