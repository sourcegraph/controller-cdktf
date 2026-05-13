package googlenetworkconnectivitygatewayadvertisedroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkConnectivityGatewayAdvertisedRouteConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#location GoogleNetworkConnectivityGatewayAdvertisedRoute#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the gateway advertised route. Route names must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#name GoogleNetworkConnectivityGatewayAdvertisedRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the spoke.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#spoke GoogleNetworkConnectivityGatewayAdvertisedRoute#spoke}
	Spoke *string `field:"required" json:"spoke" yaml:"spoke"`
	// An optional description of the gateway advertised route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#description GoogleNetworkConnectivityGatewayAdvertisedRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#id GoogleNetworkConnectivityGatewayAdvertisedRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This route's advertised IP address range.
	//
	// Must be a valid CIDR-formatted prefix.
	// If an IP address is provided without a subnet mask, it is interpreted as, for IPv4, a /32 singular IP address range, and, for IPv6, /128
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#ip_range GoogleNetworkConnectivityGatewayAdvertisedRoute#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://docs.cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#labels GoogleNetworkConnectivityGatewayAdvertisedRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The priority of this advertised route.
	//
	// You can choose a value from 0 to 65335.
	// If you don't provide a value, Google Cloud assigns a priority of 100 to the ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#priority GoogleNetworkConnectivityGatewayAdvertisedRoute#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#project GoogleNetworkConnectivityGatewayAdvertisedRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// the recipient of this advertised route Possible values: ["RECIPIENT_UNSPECIFIED", "ADVERTISE_TO_HUB"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#recipient GoogleNetworkConnectivityGatewayAdvertisedRoute#recipient}
	Recipient *string `field:"optional" json:"recipient" yaml:"recipient"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_gateway_advertised_route#timeouts GoogleNetworkConnectivityGatewayAdvertisedRoute#timeouts}
	Timeouts *GoogleNetworkConnectivityGatewayAdvertisedRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

