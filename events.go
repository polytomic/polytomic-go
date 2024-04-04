// This file was auto-generated from our API Definition.

package polytomic

import (
	time "time"
)

type EventsListRequest struct {
	OrganizationId *string    `json:"-" url:"organization_id,omitempty"`
	Type           *string    `json:"-" url:"type,omitempty"`
	StartingAfter  *time.Time `json:"-" url:"starting_after,omitempty"`
	EndingBefore   *time.Time `json:"-" url:"ending_before,omitempty"`
	Limit          *int       `json:"-" url:"limit,omitempty"`
}
