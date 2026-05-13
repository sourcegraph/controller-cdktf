package googlekmsfolderkajpolicyconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleKmsFolderKajPolicyConfigConfig struct {
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
	// The numeric folder number for which to retrieve config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_folder_kaj_policy_config#folder GoogleKmsFolderKajPolicyConfig#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// default_key_access_justification_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_folder_kaj_policy_config#default_key_access_justification_policy GoogleKmsFolderKajPolicyConfig#default_key_access_justification_policy}
	DefaultKeyAccessJustificationPolicy *GoogleKmsFolderKajPolicyConfigDefaultKeyAccessJustificationPolicy `field:"optional" json:"defaultKeyAccessJustificationPolicy" yaml:"defaultKeyAccessJustificationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_folder_kaj_policy_config#id GoogleKmsFolderKajPolicyConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_kms_folder_kaj_policy_config#timeouts GoogleKmsFolderKajPolicyConfig#timeouts}
	Timeouts *GoogleKmsFolderKajPolicyConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

