package googlenetworkservicesgrpcroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesGrpcRouteConfig struct {
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
	// Required. Service hostnames with an optional port for which this route describes traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#hostnames GoogleNetworkServicesGrpcRoute#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// Name of the GrpcRoute resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#name GoogleNetworkServicesGrpcRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#rules GoogleNetworkServicesGrpcRoute#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#description GoogleNetworkServicesGrpcRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#gateways GoogleNetworkServicesGrpcRoute#gateways}
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#id GoogleNetworkServicesGrpcRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the GrpcRoute resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#labels GoogleNetworkServicesGrpcRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#meshes GoogleNetworkServicesGrpcRoute#meshes}
	Meshes *[]*string `field:"optional" json:"meshes" yaml:"meshes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#project GoogleNetworkServicesGrpcRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#timeouts GoogleNetworkServicesGrpcRoute#timeouts}
	Timeouts *GoogleNetworkServicesGrpcRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

