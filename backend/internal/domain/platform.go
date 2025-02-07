package domain

type PlatformSiteType int

const (
	PlatformUnknown PlatformSiteType = iota
	PlatformTypeSite
	PlatformTypeCompany
	PlatformTypeSummary
)
