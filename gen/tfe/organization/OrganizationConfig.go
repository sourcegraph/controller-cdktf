package organization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#email Organization#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#name Organization#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#allow_force_delete_workspaces Organization#allow_force_delete_workspaces}.
	AllowForceDeleteWorkspaces interface{} `field:"optional" json:"allowForceDeleteWorkspaces" yaml:"allowForceDeleteWorkspaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#assessments_enforced Organization#assessments_enforced}.
	AssessmentsEnforced interface{} `field:"optional" json:"assessmentsEnforced" yaml:"assessmentsEnforced"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#collaborator_auth_policy Organization#collaborator_auth_policy}.
	CollaboratorAuthPolicy *string `field:"optional" json:"collaboratorAuthPolicy" yaml:"collaboratorAuthPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#cost_estimation_enabled Organization#cost_estimation_enabled}.
	CostEstimationEnabled interface{} `field:"optional" json:"costEstimationEnabled" yaml:"costEstimationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#id Organization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#owners_team_saml_role_id Organization#owners_team_saml_role_id}.
	OwnersTeamSamlRoleId *string `field:"optional" json:"ownersTeamSamlRoleId" yaml:"ownersTeamSamlRoleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#send_passing_statuses_for_untriggered_speculative_plans Organization#send_passing_statuses_for_untriggered_speculative_plans}.
	SendPassingStatusesForUntriggeredSpeculativePlans interface{} `field:"optional" json:"sendPassingStatusesForUntriggeredSpeculativePlans" yaml:"sendPassingStatusesForUntriggeredSpeculativePlans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#session_remember_minutes Organization#session_remember_minutes}.
	SessionRememberMinutes *float64 `field:"optional" json:"sessionRememberMinutes" yaml:"sessionRememberMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.43.0/docs/resources/organization#session_timeout_minutes Organization#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
}

