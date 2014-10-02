package main

import (
	"github.com/bndr/gopencils"
)

type IntuitionAPI struct {
	Version string
	Api     *gopencils.Resource
}

// NewIntuitionAPI constructs a new interface to the RESTful server
func NewIntuitionAPI(addr string) *IntuitionAPI {
	return &IntuitionAPI{
		Version: "v0",
		Api:     gopencils.Api(addr),
	}
}

// Health queries the current API status
func (self *IntuitionAPI) Health() (*TelepathyHealth, error) {
	health := new(TelepathyHealth)
	_, err := self.Api.Res(self.Version).Res("health", health).Get()
	return health, err
}

func (self *IntuitionAPI) Doc(resource string) (*ModulesDoc, error) {
	docs := new(ModulesDoc)
	_, err := self.Api.Res(self.Version).Res("docs/"+resource, docs).Get()
	return docs, err
}
