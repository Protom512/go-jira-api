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

// Details of the time tracking configuration.
type TimeTrackingConfiguration struct {
	// The number of hours in a working day.
	WorkingHoursPerDay float64 `json:"workingHoursPerDay"`
	// The number of days in a working week.
	WorkingDaysPerWeek float64 `json:"workingDaysPerWeek"`
	// The format that will appear on an issue's *Time Spent* field.
	TimeFormat string `json:"timeFormat"`
	// The default unit of time applied to logged time.
	DefaultUnit string `json:"defaultUnit"`
}
