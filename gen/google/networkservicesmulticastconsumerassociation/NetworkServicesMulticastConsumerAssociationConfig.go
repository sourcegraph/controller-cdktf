package networkservicesmulticastconsumerassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesMulticastConsumerAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#location NetworkServicesMulticastConsumerAssociation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A unique name for the multicast consumer association.
	//
	// The name is restricted to letters, numbers, and hyphen, with the first
	// character a letter, and the last a letter or a number. The name must not
	// exceed 48 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#multicast_consumer_association_id NetworkServicesMulticastConsumerAssociation#multicast_consumer_association_id}
	MulticastConsumerAssociationId *string `field:"required" json:"multicastConsumerAssociationId" yaml:"multicastConsumerAssociationId"`
	// The resource name of the multicast domain activation that is in the same zone as this multicast consumer association.
	//
	// Use the following format:
	// // 'projects/* /locations/* /multicastDomainActivations/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#multicast_domain_activation NetworkServicesMulticastConsumerAssociation#multicast_domain_activation}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	MulticastDomainActivation *string `field:"required" json:"multicastDomainActivation" yaml:"multicastDomainActivation"`
	// The resource name of the multicast consumer VPC network. Use following format: 'projects/{project}/locations/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#network NetworkServicesMulticastConsumerAssociation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// An optional text description of the multicast consumer association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#description NetworkServicesMulticastConsumerAssociation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#id NetworkServicesMulticastConsumerAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key-value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#labels NetworkServicesMulticastConsumerAssociation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#project NetworkServicesMulticastConsumerAssociation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_services_multicast_consumer_association#timeouts NetworkServicesMulticastConsumerAssociation#timeouts}
	Timeouts *NetworkServicesMulticastConsumerAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

