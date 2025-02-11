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
			FeedThumbnail:   "https://static.zenn.studio/images/logo-only-dark.png",
			PlatformSiteURL: "https://zenn.dev",
			FeedSiteURL:     "https://zenn.dev",
			PlatformType:    domain.PlatformTypeSite,
			FaviconURL:      "https://zenn.dev/favicon.ico",
			RssURL:          "https://zenn.dev/feed",
		},
		{
			PlatformName:    "Qiita",
			FeedName:        "Qiitaのトレンド",
			FeedDescription: `Qiitaで人気の記事一覧です。毎日5時と17時に更新しています。`,
			FeedThumbnail:   "https://cdn.qiita.com/assets/qiita-ogp-3b6fcfdd74755a85107071ffc3155898.png",
			PlatformSiteURL: "https://qiita.com",
			FeedSiteURL:     "https://qiita.com",
			PlatformType:    domain.PlatformTypeSite,
			FaviconURL:      "https://cdn.qiita.com/assets/favicons/public/production-c620d3e403342b1022967ba5e3db1aaa.ico",
			RssURL:          "https://qiita.com/popular-items/feed",
		},
		{
			PlatformName:    "Zenn",
			FeedName:        "Reactの記事一覧 | Zenn",
			FeedDescription: `Zenn は、エンジニアのための情報共有コミュニティです。知識を記録し、共有し、より多くの人に届けましょう。`,
			FeedThumbnail:   "https://static.zenn.studio/images/logo-only-dark.png",
			PlatformSiteURL: "https://zenn.dev",
			FeedSiteURL:     "https://zenn.dev/topics/react",
			PlatformType:    domain.PlatformTypeSite,
			FaviconURL:      "https://zenn.dev/favicon.ico",
			RssURL:          "https://zenn.dev/topics/react/feed",
		},
		{
			PlatformName:    "メルカリ",
			FeedName:        "メルカリエンジニアリングブログ",
			FeedDescription: `「Mercari Engineering」は メルカリのエンジニアに関する情報を、 オープンに公開・共有していくためのサイトです。`,
			FeedThumbnail:   "https://engineering.mercari.com//img/ogp/ogp_a.jpg",
			PlatformSiteURL: "https://about.mercari.com/",
			FeedSiteURL:     "https://engineering.mercari.com/blog/",
			PlatformType:    domain.PlatformTypeCompany,
			FaviconURL:      "https://engineering.mercari.com/favicon.ico",
			RssURL:          "https://engineering.mercari.com/blog/feed.xml",
		},
		{
			PlatformName:    "はてなブックマーク",
			FeedName:        "はてなブックマーク - 人気エントリー - テクノロジー",
			FeedDescription: `はてなブックマーク - 人気エントリー - テクノロジー`,
			FeedThumbnail:   "https://b.st-hatena.com/0c3a38c41aeb08c713c990efb1b369be703ea86c/images/v4/public/apple-touch-icon-precomposed.png",
			PlatformSiteURL: "https://b.hatena.ne.jp/",
			FeedSiteURL:     "https://b.hatena.ne.jp/hotentry/it",
			PlatformType:    domain.PlatformTypeSummary,
			FaviconURL:      "https://b.hatena.ne.jp//favicon.ico",
			RssURL:          "https://b.hatena.ne.jp/hotentry/it.rss",
		},
	}
}
