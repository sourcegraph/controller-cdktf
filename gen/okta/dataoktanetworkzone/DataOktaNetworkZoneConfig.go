package dataoktanetworkzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOktaNetworkZoneConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Array of locations ISO-3166-1(2) excluded. Format code: countryCode OR countryCode-regionCode. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/network_zone#dynamic_locations_exclude DataOktaNetworkZone#dynamic_locations_exclude}
	DynamicLocationsExclude *[]*string `field:"optional" json:"dynamicLocationsExclude" yaml:"dynamicLocationsExclude"`
	// ID of the network zone to retrieve, conflicts with `name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/network_zone#id DataOktaNetworkZone#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of ip service excluded. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/network_zone#ip_service_categories_exclude DataOktaNetworkZone#ip_service_categories_exclude}
	IpServiceCategoriesExclude *[]*string `field:"optional" json:"ipServiceCategoriesExclude" yaml:"ipServiceCategoriesExclude"`
	// List of ip service included. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/network_zone#ip_service_categories_include DataOktaNetworkZone#ip_service_categories_include}
	IpServiceCategoriesInclude *[]*string `field:"optional" json:"ipServiceCategoriesInclude" yaml:"ipServiceCategoriesInclude"`
	// Name of the network zone to retrieve, conflicts with `id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/data-sources/network_zone#name DataOktaNetworkZone#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

