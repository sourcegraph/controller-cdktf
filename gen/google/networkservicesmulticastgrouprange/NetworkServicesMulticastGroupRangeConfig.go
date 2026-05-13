package networkservicesmulticastgrouprange

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesMulticastGroupRangeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#location NetworkServicesMulticastGroupRange#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the multicast domain in which to create this multicast group range. Use the following format: 'projects/* /locations/global/multicastDomains/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#multicast_domain NetworkServicesMulticastGroupRange#multicast_domain}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastDomain *string `field:"required" json:"multicastDomain" yaml:"multicastDomain"`
	// A unique name for the multicast group range.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#multicast_group_range_id NetworkServicesMulticastGroupRange#multicast_group_range_id}
	MulticastGroupRangeId *string `field:"required" json:"multicastGroupRangeId" yaml:"multicastGroupRangeId"`
	// The resource name of the internal range reserved for this multicast group range.
	//
	// The internal range must be a Class D address (224.0.0.0 to 239.255.255.255)
	// and have a prefix length >= 23.
	//
	// Use the following format:
	// 'projects/* /locations/global/internalRanges/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#reserved_internal_range NetworkServicesMulticastGroupRange#reserved_internal_range}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ReservedInternalRange *string `field:"required" json:"reservedInternalRange" yaml:"reservedInternalRange"`
	// A list of consumer projects that are allowed to subscribe to the multicast IP addresses within the range defined by this MulticastGroupRange.
	//
	// The
	// project can be specified using its project ID or project number. If left
	// empty, then all consumer projects are allowed (unless
	// require_explicit_accept is set to true) once they have VPC networks
	// associated to the multicast domain. The current max length of the accept
	// list is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#consumer_accept_list NetworkServicesMulticastGroupRange#consumer_accept_list}
	ConsumerAcceptList *[]*string `field:"optional" json:"consumerAcceptList" yaml:"consumerAcceptList"`
	// An optional text description of the multicast group range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#description NetworkServicesMulticastGroupRange#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Multicast group range's distribution scope.
	//
	// Intra-zone or intra-region
	// cross-zone is supported, with default value being intra-region. Cross
	// region distribution is not supported.
	// Possible values:
	// INTRA_ZONE
	// INTRA_REGION
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#distribution_scope NetworkServicesMulticastGroupRange#distribution_scope}
	DistributionScope *string `field:"optional" json:"distributionScope" yaml:"distributionScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#id NetworkServicesMulticastGroupRange#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#labels NetworkServicesMulticastGroupRange#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#log_config NetworkServicesMulticastGroupRange#log_config}
	LogConfig *NetworkServicesMulticastGroupRangeLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#project NetworkServicesMulticastGroupRange#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Whether an empty consumer_accept_list will deny all consumer projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#require_explicit_accept NetworkServicesMulticastGroupRange#require_explicit_accept}
	RequireExplicitAccept interface{} `field:"optional" json:"requireExplicitAccept" yaml:"requireExplicitAccept"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_group_range#timeouts NetworkServicesMulticastGroupRange#timeouts}
	Timeouts *NetworkServicesMulticastGroupRangeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

