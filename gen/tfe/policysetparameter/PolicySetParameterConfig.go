package policysetparameter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicySetParameterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/policy_set_parameter#key PolicySetParameter#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/policy_set_parameter#policy_set_id PolicySetParameter#policy_set_id}.
	PolicySetId *string `field:"required" json:"policySetId" yaml:"policySetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/policy_set_parameter#id PolicySetParameter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/policy_set_parameter#sensitive PolicySetParameter#sensitive}.
	Sensitive interface{} `field:"optional" json:"sensitive" yaml:"sensitive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/policy_set_parameter#value PolicySetParameter#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

