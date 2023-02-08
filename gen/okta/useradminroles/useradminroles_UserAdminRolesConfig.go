package useradminroles

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserAdminRolesConfig struct {
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
	// User Okta admin roles - ie. ['APP_ADMIN', 'USER_ADMIN'].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_admin_roles#admin_roles UserAdminRoles#admin_roles}
	AdminRoles *[]*string `field:"required" json:"adminRoles" yaml:"adminRoles"`
	// ID of a Okta User.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_admin_roles#user_id UserAdminRoles#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// When this setting is enabled, the admins won't receive any of the default Okta administrator emails.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_admin_roles#disable_notifications UserAdminRoles#disable_notifications}
	DisableNotifications interface{} `field:"optional" json:"disableNotifications" yaml:"disableNotifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user_admin_roles#id UserAdminRoles#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

