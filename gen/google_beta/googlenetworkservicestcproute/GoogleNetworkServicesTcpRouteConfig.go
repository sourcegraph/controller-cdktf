package googlenetworkservicestcproute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesTcpRouteConfig struct {
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
	// Name of the TcpRoute resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#name GoogleNetworkServicesTcpRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#rules GoogleNetworkServicesTcpRoute#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#description GoogleNetworkServicesTcpRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	//
	// Each gateway reference should match the pattern: projects/* /locations/global/gateways/<gateway_name>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#gateways GoogleNetworkServicesTcpRoute#gateways}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#id GoogleNetworkServicesTcpRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the TcpRoute resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#labels GoogleNetworkServicesTcpRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	//
	// Each mesh reference should match the pattern: projects/* /locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#meshes GoogleNetworkServicesTcpRoute#meshes}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Meshes *[]*string `field:"optional" json:"meshes" yaml:"meshes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#project GoogleNetworkServicesTcpRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#timeouts GoogleNetworkServicesTcpRoute#timeouts}
	Timeouts *GoogleNetworkServicesTcpRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

