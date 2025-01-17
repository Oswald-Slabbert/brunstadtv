package notifications

import (
	"github.com/bcc-code/brunstadtv/backend/members"
	"github.com/bcc-code/brunstadtv/backend/sqlc"
)

// Utils contains different methods for resolving notifications
type Utils struct {
	members *members.Client
	queries *sqlc.Queries
}

// NewUtils returns a new Utils struct
func NewUtils(
	queries *sqlc.Queries,
	members *members.Client,
) *Utils {
	return &Utils{
		queries: queries,
		members: members,
	}
}
