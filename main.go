package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/Disciple-Media/terraform-provider-dme/dme"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dme.Provider})
}
