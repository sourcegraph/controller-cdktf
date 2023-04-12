package grouprole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupRoleConfig struct {
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
	// ID of group to attach admin roles to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#group_id GroupRole#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Type of Role to assign.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#role_type GroupRole#role_type}
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// When this setting is enabled, the admins won't receive any of the default Okta administrator emails.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#disable_notifications GroupRole#disable_notifications}
	DisableNotifications interface{} `field:"optional" json:"disableNotifications" yaml:"disableNotifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#id GroupRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of apps ids for the targets of the admin role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#target_app_list GroupRole#target_app_list}
	TargetAppList *[]*string `field:"optional" json:"targetAppList" yaml:"targetAppList"`
	// List of groups ids for the targets of the admin role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/group_role#target_group_list GroupRole#target_group_list}
	TargetGroupList *[]*string `field:"optional" json:"targetGroupList" yaml:"targetGroupList"`
}

