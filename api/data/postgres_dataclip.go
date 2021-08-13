package data

import (
	"github.com/davidji99/terraform-provider-herokux/api/platform"
	"time"
)

// PostgresDataClip represents a postgres data clip.
type PostgresDataClip struct {
	ID           *string                     `json:"id,omitempty"`
	Slug         *string                     `json:"slug,omitempty"`
	Title        *string                     `json:"title,omitempty"`
	TeamID       *string                     `json:"team_id,omitempty"`
	PublicSlug   *string                     `json:"public_slug,omitempty"`
	PublicSlugBy *string                     `json:"public_slug_by,omitempty"`
	UserShares   []string                    `json:"user_shares,omitempty"`
	TeamShares   []string                    `json:"team_shares,omitempty"`
	Detached     *bool                       `json:"detached,omitempty"`
	Editable     *bool                       `json:"editable,omitempty"`
	CreatedAt    *time.Time                  `json:"created_at,omitempty"`
	EditedAt     *time.Time                  `json:"edited_at,omitempty"`
	Creator      *platform.User              `json:"creator,omitempty"`
	Datasource   *PostgresDataClipDatasource `json:"datasource,omitempty"`
	Versions     []*PostgresDataClipVersion  `json:"versions,omitempty"`
}

// PostgresDataClipDatasource represents a data clip source.
type PostgresDataClipDatasource struct {
	ID             *string `json:"id,omitempty"`
	AddonID        *string `json:"addon_id,omitempty"`
	AddonName      *string `json:"addon_name,omitempty"`
	AttachmentID   *string `json:"attachment_id,omitempty"`
	AttachmentName *string `json:"attachment_name,omitempty"`
	AppID          *string `json:"app_id,omitempty"`
	AppName        *string `json:"app_name,omitempty"`
}

// PostgresDataClipVersion represents a data clip version.
type PostgresDataClipVersion struct {
	ID               *string                        `json:"id,omitempty"`
	Sql              *string                        `json:"sql,omitempty"`
	URL              *string                        `json:"url,omitempty"`
	CreatorID        *string                        `json:"creator_id,omitempty"`
	Creator          *platform.User                 `json:"creator,omitempty"`
	LatestResultSize *int                           `json:"latest_result_size,omitempty"`
	CreatedAt        *time.Time                     `json:"created_at,omitempty"`
	Result           *PostgresDataClipVersionResult `json:"result,omitempty"`
}

// PostgresDataClipVersionResult represents a data clip query result.
type PostgresDataClipVersionResult struct {
	ID            *string    `json:"id,omitempty"`
	Error         *string    `json:"error,omitempty"`
	QueryStartAt  *time.Time `json:"query_started_at,omitempty"`
	QueryFinishAt *time.Time `json:"query_finished_at,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	Duration      *int       `json:"duration,omitempty"`
}
