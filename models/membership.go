package models

type Membership struct {
	IdMembership int
	IdUser       int
	IdSeasson    int
	TypeTier     string
	PriceTier    float32
}
