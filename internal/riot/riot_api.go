package riot

import (
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
)

type RiotAPI struct {
	Client *golio.Client
}

func NewRiotAPI(token string) *RiotAPI {
	return &RiotAPI{
		Client: golio.NewClient(
			token,
			golio.WithRegion(api.RegionVietnam),
		),
	}
}
