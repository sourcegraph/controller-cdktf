package networkconnectivitymulticlouddatatransferconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityMulticloudDataTransferConfigConfig struct {
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
	// The location of the multicloud data transfer config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#location NetworkConnectivityMulticloudDataTransferConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the MulticloudDataTransferConfig resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#name NetworkConnectivityMulticloudDataTransferConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#description NetworkConnectivityMulticloudDataTransferConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#id NetworkConnectivityMulticloudDataTransferConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#labels NetworkConnectivityMulticloudDataTransferConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#project NetworkConnectivityMulticloudDataTransferConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#services NetworkConnectivityMulticloudDataTransferConfig#services}
	Services interface{} `field:"optional" json:"services" yaml:"services"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_connectivity_multicloud_data_transfer_config#timeouts NetworkConnectivityMulticloudDataTransferConfig#timeouts}
	Timeouts *NetworkConnectivityMulticloudDataTransferConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

