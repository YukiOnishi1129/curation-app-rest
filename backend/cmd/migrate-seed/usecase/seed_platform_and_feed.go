package usecase

import (
	"time"

	"github.com/YukiOnishi1129/curation-app-rest/backend/internal/domain"
)

type SeedPlatformAndFeed struct {
	PlatformName    string
	FeedName        string
	FeedDescription string
	FeedThumbnail   string
	PlatformSiteURL string
	FeedSiteURL     string
	PlatformType    domain.PlatformSiteType
	FaviconURL      string
	RssURL          string
	DeletedAt       *time.Time
}

func (u *Usecase) getSeedPlatformAndFeed() []SeedPlatformAndFeed {
	return []SeedPlatformAndFeed{
		{
			PlatformName:    "Zenn",
			FeedName:        "Zenn",
			FeedDescription: `Zenn は、エンジニアのための情報共有コミュニティです。知識を記録し、共有し、より多くの人に届けましょう。`,
			FeedThumbnail:   "https://zenn.dev/images/og/social.png",
			PlatformSiteURL: "https://zenn.dev",
			FeedSiteURL:     "https://zenn.dev",
			PlatformType:    domain.PlatformTypeSite,
			FaviconURL:      "https://zenn.dev/favicon.ico",
			RssURL:          "https://zenn.dev/feed",
		},
		{
			PlatformName:    "Qiita",
			FeedName:        "Qiita",
			FeedDescription: `Qiita は、プログラマのための技術情報共有サービスです。 プログラミングに関するTips、ノウハウ、メモを簡単に記録 & 公開することができます。`,
			FeedThumbnail:   "https://qiita.com/icons/apple-touch-icon-152x152.png",
			PlatformSiteURL: "https://qiita.com",
			FeedSiteURL:     "https://qiita.com",
			PlatformType:    domain.PlatformTypeSite,
			FaviconURL:      "https://qiita.com/icons/favicon-32x32.png",
			RssURL:          "https://qiita.com/tags/rss",
		},
	}
}
