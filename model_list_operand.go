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

// An operand that is a list of values.
type ListOperand struct {
	// The list of operand values.
	Values []JqlQueryUnitaryOperand `json:"values"`
}
