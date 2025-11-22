package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure VertexAIRAGProvider satisfies various provider interfaces.
var _ provider.Provider = &VertexAIRAGProvider{}

// VertexAIRAGProvider defines the provider implementation.
type VertexAIRAGProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// VertexAIRAGProviderModel describes the provider data model.
type VertexAIRAGProviderModel struct {
	Project     types.String `tfsdk:"project"`
	Region      types.String `tfsdk:"region"`
	AccessToken types.String `tfsdk:"access_token"`
}

func (p *VertexAIRAGProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "vertexairag"
	resp.Version = p.version
}

func (p *VertexAIRAGProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project": schema.StringAttribute{
				MarkdownDescription: "The Google Cloud Project ID.",
				Optional:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "The Google Cloud Region.",
				Optional:            true,
			},
			"access_token": schema.StringAttribute{
				MarkdownDescription: "The Google Cloud Access Token.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *VertexAIRAGProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data VertexAIRAGProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration
	config := ProviderConfig{
		Project:     os.Getenv("GOOGLE_PROJECT"),
		Region:      os.Getenv("GOOGLE_REGION"),
		AccessToken: os.Getenv("GOOGLE_ACCESS_TOKEN"),
	}

	if !data.Project.IsNull() {
		config.Project = data.Project.ValueString()
	}
	if !data.Region.IsNull() {
		config.Region = data.Region.ValueString()
	}
	if !data.AccessToken.IsNull() {
		config.AccessToken = data.AccessToken.ValueString()
	}

	resp.DataSourceData = config
	resp.ResourceData = config
}

func (p *VertexAIRAGProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRagCorpusResource,
	}
}

func (p *VertexAIRAGProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &VertexAIRAGProvider{
			version: version,
		}
	}
}

// ProviderConfig holds the configuration for the provider
type ProviderConfig struct {
	Project     string
	Region      string
	AccessToken string
}
