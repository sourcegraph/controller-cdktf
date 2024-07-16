package googlekmskeyhandle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleKmsKeyHandleConfig struct {
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
	// The location for the KeyHandle. A full list of valid locations can be found by running 'gcloud kms locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#location GoogleKmsKeyHandle#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name for the KeyHandle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#name GoogleKmsKeyHandle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Selector of the resource type where we want to protect resources. For example, 'storage.googleapis.com/Bucket OR compute.googleapis.com/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#resource_type_selector GoogleKmsKeyHandle#resource_type_selector}
	ResourceTypeSelector *string `field:"required" json:"resourceTypeSelector" yaml:"resourceTypeSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#id GoogleKmsKeyHandle#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#project GoogleKmsKeyHandle#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_key_handle#timeouts GoogleKmsKeyHandle#timeouts}
	Timeouts *GoogleKmsKeyHandleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

