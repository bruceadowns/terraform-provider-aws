// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package outposts

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	outposts_sdkv2 "github.com/aws/aws-sdk-go-v2/service/outposts"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceOutpostAsset,
			TypeName: "aws_outposts_asset",
			Name:     "Asset",
		},
		{
			Factory:  dataSourceOutpostAssets,
			TypeName: "aws_outposts_assets",
			Name:     "Assets",
		},
		{
			Factory:  dataSourceOutpost,
			TypeName: "aws_outposts_outpost",
			Name:     "Outpost",
		},
		{
			Factory:  dataSourceOutpostInstanceType,
			TypeName: "aws_outposts_outpost_instance_type",
			Name:     "Outpost Instance Type",
		},
		{
			Factory:  dataSourceOutpostInstanceTypes,
			TypeName: "aws_outposts_outpost_instance_types",
			Name:     "Outpost Instance Types",
		},
		{
			Factory:  dataSourceOutposts,
			TypeName: "aws_outposts_outposts",
			Name:     "Outposts",
		},
		{
			Factory:  dataSourceSite,
			TypeName: "aws_outposts_site",
			Name:     "Site",
		},
		{
			Factory:  dataSourceSites,
			TypeName: "aws_outposts_sites",
			Name:     "Sites",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Outposts
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*outposts_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return outposts_sdkv2.NewFromConfig(cfg,
		outposts_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
