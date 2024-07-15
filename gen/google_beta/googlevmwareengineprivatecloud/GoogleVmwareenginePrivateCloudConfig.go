package googlevmwareengineprivatecloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVmwareenginePrivateCloudConfig struct {
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
	// The location where the PrivateCloud should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#location GoogleVmwareenginePrivateCloud#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// management_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#management_cluster GoogleVmwareenginePrivateCloud#management_cluster}
	ManagementCluster *GoogleVmwareenginePrivateCloudManagementCluster `field:"required" json:"managementCluster" yaml:"managementCluster"`
	// The ID of the PrivateCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#name GoogleVmwareenginePrivateCloud#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#network_config GoogleVmwareenginePrivateCloud#network_config}
	NetworkConfig *GoogleVmwareenginePrivateCloudNetworkConfig `field:"required" json:"networkConfig" yaml:"networkConfig"`
	// User-provided description for this private cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#description GoogleVmwareenginePrivateCloud#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#id GoogleVmwareenginePrivateCloud#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#project GoogleVmwareenginePrivateCloud#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#timeouts GoogleVmwareenginePrivateCloud#timeouts}
	Timeouts *GoogleVmwareenginePrivateCloudTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vmwareengine_private_cloud#type GoogleVmwareenginePrivateCloud#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

