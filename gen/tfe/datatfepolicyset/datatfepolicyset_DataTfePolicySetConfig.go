package datatfepolicyset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataTfePolicySetConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/policy_set#name DataTfePolicySet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/policy_set#id DataTfePolicySet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The policy-as-code framework for the policy. Valid values are sentinel and opa.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/policy_set#kind DataTfePolicySet#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/policy_set#organization DataTfePolicySet#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Whether users can override this policy when it fails during a run. Only valid for OPA policies.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/d/policy_set#overridable DataTfePolicySet#overridable}
	Overridable interface{} `field:"optional" json:"overridable" yaml:"overridable"`
}

