package googleresourcemanagercapability

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleResourceManagerCapabilityConfig struct {
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
	// Capability name that should be updated on the folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_resource_manager_capability#capability_name GoogleResourceManagerCapability#capability_name}
	CapabilityName *string `field:"required" json:"capabilityName" yaml:"capabilityName"`
	// Folder on which Capability needs to be updated in the format folders/folder_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_resource_manager_capability#parent GoogleResourceManagerCapability#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Capability Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_resource_manager_capability#value GoogleResourceManagerCapability#value}
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_resource_manager_capability#id GoogleResourceManagerCapability#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_resource_manager_capability#timeouts GoogleResourceManagerCapability#timeouts}
	Timeouts *GoogleResourceManagerCapabilityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

