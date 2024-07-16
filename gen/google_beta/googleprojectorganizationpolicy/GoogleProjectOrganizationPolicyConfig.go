package googleprojectorganizationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleProjectOrganizationPolicyConfig struct {
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
	// The name of the Constraint the Policy is configuring, for example, serviceuser.services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#constraint GoogleProjectOrganizationPolicy#constraint}
	Constraint *string `field:"required" json:"constraint" yaml:"constraint"`
	// The project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#project GoogleProjectOrganizationPolicy#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// boolean_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#boolean_policy GoogleProjectOrganizationPolicy#boolean_policy}
	BooleanPolicy *GoogleProjectOrganizationPolicyBooleanPolicy `field:"optional" json:"booleanPolicy" yaml:"booleanPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#id GoogleProjectOrganizationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// list_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#list_policy GoogleProjectOrganizationPolicy#list_policy}
	ListPolicy *GoogleProjectOrganizationPolicyListPolicy `field:"optional" json:"listPolicy" yaml:"listPolicy"`
	// restore_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#restore_policy GoogleProjectOrganizationPolicy#restore_policy}
	RestorePolicy *GoogleProjectOrganizationPolicyRestorePolicy `field:"optional" json:"restorePolicy" yaml:"restorePolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#timeouts GoogleProjectOrganizationPolicy#timeouts}
	Timeouts *GoogleProjectOrganizationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Version of the Policy. Default version is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_project_organization_policy#version GoogleProjectOrganizationPolicy#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

