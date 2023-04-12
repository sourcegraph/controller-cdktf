package datatfeoutputs

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfeOutputsConfig struct {
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
	// The workspace to fetch the remote state from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/outputs#workspace DataTfeOutputs#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// The organization to fetch the remote state from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/outputs#organization DataTfeOutputs#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/outputs#values DataTfeOutputs#values}.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
}

