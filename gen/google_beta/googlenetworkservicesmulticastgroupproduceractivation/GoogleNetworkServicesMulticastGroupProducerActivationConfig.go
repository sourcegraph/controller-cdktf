package googlenetworkservicesmulticastgroupproduceractivation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesMulticastGroupProducerActivationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#location GoogleNetworkServicesMulticastGroupProducerActivation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A unique name for the multicast group producer activation.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#multicast_group_producer_activation_id GoogleNetworkServicesMulticastGroupProducerActivation#multicast_group_producer_activation_id}
	MulticastGroupProducerActivationId *string `field:"required" json:"multicastGroupProducerActivationId" yaml:"multicastGroupProducerActivationId"`
	// The resource name of the multicast group range activationcreated by the admin in the same zone as this multicast group producer activation.
	//
	// Use the
	// following format:
	// // 'projects/* /locations/* /multicastGroupRangeActivations/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#multicast_group_range_activation GoogleNetworkServicesMulticastGroupProducerActivation#multicast_group_range_activation}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastGroupRangeActivation *string `field:"required" json:"multicastGroupRangeActivation" yaml:"multicastGroupRangeActivation"`
	// The resource name of the multicast producer association that is in the same zone as this multicast group producer activation.
	//
	// Use the following format:
	// 'projects/* /locations/* /multicastProducerAssociations/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#multicast_producer_association GoogleNetworkServicesMulticastGroupProducerActivation#multicast_producer_association}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastProducerAssociation *string `field:"required" json:"multicastProducerAssociation" yaml:"multicastProducerAssociation"`
	// An optional text description of the multicast group producer activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#description GoogleNetworkServicesMulticastGroupProducerActivation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#id GoogleNetworkServicesMulticastGroupProducerActivation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#labels GoogleNetworkServicesMulticastGroupProducerActivation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#project GoogleNetworkServicesMulticastGroupProducerActivation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_services_multicast_group_producer_activation#timeouts GoogleNetworkServicesMulticastGroupProducerActivation#timeouts}
	Timeouts *GoogleNetworkServicesMulticastGroupProducerActivationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

