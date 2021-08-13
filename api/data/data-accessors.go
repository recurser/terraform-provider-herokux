// Copyright 2020
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by gen-accessors; DO NOT EDIT.
package data

import (
	"time"
)

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetCreatedAt() time.Time {
	if p == nil || p.CreatedAt == nil {
		return time.Time{}
	}
	return *p.CreatedAt
}

// GetDatasource returns the Datasource field.
func (p *PostgresDataClip) GetDatasource() *PostgresDataClipDatasource {
	if p == nil {
		return nil
	}
	return p.Datasource
}

// GetDetached returns the Detached field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetDetached() bool {
	if p == nil || p.Detached == nil {
		return false
	}
	return *p.Detached
}

// GetEditable returns the Editable field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetEditable() bool {
	if p == nil || p.Editable == nil {
		return false
	}
	return *p.Editable
}

// GetEditedAt returns the EditedAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetEditedAt() time.Time {
	if p == nil || p.EditedAt == nil {
		return time.Time{}
	}
	return *p.EditedAt
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetID() string {
	if p == nil || p.ID == nil {
		return ""
	}
	return *p.ID
}

// GetPublicSlug returns the PublicSlug field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetPublicSlug() string {
	if p == nil || p.PublicSlug == nil {
		return ""
	}
	return *p.PublicSlug
}

// GetPublicSlugBy returns the PublicSlugBy field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetPublicSlugBy() string {
	if p == nil || p.PublicSlugBy == nil {
		return ""
	}
	return *p.PublicSlugBy
}

// GetSlug returns the Slug field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetSlug() string {
	if p == nil || p.Slug == nil {
		return ""
	}
	return *p.Slug
}

// GetTeamID returns the TeamID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetTeamID() string {
	if p == nil || p.TeamID == nil {
		return ""
	}
	return *p.TeamID
}

// HasTeamShares checks if PostgresDataClip has any TeamShares.
func (p *PostgresDataClip) HasTeamShares() bool {
	if p == nil || p.TeamShares == nil {
		return false
	}
	if len(p.TeamShares) == 0 {
		return false
	}
	return true
}

// GetTitle returns the Title field if it's non-nil, zero value otherwise.
func (p *PostgresDataClip) GetTitle() string {
	if p == nil || p.Title == nil {
		return ""
	}
	return *p.Title
}

// HasUserShares checks if PostgresDataClip has any UserShares.
func (p *PostgresDataClip) HasUserShares() bool {
	if p == nil || p.UserShares == nil {
		return false
	}
	if len(p.UserShares) == 0 {
		return false
	}
	return true
}

// HasVersions checks if PostgresDataClip has any Versions.
func (p *PostgresDataClip) HasVersions() bool {
	if p == nil || p.Versions == nil {
		return false
	}
	if len(p.Versions) == 0 {
		return false
	}
	return true
}

// GetAddonID returns the AddonID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAddonID() string {
	if p == nil || p.AddonID == nil {
		return ""
	}
	return *p.AddonID
}

// GetAddonName returns the AddonName field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAddonName() string {
	if p == nil || p.AddonName == nil {
		return ""
	}
	return *p.AddonName
}

// GetAppID returns the AppID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAppID() string {
	if p == nil || p.AppID == nil {
		return ""
	}
	return *p.AppID
}

// GetAppName returns the AppName field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAppName() string {
	if p == nil || p.AppName == nil {
		return ""
	}
	return *p.AppName
}

// GetAttachmentID returns the AttachmentID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAttachmentID() string {
	if p == nil || p.AttachmentID == nil {
		return ""
	}
	return *p.AttachmentID
}

// GetAttachmentName returns the AttachmentName field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetAttachmentName() string {
	if p == nil || p.AttachmentName == nil {
		return ""
	}
	return *p.AttachmentName
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipDatasource) GetID() string {
	if p == nil || p.ID == nil {
		return ""
	}
	return *p.ID
}

// GetCreatedAt returns the CreatedAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetCreatedAt() time.Time {
	if p == nil || p.CreatedAt == nil {
		return time.Time{}
	}
	return *p.CreatedAt
}

// GetCreatorID returns the CreatorID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetCreatorID() string {
	if p == nil || p.CreatorID == nil {
		return ""
	}
	return *p.CreatorID
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetID() string {
	if p == nil || p.ID == nil {
		return ""
	}
	return *p.ID
}

// GetLatestResultSize returns the LatestResultSize field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetLatestResultSize() int {
	if p == nil || p.LatestResultSize == nil {
		return 0
	}
	return *p.LatestResultSize
}

// GetResult returns the Result field.
func (p *PostgresDataClipVersion) GetResult() *PostgresDataClipVersionResult {
	if p == nil {
		return nil
	}
	return p.Result
}

// GetSql returns the Sql field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetSql() string {
	if p == nil || p.Sql == nil {
		return ""
	}
	return *p.Sql
}

// GetURL returns the URL field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersion) GetURL() string {
	if p == nil || p.URL == nil {
		return ""
	}
	return *p.URL
}

// GetCompletedAt returns the CompletedAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetCompletedAt() time.Time {
	if p == nil || p.CompletedAt == nil {
		return time.Time{}
	}
	return *p.CompletedAt
}

// GetDuration returns the Duration field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetDuration() int {
	if p == nil || p.Duration == nil {
		return 0
	}
	return *p.Duration
}

// GetError returns the Error field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetError() string {
	if p == nil || p.Error == nil {
		return ""
	}
	return *p.Error
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetID() string {
	if p == nil || p.ID == nil {
		return ""
	}
	return *p.ID
}

// GetQueryFinishAt returns the QueryFinishAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetQueryFinishAt() time.Time {
	if p == nil || p.QueryFinishAt == nil {
		return time.Time{}
	}
	return *p.QueryFinishAt
}

// GetQueryStartAt returns the QueryStartAt field if it's non-nil, zero value otherwise.
func (p *PostgresDataClipVersionResult) GetQueryStartAt() time.Time {
	if p == nil || p.QueryStartAt == nil {
		return time.Time{}
	}
	return *p.QueryStartAt
}

// GetAddonName returns the AddonName field if it's non-nil, zero value otherwise.
func (p *Privatelink) GetAddonName() string {
	if p == nil || p.AddonName == nil {
		return ""
	}
	return *p.AddonName
}

// GetAddonUUID returns the AddonUUID field if it's non-nil, zero value otherwise.
func (p *Privatelink) GetAddonUUID() string {
	if p == nil || p.AddonUUID == nil {
		return ""
	}
	return *p.AddonUUID
}

// HasAllowedAccounts checks if Privatelink has any AllowedAccounts.
func (p *Privatelink) HasAllowedAccounts() bool {
	if p == nil || p.AllowedAccounts == nil {
		return false
	}
	if len(p.AllowedAccounts) == 0 {
		return false
	}
	return true
}

// GetAppName returns the AppName field if it's non-nil, zero value otherwise.
func (p *Privatelink) GetAppName() string {
	if p == nil || p.AppName == nil {
		return ""
	}
	return *p.AppName
}

// HasConnections checks if Privatelink has any Connections.
func (p *Privatelink) HasConnections() bool {
	if p == nil || p.Connections == nil {
		return false
	}
	if len(p.Connections) == 0 {
		return false
	}
	return true
}

// GetServiceName returns the ServiceName field if it's non-nil, zero value otherwise.
func (p *Privatelink) GetServiceName() string {
	if p == nil || p.ServiceName == nil {
		return ""
	}
	return *p.ServiceName
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (p *Privatelink) GetStatus() string {
	if p == nil || p.Status == nil {
		return ""
	}
	return *p.Status
}

// GetARN returns the ARN field if it's non-nil, zero value otherwise.
func (p *PrivatelinkAllowedAccounts) GetARN() string {
	if p == nil || p.ARN == nil {
		return ""
	}
	return *p.ARN
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (p *PrivatelinkAllowedAccounts) GetStatus() string {
	if p == nil || p.Status == nil {
		return ""
	}
	return *p.Status
}

// GetEndpointID returns the EndpointID field if it's non-nil, zero value otherwise.
func (p *PrivatelinkConnections) GetEndpointID() string {
	if p == nil || p.EndpointID == nil {
		return ""
	}
	return *p.EndpointID
}

// GetHostname returns the Hostname field if it's non-nil, zero value otherwise.
func (p *PrivatelinkConnections) GetHostname() string {
	if p == nil || p.Hostname == nil {
		return ""
	}
	return *p.Hostname
}

// GetOwnerARN returns the OwnerARN field if it's non-nil, zero value otherwise.
func (p *PrivatelinkConnections) GetOwnerARN() string {
	if p == nil || p.OwnerARN == nil {
		return ""
	}
	return *p.OwnerARN
}

// GetStatus returns the Status field if it's non-nil, zero value otherwise.
func (p *PrivatelinkConnections) GetStatus() string {
	if p == nil || p.Status == nil {
		return ""
	}
	return *p.Status
}
