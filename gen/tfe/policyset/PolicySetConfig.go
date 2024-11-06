package policyset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PolicySetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#name PolicySet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#description PolicySet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#global PolicySet#global}.
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#id PolicySet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#kind PolicySet#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#organization PolicySet#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#overridable PolicySet#overridable}.
	Overridable interface{} `field:"optional" json:"overridable" yaml:"overridable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#policies_path PolicySet#policies_path}.
	PoliciesPath *string `field:"optional" json:"policiesPath" yaml:"policiesPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#policy_ids PolicySet#policy_ids}.
	PolicyIds *[]*string `field:"optional" json:"policyIds" yaml:"policyIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#slug PolicySet#slug}.
	Slug *map[string]*string `field:"optional" json:"slug" yaml:"slug"`
	// vcs_repo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#vcs_repo PolicySet#vcs_repo}
	VcsRepo *PolicySetVcsRepo `field:"optional" json:"vcsRepo" yaml:"vcsRepo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/policy_set#workspace_ids PolicySet#workspace_ids}.
	WorkspaceIds *[]*string `field:"optional" json:"workspaceIds" yaml:"workspaceIds"`
}

