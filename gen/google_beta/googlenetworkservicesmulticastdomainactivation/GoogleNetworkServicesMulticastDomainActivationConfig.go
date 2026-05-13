package googlenetworkservicesmulticastdomainactivation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesMulticastDomainActivationConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#location GoogleNetworkServicesMulticastDomainActivation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the multicast domain to activate. Use the following format: 'projects/* /locations/global/multicastDomains/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#multicast_domain GoogleNetworkServicesMulticastDomainActivation#multicast_domain}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastDomain *string `field:"required" json:"multicastDomain" yaml:"multicastDomain"`
	// A unique name for the multicast domain activation.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#multicast_domain_activation_id GoogleNetworkServicesMulticastDomainActivation#multicast_domain_activation_id}
	MulticastDomainActivationId *string `field:"required" json:"multicastDomainActivationId" yaml:"multicastDomainActivationId"`
	// An optional text description of the multicast domain activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#description GoogleNetworkServicesMulticastDomainActivation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Option to allow disabling placement policy for multicast infrastructure.
	//
	// Only applicable if the activation is for a domain associating with a
	// multicast domain group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#disable_placement_policy GoogleNetworkServicesMulticastDomainActivation#disable_placement_policy}
	DisablePlacementPolicy interface{} `field:"optional" json:"disablePlacementPolicy" yaml:"disablePlacementPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#id GoogleNetworkServicesMulticastDomainActivation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#labels GoogleNetworkServicesMulticastDomainActivation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#project GoogleNetworkServicesMulticastDomainActivation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#timeouts GoogleNetworkServicesMulticastDomainActivation#timeouts}
	Timeouts *GoogleNetworkServicesMulticastDomainActivationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_domain_activation#traffic_spec GoogleNetworkServicesMulticastDomainActivation#traffic_spec}
	TrafficSpec *GoogleNetworkServicesMulticastDomainActivationTrafficSpec `field:"optional" json:"trafficSpec" yaml:"trafficSpec"`
}

