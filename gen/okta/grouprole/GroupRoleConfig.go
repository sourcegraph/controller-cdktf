package grouprole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleConfig struct {
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
	// ID of group to attach admin roles to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#group_id GroupRole#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Admin role assigned to the group.
	//
	// It can be any one of the following values:
	// 	"API_ADMIN",
	// 	"APP_ADMIN",
	// 	"CUSTOM",
	// 	"GROUP_MEMBERSHIP_ADMIN",
	// 	"HELP_DESK_ADMIN",
	// 	"MOBILE_ADMIN",
	// 	"ORG_ADMIN",
	// 	"READ_ONLY_ADMIN",
	// 	"REPORT_ADMIN",
	// 	"SUPER_ADMIN",
	// 	"USER_ADMIN"
	// 	. See [API Docs](https://developer.okta.com/docs/reference/api/roles/#role-types).
	// 	- "USER_ADMIN" is the Group Administrator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#role_type GroupRole#role_type}
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// When this setting is enabled, the admins won't receive any of the default Okta administrator emails.
	//
	// These admins also won't have access to contact Okta Support and open support cases on behalf of your org.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#disable_notifications GroupRole#disable_notifications}
	DisableNotifications interface{} `field:"optional" json:"disableNotifications" yaml:"disableNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#id GroupRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource Set ID. Required for role_type = `CUSTOM`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#resource_set_id GroupRole#resource_set_id}
	ResourceSetId *string `field:"optional" json:"resourceSetId" yaml:"resourceSetId"`
	// Role ID. Required for role_type = `CUSTOM`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#role_id GroupRole#role_id}
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// A list of app names (name represents set of app instances, like 'salesforce' or 'facebook'), or a combination of app name and app instance ID (like 'facebook.0oapsqQ6dv19pqyEo0g3') you would like as the targets of the admin role. - Only supported when used with the role type `APP_ADMIN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#target_app_list GroupRole#target_app_list}
	TargetAppList *[]*string `field:"optional" json:"targetAppList" yaml:"targetAppList"`
	// A list of group IDs you would like as the targets of the admin role.
	//
	// - Only supported when used with the role types: `GROUP_MEMBERSHIP_ADMIN`, `HELP_DESK_ADMIN`, or `USER_ADMIN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/group_role#target_group_list GroupRole#target_group_list}
	TargetGroupList *[]*string `field:"optional" json:"targetGroupList" yaml:"targetGroupList"`
}

