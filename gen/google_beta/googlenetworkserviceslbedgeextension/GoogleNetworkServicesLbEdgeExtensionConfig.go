package googlenetworkserviceslbedgeextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesLbEdgeExtensionConfig struct {
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
	// extension_chains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#extension_chains GoogleNetworkServicesLbEdgeExtension#extension_chains}
	ExtensionChains interface{} `field:"required" json:"extensionChains" yaml:"extensionChains"`
	// A list of references to the forwarding rules to which this service extension is attached.
	//
	// At least one forwarding rule is required. Only one LbEdgeExtension resource can be associated with a forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#forwarding_rules GoogleNetworkServicesLbEdgeExtension#forwarding_rules}
	ForwardingRules *[]*string `field:"required" json:"forwardingRules" yaml:"forwardingRules"`
	// All forwarding rules referenced by this extension must share the same load balancing scheme. Possible values: ["EXTERNAL_MANAGED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#load_balancing_scheme GoogleNetworkServicesLbEdgeExtension#load_balancing_scheme}
	LoadBalancingScheme *string `field:"required" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// The location of the edge extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#location GoogleNetworkServicesLbEdgeExtension#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the LbEdgeExtension resource in the following format: projects/{project}/locations/{location}/lbEdgeExtensions/{lbEdgeExtensions}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#name GoogleNetworkServicesLbEdgeExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#description GoogleNetworkServicesLbEdgeExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#id GoogleNetworkServicesLbEdgeExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with the LbEdgeExtension resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#labels GoogleNetworkServicesLbEdgeExtension#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#project GoogleNetworkServicesLbEdgeExtension#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_lb_edge_extension#timeouts GoogleNetworkServicesLbEdgeExtension#timeouts}
	Timeouts *GoogleNetworkServicesLbEdgeExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

