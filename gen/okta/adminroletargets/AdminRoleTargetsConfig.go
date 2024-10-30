package adminroletargets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdminRoleTargetsConfig struct {
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
	// Type of the role that is assigned to the user and supports optional targets. See [API Docs](https://developer.okta.com/docs/reference/api/roles/#role-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_targets#role_type AdminRoleTargets#role_type}
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// User associated with the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_targets#user_id AdminRoleTargets#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// List of app names (name represents set of app instances) or a combination of app name and app instance ID (like 'salesforce' or 'facebook.0oapsqQ6dv19pqyEo0g3').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_targets#apps AdminRoleTargets#apps}
	Apps *[]*string `field:"optional" json:"apps" yaml:"apps"`
	// List of group IDs. Conflicts with apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_targets#groups AdminRoleTargets#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/admin_role_targets#id AdminRoleTargets#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

