package networkzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkZoneConfig struct {
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
	// Name of the Network Zone Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#name NetworkZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the Network Zone - can be `IP`, `DYNAMIC` or `DYNAMIC_V2` only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#type NetworkZone#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of asns included.
	//
	// Format of each array value: a string representation of an ASN numeric value. Use with type `DYNAMIC` or `DYNAMIC_V2`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#asns NetworkZone#asns}
	Asns *[]*string `field:"optional" json:"asns" yaml:"asns"`
	// Array of locations ISO-3166-1(2) included. Format code: countryCode OR countryCode-regionCode. Use with type `DYNAMIC` or `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#dynamic_locations NetworkZone#dynamic_locations}
	DynamicLocations *[]*string `field:"optional" json:"dynamicLocations" yaml:"dynamicLocations"`
	// Array of locations ISO-3166-1(2) excluded. Format code: countryCode OR countryCode-regionCode. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#dynamic_locations_exclude NetworkZone#dynamic_locations_exclude}
	DynamicLocationsExclude *[]*string `field:"optional" json:"dynamicLocationsExclude" yaml:"dynamicLocationsExclude"`
	// Type of proxy being controlled by this dynamic network zone - can be one of `Any`, `TorAnonymizer` or `NotTorAnonymizer`.
	//
	// Use with type `DYNAMIC`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#dynamic_proxy_type NetworkZone#dynamic_proxy_type}
	DynamicProxyType *string `field:"optional" json:"dynamicProxyType" yaml:"dynamicProxyType"`
	// Array of values in CIDR/range form depending on the way it's been declared (i.e. CIDR will contain /suffix). Please check API docs for examples. Use with type `IP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#gateways NetworkZone#gateways}
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#id NetworkZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of ip service excluded. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#ip_service_categories_exclude NetworkZone#ip_service_categories_exclude}
	IpServiceCategoriesExclude *[]*string `field:"optional" json:"ipServiceCategoriesExclude" yaml:"ipServiceCategoriesExclude"`
	// List of ip service included. Use with type `DYNAMIC_V2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#ip_service_categories_include NetworkZone#ip_service_categories_include}
	IpServiceCategoriesInclude *[]*string `field:"optional" json:"ipServiceCategoriesInclude" yaml:"ipServiceCategoriesInclude"`
	// Array of values in CIDR/range form depending on the way it's been declared (i.e. CIDR will contain /suffix). Please check API docs for examples. Can not be set if `usage` is set to `BLOCKLIST`. Use with type `IP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#proxies NetworkZone#proxies}
	Proxies *[]*string `field:"optional" json:"proxies" yaml:"proxies"`
	// Network Status - can either be `ACTIVE` or `INACTIVE` only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#status NetworkZone#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Usage of the Network Zone - can be either `POLICY` or `BLOCKLIST`. By default, it is `POLICY`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone#usage NetworkZone#usage}
	Usage *string `field:"optional" json:"usage" yaml:"usage"`
}

