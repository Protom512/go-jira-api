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
import (
	"time"
)

// An audit record.
type AuditRecordBean struct {
	// The ID of the audit record.
	Id int64 `json:"id,omitempty"`
	// The summary of the audit record.
	Summary string `json:"summary,omitempty"`
	// The URL of the computer where the creation of the audit record was initiated.
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Deprecated, use `authorAccountId` instead. The key of the user who created the audit record.
	AuthorKey string `json:"authorKey,omitempty"`
	// The date and time on which the audit record was created.
	Created time.Time `json:"created,omitempty"`
	// The category of the audit record. For a list of these categories, see the help article [Auditing in Jira applications](https://confluence.atlassian.com/x/noXKM).
	Category string `json:"category,omitempty"`
	// The event the audit record originated from.
	EventSource string `json:"eventSource,omitempty"`
	// The description of the audit record.
	Description string `json:"description,omitempty"`
	ObjectItem *AssociatedItemBean `json:"objectItem,omitempty"`
	// The list of values changed in the record event.
	ChangedValues []ChangedValueBean `json:"changedValues,omitempty"`
	// The list of items associated with the changed record.
	AssociatedItems []AssociatedItemBean `json:"associatedItems,omitempty"`
}
