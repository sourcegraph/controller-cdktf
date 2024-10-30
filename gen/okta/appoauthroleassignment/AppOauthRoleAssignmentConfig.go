package appoauthroleassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppOauthRoleAssignmentConfig struct {
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
	// Client ID for the role to be assigned to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth_role_assignment#client_id AppOauthRoleAssignment#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Role type to assign.
	//
	// This can be one of the standard Okta roles, such as `HELP_DESK_ADMIN`, or `CUSTOM`. Using custom requires the `resource_set` and `role` attributes to be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth_role_assignment#type AppOauthRoleAssignment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Resource set for the custom role to assign, must be the ID of the created resource set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth_role_assignment#resource_set AppOauthRoleAssignment#resource_set}
	ResourceSet *string `field:"optional" json:"resourceSet" yaml:"resourceSet"`
	// Custom Role ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_oauth_role_assignment#role AppOauthRoleAssignment#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

