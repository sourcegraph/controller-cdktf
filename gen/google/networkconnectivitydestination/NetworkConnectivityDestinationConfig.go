package networkconnectivitydestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityDestinationConfig struct {
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
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#endpoints NetworkConnectivityDestination#endpoints}
	Endpoints interface{} `field:"required" json:"endpoints" yaml:"endpoints"`
	// The IP prefix that represents your workload on another CSP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#ip_prefix NetworkConnectivityDestination#ip_prefix}
	IpPrefix *string `field:"required" json:"ipPrefix" yaml:"ipPrefix"`
	// The location of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#location NetworkConnectivityDestination#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The multicloud data transfer config of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#multicloud_data_transfer_config NetworkConnectivityDestination#multicloud_data_transfer_config}
	MulticloudDataTransferConfig *string `field:"required" json:"multicloudDataTransferConfig" yaml:"multicloudDataTransferConfig"`
	// The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#name NetworkConnectivityDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#description NetworkConnectivityDestination#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#id NetworkConnectivityDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#labels NetworkConnectivityDestination#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#project NetworkConnectivityDestination#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_destination#timeouts NetworkConnectivityDestination#timeouts}
	Timeouts *NetworkConnectivityDestinationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

