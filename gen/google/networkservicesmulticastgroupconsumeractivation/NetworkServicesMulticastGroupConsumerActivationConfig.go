package networkservicesmulticastgroupconsumeractivation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesMulticastGroupConsumerActivationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#location NetworkServicesMulticastGroupConsumerActivation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the multicast consumer association that is in the same zone as this multicast group consumer activation.
	//
	// Use the following format:
	// 'projects/* /locations/* /multicastConsumerAssociations/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#multicast_consumer_association NetworkServicesMulticastGroupConsumerActivation#multicast_consumer_association}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastConsumerAssociation *string `field:"required" json:"multicastConsumerAssociation" yaml:"multicastConsumerAssociation"`
	// A unique name for the multicast group consumer activation.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#multicast_group_consumer_activation_id NetworkServicesMulticastGroupConsumerActivation#multicast_group_consumer_activation_id}
	MulticastGroupConsumerActivationId *string `field:"required" json:"multicastGroupConsumerActivationId" yaml:"multicastGroupConsumerActivationId"`
	// The resource name of the multicast group range activation created by the admin in the same zone as this multicast group consumer activation.
	//
	// Use the
	// following format:
	// // 'projects/* /locations/* /multicastGroupRangeActivations/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#multicast_group_range_activation NetworkServicesMulticastGroupConsumerActivation#multicast_group_range_activation}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastGroupRangeActivation *string `field:"required" json:"multicastGroupRangeActivation" yaml:"multicastGroupRangeActivation"`
	// An optional text description of the multicast group consumer activation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#description NetworkServicesMulticastGroupConsumerActivation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#id NetworkServicesMulticastGroupConsumerActivation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#labels NetworkServicesMulticastGroupConsumerActivation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#log_config NetworkServicesMulticastGroupConsumerActivation#log_config}
	LogConfig *NetworkServicesMulticastGroupConsumerActivationLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#project NetworkServicesMulticastGroupConsumerActivation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_consumer_activation#timeouts NetworkServicesMulticastGroupConsumerActivation#timeouts}
	Timeouts *NetworkServicesMulticastGroupConsumerActivationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

