package googlenetworkservicesmulticastgrouprangeactivation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesMulticastGroupRangeActivationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#location GoogleNetworkServicesMulticastGroupRangeActivation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of a multicast domain activation that is in the same zone as this multicast group.
	//
	// Use the following format:
	// 'projects/* /locations/* /multicastDomainActivations/*'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#multicast_domain_activation GoogleNetworkServicesMulticastGroupRangeActivation#multicast_domain_activation}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastDomainActivation *string `field:"required" json:"multicastDomainActivation" yaml:"multicastDomainActivation"`
	// The resource name of the global multicast group range for the group. Use the following format: 'projects/* /locations/global/multicastGroupRanges/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#multicast_group_range GoogleNetworkServicesMulticastGroupRangeActivation#multicast_group_range}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastGroupRange *string `field:"required" json:"multicastGroupRange" yaml:"multicastGroupRange"`
	// A unique name for the multicast group range activation.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#multicast_group_range_activation_id GoogleNetworkServicesMulticastGroupRangeActivation#multicast_group_range_activation_id}
	MulticastGroupRangeActivationId *string `field:"required" json:"multicastGroupRangeActivationId" yaml:"multicastGroupRangeActivationId"`
	// An optional text description of the multicast group range activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#description GoogleNetworkServicesMulticastGroupRangeActivation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#id GoogleNetworkServicesMulticastGroupRangeActivation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#labels GoogleNetworkServicesMulticastGroupRangeActivation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#log_config GoogleNetworkServicesMulticastGroupRangeActivation#log_config}
	LogConfig *GoogleNetworkServicesMulticastGroupRangeActivationLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#project GoogleNetworkServicesMulticastGroupRangeActivation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_range_activation#timeouts GoogleNetworkServicesMulticastGroupRangeActivation#timeouts}
	Timeouts *GoogleNetworkServicesMulticastGroupRangeActivationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

