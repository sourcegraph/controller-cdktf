package googleaccesscontextmanageraccesslevel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAccessContextManagerAccessLevelConfig struct {
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
	// Resource name for the Access Level.
	//
	// The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#name GoogleAccessContextManagerAccessLevel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#parent GoogleAccessContextManagerAccessLevel#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Human readable title. Must be unique within the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#title GoogleAccessContextManagerAccessLevel#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#basic GoogleAccessContextManagerAccessLevel#basic}
	Basic *GoogleAccessContextManagerAccessLevelBasic `field:"optional" json:"basic" yaml:"basic"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#custom GoogleAccessContextManagerAccessLevel#custom}
	Custom *GoogleAccessContextManagerAccessLevelCustom `field:"optional" json:"custom" yaml:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#description GoogleAccessContextManagerAccessLevel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#id GoogleAccessContextManagerAccessLevel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_access_level#timeouts GoogleAccessContextManagerAccessLevel#timeouts}
	Timeouts *GoogleAccessContextManagerAccessLevelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

