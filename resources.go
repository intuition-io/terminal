package main

type Trades struct {
	workers []interface{}
}

type ComponentInfos struct {
	Name    string
	Version string
	State   interface{}
}

type TelepathyHealth struct {
	Components []ComponentInfos
	Version    string
}

type ModulesDoc struct {
	Version   string
	Resources []interface{}
}
