package main

import (
	"github.com/bndr/gopencils"
)

type IntuitionAPI struct {
	Version string
	Api     *gopencils.Resource
}

func NewIntuitionAPI(addr string) *IntuitionAPI {
	return &IntuitionAPI{
		Version: "v0",
		Api:     gopencils.Api(addr),
	}
}

// FIXME Doesn't understand 'true' value
func (self *IntuitionAPI) Health() (*TelepathyHealth, error) {
	health := new(TelepathyHealth)
	_, err := self.Api.Res(self.Version).Res("health", health).Get()
	return health, err
}
