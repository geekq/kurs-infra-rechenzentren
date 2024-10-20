package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"infrastructure-as-code.de/example/terraform-provider-apt-package/internal/provider"
)

var (
	version string = "1.0.0"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "infrastructure-as-code.de/example/apt-package",
	})
}
