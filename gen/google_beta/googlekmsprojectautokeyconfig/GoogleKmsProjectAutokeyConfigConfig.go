package googlekmsprojectautokeyconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleKmsProjectAutokeyConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_project_autokey_config#id GoogleKmsProjectAutokeyConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// How Autokey determines which project to use when provisioning CMEK keys. Possible values: ["RESOURCE_PROJECT", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_project_autokey_config#key_project_resolution_mode GoogleKmsProjectAutokeyConfig#key_project_resolution_mode}
	KeyProjectResolutionMode *string `field:"optional" json:"keyProjectResolutionMode" yaml:"keyProjectResolutionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_project_autokey_config#project GoogleKmsProjectAutokeyConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_project_autokey_config#timeouts GoogleKmsProjectAutokeyConfig#timeouts}
	Timeouts *GoogleKmsProjectAutokeyConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

