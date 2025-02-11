package usecase

import "context"

type BatchMigrateSeedInterface interface {
	BatchMigrateSeed(ctx context.Context) error
}

func (u *Usecase) BatchMigrateSeed(ctx context.Context) error {
	// seedPlatformAndFeed := u.getSeedPlatformAndFeed()
	// for _, seed := range seedPlatformAndFeed {
	// 	platformID, err := u.platformRepository.CreatePlatform(ctx, seed.PlatformName, seed.PlatformSiteURL, seed.PlatformType, seed.FaviconURL)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	_, err = u.feedRepository.CreateFeed(ctx, seed.FeedName, seed.FeedDescription, seed.FeedThumbnail, seed.FeedSiteURL, seed.RssURL, platformID)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
