package persistence

import (
	"context"

	"github.com/YukiOnishi1129/curation-app-rest/backend/internal/domain"
)

type platformRepository struct {
	queries Queries
}

func NewPlatformRepository(queries Queries) domain.PlatformRepository {
	return &platformRepository{
		queries: queries,
	}
}

func (r *platformRepository) CreatePlatform(ctx context.Context, arg domain.CreatePlatformsParams) (*domain.Platform, error) {
	result, err := r.queries.CreatePlatforms(ctx, CreatePlatformsParams{
		Name:             arg.Name,
		SiteUrl:          arg.SiteURL,
		PlatformSiteType: arg.PlatformSiteType,
		FaviconUrl:       arg.FaviconURL,
	})
	if err != nil {
		return nil, err
	}
	output := &domain.Platform{
		ID:               result.ID.String(),
		Name:             result.Name,
		SiteURL:          result.SiteUrl,
		PlatformSiteType: result.PlatformSiteType,
		FaviconURL:       result.FaviconUrl,
		CreatedAt:        result.CreatedAt.Time,
		UpdatedAt:        result.UpdatedAt.Time,
	}
	if result.DeletedAt.Valid {
		output.DeletedAt = &result.DeletedAt.Time
	}
	return output, nil
}
