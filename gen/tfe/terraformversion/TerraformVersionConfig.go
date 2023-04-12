package terraformversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TerraformVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#sha TerraformVersion#sha}.
	Sha *string `field:"required" json:"sha" yaml:"sha"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#url TerraformVersion#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#version TerraformVersion#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#beta TerraformVersion#beta}.
	Beta interface{} `field:"optional" json:"beta" yaml:"beta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#deprecated TerraformVersion#deprecated}.
	Deprecated interface{} `field:"optional" json:"deprecated" yaml:"deprecated"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#deprecated_reason TerraformVersion#deprecated_reason}.
	DeprecatedReason *string `field:"optional" json:"deprecatedReason" yaml:"deprecatedReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#enabled TerraformVersion#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#id TerraformVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/terraform_version#official TerraformVersion#official}.
	Official interface{} `field:"optional" json:"official" yaml:"official"`
}

