package networkzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkZoneConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#name NetworkZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the Network Zone - can either be IP or DYNAMIC only.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#type NetworkZone#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Format of each array value: a string representation of an ASN numeric value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#asns NetworkZone#asns}
	Asns *[]*string `field:"optional" json:"asns" yaml:"asns"`
	// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#dynamic_locations NetworkZone#dynamic_locations}
	DynamicLocations *[]*string `field:"optional" json:"dynamicLocations" yaml:"dynamicLocations"`
	// Type of proxy being controlled by this network zone.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#dynamic_proxy_type NetworkZone#dynamic_proxy_type}
	DynamicProxyType *string `field:"optional" json:"dynamicProxyType" yaml:"dynamicProxyType"`
	// Array of values in CIDR/range form depending on the way it's been declared (i.e. CIDR will contain /suffix). Please check API docs for examples.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#gateways NetworkZone#gateways}
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#id NetworkZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Array of values in CIDR/range form depending on the way it's been declared (i.e. CIDR will contain /suffix). Please check API docs for examples.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#proxies NetworkZone#proxies}
	Proxies *[]*string `field:"optional" json:"proxies" yaml:"proxies"`
	// Zone's purpose: POLICY or BLOCKLIST.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/network_zone#usage NetworkZone#usage}
	Usage *string `field:"optional" json:"usage" yaml:"usage"`
}

