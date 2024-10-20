package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &AptProvider{}
}

type AptProvider struct{}

func (p *AptProvider) DataSources(context.Context) []func() datasource.DataSource {
	// we can implement later and list apt packages installed
	// return an empty slice for now
	return []func() datasource.DataSource{}
}

func (p *AptProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "apt_package"
}

func (p *AptProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	// No global provider-level configuration needed
}

// Configure is used to initialize the provider configuration
func (p *AptProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Since we do not need any authentication or any other global
	// configuration, we do not need to do anything here.
}

func (p *AptProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAptPackageResource,
	}
}
