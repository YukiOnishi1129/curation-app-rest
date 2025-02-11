package domain

import (
	"context"
	"time"
)

type Platform struct {
	ID               string
	Name             string
	SiteURL          string
	PlatformSiteType int32
	FaviconURL       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

type CreatePlatformsParams struct {
	Name             string
	SiteURL          string
	PlatformSiteType int32
	FaviconURL       string
}

type PlatformRepository interface {
	CreatePlatform(ctx context.Context, arg CreatePlatformsParams) (*Platform, error)
}
